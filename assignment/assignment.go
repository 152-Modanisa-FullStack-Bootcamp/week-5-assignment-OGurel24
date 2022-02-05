package assignment

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func AddUint32(x, y uint32) (uint32, bool) {

	if int64(x)+int64(y) > math.MaxUint32 {
		return x+y, true
	}
	return x + y, false
}

func CeilNumber(f float64) float64 {
	intPart, floatPart := math.Modf(f)

	if floatPart > 0.75 {
		return intPart + 1
	} else if floatPart > 0.5 {
		return intPart + 0.75
	} else if floatPart > 0.25 {
		return intPart + 0.5
	} else if floatPart > 0 {
		return intPart + 0.25
	} else {
		return intPart
	}
}

func AlphabetSoup(s string) string {
	//LengthOfWord := len(s)
	a := strings.Split(s, "")
	b := strings.Join(a, ",")
	sort.Strings(strings.Split(b, ""))
	fmt.Println(b)

	fmt.Println(len(a))
	/*
		for i := range s {

		}*/
	return s

	//return fmt.Sprint(LengthOfWord)
}

func StringMask(s string, n uint) string {
	var NewSlice []rune
	chars := []rune(s)

	if len(s) <= 1 {
		return "*"
	}
	if int(n) >= len(chars) {
		for i := 0; i < len(chars); i++ {
			NewSlice = append(NewSlice, '*')
		}
	} else {
		for i := uint(0); i < n; i++ {
			NewSlice = append(NewSlice, chars[i])
		}
		for i := int(n); i < len(chars); i++ {
			NewSlice = append(NewSlice, '*')
		}
	}

	return string(NewSlice)
}

func WordSplit(arr [2]string) string {
	chars := []rune(arr[0])

	for i := range chars {
		FirstPart := arr[0][0:i]
		SecondPart := arr[0][i:]
		if strings.Contains(arr[1], FirstPart) && strings.Contains(arr[1], SecondPart) {
			return fmt.Sprintf("%s,%s", FirstPart, SecondPart)
		}
	}
	return "not possible"
}

func VariadicSet(i ...interface{}) []interface{} {
	return nil
}
