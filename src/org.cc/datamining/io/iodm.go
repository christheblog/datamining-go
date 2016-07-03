package io

import (
	"bufio"
    "io"
    "os"
    "strings"
    "strconv"
    
    "org.cc/datamining/data"
)


func UReadCsv(path string, sep string, comment string) (*data.USet,error) {
	return ReadUSet(path, 
					func(l string) (bool) { return IsComment(l,comment) || IsEmpty(l) },
					func(l string) (*data.UVector,error) { return ToUVector(Split(l,sep)) }) 
}

// Read an unsupervised set from a file
func ReadUSet(path string, 
			  ignore func(string) (bool),
			  convert func(string) (*data.UVector,error)) (*data.USet,error) {
	f, err := os.Open(path)
	if err!=nil {
		return nil, err
	}
	
	res := make([]data.UVector,0)
	r := bufio.NewReader(f)
	
	for {
		line, _, err := r.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil{
			return nil,err			
		} else {
			sline := string(line)
			if !ignore(sline) {
				v, err := convert(sline)
				if err!=nil {
					return nil, err
				}
				res = append(res,*v)
			}			
		}
	}
	return data.USetOf(res), nil
}

// Read an unsupervised set from a file
func ReadSSet(path string, 
			  ignore func(string) (bool),
			  convert func(string) (*data.SVector,error)) (*data.SSet,error) {
	
	f, err := os.Open(path)
	if err!=nil {
		return nil, err
	}
	
	res := make([]data.SVector,0)
	r := bufio.NewReader(f)
	
	for {
		line, _, err := r.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil{
			return nil,err			
		} else {
			sline := string(line)
			if !ignore(sline) {
				v, err := convert(sline)
				if err!=nil {
					return nil, err
				}
				res = append(res,*v)
			}			
		}
	}
	return data.SSetOf(res), nil
}


// Trim a value
func Trim(s string) string {
	return strings.TrimRight(strings.TrimLeft(s," \t")," \t")
} 

// Split a line according to the provided separator
func Split(line string, sep string) []string{
	res := strings.Split(line,sep)
	return res
}

// Ignore lines starting with comment prefix
func IsComment(line string, commentPrefix string) bool {
	return strings.HasPrefix(strings.TrimLeft(line," \t"),commentPrefix)
}

// Ignore empty lines
func IsEmpty(line string) bool {
	return len(strings.TrimRight(strings.TrimLeft(line," \t")," \t")) == 0
}



// To SVector, assuming the class is the last element

func ToSVector(fields []string) (*data.SVector, error) {
	res := make([]float64, len(fields)-1)
	for i:=0; i < len(fields)-1; i++ {
		value, err := strconv.ParseFloat(Trim(fields[i]),64)
		if err != nil {
			return nil, err
		}
		res[i] = value
	}
	svec := data.SVectorOf(res, Trim(fields[len(fields)-1]))
	return &svec, nil
}


// To UVector, assuming clas is the last element

func ToUVector(fields []string) (*data.UVector, error) {
	res := make([]float64, len(fields))
	for i:=0; i < len(fields); i++ {
		value, err := strconv.ParseFloat(Trim(fields[i]),64)
		if err != nil {
			return nil, err
		}
		res[i] = value
	}
	svec := data.VectorOf(res)
	return &svec, nil
}
