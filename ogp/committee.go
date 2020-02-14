package ogp

import "sort"

// 投票を行うグループ
type Committee struct {
	Validators []*Validator
}

func NewCommittee(validators []*Validator) *Committee {
	return &Committee{
		Validators: validators,
	}
}

func (c *Committee) FindValidatorByID(id int) *Validator {
	for i := 0; i < len(c.Validators); i++ {
		if c.Validators[i].ID == id {
			return c.Validators[i]
		}
	}
	panic("Not Found")
}

func (c *Committee) ApproximateUtility(iter int) float64 {
	accs := make([]float64, len(c.Validators))
	for i := 0; i < len(c.Validators); i++ {
		accs[i] = c.Validators[i].Accuracy
	}
	pav := Pav(accs, iter)
	pri := Pri(accs, iter)
	// fmt.Println("-----------")
	// fmt.Println(pav)
	// fmt.Println(pri)
	return (1.0-0.5)*(1.0+0.01)*pav + 0.5*(1.0+12.0)*pri
}

func (c *Committee) UtilityFor3Validator() float64 {
	if len(c.Validators) != 3 {
		panic("Not 3 Validators")
	}
	a0 := c.Validators[0].Accuracy
	a1 := c.Validators[1].Accuracy
	a2 := c.Validators[2].Accuracy
	p := a0*a1 + a1*a2 + a2*a0 - 2*a0*a1*a2
	return Ecw(p, p)
}

type Committees struct {
	Committees []*Committee
	Validators []*Validator
}

func NewCommittees(committees []*Committee, validators []*Validator) *Committees {
	return &Committees{
		Committees: committees,
		Validators: validators,
	}
}

// validators: need to be sorted
func NewCommitteesByStake(validators []*Validator, committeeCount int) *Committees {
	selectedProbability := make([]float64, len(validators))
	for i := 0; i < len(selectedProbability); i++ {
		selectedProbability[i] = validators[i].Stake
	}

	committees := make([]*Committee, committeeCount)
	for i := 0; i < committeeCount; i++ {
		ps := make([]float64, len(validators))
		copy(ps, selectedProbability)
		indexs := WeightedChoice(len(validators), len(validators)/2, ps)
		sort.Ints(indexs)
		vs := make([]*Validator, len(indexs))
		for j := 0; j < len(indexs); j++ {
			vs[j] = validators[indexs[j]]
		}
		committees[i] = NewCommittee(vs)
	}

	return NewCommittees(committees, validators)
}

func (c *Committees) SelectedCountByValidator() []int {
	count := make([]int, len(c.Validators))
	for i := 0; i < len(c.Committees); i++ {
		for j := 0; j < len(c.Committees[0].Validators); j++ {
			count[c.Committees[i].Validators[j].ID]++
		}
	}
	return count
}

type Performance struct {
	Utility          float64
	Decentralization float64
	Profits          []float64
}

func NewPerformance(utility, decentralization float64, profits []float64) *Performance {
	return &Performance{
		Utility:          utility,
		Decentralization: decentralization,
		Profits:          profits,
	}
}

// 効用について
// iter=0なら数学的に計算する
// そうでないなら数値シミュレーション
// CommitteeSizeが4以下なら数学的にO(1)で計算する
func (c *Committees) Performance(iter int) *CommitteesPerformance {
	committeeSize := len(c.Committees[0].Validators)
	profits := make([]float64, len(c.Validators))

	// ValidatorIDの中から最小のものを求める
	diffID := 100000000000000000
	for i := 0; i < len(c.Validators); i++ {
		if c.Validators[i].ID < diffID {
			diffID = c.Validators[i].ID
		}
	}

	all := 0.0
	for i := 0; i < len(c.Committees); i++ {
		u := 0.0
		if iter == 0 {
			u = c.Committees[i].UtilityFor3Validator()
		} else {
			u = c.Committees[i].ApproximateUtility(iter)
		}
		all += u
		for j := 0; j < committeeSize; j++ {
			profits[c.Committees[i].Validators[j].ID-diffID] += u / float64(committeeSize)
		}
	}

	decentralization := Decentralization(profits)
	isIncentiveCompatible := IsSorterdByASC(profits)

	p := NewPerformance(all, decentralization, profits)
	return NewCommitteesPerformance(c, p, isIncentiveCompatible)
}

type CommitteesPerformance struct {
	Committees            *Committees
	Performance           *Performance
	IsIncentiveCompatible bool
}

func NewCommitteesPerformance(committees *Committees, performance *Performance, isIncentiveCompatible bool) *CommitteesPerformance {
	return &CommitteesPerformance{
		Committees:            committees,
		Performance:           performance,
		IsIncentiveCompatible: isIncentiveCompatible,
	}
}
