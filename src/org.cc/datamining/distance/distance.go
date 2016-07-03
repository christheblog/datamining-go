package distance

import (
	"math"
	"org.cc/datamining/data"
)

// Distances

type DistanceFn func(data.Vector, data.Vector) float64

// Assuming v1 and v2 have the same length
func SqEuclidean(v1 data.Vector, v2 data.Vector) float64 {
	res := 0.0
	for i := 0; i < v1.Len(); i++ {
		res += math.Pow(v2.ElemAt(i)-v1.ElemAt(i), 2.0)
	}
	return res
}

func Euclidean(v1 data.Vector, v2 data.Vector) float64 {
	return math.Sqrt(SqEuclidean(v1, v2))
}


// Utils

// Find the closest Vector amongst a set of Vectors according to the provided distance function
func Closest(distfn DistanceFn, target data.Vector, vecsptr *[]data.Vector) (data.Vector, int, float64) {
	if len(*vecsptr)==0 {
		return nil,-1,0.0 
	}
	vecs := *vecsptr
	closest, windist := 0, math.MaxFloat64
	for i:=0; i < len(vecs); i++ {
		v := vecs[i]
		dist := distfn(target, v)
		if dist < windist {
			closest, windist = i, dist
		}
	}
	return vecs[closest], closest, windist
}


// Vector equality using a distance function. 
// If distance is less than epsilon, vectors are considered equals
func Eq(distfn DistanceFn, v1 data.Vector, v2 data.Vector, epsilon float64) bool {
	return distfn(v1,v2) <= epsilon 
}
