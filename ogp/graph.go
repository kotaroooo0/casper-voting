package ogp

// import (
// 	"fmt"
// 	"image/color"
// 	"math/rand"

// 	"gonum.org/v1/plot"
// 	"gonum.org/v1/plot/plotter"
// 	"gonum.org/v1/plot/plotutil"
// 	"gonum.org/v1/plot/vg"
// )

// func Grouping5555FullSearch() {
// 	// バリデータそれぞれの検証精度
// 	// validators := NewValidators()
// 	validators := make([]*Validator, 6)
// 	validators[0] = NewValidator(0, 0.99, 0.0)
// 	validators[1] = NewValidator(0, 0.9, 0.0)
// 	validators[2] = NewValidator(0, 0.8, 0.0)
// 	validators[3] = NewValidator(0, 0.7, 0.0)
// 	validators[4] = NewValidator(0, 0.6, 0.0)
// 	validators[5] = NewValidator(0, 0.51, 0.0)

// 	// DFSにより全ての組み合わせを列挙
// 	allCommittees = createAllCommittees()

// 	// インセンティブ整合を満たすものとそうでないものを区別
// 	notIncentiveCompatiblePerformance := make([]CommitteePerformance, 0)
// 	incentiveCompatiblePerformance := make([]CommitteePerformance, 0)
// 	for i := 0; i < len(allCommittees); i++ {
// 		p := performanceByCommittee(allCommittees[i], accuracy)
// 		if p.isIncentiveCompatible {
// 			incentiveCompatiblePerformance = append(incentiveCompatiblePerformance, p)
// 		} else {
// 			notIncentiveCompatiblePerformance = append(notIncentiveCompatiblePerformance, p)
// 		}
// 	}

// 	incentiveCompatiblePts := make(plotter.XYs, len(incentiveCompatiblePerformance))
// 	for i := 0; i < len(incentiveCompatiblePerformance); i++ {
// 		incentiveCompatiblePts[i].X = incentiveCompatiblePerformance[i].utility
// 		incentiveCompatiblePts[i].Y = incentiveCompatiblePerformance[i].decentralization
// 	}
// 	// notIncentiveCompatiblePts := make(plotter.XYs, len(notIncentiveCompatiblePerformance))
// 	// for i := 0; i < len(notIncentiveCompatiblePerformance); i++ {
// 	// 	notIncentiveCompatiblePts[i].X = notIncentiveCompatiblePerformance[i].utility
// 	// 	notIncentiveCompatiblePts[i].Y = notIncentiveCompatiblePerformance[i].decentralization
// 	// }

// 	p, _ := plot.New()
// 	p.X.Label.Text = "utility"
// 	p.X.Label.Font.Size = 12
// 	p.Y.Label.Text = "variance"
// 	p.Legend.Top = true
// 	p.Legend.Left = true

// 	incentiveCompatiblePlot, _ := plotter.NewScatter(incentiveCompatiblePts)
// 	incentiveCompatiblePlot.GlyphStyle.Color = color.Black
// 	incentiveCompatiblePlot.GlyphStyle.Shape = plotutil.Shape(4)
// 	p.Add(incentiveCompatiblePlot)
// 	p.Legend.Add("Incentive Compatible", incentiveCompatiblePlot)

// 	// notIncentiveCompatiblePlot, _ := plotter.NewScatter(notIncentiveCompatiblePts)
// 	// notIncentiveCompatiblePlot.GlyphStyle.Color = plotutil.Color(0)
// 	// notIncentiveCompatiblePlot.GlyphStyle.Shape = plotutil.Shape(4)
// 	// p.Add(notIncentiveCompatiblePlot)
// 	// p.Legend.Add("Not Incentive Compatible", notIncentiveCompatiblePlot)

// 	p.Save(6*vg.Inch, 6*vg.Inch, "plot_iinkai"+floatSliceToString(accuracy)+".png")

// 	fmt.Println(paretoOptimalsByPerformances(incentiveCompatiblePerformance))
// }

// func groupingRandom() {
// 	accuracy := createValidatorAccuracy()        // バリデータそれぞれの検証精度
// 	countByValidator := createCountByValidator() // それぞれのバリデータが何回選ばれるか

// 	p, _ := plot.New()
// 	p.X.Label.Text = "utility"
// 	p.Y.Label.Text = "variance"
// 	p.Legend.Top = true
// 	p.Legend.Left = true

// 	// ランダムのプロット
// 	for l := 0; l < countKind; l++ {
// 		notIncentiveCompatiblePerformance := make([]CommitteePerformance, 0)
// 		incentiveCompatiblePerformance := make([]CommitteePerformance, 0)
// 		for k := 0; k < iter; k++ {
// 			// 誰を何回選ぶかを設定
// 			count := make([]int, 6)
// 			copy(count, countByValidator[l])

// 			performance := randomMethodPerformance(accuracy, count)

