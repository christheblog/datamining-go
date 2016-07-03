package classifier

import(
	"org.cc/datamining/data"
)

// Classifier

type Classifier interface {
	Classify(data.Vector) (string, error)
}


// Estimator

type Estimator interface {
	Estimate(data.Vector) (float64, error)
}
