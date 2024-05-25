package main

import (
	f "fmt"
	"slices"
)

var print = f.Printf

func main() {
	//parse array
	var arr = [...]int{1, 2, 3, 4, 5, 10: 11, 11: 12, 2, 3, 4, 5}
	print("%v %T\n", arr, arr)
	//on entering the any value are define the position the value without any position
	//will be started after the postion specified in the sparse array
	a := []int{1, 2, 3, 4, 5}
	b := []int{1, 2, 3, 4, 5}
	print("%t\n", slices.Equal(a, b))

	b = append(a, b...)
	print("%v\n", b)

	//go has three way to declare a slice
	//Using a slice literal
	d := []int{1, 2, 3, 4, 5}
	//Using a nil slice
	var e []int
	//Using an empty slice
	f := make([]int, 5, 9)

	print("%v %v %v\n", d, e, f)
	print("%d %d %d\n", len(d), len(e), len(f))
	print("%d %d %d\n", cap(d), cap(e), cap(f))

	//clearing a slice funtion "clear" has been added since go 1.21
	clear(d)
	print("%v %d %d\n", d, len(d), cap(d))

	g := []int{}
	print("%t %d %d\n", g == nil, len(g), cap(g))

	//g = make([]int, 0, 0) both make([]int, 0, 0) and make([]int, 0) make a zero len and zero cap slice
	g = make([]int, 0)
	print("%t %d %d\n", g == nil, len(g), cap(g))
	//Point to ponder a zero len and cap slice is not a nil slice!
	//so there are two ways to make a zero len and cap slice
	//using an empty slice literal or using a map

	h := []string{"a", "b", "c", "d"}
	i := h[:2]
	print("h len := %d h cap := %d i len := %d i cap := %d\n", len(h), cap(h), len(i), cap(i))
	test_slice()
	test_slice_two()
	test_copy()
	test_slice_to_array()
	test_strings()
	test_maps()
	test_struct()
}

func test_slice() {
	//use a three part slice expression to limit the capacity of the slice to the length of the slice
	//in this way any addition of the slice will clear a new backing array and will not modify the original array
	//But this works only for addtion in deletion ie using the clear function will also modify the original slice as well
	x := make([]string, 0, 5)
	x = append(x, "a", "b", "c", "d")
	print("%v\n", x)
	y := x[:2:2]
	z := x[2:4:4]
	y = append(y, "i", "j", "k")
	z = append(z, "y")
	x = append(x, "x")
	print("x := %v\ny := %v\nz := %v\n", x, y, z)
	print("x := %d\ny := %d\nz := %d\n", cap(x), cap(y), cap(z))

}

func test_slice_two() {
	//Beware to call the clear function on a slice if the backing array is shared by two or more slices
	//since the backing array will the cleared by the starting and ending position mentioned slice passed
	//into the clear function and all the remainging slice will also seen the changes made
	x := make([]string, 0, 5)
	x = append(x, "a", "b", "c", "d")
	print("%v\n", x)
	y := x[:2]
	z := y[:]
	clear(y)
	print("x = %v\ny = %v\ny = %v\n", x, y, z)
	//here the on clearing the since y the values from x and z are also deleted since the both share the
	//same backing array

}

func test_copy() {
	x := []int{1, 2, 3, 4}
	y := make([]int, 4)
	copy(y, x[:1])
	print("%v\n", y)
	//Copy function take a destination slice and a source slice and copy the value into the destination slice
	//on the basis of which ever one is the smaller, hence if we want to copy only the first two elements of the slice
	//then we have two options either we can make the destination slice of size two or we can make the source slice a slice of
	//size of two both are valid.
	y = make([]int, 1)
	copy(y, x)
	print("%v\n", y)

	x = []int{1, 2, 3, 4, 5}
	copy(x[:3], x[1:])
	print("%v\n", x)
	//here we are using the copy function on the same array to replace the value present in the same array
	//ie we can use the copy function on the same array
	//this can be used to copy one part of the array on another part if the array

}

func test_slice_to_array() {
	x := []int{1, 2, 3, 4, 5}
	y := [4]int(x) // convering a slice to an arrray, they do not share memory

	z := x[:2]
	print("%v %v %v\n", x, y, z)
	z[1] = 10000
	print("%v %v %v\n", x, y, z)
	m := [2]int(x)
	print("%v\n", m)
	//the below code panics at run time since the size of slice is smaller than the size of the array
	//and this is not allowed in go, ie the size of the array is allowed to be smaller but cannot be bigger
	//than the size of the slice being used to initialized it
	//p := [6]int(x)
	//print("%v\n", p)
}

func test_strings() {
	//s := "q2eqwe"
	//s[0] = "ASD"
	//in go a character is present by a rune in its proper form ie if it not ascii and needs more than a
	//byte to represent in memory like 🌞 needs more than a byte to be presented in the memory
	//it is a unicode code point and they take any were from 1 to 4 bytes to be present in the memory

	//var ch byte = '🌞'
	var ch rune = '🌞'
	var y = 'A'
	print("%s %s\n", string(ch), string(y)) // since ascii is a subset of unicode, string can be both a array of rune and an
	//array of bytes both are valid
	var s string = "Hello, 🌞"
	var bs []byte = []byte(s)
	var rs []rune = []rune(s)
	print("%v\n%v\n", bs, rs)

}

func test_maps() {
	//nil map
	var m map[string]int
	print("%t %v\n", m == nil, len(m))
	//empty map
	m = map[string]int{}
	print("%t %v\n", m == nil, len(m))

	//empty map with a capacity
	m = make(map[string]int, 10)
	print("%t %v\n", m == nil, len(m))

	// map have delete and clear fuctions that delete key value and empty a map respectively
	// in go 1.21 added maps package containing Equal functions for comparing two maps
	// when a key value is not present in the map, it returns the zero value for the map value
	//ie 0 for int, "" for string 0.0 for float32 or float64
	m = map[string]int{}
	print("%d\n", m["sdfsdf"])

	mf32 := map[string]float32{}
	print("%f\n", mf32["sdfsdf"])

	mf64 := map[string]float64{}
	print("%f\n", mf64["sdfsdf"])

	mb := map[string]bool{}
	print("%t\n", mb["sdfsdf"])

	ms := map[string]string{}
	print("%s\n", ms["sdfsdf"])

	mstruct := map[int]struct{}{}
	print("%v\n", mstruct[12])

	//in map the key should always be a comparable ie it cannot have nillable items as keys

}

func test_struct() {
	type person struct {
		name string
		age  int
		pet  string
	}
	// both declaring have the same value ie the default zero value for each member in the struct
	var john person
	doe := person{}

	print("%v\n%v\n", john, doe)

	seth := person{
		name: "seth rogen",
		age:  66,
	}

	print("%v\n", seth)
}
