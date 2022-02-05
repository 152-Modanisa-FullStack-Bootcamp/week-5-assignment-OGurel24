package assignment

import (
	_ "fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddUint32(t *testing.T) {
	/*
		Sum uint32 numbers, return uint32 sum value and boolean overflow flag
		cases need to pass:
			math.MaxUint32, 1 => 0, true
			1, 1 => 2, false
			42, 2701 => 2743, false
			42, math.MaxUint32 => 41, true
			4294967290, 5 => 4294967295, false
			4294967290, 6 => 0, true
			4294967290, 10 => 4, true
	*/
	m := make(map[float64]float64)
	m[42.42] = 42.50
	m[42] = 42
	m[42.01] = 42.25
	m[42.24] = 42.25
	m[42.25] = 42.25
	m[42.26] = 42.50
	m[42.55] = 42.75
	m[42.75] = 42.75
	m[42.76] = 43
	m[42.99] = 43
	m[43.13] = 43.25


	sum, overflow := AddUint32(math.MaxUint32, 1)

	assert.Equal(t, uint32(0), sum)
	assert.True(t, overflow)
}

func TestCeilNumber(t *testing.T) {

	m := make(map[float64]float64)
	m[42.42] = 42.50
	m[42] = 42
	m[42.01] = 42.25
	m[42.24] = 42.25
	m[42.25] = 42.25
	m[42.26] = 42.50
	m[42.55] = 42.75
	m[42.75] = 42.75
	m[42.76] = 43
	m[42.99] = 43
	m[43.13] = 43.25

	for k, v := range m {
		assert.Equal(t, CeilNumber(k), v)
	}
}

func TestAlphabetSoup(t *testing.T) {
	m := make(map[string]string)

	m["hello"] = "ehllo"
	m[""] = ""
	m["h"] = "h"
	m["ab"] = "ab"
	m["ba"] = "ab"
	m["bac"] = "abc"
	m["cba"] = "abc"




	for k, v := range m {
		assert.Equal(t, AlphabetSoup(m[k]), v)
	}
}

func TestStringMask(t *testing.T) {
	/*
		Replace after n(uint) character of string with '*' character.
		cases need to pass:
			"!mysecret*", 2 => "!m********"
			"", n(any positive number) => "*"
			"a", 1 => "*"
			"string", 0 => "******"
			"string", 3 => "str***"
			"string", 5 => "strin*"
			"string", 6 => "******"
			"string", 7(bigger than len of "string") => "******"
			"s*r*n*", 3 => "s*r***"
	*/
	result := StringMask("!mysecret*", 2)

	assert.Equal(t, "!m********", result)
}

func TestWordSplit(t *testing.T) {
	words := "apple,bat,cat,goodbye,hello,yellow,why"
	/*
		Your goal is to determine if the first element in the array can be split into two words,
		where both words exist in the dictionary(words variable) that is provided in the second element of array.

		cases need to pass:
			[2]string{"hellocat",words} => hello,cat
			[2]string{"catbat",words} => cat,bat
			[2]string{"yellowapple",words} => yellow,apple
			[2]string{"",words} => not possible
			[2]string{"notcat",words} => not possible
			[2]string{"bootcamprocks!",words} => not possible
	*/
	result := WordSplit([2]string{"hellocat", words})

	assert.Equal(t, "hello,cat", result)
}

func TestVariadicSet(t *testing.T) {
	/*
		FINAL BOSS ALERT :)
		Tip: Learn and apply golang variadic functions(search engine -> "golang variadic function" -> WOW You can really dance! )

		Convert inputs to set(no duplicate element)
		cases need to pass:
			4,2,5,4,2,4 => []interface{4,2,5}
			"bootcamp","rocks!","really","rocks! => []interface{"bootcamp","rocks!","really"}
			1,uint32(1),"first",2,uint32(2),"second",1,uint32(2),"first" => []interface{1,uint32(1),"first",2,uint32(2),"second"}
	*/
	set := VariadicSet(4, 2, 5, 4, 2, 4)

	assert.Equal(t, []interface{}{4, 2, 5}, set)
}
