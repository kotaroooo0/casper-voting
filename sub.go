package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"

	"gonum.org/v1/gonum/stat"

	"github.com/kotaroooo0/casper-voting/ogp"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func describeVari(v, u float64, size int) {
	indiAve := u / float64(size)
	di := math.Sqrt(v)
	fmt.Println("------------")
	fmt.Println(u)
	fmt.Println(indiAve)
	fmt.Println(di)
	fmt.Println(di / indiAve)
}

func subMain() {
	rand.Seed(time.Now().UnixNano()) // 乱数の調整

	fmt.Printf("%#v\n", stat.Variance([]float64{0.648, 0.648, 0.648, 0.0, 0.0, 0.0}, nil))
	fmt.Printf("%#v\n", p(0.9, 0.9, 0.9))
	fmt.Printf("%#v\n", p(0.6, 0.6, 0.6))
	fmt.Printf("%#v\n", p(0.9, 0.9, 0.9)+p(0.6, 0.6, 0.6))
	fmt.Printf("%#v\n", p(0.9, 0.9, 0.6))
	fmt.Printf("%#v\n", p(0.9, 0.6, 0.6))
	fmt.Printf("%#v\n", p(0.9, 0.9, 0.6)+p(0.9, 0.6, 0.6))
	fmt.Printf("%#v\n", di3(0.9, 0.9, 0.9, 0.6, 0.6, 0.6))
	fmt.Printf("%#v\n", di3(0.9, 0.9, 0.6, 0.9, 0.6, 0.6))
	return

	describeVari(0.006369, 69.79, 300)
	describeVari(0.0000533, 67.86, 300)
	describeVari(0.000002759, 69.24, 300)
	describeVari(0.0000009734, 69.31, 300)

	describeVari(0.006477, 69.85, 300)
	describeVari(0.0000492, 67.94, 300)
	describeVari(0.000002705, 69.15, 300)
	describeVari(0.000001076, 69.30, 300)

	describeVari(0.006528, 68.8, 300)
	describeVari(0.00006307, 67.67, 300)
	describeVari(0.000004152, 68.72, 300)
	describeVari(0.0000005906, 68.94, 300)

	describeVari(0.007177, 69.6, 300)
	describeVari(0.00004991, 67.93, 300)
	describeVari(0.00001268, 68.31, 300)
	describeVari(0.000001582, 69.27, 300)

	describeVari(0.006733, 69.89, 300)
	describeVari(0.00004825, 67.96, 300)
	describeVari(0.00001181, 68.40, 300)
	describeVari(0.00000189, 69.31, 300)

	describeVari(0.007391, 68.96, 300)
	describeVari(0.0000492, 67.94, 300)
	describeVari(0.00001267, 68.34, 300)
	describeVari(0.000001777, 69.24, 300)
	return

	// x1, x2 := ogp.MultivariateNormal(0.8, 0.8, 0.05, 100000)

	// v1,v2,v3の大小関係についての検証
	// a, b, c, d := 0.99, 0.8, 0.7, 0.51
	// fmt.Println("v1")
	// fmt.Println(v1(a, b, c, d))
	// fmt.Println("v2")
	// fmt.Println(v2(a, b, c, d))
	// fmt.Println("v3")
	// fmt.Println(v3(a, b, c, d))

	// 構成法の効用についての検証
	// accs := randomFloat(6)
	// a, b, c, d, e, f := accs[0], accs[1], accs[2], accs[3], accs[4], accs[5]
	// fmt.Println(accs)

	// c1 := pd(a, b, c, d, e, f)

	// c2 := pd(a, b, d, c, e, f)
	// c3 := pd(a, b, e, c, d, f)
	// c4 := pd(a, b, f, c, d, e)

	// c5 := pd(a, c, d, b, e, f)
	// c6 := pd(a, c, e, b, d, f)
	// c7 := pd(a, c, f, b, d, e)

	// c8 := pd(a, d, e, b, c, f)
	// c9 := pd(a, d, f, b, c, e)
	// c10 := pd(a, c, f, b, d, e)

	// p1 := pu(a, b, c, d, e, f)
	// p2 := pu(a, b, d, c, e, f)
	// p3 := pu(a, b, e, c, d, f)
	// p4 := pu(a, b, f, c, d, e)
	// p5 := pu(a, c, d, b, e, f)
	// p6 := pu(a, c, e, b, d, f)
	// p7 := pu(a, c, f, b, d, e)
	// p8 := pu(a, d, e, b, c, f)
	// p9 := pu(a, d, f, b, c, e)
	// p10 := pu(a, c, f, b, d, e)
	// d1 := di(a, b, c, d, e, f)
	// d2 := di(a, b, d, c, e, f)
	// d3 := di(a, b, e, c, d, f)
	// d4 := di(a, b, f, c, d, e)
	// d5 := di(a, c, d, b, e, f)
	// d6 := di(a, c, e, b, d, f)
	// d7 := di(a, c, f, b, d, e)
	// d8 := di(a, d, e, b, c, f)
	// d9 := di(a, d, f, b, c, e)
	// d10 := di(a, c, f, b, d, e)

	// fmt.Println(p1)
	// fmt.Println(d1)
	// fmt.Println("--------")
	// fmt.Println(p2)
	// fmt.Println(d2)
	// fmt.Println(p3)
	// fmt.Println(d3)
	// fmt.Println(p4)
	// fmt.Println(d4)
	// fmt.Println("--------")
	// fmt.Println(p5)
	// fmt.Println(d5)
	// fmt.Println(p6)
	// fmt.Println(d6)
	// fmt.Println(p7)
	// fmt.Println(d7)
	// fmt.Println("--------")
	// fmt.Println(p8)
	// fmt.Println(d8)
	// fmt.Println(p9)
	// fmt.Println(d9)
	// fmt.Println(p10)
	// fmt.Println(d10)

	// maxPcount := map[string]int{
	// 	"c1":  0,
	// 	"c2":  0,
	// 	"c3":  0,
	// 	"c4":  0,
	// 	"c5":  0,
	// 	"c6":  0,
	// 	"c7":  0,
	// 	"c8":  0,
	// 	"c9":  0,
	// 	"c10": 0,
	// }

	// minDcount := map[string]int{
	// 	"c1":  0,
	// 	"c2":  0,
	// 	"c3":  0,
	// 	"c4":  0,
	// 	"c5":  0,
	// 	"c6":  0,
	// 	"c7":  0,
	// 	"c8":  0,
	// 	"c9":  0,
	// 	"c10": 0,
	// }

	// count := 0
	// sum := 0.0
	// iter := 1000
	// for i := 0; i < iter; i++ {
	accs := randomFloat(6)
	a, b, c, d, e, f := accs[0], accs[1], accs[2], accs[3], accs[4], accs[5]
	p1 := ava3(a, b, c, d, e, f)
	p2 := ava3(a, b, d, c, e, f)
	p3 := ava3(a, b, e, c, d, f)
	p4 := ava3(a, b, f, c, d, e)
	p5 := ava3(a, c, d, b, e, f)
	p6 := ava3(a, c, e, b, d, f)
	p7 := ava3(a, c, f, b, d, e)
	p8 := ava3(a, d, e, b, c, f)
	p9 := ava3(a, d, f, b, c, e)
	p10 := ava3(a, e, f, b, c, d)
	d1 := di3(a, b, c, d, e, f)
	d2 := di3(a, b, d, c, e, f)
	d3 := di3(a, b, e, c, d, f)
	d4 := di3(a, b, f, c, d, e)
	d5 := di3(a, c, d, b, e, f)
	d6 := di3(a, c, e, b, d, f)
	d7 := di3(a, c, f, b, d, e)
	d8 := di3(a, d, e, b, c, f)
	d9 := di3(a, d, f, b, c, e)
	d10 := di3(a, e, f, b, c, d)
	ps := []float64{p1, p2, p3, p4, p5, p6, p7, p8, p9, p10}
	ds := []float64{d1, d2, d3, d4, d5, d6, d7, d8, d9, d10}

	// md := map[string]float64{
	// 	"c1":  d1,
	// 	"c2":  d2,
	// 	"c3":  d3,
	// 	"c4":  d4,
	// 	"c5":  d5,
	// 	"c6":  d6,
	// 	"c7":  d7,
	// 	"c8":  d8,
	// 	"c9":  d9,
	// 	"c10": d10,
	// }

	// mp := map[string]float64{
	// 	"c1":  p1,
	// 	"c2":  p2,
	// 	"c3":  p3,
	// 	"c4":  p4,
	// 	"c5":  p5,
	// 	"c6":  p6,
	// 	"c7":  p7,
	// 	"c8":  p8,
	// 	"c9":  p9,
	// 	"c10": p10,
	// }
	// maxPcount[maxC(mp)]++
	// minDcount[maxC(mp)]++
	// 	if maxC(mp) == "c10" {
	// 		fmt.Printf("%#v\n", accs)
	// 	}
	// 	if minC(md) == maxC(mp) {
	// 		count++
	// 	}
	// 	sum += stat.Correlation(ps, ds, nil)

	// }
	// fmt.Println(maxPcount)
	// fmt.Println(minDcount)
	// fmt.Println(count)
	// fmt.Println(sum / float64(iter))

	// return

	// fmt.Println(stat.Correlation(ps, ds, nil))

	p, _ := plot.New()
	pts := make(plotter.XYs, 1)
	pts[0].X = ps[0]
	pts[0].Y = ds[0]
	plot, _ := plotter.NewScatter(pts)
	plot.GlyphStyle.Color = plotutil.Color(0)
	plot.GlyphStyle.Shape = plotutil.Shape(0)
	p.Legend.Add("1", plot)
	p.Add(plot)

	pts = make(plotter.XYs, 1)
	pts[0].X = ps[1]
	pts[0].Y = ds[1]
	plot, _ = plotter.NewScatter(pts)
	plot.GlyphStyle.Color = plotutil.Color(1)
	plot.GlyphStyle.Shape = plotutil.Shape(1)
	p.Legend.Add("2", plot)
	p.Add(plot)

	pts = make(plotter.XYs, 1)
	pts[0].X = ps[2]
	pts[0].Y = ds[2]
	plot, _ = plotter.NewScatter(pts)
	plot.GlyphStyle.Color = plotutil.Color(2)
	plot.GlyphStyle.Shape = plotutil.Shape(2)
	p.Legend.Add("3", plot)
	p.Add(plot)

	pts = make(plotter.XYs, 1)
	pts[0].X = ps[3]
	pts[0].Y = ds[3]
	plot, _ = plotter.NewScatter(pts)
	plot.GlyphStyle.Color = plotutil.Color(3)
	plot.GlyphStyle.Shape = plotutil.Shape(3)
	p.Legend.Add("4", plot)
	p.Add(plot)

	pts = make(plotter.XYs, 1)
	pts[0].X = ps[4]
	pts[0].Y = ds[4]
	plot, _ = plotter.NewScatter(pts)
	plot.GlyphStyle.Color = plotutil.Color(4)
	plot.GlyphStyle.Shape = plotutil.Shape(4)
	p.Legend.Add("5", plot)
	p.Add(plot)

	pts = make(plotter.XYs, 1)
	pts[0].X = ps[5]
	pts[0].Y = ds[5]
	plot, _ = plotter.NewScatter(pts)
	plot.GlyphStyle.Color = plotutil.Color(5)
	plot.GlyphStyle.Shape = plotutil.Shape(5)
	p.Legend.Add("6", plot)
	p.Add(plot)

	pts = make(plotter.XYs, 1)
	pts[0].X = ps[6]
	pts[0].Y = ds[6]
	plot, _ = plotter.NewScatter(pts)
	plot.GlyphStyle.Color = plotutil.Color(6)
	plot.GlyphStyle.Shape = plotutil.Shape(6)
	p.Legend.Add("7", plot)
	p.Add(plot)

	pts = make(plotter.XYs, 1)
	pts[0].X = ps[7]
	pts[0].Y = ds[7]
	plot, _ = plotter.NewScatter(pts)
	plot.GlyphStyle.Color = plotutil.Color(7)
	plot.GlyphStyle.Shape = plotutil.Shape(7)
	p.Legend.Add("8", plot)
	p.Add(plot)

	pts = make(plotter.XYs, 1)
	pts[0].X = ps[8]
	pts[0].Y = ds[8]
	plot, _ = plotter.NewScatter(pts)
	plot.GlyphStyle.Color = plotutil.Color(8)
	plot.GlyphStyle.Shape = plotutil.Shape(8)
	p.Legend.Add("9", plot)
	p.Add(plot)

	pts = make(plotter.XYs, 1)
	pts[0].X = ps[9]
	pts[0].Y = ds[9]
	plot, _ = plotter.NewScatter(pts)
	plot.GlyphStyle.Color = plotutil.Color(9)
	plot.GlyphStyle.Shape = plotutil.Shape(9)
	p.Legend.Add("10", plot)
	p.Add(plot)
	p.Legend.Left = true
	p.X.Label.Text = "sum of verification accuracy"
	p.Y.Label.Text = "variance"
	p.Save(6*vg.Inch, 6*vg.Inch, ogp.FloatSliceToString(accs)+".png")

	return

	// fmt.Println(c1)
	// fmt.Println("--------")
	// fmt.Println(c2)
	// fmt.Println(c3)
	// fmt.Println(c4)
	// fmt.Println("--------")
	// fmt.Println(c5)
	// fmt.Println(c6)
	// fmt.Println(c7)
	// fmt.Println("--------")
	// fmt.Println(c8)
	// fmt.Println(c9)
	// fmt.Println(c10)

	// if c1 <= c2 && c2 <= c3 && c3 <= c4 {
	// 	fmt.Println("c1c2c3c4")
	// }
	// if c1 <= c5 && c5 <= c6 && c6 <= c7 {
	// 	fmt.Println("c1c5c6c7")
	// }
	// if c2 <= c5 && c3 <= c6 {
	// 	fmt.Println("c2c3c5c6")
	// }
	// m := map[string]float64{
	// 	"c1":  c1,
	// 	"c2":  c2,
	// 	"c3":  c3,
	// 	"c4":  c4,
	// 	"c5":  c5,
	// 	"c6":  c6,
	// 	"c7":  c7,
	// 	"c8":  c8,
	// 	"c9":  c9,
	// 	"c10": c10,
	// }

	// count := map[string]int{
	// 	"c1":  0,
	// 	"c2":  0,
	// 	"c3":  0,
	// 	"c4":  0,
	// 	"c5":  0,
	// 	"c6":  0,
	// 	"c7":  0,
	// 	"c8":  0,
	// 	"c9":  0,
	// 	"c10": 0,
	// }

	// for i := 0; i < 5000000; i++ {
	// 	accs := randomFloat(6)
	// 	a, b, c, d, e, f := accs[0], accs[1], accs[2], accs[3], accs[4], accs[5]
	// 	c1 := pd(a, b, c, d, e, f)
	// 	c2 := pd(a, b, d, c, e, f)
	// 	c3 := pd(a, b, e, c, d, f)
	// 	c4 := pd(a, b, f, c, d, e)
	// 	c5 := pd(a, c, d, b, e, f)
	// 	c6 := pd(a, c, e, b, d, f)
	// 	c7 := pd(a, c, f, b, d, e)
	// 	c8 := pd(a, d, e, b, c, f)
	// 	c9 := pd(a, d, f, b, c, e)
	// 	c10 := pd(a, c, f, b, d, e)
	// 	m := map[string]float64{
	// 		"c1":  c1,
	// 		"c2":  c2,
	// 		"c3":  c3,
	// 		"c4":  c4,
	// 		"c5":  c5,
	// 		"c6":  c6,
	// 		"c7":  c7,
	// 		"c8":  c8,
	// 		"c9":  c9,
	// 		"c10": c10,
	// 	}
	// 	// count[maxC(m)]++

	// 	if maxC(m) == "c2" {
	// 		fmt.Println(accs)
	// 		fmt.Println(c1)
	// 		fmt.Println("--------")
	// 		fmt.Println(c2)
	// 		fmt.Println(c3)
	// 		fmt.Println(c4)
	// 		fmt.Println("--------")
	// 		fmt.Println(c5)
	// 		fmt.Println(c6)
	// 		fmt.Println(c7)
	// 		fmt.Println("--------")
	// 		fmt.Println(c8)
	// 		fmt.Println(c9)
	// 		fmt.Println(c10)

	// 	}
	// }
	// fmt.Println(count)
}

