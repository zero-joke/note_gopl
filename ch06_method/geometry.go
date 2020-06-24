package geometry

import "math"

type Point struct{ x, y float64 }

// traditional function
func Distance(p, q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

// 在函数声明时，在其名字之前放上一个变量，即是一个方法。
// 这个附加的参数会将该函数附加到这 种类型上，即相当于为这种类型定义了一个独占的方法。
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

func useDistance() {
	p := Point{1, 2}
	q := Point{3, 4}
	p.Distance(q)
}

// 基于指针对象的方法
func (p *Point) scaleBy(factor float64) {
	p.x *= factor
	p.y *= factor
}

func useScaleBy() {
	r := &Point{1, 2}
	r.scaleBy(2)

	// 编译器会隐式地帮我们用&p去调用ScaleBy这个方法
	p := Point{3, 4}
	p.scaleBy(2)
}