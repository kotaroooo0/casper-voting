package ogp

import (
	"fmt"
	"math/rand"
	"strconv"
)

func Contains(s []int, e int) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}

func ContainsString(s []string, e string) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}

func Chunks(l []float64, n int) [][]float64 {
	res := make([][]float64, 0)
	for i := 0; i < len(l); i += n {
		fromIdx := i
		toIdx := i + n
		if toIdx > len(l) {
			toIdx = len(l)
		}
		res = append(res, l[fromIdx:toIdx])
	}
	return res
}

func ChunksValidators(l []*Validator, n int) [][]*Validator {
	res := make([][]*Validator, 0)
	for i := 0; i < len(l); i += n {
		fromIdx := i
		toIdx := i + n
		if toIdx > len(l) {
			toIdx = len(l)
		}
		res = append(res, l[fromIdx:toIdx])
	}
	return res
}

func ArrayMaxRand(s []int) int {
	max := -1
	for _, v := range s {
		if v > max {
			max = v
		}
	}

	maxIdxs := []int{}
	for i, v := range s {
		if v == max {
			maxIdxs = append(maxIdxs, i)
		}
	}
	return maxIdxs[rand.Intn(len(maxIdxs))]
}

func IsSorterdByASC(s []float64) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] < s[i+1] {
			return false
		}
	}
	return true
}

func IsSorterdIntByASC(s []int) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] < s[i+1] {
			return false
		}
	}
	return true
}

func IntSliceToString(s []int) string {
	res := ""
	for i := 0; i < len(s); i++ {
		res += strconv.Itoa(s[i]) + "/"
	}
	return res
}

func FloatSliceToString(s []float64) string {
	res := ""
	for i := 0; i < len(s); i++ {
		res += fmt.Sprint(s[i])
	}
	return res
}

func WeightedChoice(v, size int, w []float64) []int {
	// v を slice　に変換
	// ex) 5 -> [0, 1, 2, 3, 4]
	vs := make([]int, 0, v)
	for i := 0; i < v; i++ {
		vs = append(vs, i)
	}

	// weightの合計値を計算
	var sum float64
	for _, v := range w {
		sum += v
	}

	result := make([]int, 0, size)
	for i := 0; i < size; i++ {
		r := rand.Float64() * sum

		for j, v := range vs {
			r -= w[j]
			if r < 0 {
				result = append(result, v)

				// weightの合計値から選ばれたアイテムのweightを引く
				sum -= w[j]

				// 選択されたアイテムと重みを排除
				w = append(w[:j], w[j+1:]...)
				vs = append(vs[:j], vs[j+1:]...)

				break
			}
		}
	}
	return result
}