func maxC(m map[string]float64) string {
	arr := make([]float64, 0)
	for _, v := range m {
		arr = append(arr, v)
	}

	sort.Sort(sort.Reverse(sort.Float64Slice(arr)))

	for k, v := range m {
		if v == arr[0] {
			return k
		}
	}

	panic("awawawawa")
}

func minC(m map[string]float64) string {
	arr := make([]float64, 0)
	for _, v := range m {
		arr = append(arr, v)
	}

	sort.Sort(sort.Float64Slice(arr))

	for k, v := range m {
		if v == arr[0] {
			return k
		}
	}

	panic("awawawawa")
}

func randomFloat(size int) []float64 {
	accs := make([]float64, size)
	acc := 0.0
	for i := 0; i < size; i++ {
		for {
			acc = rand.Float64()
			// acc = rand.NormFloat64()*0.2 + 0.8
			// 精度は0.5以下、1.0以上にはならない
			if acc > 0.5 && acc <= 1.0 && !contains(accs, acc) {
				break
			}
		}
		accs[i] = math.Round(acc*100) / 100
		// accs[i] = acc
	}

	sort.Sort(sort.Reverse(sort.Float64Slice(accs)))

	return []float64{0.94, 0.77, 0.72, 0.71, 0.7, 0.69}
	// return []float64{0.9, 0.8, 0.75, 0.7, 0.65, 0.6}
	return accs
}

func contains(s []float64, e float64) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}

func v1(a, b, c, d float64) float64 {
	return ((a + b - c - d) / 8) * ((a + b - c - d) / 8)
}

func v2(a, b, c, d float64) float64 {
	return ((a + c - b - d) / 8) * ((a + c - b - d) / 8)
}

func v3(a, b, c, d float64) float64 {
	return ((a + d - b - c) / 8) * ((a + d - b - c) / 8)
}

func p(a, b, c float64) float64 {
	return a*b + b*c + c*a - 2*a*b*c
}

func pd(a, b, c, d, e, f float64) float64 {
	return (p(a, b, c) + p(d, e, f)) / 2
}

func ava3(a, b, c, d, e, f float64) float64 {
	return (p(a, b, c) + p(d, e, f)) / 2
}

func di3(a, b, c, d, e, f float64) float64 {
	return (p(a, b, c) - p(d, e, f)) * (p(a, b, c) - p(d, e, f)) / 36
}
