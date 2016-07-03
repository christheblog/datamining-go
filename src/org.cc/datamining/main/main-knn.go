package main

import (
	"fmt"
	"os"

	"org.cc/datamining/algo/knn"
	"org.cc/datamining/data"
	"org.cc/datamining/distance"
	"org.cc/datamining/io"
)

//func main() {
func mainKnn() {
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
	
	if err != nil {
		panic(err)
	}
	

	fmt.Printf("File %s contains %d lines of data.\n", args[0], (*dat).Len())
	// Doing a Knn
	neighbours := knn.Knn(dat, distance.Euclidean, data.VectorOf([]float64 {5,3.3,1.4,0.2} ), 7)
	for i := 0; i < len(neighbours); i++ {
		fmt.Printf("%v - Class %s at distance %f\n", neighbours[i], neighbours[i].Vec.Class, neighbours[i].Dist)
	}
	// Using the neighbours to identify the class
	fmt.Printf("\nPredicted class = %s", knn.ClassFrom(neighbours))
	fmt.Printf("\nPredicted weighted class = %s", knn.WeightedClassFrom(neighbours))
}
