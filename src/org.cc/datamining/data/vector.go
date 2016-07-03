package data

import (
	"math"
)


// Vector

type Vector interface {
	ElemAt(i int) float64
	Len() int
}

func Empty(n int) Vector {
	return UVector{make([]float64, n)}
}

func Norm(v Vector) float64 {
	res := 0.0
	for i := 0; i < v.Len(); i++ {
		res += v.ElemAt(i) * v.ElemAt(i)
	}
	return math.Sqrt(res)
}

func Dot(v1, v2 Vector) float64 {
	res := 0.0
	for i := 0; i < v1.Len(); i++ {
		res += v1.ElemAt(i) * v2.ElemAt(i)
	}
	return res
}

func Normalize(v Vector) Vector {
	res := make([]float64, v.Len())
	norm := Norm(v)
	for i := 0; i < v.Len(); i++ {
		res[i] = v.ElemAt(i) / norm
	}
	return UVector{res}
}

func Add(v1 Vector, v2 Vector) Vector {
	res := make([]float64, v1.Len())
	for i := 0; i < v1.Len(); i++ {
		res[i] = v1.ElemAt(i) + v2.ElemAt(i)
	}
	return UVector{res}
}

func Mult(v Vector, scalar float64) Vector {
	res := make([]float64, v.Len(), v.Len())
	for i := 0; i < v.Len(); i++ {
		res[i] = v.ElemAt(i) * scalar
	}
	return UVector{res}
}


// Equality of Vectors
func Eq(v1 Vector, v2 Vector) bool {
	if v1.Len()!=v2.Len() {
		return false
	} else {
		for i:=0; i< v1.Len(); i++ {
			if v1.ElemAt(i) != v2.ElemAt(i) {
				return false
			}
		}
		return true
	}
}

func NotEq(v1 Vector, v2 Vector) bool {
	return !Eq(v1,v2)
}

