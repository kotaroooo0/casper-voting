package ogp

import (
	"math/rand"
	"sort"

	"gonum.org/v1/gonum/stat"
)

// ------------------------------ Utility -----------------------------

// 精度から意思決定
// 提案されたブロックに対して正しい投票ができるかどうか
func Accuracy2Dicision(accuracy float64, proposedBlock bool) bool {
	if accuracy > rand.Float64() {
		return proposedBlock
	} else {
		return !proposedBlock
	}
}

// n人の時の効用
// 2/3以上の承認でブロックを承認するというルール
// 人数が多い時は数学的に計算することができないのでモンテカルロシュミレーションによる近似
func ApproximateUtility(accuracy []float64, iter int) float64 {
	pav := Pav(accuracy, iter)
	pri := Pri(accuracy, iter)
	return Ecw(pav, pri)
}

// 多数決による意思決定
func MajorityRules(accuracy []float64, proposedBlock bool) bool {
	size := len(accuracy)
	decisions := make([]bool, len(accuracy))
	for i := 0; i < size; i++ {
		decisions[i] = Accuracy2Dicision(accuracy[i], proposedBlock)
	}

	sum := 0
	for i := 0; i < size; i++ {
		if decisions[i] {
			sum++
		} else {
			sum--
		}
	}

	if float64(sum) >= (2.0*2.0/3.0-1.0)*float64(size) {
		return true
	}
	return false
}

// probability of accepting valid block
func Pav(accuracy []float64, iter int) float64 {
	opinions := make([]bool, iter)
	for i := 0; i < iter; i++ {
		opinions[i] = MajorityRules(accuracy, true)
	}

	approveCount := 0
	for i := 0; i < iter; i++ {
		if opinions[i] {
			approveCount++
		}
	}

	return float64(approveCount) / float64(iter)
}

// probability of rejecting invalid block
func Pri(accuracy []float64, iter int) float64 {
	opinions := make([]bool, iter)
	for i := 0; i < iter; i++ {
		opinions[i] = MajorityRules(accuracy, false)
	}

	rejectCount := 0
	for i := 0; i < iter; i++ {
		if !opinions[i] {
			rejectCount++
		}
	}

	return float64(rejectCount) / float64(iter)
}

// expected collective walfare
func Ecw(pav, pri float64) float64 {
	return (1.0-0.5)*(1.0+0.01)*pav + 0.5*(1.0+12)*pri
}

// ------------------------------ Decentralization ------------------------------

// ジニ係数
func Gini(profits []float64) float64 {

	size := len(profits)
	if size == 1 {
		return 1.0
	}

	sort.Float64s(profits)

	var profitSum float64 = 0
	for i := 0; i < len(profits); i++ {
		profitSum += profits[i]
	}

	profitRate := make([]float64, size)
	for i := 0; i < len(profits); i++ {
		profitRate[i] = profits[i] / profitSum
	}

	cumulativeProfits := make([]float64, size)
	for i := 0; i < len(profits); i++ {
		cumulativeProfit := 0.0
		for j := 0; j < i+1; j++ {
			cumulativeProfit += profitRate[j]
		}
		cumulativeProfits[i] = cumulativeProfit
	}

	res := cumulativeProfits[0] / float64(size) / 2
	for i := 1; i < len(profits); i++ {
		res += (cumulativeProfits[i-1] + cumulativeProfits[i]) / float64(size) / 2
	}

	return 1.0 - res*2
}

// ジニ係数は大きいほど偏りが大きいことを示しており、偏りが大きいということは中央集権であることを示している
// ジニ係数が小さければ小さいほど良いということになる
func DecentralizationByGini(profits []float64) float64 {
	return Gini(profits)
}

// 分散は大きいほど偏りが大きいことを示しており、偏りが大きいということは中央集権であることを示している
// 分散が小さければ小さいほど良いということになる
func DecentralizationByVariance(profits []float64) float64 {
	return stat.Variance(profits, nil)
}

// どちらの指標を用いるか決定する
func Decentralization(profits []float64) float64 {
	return DecentralizationByVariance(profits)
}