// 			if performance.isIncentiveCompatible {
// 				incentiveCompatiblePerformance = append(incentiveCompatiblePerformance, performance)
// 			} else {
// 				notIncentiveCompatiblePerformance = append(notIncentiveCompatiblePerformance, performance)
// 			}
// 		}

// 		incentiveCompatiblePts := make(plotter.XYs, len(incentiveCompatiblePerformance))
// 		for i := 0; i < len(incentiveCompatiblePerformance); i++ {
// 			incentiveCompatiblePts[i].X = incentiveCompatiblePerformance[i].utility
// 			incentiveCompatiblePts[i].Y = incentiveCompatiblePerformance[i].decentralization
// 		}

// 		notIncentiveCompatiblePts := make(plotter.XYs, len(notIncentiveCompatiblePerformance))
// 		for i := 0; i < len(notIncentiveCompatiblePerformance); i++ {
// 			notIncentiveCompatiblePts[i].X = notIncentiveCompatiblePerformance[i].utility
// 			notIncentiveCompatiblePts[i].Y = notIncentiveCompatiblePerformance[i].decentralization
// 		}

// 		notIncentiveCompatiblePlot, _ := plotter.NewScatter(notIncentiveCompatiblePts)
// 		notIncentiveCompatiblePlot.GlyphStyle.Color = plotutil.Color(l)
// 		notIncentiveCompatiblePlot.GlyphStyle.Shape = plotutil.Shape(4)
// 		p.Add(notIncentiveCompatiblePlot)
// 		p.Legend.Add("Not Incentive Compatible"+intSliceToString(countByValidator[l]), notIncentiveCompatiblePlot)

// 		incentiveCompatiblePlot, _ := plotter.NewScatter(incentiveCompatiblePts)
// 		incentiveCompatiblePlot.GlyphStyle.Color = color.Black
// 		incentiveCompatiblePlot.GlyphStyle.Shape = plotutil.Shape(4)
// 		p.Add(incentiveCompatiblePlot)
// 		p.Legend.Add("Incentive Compatible"+intSliceToString(countByValidator[l]), incentiveCompatiblePlot)

// 	}

// 	// ベースラインのプロット
// 	blmpts := make(plotter.XYs, 4)
// 	for i := 0; i < 4; i++ {
// 		count := make([]int, 6)
// 		copy(count, countByValidator[i])

// 		performance := baselineMethodPerformance(accuracy, count)
// 		blmpts[i].X = performance.utility
// 		blmpts[i].Y = performance.decentralization
// 	}
// 	blplot, _ := plotter.NewScatter(blmpts)
// 	blplot.GlyphStyle.Color = color.Black
// 	blplot.GlyphStyle.Shape = plotutil.Shape(5)
// 	p.Add(blplot)
// 	p.Legend.Add("baseline", blplot)

// 	// 混合法のプロット
// 	mgmpts := make(plotter.XYs, countKind)
// 	for i := 0; i < countKind; i++ {
// 		count := make([]int, 6)
// 		copy(count, countByValidator[i])

// 		performance := mixGroupingMethodPerformance(accuracy, count)
// 		mgmpts[i].X = performance.utility
// 		mgmpts[i].Y = performance.decentralization
// 	}
// 	mgplot, _ := plotter.NewScatter(mgmpts)
// 	mgplot.GlyphStyle.Color = plotutil.Color(4)
// 	mgplot.GlyphStyle.Shape = plotutil.Shape(5)
// 	p.Add(mgplot)
// 	p.Legend.Add("mix", mgplot)

// 	p.Save(6*vg.Inch, 6*vg.Inch, "plot"+floatSliceToString(accuracy)+".png")
// }

// // ------------------------------ Method ------------------------------
// // ランダムに構成する方法
// // TODO: 偏りがあるしバグがありそう
// func randomMethodPerformance(accuracy []float64, count []int) CommitteePerformance {
// 	// それぞれのCommitteeが後何人受け入れられるか
// 	restCount := []int{3, 3, 3, 3, 3, 3, 3, 3, 3, 3}

// 	// Committeeに誰を選出したかを保存
// 	committees := make([][]int, 10)
// 	for i := 0; i < 10; i++ {
// 		committees[i] = make([]int, 0)
// 	}

// 	// ランダムにCommitteeを構成
// 	// うまく構成できるようにあまり使われていないバリデータを優先する
// 	for i := 0; i < 10; i++ {
// 		for j := 0; j < restCount[i]; j++ {
// 			randInt := arrayMaxRand(count)
// 			for count[randInt] <= 0 || contains(committees[i], randInt) {
// 				randInt = rand.Intn(6)
// 			}
// 			count[randInt]--
// 			committees[i] = append(committees[i], randInt)
// 		}
// 	}

// 	return performanceByCommittee(committees, accuracy)
// }

// // 必ずインセンティブ整合性がある方法(ベースライン)
// func baselineMethodPerformance(accuracy []float64, count []int) CommitteePerformance {
// 	committees := make([][]int, 10)
// 	for i := 0; i < 10; i++ {
// 		committees[i] = make([]int, 3)
// 	}

