package assignment

import (
	_ "fmt"
	_ "math"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddUint32(t *testing.T) {
	m := make(map[string]string)
	m["1,1"] = "2,false"
	m["42,2701"] = "2743,false"
	//m["42, math.MaxUint32"] = "41,true"
	m["4294967290,5"] = "4294967295,false"
	m["4294967290,6"] = "0,true"
	m["4294967290,10"] = "4,true"

	for k, v := range m {
		Arguments := strings.Split(k, ",")
		FirstArgument,_ := strconv.Atoi(Arguments[0])
		SecondArgument,_ := strconv.Atoi(Arguments[1])

		SeperatedResult:=strings.Split(v,",")
		FirstResult,_ :=strconv.Atoi(SeperatedResult[0])
		SecondResult,_ := strconv.ParseBool(SeperatedResult[1])

		FirstOutput,SecondOutput := AddUint32(uint32(uint(FirstArgument)),uint32(SecondArgument))
		assert.Equal(t, uint32(FirstOutput), uint32(FirstResult))
		assert.Equal(t, SecondOutput, SecondResult)

	}
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
	m := make(map[string]string)
	m["!mysecret*,2"] = "!m********"
	m[",5"] = "*"
	m["a,1"] =  "*"
	m["string,0"] =  "******"
	m["string,3"] =  "str***"
	m["string,5"] =  "strin*"
	m["string,6"] =  "******"
	m["string,7"] =  "******"
	m["s*r*n*,3"] =  "s*r***"

	for k, v := range m {
		Arguments := strings.Split(k, ",")
		StringArgument := Arguments[0]
		NumberArg, _ := strconv.Atoi(Arguments[1])

		result := StringMask(StringArgument, uint(NumberArg))
		assert.Equal(t, result, v)
	}

}

func TestWordSplit(t *testing.T) {
	words := "apple,bat,cat,goodbye,hello,yellow,why"

	m := make(map[string]string)
	m["hellocat"] = "hello,cat"
	m["catbat"] = "cat,bat"
	m["yellowapple"] = "yellow,apple"
	m[",words"] = "not possible"
	m["notcat"] = "not possible"
	m["bootcamprocks!"] = "not possible"

	for k, v := range m {
		var Arguments [2]string
		Arguments[0] = k
		Arguments[1] = words

		result := WordSplit(Arguments)
		assert.Equal(t, result, v)
	}
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
