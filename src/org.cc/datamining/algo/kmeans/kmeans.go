package kmeans

import(
	"org.cc/datamining/distance"
	"org.cc/datamining/data"
)


// Standard K-means algorithm
func KMeans(usptr data.Dataset, 
			distfn distance.DistanceFn,
			init func(*data.Dataset,int) ([]data.Vector),
			stop func(int,[]data.Vector,[]data.Vector) (bool), // stop criterion
			k int) []data.Vector {
	nextfn := nextIter
	return kmeansAlgo(&usptr,distfn,init,nextfn,stop,k)
}

// Parallel K-means
func ParKMeans(usptr data.Dataset, 
			   distfn distance.DistanceFn,
			   init func(*data.Dataset,int) ([]data.Vector),
			   stop func(int,[]data.Vector,[]data.Vector) (bool), // stop criterion
			   k int,
			   threshold int) []data.Vector {
	nextfn := func(uset *data.Dataset, distfn distance.DistanceFn, centroids []data.Vector) ([]VecWithCount) { return parNextIter(&usptr,distfn,centroids,threshold) }
	return kmeansAlgo(&usptr,distfn,init,nextfn,stop,k)
}


// K-Means algorithm structure
func kmeansAlgo(usptr *data.Dataset, 
			    distfn distance.DistanceFn,
			    init func(*data.Dataset,int) ([]data.Vector),
			    next func(*data.Dataset,distance.DistanceFn,[]data.Vector) ([]VecWithCount), // next iteration function
			    stop func(int,[]data.Vector,[]data.Vector) (bool), // stop criterion
			    k int) []data.Vector {
	// Initializing centroids
	centroids := init(usptr,k)
	// Looping
	for iter, cont := 0, true; cont; iter++ {
		prev := centroids
		centroids = toCentroids(next(usptr,distfn,centroids))
		cont = !stop(iter,prev,centroids)
	}
	return centroids
}

// Initialize centroids with the k first vectors
func InitWithFirstVectors(usptr *data.Dataset, k int) ([]data.Vector) {
	us := *usptr
	res:= make([]data.Vector,k)
	for i:=0; i < k; i++ {
		res[i] = us.VectorAt(i)
	}
	return res
}

// Max iteration criterion
func MaxIter(max int) (func(int,[]data.Vector,[]data.Vector) (bool)) {
	return func(n int,_ []data.Vector,_ []data.Vector) (bool) { return n > max }
}

// Centroids are stable (at epslion)
func StableIter(epsilon float64) (func(int,[]data.Vector,[]data.Vector) (bool)) {
	return func(n int, prev []data.Vector, curr []data.Vector) (bool) { return false } // TODO
}



// Aggregated Vector with count 
type VecWithCount struct {
	data.Vector
	int
}
func toCentroids( vecs []VecWithCount) []data.Vector {
	res := make([]data.Vector,len(vecs))
	for i := range vecs {
		res[i] = data.Mult(vecs[i].Vector, 1.0 / float64(vecs[i].int))
	}
	return res
}

// Compute one iteration of KMeans algorithm
func nextIter(usptr *data.Dataset, distfn distance.DistanceFn, centers []data.Vector) []VecWithCount {
	us := *usptr
	k := len(centers)
	dim := centers[0].Len()
	centroids := make([]VecWithCount, k)
	for i := range centroids { centroids[i] =  VecWithCount{ data.Empty(dim), 0} }
	for i:=0; i < us.Len(); i++ {
		vec := us.VectorAt(i)
		// Nearest-neighbour
		index := nearest(distfn,vec,centers)
		centroids[index] = VecWithCount { data.Add(centroids[index].Vector,vec), centroids[index].int + 1 }
	}
	return centroids
}

// Finds the nearest centroid for vec according to provided distance function
func nearest(distfn distance.DistanceFn, vec data.Vector, centroids []data.Vector) int {
	_, res, _ := distance.Closest(distfn, vec, &centroids)
	return res
}


// Parallel version of k-means

// Divide and conquer k-mean next iteration 
func parNextIter(usptr *data.Dataset, distfn distance.DistanceFn, centers []data.Vector, threshold int) []VecWithCount {
	if (*usptr).Len() < threshold {
		return nextIter(usptr,distfn,centers)
	} else {
		us := *usptr
		left, right := data.Split(us)
		out := make(chan []VecWithCount)
		go func(out chan<- []VecWithCount) { out <- parNextIter(&left, distfn, centers, threshold) }(out)
		go func(out chan<- []VecWithCount) { out <- parNextIter(&right, distfn, centers, threshold) }(out)
		part1, part2 := <-out, <-out
		close(out)
		return mergekmeans(part1, part2)
	}
}

func mergekmeans(xs []VecWithCount, ys []VecWithCount) []VecWithCount {
	res := make([]VecWithCount,len(xs))
	for i := range xs {
		x, y := xs[i], ys[i]
		res[i] = VecWithCount { data.Add(x.Vector,y.Vector), x.int + y.int }
	}
	return res
}
