package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64){
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64){
	v.X = v.X * f
	v.Y = v.Y * f
}

func main(){
	// v.Scale(5) 는 v 가 포인터가 아니라 값인데도 
	// 포인터 리시버가 있는 메서드는 자동으로 호출된다.
	// 즉, Scale 메서드가 포인터 리시버를 가졌기 때문에 
	// 편의상 Go는 v.Scale(5) 라는 것을 (&v).Scale(5) 로 해석한다.
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}