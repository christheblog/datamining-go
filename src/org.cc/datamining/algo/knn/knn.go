package knn

import (
	"sort"
	
	"org.cc/datamining/distance"
	"org.cc/datamining/data"
)


// K nearest-neighbours algorithm
func Knn(ssptr *data.SSet, distfn distance.DistanceFn, v data.Vector, k int) []VecWithDist {
	ss := *ssptr
	res := make([]VecWithDist, 0)
	threshold := distfn(ss.VectorAt(0), v)
	for i := 0; i < ss.Len(); i++ {
		vec := ss.SVectorAt(i)
		dist := distfn(vec, v)
		// We should use a kind of Bounded Heap here
		if len(res) < k {
			res = append(res, VecWithDist{vec, dist})
			sort.Sort(ByDistance(res))
			threshold = res[len(res)-1].Dist
		} else if dist <= threshold {
			res[k-1] = VecWithDist{vec, dist}
			sort.Sort(ByDistance(res))
			threshold = res[k-1].Dist
		}
	}
	return res
}


// Divide-and-conquer parallel algorithm
func ParKnn(ssptr *data.SSet, distfn distance.DistanceFn, v data.Vector, k int, threshold int) []VecWithDist {
	ss := *ssptr
	if ss.Len() < threshold {
		return Knn(ssptr, distfn, v, k)
	} else {
		left, right := ss.Split()
		out := make(chan []VecWithDist)
		go func(out chan<- []VecWithDist) { out <- ParKnn(&left, distfn, v, k, threshold) }(out)
		go func(out chan<- []VecWithDist) { out <- ParKnn(&right, distfn, v, k, threshold) }(out)
		part1, part2 := <-out, <-out
		close(out)
		return mergeknn(k, part1, part2)
	}
}

// Merges results from 2 Knn calls
func mergeknn(k int, s1 []VecWithDist, s2 []VecWithDist) []VecWithDist {
	// min function for integers ... cannot find it in the math package
	min := func(i1 int, i2 int) int {
		if i1 < i2 { 
			return i1 
		} else { 
			return i2 
		}
	}

	temp := make([]VecWithDist, 0)
	res := make([]VecWithDist, 0)
	// TODO A slice merge-sort like function here ...
	temp = append(temp, s1...)
	temp = append(temp, s2...)
	sort.Sort(ByDistance(temp))
	// Copying only first k elements
	res = append(res, temp[0:min(k, len(temp))]...)
	return res
}

// Compute a class from a list of Supervised Vectors by counting classes
func ClassFrom(s []VecWithDist) (string) {
	counter := make(map[string]int)
	max := 0
	winner := s[0].Vec.Class
	for i := range s {
		cl := s[i].Vec.Class
		counter[cl] = counter[cl] + 1
		if counter[cl] > max {
			max = counter[cl]
			winner = cl 		
		} 
	}
	return winner
}

// Compute a class from a list of Supervised Vectors, by weighting by the distance
func WeightedClassFrom(s []VecWithDist) (string) {
	counter := make(map[string]float64)
	max := 0.0
	winner := s[0].Vec.Class
	for i := range s {
		cl := s[i].Vec.Class
		counter[cl] = counter[cl] + (1.0 / s[i].Dist) 
		if counter[cl] > max {
			max = counter[cl]
			winner = cl
		}
	}
	return winner
}


type VecWithDist struct {
	Vec  data.SVector
	Dist float64
}

type ByDistance []VecWithDist
func (a ByDistance) Len() int           { return len(a) }
func (a ByDistance) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDistance) Less(i, j int) bool { return a[i].Dist < a[j].Dist }
