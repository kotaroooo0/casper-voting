package ogp

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"

	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/stat"
)

type Validator struct {
	ID       int
	Accuracy float64
	Stake    float64
}

func NewValidator(id int, accuracy, stake float64) *Validator {
	return &Validator{
		ID:       id,
		Accuracy: accuracy,
		Stake:    stake,
	}
}

func NewValidatorsByRand(size int, isSorted bool) []*Validator {
	accs := make([]float64, size)
	acc := 0.0
	for i := 0; i < size; i++ {
		for {
			acc = rand.Float64()
			// 精度は0.5以下にはならない
			if acc > 0.5 {
				break
			}
		}
		accs[i] = acc
	}

	// ソートするかどうか
	if isSorted {
		sort.Sort(sort.Reverse(sort.Float64Slice(accs)))
	}

	vs := make([]*Validator, size)
	for i := 0; i < size; i++ {
		vs[i] = NewValidator(i, accs[i], 0.0)
	}
	return vs
}

func NewValidatorsByNorm(size int, isSorted bool) []*Validator {
	accs := make([]float64, size)
	acc := 0.0
	for i := 0; i < size; i++ {
		for {
			acc = rand.NormFloat64()*0.1 + 0.8
			// 精度は0.5以下、1.0以上にはならない
			if acc > 0.5 && acc <= 1.0 {
				break
			}
		}
		accs[i] = acc
	}

	// ソートするかどうか
	if isSorted {
		sort.Sort(sort.Reverse(sort.Float64Slice(accs)))
	}

	vs := make([]*Validator, size)
	for i := 0; i < size; i++ {
		vs[i] = NewValidator(i, accs[i], 0.0)
	}
	return vs
}

// size: バリデータの人数
// isSorted: accuracy順にソートするかどうか
// dist: 0->Rand, 1->Norm
// corr: 0->比例, 1->反比例, 2->無相関
func NewValidators(size int, isSorted bool, dist, corr int) []*Validator {
	accs := make([]float64, size)
	acc := 0.0
	for i := 0; i < size; i++ {
		for {
			if dist == 0 {
				// Rand
				acc = rand.Float64()
			} else if dist == 1 {
				// Norm
				acc = rand.NormFloat64()*0.1 + 0.75
			} else {
				panic("Invalid Distribution")
			}
			// 精度は0.5以下、1.0以上にはならない
			if acc > 0.5 && acc <= 1.0 {
				break
			}
		}
		accs[i] = acc
	}

	// ソートするかどうか
	if isSorted {
		sort.Sort(sort.Reverse(sort.Float64Slice(accs)))
	}

	// TODO: 相関を指定する
	// stakes := make([]float64, size)

	vs := make([]*Validator, size)
	for i := 0; i < size; i++ {
		vs[i] = NewValidator(i, accs[i], 0.0)
	}
	return vs
}

func NewValidatorsWithStake(size int, isSorted bool, corr, mean, std float64) []*Validator {
	accs, stakes := MultivariateNormal(corr, mean, std, size*10)
	oks := make([][]float64, size)
	j := 0
	for i := 0; i < size; i++ {
		for {
			j++
			if accs[j] > 0.5 && accs[j] <= 1.0 && stakes[j] > 0.5 && stakes[j] <= 1.0 {
				break
			}
		}
		oks[i] = []float64{accs[j], stakes[j]}
	}

	// ソートするかどうか
	if isSorted {
		sort.Slice(oks, func(i, j int) bool {
			return oks[i][0] > oks[j][0]
		})
	}

	vs := make([]*Validator, size)
	for i := 0; i < size; i++ {
		vs[i] = NewValidator(i, oks[i][0], oks[i][1])
	}

	return vs
}

func (v *Validator) P2Decision(proposedBlock bool) bool {
	if v.Accuracy > rand.Float64() {
		return proposedBlock
	} else {
		return !proposedBlock
	}
}

// 二次元正規分布を返す
// 相関係数: corr, 平均: mean, 標準偏差: std, サンプル数: size
func MultivariateNormal(corr, mean, std float64, size int) ([]float64, []float64) {
	x1 := make([]float64, size)
	x2 := make([]float64, size)

	for i := 0; i < size; i++ {
		y, _ := MultiNorm(mat.NewVecDense(2, []float64{0.0, 0.0}),
			mat.NewSymDense(2, []float64{1.0, corr, corr, 1.0}),
		)

		x1[i] = y.At(0, 0)*std + mean
		x2[i] = y.At(1, 0)*std + mean
	}
	fmt.Println("Correlation")
	fmt.Println(stat.Correlation(x1, x2, nil))
	fmt.Println("Mean Variance")
	fmt.Println(stat.MeanVariance(x1, nil))
	_, v := stat.MeanVariance(x1, nil)
	fmt.Println("Std")
	fmt.Println(math.Sqrt(v))
	fmt.Println("Mean Variance")
	fmt.Println(stat.MeanVariance(x2, nil))
	_, v = stat.MeanVariance(x2, nil)
	fmt.Println("Std")
	fmt.Println(math.Sqrt(v))
	return x1, x2
}

// MultiNorm returns multi-dimension normally distributed VecDense
// with average vector u and covariance matrix S.
func MultiNorm(u *mat.VecDense, S *mat.SymDense) (*mat.VecDense, error) {
	rand.Seed(time.Now().UnixNano())

	n, _ := S.Dims()
	x := make([]float64, n)
	for i := range x {
		x[i] = rand.NormFloat64()
	}

	y := mat.NewVecDense(len(x), x)

	var chol mat.Cholesky
	if ok := chol.Factorize(S); !ok {
		return nil, fmt.Errorf("covariance matrix must be poositive defined")
	}

	var L mat.TriDense
	chol.LTo(&L)

	y.MulVec(&L, y)
	y.AddVec(y, u)

	return y, nil
}

func NewValidatorsWithArbitrar(size int, isSorted bool, attackedRate float64, normalProbability float64) []*Validator {
	validators := make([]*Validator, size)

	attackedSize := int(float64(size) * attackedRate)
	nonAttackedSize := size - attackedSize
	for i := 0; i < nonAttackedSize; i++ {
		validators[i] = NewValidator(i, normalProbability, 0.51)
		// validators[i] = NewValidator(i, normalProbability, normalProbability)
		// if i%2 == 1 {
		// 	validators[i] = NewValidator(i, normalProbability, 0.51)
		// } else {
		// 	validators[i] = NewValidator(i, normalProbability, normalProbability)
		// }
	}
	for i := nonAttackedSize; i < size; i++ {
		validators[i] = NewValidator(i, 0.51, normalProbability)
		// validators[i] = NewValidator(i, 0.51, 0.51)
		// if i%2 == 1 {
		// 	validators[i] = NewValidator(i, 0.51, 0.51)
		// } else {
		// 	validators[i] = NewValidator(i, 0.51, normalProbability)

		// }
	}
	return validators
}
