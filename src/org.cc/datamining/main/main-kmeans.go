package main

import (
	"fmt"
	"os"

	"org.cc/datamining/algo/kmeans"
	"org.cc/datamining/data"
	"org.cc/datamining/distance"
	"org.cc/datamining/io"
)


func main() {
//func mainKMeans() {
	args := os.Args[1:]
	fmt.Println("Arguments")
	fmt.Println(args)

	// Checking arguments
	if len(args) != 1 {
		fmt.Println("Program accept a csv file as an argument - where the last item is a class")
		os.Exit(1)
	}

	// Reading file
	filename := args[0]
	fmt.Printf("Reading file %s ...\n", args[0])
	dat, err := io.ReadSSet(filename,
		io.IsEmpty, // skips empty lines if any
		func(line string) (*data.SVector, error) {
			splitted := io.Split(line, ",")
			return io.ToSVector(splitted)
		})
	if err != nil { panic(err) }
	

	fmt.Printf("File %s contains %d lines of data.\n", args[0], (*dat).Len())
	// Doing a Kmeans
	fmt.Printf("---> Sequential\n")
	centroids := kmeans.KMeans(dat, distance.Euclidean, kmeans.InitWithFirstVectors, kmeans.MaxIter(50), 3)
	for i := range centroids {
		fmt.Printf("Centroid %v\n",centroids[i])
	}
	// Doing a ParKmeans
	fmt.Printf("---> Parallel\n")
	centroids = kmeans.ParKMeans(dat, distance.Euclidean, kmeans.InitWithFirstVectors, kmeans.MaxIter(50), 3,5)
	for i := range centroids {
		fmt.Printf("Centroid %v\n",centroids[i])
	}
	
}

