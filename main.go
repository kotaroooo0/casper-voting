package main

import (
	"fmt"
	"image/color"
	"math/rand"
	"strconv"
	"time"

	"github.com/kotaroooo0/casper-voting/ogp"
	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

const (
	committeeCount   = 10  // Committee形成を何回繰り返すか
	allValidatorSize = 300 // 全てのバリデータ数 6の倍数とする
	committeeSize    = 3   // 分割されたCommitteeあたりの人数
)

func main() {
	rand.Seed(time.Now().UnixNano()) // 乱数の調整
	Grouping5555FullSearch()
	// plotRealSizeHist()
	// groupingRealSize()
	// plotParetoOptimalsBySelectedCount()
}

func Grouping5555FullSearch() {
	// バリデータそれぞれの検証精度
	validators := make([]*ogp.Validator, 6)
	validators[0] = ogp.NewValidator(0, 0.99, 0.0)
	validators[1] = ogp.NewValidator(1, 0.9, 0.0)
	validators[2] = ogp.NewValidator(2, 0.8, 0.0)
	validators[3] = ogp.NewValidator(3, 0.7, 0.0)
	validators[4] = ogp.NewValidator(4, 0.6, 0.0)
	validators[5] = ogp.NewValidator(5, 0.51, 0.0)

	// DFSにより全ての組み合わせを列挙
	allCommittees = createAllCommittees(createCommitteeCombination(6))

	// インセンティブ整合を満たすものとそうでないものを区別
	incentiveCompatiblePerformance := make([]*ogp.CommitteesPerformance, 0)
	for i := 0; i < len(allCommittees); i++ {
		cs := make([]*ogp.Committee, 0) // Committees作って
		for j := 0; j < len(allCommittees[0]); j++ {
			vs := make([]*ogp.Validator, 0)
			for k := 0; k < len(allCommittees[0][0]); k++ {
				vs = append(vs, validators[allCommittees[i][j][k]])
			}
			cs = append(cs, ogp.NewCommittee(vs))
		}
		comm := ogp.NewCommittees(cs, validators)

		p := comm.Performance(0)
		if p.IsIncentiveCompatible {
			incentiveCompatiblePerformance = append(incentiveCompatiblePerformance, p)
		}
	}

	us := make([]float64, len(incentiveCompatiblePerformance))
	ds := make([]float64, len(incentiveCompatiblePerformance))
	incentiveCompatiblePts := make(plotter.XYs, len(incentiveCompatiblePerformance))
	for i := 0; i < len(incentiveCompatiblePerformance); i++ {
		incentiveCompatiblePts[i].X = incentiveCompatiblePerformance[i].Performance.Utility
		incentiveCompatiblePts[i].Y = incentiveCompatiblePerformance[i].Performance.Decentralization
		us[i] = incentiveCompatiblePerformance[i].Performance.Utility
		ds[i] = incentiveCompatiblePerformance[i].Performance.Decentralization
	}

	p, _ := plot.New()
	p.X.Label.Text = "Utility"
	p.X.Label.Font.Size = 30
	p.Y.Label.Text = "Variance"
	p.Y.Label.Font.Size = 30

	incentiveCompatiblePlot, _ := plotter.NewScatter(incentiveCompatiblePts)
	incentiveCompatiblePlot.GlyphStyle.Color = color.Black
	incentiveCompatiblePlot.GlyphStyle.Shape = plotutil.Shape(4)
	p.Add(incentiveCompatiblePlot)

	p.Save(6*vg.Inch, 6*vg.Inch, "plot_iinkai555555.png")
}

func plotRealSizeHist() {
	rand.Seed(time.Now().UnixNano()) // 乱数の調整

	// 既存手法
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	v := make([]float64, 0)
	for i := 0; i < 100; i++ {
		validators := ogp.NewValidatorsWithStake(allValidatorSize, true, 0.9, 0.7, 0.15)
		committees := ogp.NewCommitteesByStake(validators, committeeCount)
		p := committees.Performance(100).Performance
		v = append(v, p.Profits...)
	}
	h, err := plotter.NewHist(plotter.Values(v), 11)
	if err != nil {
		panic(err)
	}
	h.Normalize(1)
	p.Add(h)
	file := "hist_existing.png"
	if err := p.Save(10*vg.Inch, 6*vg.Inch, file); err != nil {
		panic(err)
	}

	// // 提案手法
	// ValidatorをN分割して、分割を1人のValidatorとして扱う
	for l := 2; l < 7; l += 2 {
		p, err := plot.New()
		if err != nil {
			panic(err)
		}
		v := make([]float64, 0)

		for m := 0; m < 30; m++ {
			validators := ogp.NewValidatorsWithStake(allValidatorSize, true, 0.9, 0.7, 0.15)
			diviedGroups := DividingToNGroups(validators, l)
			representValidators := RepresentNVadalidators(diviedGroups, l)
			betterCommittees := FindBetterCommittees(representValidators, createAllCommittees(createCommitteeCombination(l)))
			for k := 0; k < len(betterCommittees); k++ {
				realSizeByDivided := make([][]*ogp.Validator, committeeCount)
				for i := 0; i < 10; i++ {
					for j := 0; j < len(betterCommittees[k].Committees.Committees[i].Validators); j++ {
						realSizeByDivided[i] = append(realSizeByDivided[i], diviedGroups[betterCommittees[k].Committees.Committees[i].Validators[j].ID]...)
					}
				}
				cs := make([]*ogp.Committee, committeeCount)
				for i := 0; i < committeeCount; i++ {
					cs[i] = ogp.NewCommittee(realSizeByDivided[i])
				}
				realCommittees := ogp.NewCommittees(cs, validators)
				realPerformance := realCommittees.Performance(100)
				v = append(v, realPerformance.Performance.Profits...)
			}
		}
		v = append(v, 0.0)
		v = append(v, 0.44)
		h, err := plotter.NewHist(plotter.Values(v), 11)
		if err != nil {
			panic(err)
		}
		h.Normalize(1)
		// h.DataRange(0.0, 4.4, 0, 1000)
		p.Add(h)
		file := "hist_split" + strconv.Itoa(l) + ".png"
		if err := p.Save(10*vg.Inch, 6*vg.Inch, file); err != nil {
			panic(err)
		}
	}
}

// 現実の状態のような大きい数での検証について
func groupingRealSize() {
	rand.Seed(time.Now().UnixNano()) // 乱数の調整

	// サイズ: AllValidatorSizeの集団を作る
	validators := ogp.NewValidatorsWithStake(allValidatorSize, true, 0.9, 0.8, 0.3)
	// validators := ogp.NewValidatorsByNorm(allValidatorSize, true)
	// validators := ogp.NewValidatorsWithArbitrar(allValidatorSize, true, 0.3, 0.8)

	ps := make([]float64, len(validators))
	ss := make([]float64, len(validators))
	for i := 0; i < len(validators); i++ {
		ps[i] = validators[i].Accuracy
		ss[i] = validators[i].Stake
	}
	fmt.Println(stat.Correlation(ps, ss, nil))

	fmt.Println("Existing Method")
	count := make([]int, allValidatorSize)
	util := 0.0
	dece := 0.0
	for i := 0; i < 100; i++ {
		committees := ogp.NewCommitteesByStake(validators, committeeCount)
		for j := 0; j < len(committees.Committees); j++ {
			for l := 0; l < allValidatorSize/2; l++ {
				count[committees.Committees[j].Validators[l].ID]++
			}
		}
		p := committees.Performance(100).Performance
		util += p.Utility / 100.0
		dece += p.Decentralization / 100.0
	}

	fmt.Println(util)
	fmt.Println(dece)
	fmt.Println("------------------")

	// 提案手法
	// ValidatorをN分割して、分割を1人のValidatorとして扱う
	for l := 2; l < 7; l += 2 {
		diviedGroups := DividingToNGroups(validators, l)
		representValidators := RepresentNVadalidators(diviedGroups, l)
		betterCommittees := FindBetterCommittees(representValidators, createAllCommittees(createCommitteeCombination(l)))

		for k := 0; k < len(betterCommittees); k++ {
			fmt.Println("----------------------")
			realSizeByDivided := make([][]*ogp.Validator, committeeCount)
			for i := 0; i < 10; i++ {
				for j := 0; j < len(betterCommittees[k].Committees.Committees[i].Validators); j++ {
					realSizeByDivided[i] = append(realSizeByDivided[i], diviedGroups[betterCommittees[k].Committees.Committees[i].Validators[j].ID]...)
				}
			}
			cs := make([]*ogp.Committee, committeeCount)
			for i := 0; i < committeeCount; i++ {
				cs[i] = ogp.NewCommittee(realSizeByDivided[i])
			}
			realCommittees := ogp.NewCommittees(cs, validators)
			realPerformance := realCommittees.Performance(100)
			fmt.Println("realPerformance")
			fmt.Println(realPerformance)
			fmt.Println(realPerformance.Performance.Utility)
			fmt.Println(realPerformance.Performance.Decentralization)
		}
	}
}

func plotParetoOptimalsBySelectedCount() {
	// バリデータ
	validators := ogp.NewValidatorsByNorm(6, true)

	// DFSにより全ての組み合わせを列挙
	enumeratedCommittees := CreateEnumeratedCommittees()

	// インセンティブ整合を満たすものだけ
	betterCommittees := FindBetterCommittees(validators, enumeratedCommittees)

	// カウントによって色分けする
	countKind := make([]string, 0)
	for i := 0; i < len(betterCommittees); i++ {
		temp := ogp.IntSliceToString(betterCommittees[i].Committees.SelectedCountByValidator())
		if !ogp.ContainsString(countKind, temp) {
			countKind = append(countKind, temp)
		}
	}

	p, _ := plot.New()
	p.X.Label.Text = "utility"
	p.Y.Label.Text = "variance"
	p.Legend.Top = true
	p.Legend.Left = true

	for i := 0; i < len(countKind); i++ {
		tempBetterCommittees := make([]*ogp.CommitteesPerformance, 0)
		for j := 0; j < len(betterCommittees); j++ {
			temp := ogp.IntSliceToString(betterCommittees[j].Committees.SelectedCountByValidator())
			if temp == countKind[i] {
				tempBetterCommittees = append(tempBetterCommittees, betterCommittees[j])
			}
		}

		pts := make(plotter.XYs, len(tempBetterCommittees))
		for j := 0; j < len(tempBetterCommittees); j++ {
			pts[j].X = tempBetterCommittees[j].Performance.Utility
			pts[j].Y = tempBetterCommittees[j].Performance.Decentralization
		}

		plot, _ := plotter.NewScatter(pts)
		plot.GlyphStyle.Color = plotutil.Color(i)
		plot.GlyphStyle.Shape = plotutil.Shape(i)
		p.Add(plot)
		p.Legend.Add(countKind[i], plot)
	}

	accs := make([]float64, len(validators))
	for i := 0; i < len(validators); i++ {
		accs = append(accs, validators[i].Accuracy)
	}
	fmt.Println(strconv.Itoa(committeeCount) + "-" + ogp.FloatSliceToString(accs))
	p.Save(6*vg.Inch, 6*vg.Inch, "img/a"+strconv.Itoa(committeeCount)+"-"+ogp.FloatSliceToString(accs)+".png")
}

func RepresentNVadalidators(validators [][]*ogp.Validator, n int) []*ogp.Validator {
	represent := make([]*ogp.Validator, n)
	size := len(validators[0])
	for i := 0; i < n; i++ {
		acc := 0.0
		for j := 0; j < size; j++ {
			acc += validators[i][j].Accuracy / float64(size)
		}
		represent[i] = ogp.NewValidator(i, acc, 0.0)
	}
	return represent
}

func DividingToNGroups(validators []*ogp.Validator, n int) [][]*ogp.Validator {
	size := len(validators)
	if size%n != 0 {
		panic("Validators is not divided by N")
	}

	groups := make([][]*ogp.Validator, 0)
	for i := 0; i < n; i++ {
		groups = append(groups, validators[i*size/n:(i+1)*size/n])
	}
	return groups
}

// 重複した要素を削除して返却
func RemoveDuplicateCommitteesPerformance(committeesPerformance []*ogp.CommitteesPerformance) []*ogp.CommitteesPerformance {
	results := make([]*ogp.CommitteesPerformance, 0, len(committeesPerformance))
	encountered := map[float64]bool{}
	for i := 0; i < len(committeesPerformance); i++ {
		if !encountered[committeesPerformance[i].Performance.Utility] {
			encountered[committeesPerformance[i].Performance.Utility] = true
			results = append(results, committeesPerformance[i])
		}
	}
	return results
}

func ParetoOptimalsByPerformances(committeesPerformance []*ogp.CommitteesPerformance) []*ogp.CommitteesPerformance {
	paretoOptimalPoints := make([]*ogp.CommitteesPerformance, 0)
	for i := 0; i < len(committeesPerformance); i++ {

		isOptimal := true
		for j := 0; j < len(committeesPerformance); j++ {
			if i == j {
				continue
			}
			if committeesPerformance[i].Performance.Utility < committeesPerformance[j].Performance.Utility && committeesPerformance[i].Performance.Decentralization > committeesPerformance[j].Performance.Decentralization {
				isOptimal = false
				break
			}
			if len(paretoOptimalPoints) > 0 && committeesPerformance[i].Performance.Utility == committeesPerformance[j].Performance.Utility && committeesPerformance[i].Performance.Decentralization == committeesPerformance[j].Performance.Decentralization {
				isOptimal = false
				break
			}
		}
		if isOptimal {
			paretoOptimalPoints = append(paretoOptimalPoints, committeesPerformance[i])
		}
	}
	return paretoOptimalPoints
}

var allCommittees [][][]int
var committeeCombination [][][]int
var enumeratedCommittee [][]int

// Committeeの組み合わせを全列挙するDFS
func dfs(restCount int, addedCommittees, existingCommittees [][]int) {
	restCount--
	existingCommittees = append(existingCommittees, addedCommittees...)
	if restCount == 0 {
		allCommittees = append(allCommittees, existingCommittees)
	} else {
		for i := 0; i < len(committeeCombination); i++ {
			dfs(restCount, committeeCombination[i], existingCommittees)
		}
	}
}

// TODO: 枝刈り
func Dfs(restCount int, addedCommittee []int, existingCommittees [][]int) {
	restCount--
	existingCommittees = append(existingCommittees, addedCommittee)
	if restCount == 0 {
		allCommittees = append(allCommittees, existingCommittees)
	} else {
		for i := 0; i < len(enumeratedCommittee); i++ {
			Dfs(restCount, enumeratedCommittee[i], existingCommittees)
		}
	}
}

func createAllCommittees(combinations [][][]int) [][][]int {
	committeeCombination = combinations
	// DFSにより全ての組み合わせを列挙
	allCommittees = make([][][]int, 0)
	for i := 0; i < len(combinations); i++ {
		dfs(committeeCount/2, combinations[i], make([][]int, 0))
	}
	return allCommittees
}

func CreateEnumeratedCommittees() [][][]int {
	// committeeの形成の種類
	enumeratedCommittee = CreateEnumeratedCommittee()

	// DFSにより全ての組み合わせを列挙
	allCommittees = make([][][]int, 0)
	for i := 0; i < len(enumeratedCommittee); i++ {
		Dfs(committeeCount, enumeratedCommittee[i], make([][]int, 0))
	}
	return allCommittees
}

// n->2/nのときの組み合わせをハードコーディング
func createCommitteeCombination(n int) [][][]int {
	committeeCombination := make([][][]int, 0)
	if n == 6 {
		committeeCombination = append(committeeCombination, [][]int{{0, 1, 2}, {3, 4, 5}})
		committeeCombination = append(committeeCombination, [][]int{{0, 1, 3}, {2, 4, 5}})
		committeeCombination = append(committeeCombination, [][]int{{0, 1, 4}, {2, 3, 5}})
		committeeCombination = append(committeeCombination, [][]int{{0, 1, 5}, {2, 3, 4}})
		committeeCombination = append(committeeCombination, [][]int{{0, 2, 3}, {1, 4, 5}})
		committeeCombination = append(committeeCombination, [][]int{{0, 2, 4}, {1, 3, 5}})
		committeeCombination = append(committeeCombination, [][]int{{0, 2, 5}, {1, 3, 4}})
		committeeCombination = append(committeeCombination, [][]int{{0, 3, 4}, {1, 2, 5}})
		committeeCombination = append(committeeCombination, [][]int{{0, 3, 5}, {1, 2, 4}})
		committeeCombination = append(committeeCombination, [][]int{{0, 4, 5}, {1, 2, 3}})
	} else if n == 4 {
		committeeCombination = append(committeeCombination, [][]int{{0, 1}, {2, 3}})
		committeeCombination = append(committeeCombination, [][]int{{0, 2}, {1, 3}})
		committeeCombination = append(committeeCombination, [][]int{{0, 3}, {1, 2}})
	} else if n == 2 {
		committeeCombination = append(committeeCombination, [][]int{{0}, {1}})
	} else {
		panic("Invalid N")
	}
	return committeeCombination
}

// 6 -> 3の同じ数でない時の組み合わせ 6C3
// TODO: ハードコーディングださい
func CreateEnumeratedCommittee() [][]int {
	committeeCombination := make([][]int, 0)
	committeeCombination = append(committeeCombination, []int{0, 1, 2})
	committeeCombination = append(committeeCombination, []int{0, 1, 3})
	committeeCombination = append(committeeCombination, []int{0, 1, 4})
	committeeCombination = append(committeeCombination, []int{0, 1, 5})
	committeeCombination = append(committeeCombination, []int{0, 2, 3})
	committeeCombination = append(committeeCombination, []int{0, 2, 4})
	committeeCombination = append(committeeCombination, []int{0, 2, 5})
	committeeCombination = append(committeeCombination, []int{0, 3, 4})
	committeeCombination = append(committeeCombination, []int{0, 3, 5})
	committeeCombination = append(committeeCombination, []int{0, 4, 5})
	committeeCombination = append(committeeCombination, []int{3, 4, 5})
	committeeCombination = append(committeeCombination, []int{2, 4, 5})
	committeeCombination = append(committeeCombination, []int{2, 3, 5})
	committeeCombination = append(committeeCombination, []int{2, 3, 4})
	committeeCombination = append(committeeCombination, []int{1, 4, 5})
	committeeCombination = append(committeeCombination, []int{1, 3, 5})
	committeeCombination = append(committeeCombination, []int{1, 3, 4})
	committeeCombination = append(committeeCombination, []int{1, 2, 5})
	committeeCombination = append(committeeCombination, []int{1, 2, 4})
	committeeCombination = append(committeeCombination, []int{1, 2, 3})
	return committeeCombination
}

// 強い人が弱い人より少ない数しか選ばれていない場合弾く
// 最も少ない人が0の場合も弾く
func RemoveInvalidCommittees(committees [][][]int) [][][]int {
	res := make([][][]int, 0)
	for i := 0; i < len(committees); i++ {
		count := make([]int, 2*committeeSize)
		for j := 0; j < committeeCount; j++ {
			for k := 0; k < committeeSize; k++ {
				count[committees[i][j][k]]++
			}
		}

		if ogp.IsSorterdIntByASC(count) && count[2*committeeSize-1] != 0 {
			res = append(res, committees[i])
		}
	}
	return res
}

// validatorsに対して、allCommitteesの中から良さげなCommitteeを取得する
func FindBetterCommittees(validators []*ogp.Validator, allCommittees [][][]int) []*ogp.CommitteesPerformance {
	// インセンティブ整合を満たすものとそうでないものを区別
	incentiveCompatiblePerformance := make([]*ogp.CommitteesPerformance, 0)

	for i := 0; i < len(allCommittees); i++ {
		cs := make([]*ogp.Committee, 0) // Committees作って
		for j := 0; j < len(allCommittees[0]); j++ {
			vs := make([]*ogp.Validator, 0)
			for k := 0; k < len(allCommittees[0][0]); k++ {
				vs = append(vs, validators[allCommittees[i][j][k]])
			}
			cs = append(cs, ogp.NewCommittee(vs))
		}
		comm := ogp.NewCommittees(cs, validators)

		iter := 1000
		if len(allCommittees[0][0]) == 3 {
			iter = 0
		}
		p := comm.Performance(iter)
		if p.IsIncentiveCompatible {
			incentiveCompatiblePerformance = append(incentiveCompatiblePerformance, p)
		}
	}

	// パレート最適のものを出力する
	return ParetoOptimalsByPerformances(RemoveDuplicateCommitteesPerformance(incentiveCompatiblePerformance))
}
