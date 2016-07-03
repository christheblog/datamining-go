package data

// Dataset

type Dataset interface {
	VectorAt(i int) Vector
	Len() int
}


// Sliced view of a Dataset
type SetSlice struct {
	Source Dataset
	Start int
	End int // exclusive
}
// Dataset
func (self SetSlice) Len() (int) {	return (self.End - self.Start) }
func (self SetSlice) VectorAt(i int) (Vector) {	return self.Source.VectorAt(self.Start + i) }
// Splitting a Dataset
func Split(s Dataset) (Dataset,Dataset) {
	l := s.Len()
	if l==0 {
		return SetSlice{s,0,0}, SetSlice{s,0,0}
	} else {
		return SetSlice{s,0,l/2}, SetSlice{s,l/2,l}
	}
}

// TODO shuffled set



// Unsupervised

type UVector struct {
	elems []float64
}
func (self UVector) ElemAt(i int) float64 { return self.elems[i] }
func (self UVector) Len() int             { return len(self.elems) }

func VectorOf(elts []float64) UVector {
	return UVector{ elts }
}

type USet struct{ elems []UVector }
func (self USet) VectorAt(i int) Vector 	{ return self.elems[i] }
func (self USet) Len() int               	{ return len(self.elems) }
func (self USet) UVectorAt(i int) UVector 	{ return self.elems[i] }
func (self USet) Split() (USet, USet) {
	return USet{self.elems[0 : self.Len()/2]}, USet{self.elems[self.Len()/2:]}
}



func USetOf(vecs []UVector) (*USet) {
	return &USet{ vecs }
}


// Supervised

type SVector struct {
	elems []float64
	Class string
}
func (self SVector) ElemAt(i int) float64 { return self.elems[i] }
func (self SVector) Len() int             { return len(self.elems) }

func SVectorOf(elts []float64, cl string) SVector {
	return SVector{ elts, cl }
}


type SSet struct{ elems []SVector }
func (self SSet) VectorAt(i int) Vector { return self.elems[i] }
func (self SSet) Len() int               { return len(self.elems) }
func (self SSet) SVectorAt(i int) SVector { return self.elems[i] }
func (self SSet) Split() (SSet, SSet) {
	return SSet{self.elems[0 : self.Len()/2]}, SSet{self.elems[self.Len()/2:]}
}

func SSetOf(vecs []SVector) (*SSet) {
	return &SSet{ vecs }
}