// 	for i := 0; i < 3; i++ {
// 		for j := 0; j < count[i]; j++ {
// 			committees[j][i] = i
// 		}
// 	}

// 	for i := 0; i < 3; i++ {
// 		for j := 0; j < 4; j++ {
// 			count[i+3]--
// 			committees[9-j][i] = i + 3
// 		}
// 	}

// 	restCount := count[3] + count[4] + count[5]
// 	if restCount == 1 {
// 		committees[5][2] = 3
// 	} else if restCount == 2 {
// 		committees[5][1] = 3
// 		committees[5][2] = 4
// 	} else if restCount == 3 {
// 		committees[5][0] = 3
// 		committees[5][1] = 4
// 		committees[5][2] = 5
// 	}

// 	return performanceByCommittee(committees, accuracy)
// }

// // それなりにバランスよく混ぜた方法
// func mixGroupingMethodPerformance(accuracy []float64, count []int) CommitteePerformance {
// 	committees := make([][]int, 10)
// 	for i := 0; i < 10; i++ {
// 		committees[i] = make([]int, 3)
// 	}

// 	countIdx := 0
// 	for i := 0; i < 3; i++ {
// 		if i == 0 {
// 			for j := 0; j < 10; j++ {
// 				if count[countIdx] == 0 {
// 					countIdx++
// 				}
// 				committees[j][i] = countIdx
// 				count[countIdx]--
// 			}
// 		} else if i == 1 {
// 			if count[countIdx] == 0 {
// 				countIdx++
// 				for j := 0; j < 10; j++ {
// 					if count[countIdx] == 0 {
// 						countIdx++
// 					}
// 					committees[9-j][i] = countIdx
// 					count[countIdx]--
// 				}
// 			} else {
// 				// 1人が10回以上選ばれることはない状態のみを考えている
// 				// 10回以上選ばれるならmin(10, restCount)を考えないといけない
// 				restCount := count[countIdx]
// 				for j := 0; j < restCount; j++ {
// 					committees[j][i] = countIdx
// 				}
// 				countIdx++

// 				for j := 0; j < 10-restCount; j++ {
// 					if count[countIdx] == 0 {
// 						countIdx++
// 					}
// 					committees[9-j][i] = countIdx
// 					count[countIdx]--
// 				}
// 			}
// 		} else if i == 2 {
// 			if count[countIdx] == 0 {
// 				countIdx++
// 				for j := 0; j < 10; j++ {
// 					if count[countIdx] == 0 {
// 						countIdx++
// 					}
// 					committees[j][i] = countIdx
// 					count[countIdx]--
// 				}
// 			} else {
// 				// 1人が10回以上選ばれることはない状態のみを考えている
// 				// 10回以上選ばれるならmin(10, restCount)を考えないといけない
// 				restCount := count[countIdx]
// 				for j := 0; j < restCount; j++ {
// 					committees[9-j][i] = countIdx
// 				}
// 				countIdx++

// 				for j := 0; j < 10-restCount; j++ {
// 					if count[countIdx] == 0 {
// 						countIdx++
// 					}
// 					committees[j][i] = countIdx
// 					count[countIdx]--
// 				}
// 			}
// 		}
// 	}

// 	return performanceByCommittee(committees, accuracy)
// }

// // 良さげなCommitteeを何個かハードコーディング
// func betterCommittees() [][][]int {
// 	betterCommittees := make([][][]int, 0)
// 	betterCommittees = append(betterCommittees,
// 		[][]int{
// 			[]int{0, 1, 4},
// 			[]int{2, 3, 5},
// 			[]int{0, 1, 5},
// 			[]int{2, 3, 4},
// 			[]int{0, 4, 5},
// 			[]int{1, 2, 3},
// 			[]int{0, 2, 3},
// 			[]int{1, 4, 5},
// 			[]int{0, 2, 5},
// 			[]int{1, 3, 4},
// 		})

// 	betterCommittees = append(betterCommittees,
// 		[][]int{
// 			[]int{0, 1, 4},
// 			[]int{2, 3, 5},
// 			[]int{0, 1, 4},
// 			[]int{2, 3, 5},
// 			[]int{0, 2, 3},
// 			[]int{1, 4, 5},
// 			[]int{0, 2, 5},
// 			[]int{1, 3, 4},
// 			[]int{0, 4, 5},
// 			[]int{1, 2, 3},
// 		})

// 	betterCommittees = append(betterCommittees,
// 		[][]int{
// 			[]int{0, 1, 2},
// 			[]int{3, 4, 5},
// 			[]int{0, 3, 4},
// 			[]int{1, 2, 5},
// 			[]int{0, 3, 4},
// 			[]int{1, 2, 5},
// 			[]int{0, 3, 4},
// 			[]int{1, 2, 5},
// 			[]int{0, 3, 4},
// 			[]int{1, 2, 5},
// 		})
// 	return betterCommittees
// }
