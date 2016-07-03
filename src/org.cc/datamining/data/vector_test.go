package data

import (
    "testing"
)

func TestEmpty(t *testing.T) {
	empty := Empty(10)
	if empty.Len() != 0 {
		t.Error("Length of Empty(10) is not 10, it is ",empty.Len())
	}
	for i:=0; i < empty.Len(); i++ {
		if empty.ElemAt(i)!=0 {
			t.Errorf("Element at %i is not 0 as expected, but %f",i,empty.ElemAt(i))
		}
	}
}

func TestNorm(t *testing.T) {
	v := VectorOf([]float64{ 1.0, 2.0, 3.0 })
	expected := 3.74166
	if v.Len() != 3 {
		t.Error("Length of vector is not 3, it is ",v.Len())
	}
	
	if Norm(v)!= 3.74166 {
		t.Errorf("Norm is not %f as expected, but %f",expected, Norm(v))
	}
}

func TestAdd(t *testing.T) {
	v1 := VectorOf([]float64{ 1.0, 2.0, 3.0 })
	v2 := VectorOf([]float64{ 3.0, 2.0, 1.0 })
	expected := VectorOf([]float64{ 4.0, 4.0, 4.0 })
	if Eq(Add(v1,v2), expected) {
		t.Error("Error while adding 2 vectors")
	}
}

