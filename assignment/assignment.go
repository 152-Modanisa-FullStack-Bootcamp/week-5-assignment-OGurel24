package assignment

import (
	"math"
)

func AddUint32(x, y uint32) (uint32, bool) {
	return 0, false
}

func CeilNumber(f float64) float64 {
	intPart,floatPart := math.Modf(f)

	if floatPart>0.75 {
		return intPart+1
	}else if floatPart>0.5{
		return intPart+0.75
	}else if floatPart>0.25{
		return intPart+0.5
	}else if floatPart>0{
		return intPart+0.25
	}else{
		return intPart
	}
}

func AlphabetSoup(s string) string {
	return ""
}

func StringMask(s string, n uint) string {
	return ""
}

func WordSplit(arr [2]string) string {
	return ""
}

func VariadicSet(i ...interface{}) []interface{} {
	return nil
}
