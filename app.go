package main

import (
	"fmt"
	"math"
	"reflect"
)

const (
	a = 2 + iota
	b
	c
)

type SomeInterface interface {
	Print() (ok bool)
}

type SomeOtherInterface interface {
	Read(data any, size int, len any) (res string, ok bool)
}

type Comments struct {
	Comment string
	Place   string
}

type Point struct {
	X        int
	Y        int
	Comments Comments
}

func (p Point) Print() (ok bool) {
	fmt.Printf("%#+v\n", p)
	return
}

func sum(a int, b ...int) (result int, ok bool) {
	result = a
	for _, val := range b {
		result += val
	}
	return
}

func inc(a *int) {
	*a++
}

func (p Point) Abs(a int) float64 {
	return math.Abs(float64(p.X*p.X + p.Y*p.Y + a))
}

func main() {
	i := 1

	var at SomeInterface = Point{}
	//var y interface{} = at

	someVar, ok := at.(Point)
	fmt.Println(someVar, ok)
	fmt.Println(reflect.TypeOf(at))
	inc(&i)
	fmt.Println("i = ", i)
	//b := []int{1, 2, 3, 4}
	res, ok := sum(5, 1, 2, 3)
	fmt.Println("Sum = ", res, ok)

	//for i := 0; i < 10; i++ {
	//	if i%2 == 0 {
	//		fmt.Println(i, " Is odd")
	//	} else {
	//		fmt.Println(i, " Is even")
	//	}
	//}
	//const a = 50

	var str string
	//str = "Hello"
	var p Point
	p.X = 1
	p.Y = 5
	p.Comments.Comment = "Test"
	p.Comments.Place = "Russia"

	fmt.Println(p.Abs(50))
	fmt.Println(p)
	m := map[string]int{"one": 1, "two": 2}
	fmt.Println(m)
	delete(m, "one")
	fmt.Println(m)
	var point = [5]float32{1, 2, 3, 4, 5}
	var s []int
	s1 := make([]int, 0, 5)
	fmt.Println(cap(s1), len(s1), s1)
	//s = append(s, 5)
	fmt.Println("fsdf\n", cap(s), len(s), s)
	fmt.Println(point)
	fmt.Println(point[2:4])
	fmt.Println(a)
	fmt.Println(c)
	fmt.Println(str)
	switch str {
	case "Hello":
		fmt.Println("It's a hello msg")
		fallthrough
	case "tmp":
		fmt.Println(55)
		fallthrough
	default:
		fmt.Println("Base")

	}

	fmt.Println("Hello world")
}
