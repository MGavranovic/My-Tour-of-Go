package main

import (
	"fmt"
	"math"
	"reflect"

	// "math/cmplx"
	"math/rand"
	"runtime"
	"strings"
	"time"
)

// type assigned by me
var b float64 = 3.2;
// this can be done only when type is present, if type is omitted, the value needs to be assigned
// var test1 = 1
// var test2 > expecting either type or initializer
var test int

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// STRUCTS
type TestStruct struct{
	X int
	Y int
}


func main() {
	// type is implicit > := can only be used inside of functions
	a := 2;

	///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// POINTERS
	pointerToa := &a // pointerToa will point to a's memory location
	fmt.Println("pointer memory location ************", pointerToa) // this will be location in the memory that a holds
	// this will read the value from the memory location
	fmt.Println("pointer reading value through memory location ************", *pointerToa) 
	// value of a before changing through pointer
	fmt.Println("value of a var before it was changed through it's pointer", a)
	// changing value through a pointer
	*pointerToa = 13
	fmt.Println("value of a var after it was changed through it's pointer", a)

	// TODO: testing if push works
	///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// type from initializer
	var c, d = 7455555555555548942, "string d";
	test = 12
	// if I want to convert a variable to a dif type, it needs to be done explicitly T(v) type(value)
	var test3 = float64(test)
	// if a variable is declared with no value it wll get it's default values, 0, false, ""
	var ifNoValue int

	fmt.Println("Random number generated: ", rand.Float32())

	fmt.Println(add(1,2))

	fmt.Println(split(17))

	fmt.Printf("a shorthand declared var %v of type %T \n", a, a)
	fmt.Printf("b var declared var %v of type %T \n", b, b)
	fmt.Printf("c var declared with type from initializer var %v of type %T \n", c, c)
	fmt.Printf("c var declared with type from initializer var %v of type %T \n", d, d)

	fmt.Printf("test var declared with type assigned value later=%v of type %T \n", test, test)
	fmt.Printf("test var turned into float=%v of type %T \n", test3, test3)

	// fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	// fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	// fmt.Printf("Type: %T Value: %v\n", z, z)
	fmt.Printf("Type: %T Value: %v\n", ifNoValue, ifNoValue)

	///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// >>> Struct <<<
	fmt.Printf("Printing a regular struct of type: (%T) with value: %v\n", TestStruct{1,2}, TestStruct{1,2})
	// accessing structs field
	v := TestStruct{1,2}
	v.X = 3
	fmt.Printf("Printing a field from struct of type: (%T) with value: %v\n", v.X, v.X)
	// pointers to struct
	pointerToStruct := &v
	// when accessing the value of a field through pointer, we would write it like this (*pointerToStruct).X, but Go allows pointerToStruct.X
	pointerToStruct.X = 1e9
	fmt.Printf("Printing a field from struct of type: (%T) with value: (%v) which was accessed through it's pointer\n", pointerToStruct.X, pointerToStruct.X)
	// Struck literals
	// struct can be declared with no value which would assign the variable it's zero value in this case 0
	// or can be declared with some fields and others left for implicit assignment
	v1 := TestStruct{}
	v2 := TestStruct{X: 1}
	fmt.Println("struct v1", v1, "struct v2", v2)

	// Trying to loop over a struct
	valOfStructV := reflect.ValueOf(v)
	fmt.Printf("trying to loop over struct using reflect*****\nafter that we get a value %v of type %T*****\n", valOfStructV, valOfStructV)
	fmt.Println("NumField()", valOfStructV.NumField())
	for i := 0; i < valOfStructV.NumField(); i++ {
		field := valOfStructV.Field(i)
		fmt.Printf("Field %d: %v\n", i, field.Interface() )
		// using just field instead of field.Interface() in this case would give the same value, but the value from field is of type
		// reflect.Value which isn't the actual value but the objects which holds the metadata and the actual value
		// field.Interface() extracts the actual value from reflect.Value
	}

	///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// ARRAYS
	// [n]T n-items T-type
	// arrays are static in length
	var arr [5]int
	for i := 0; i < len(arr); i++{
		arr[i] = i
	}
	fmt.Println("Array fixed size", arr)

	///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// SLICES
	// []T type
	// NOTE:
	// dynamically sized, flexible view into the elements of an array
	// slice []int = arr[1:2] 1 and 2 being indexes
	var sliceFromArr []int = arr[2:4]
	fmt.Println("printing a slice from arr",sliceFromArr)
	arr[2]= 98
	for i := 0; i < len(sliceFromArr); i++ {
		fmt.Println("printing slice elements using for loop",sliceFromArr[i])
	}

	// NOTE: 
	// slice doesn't store data, it's just describes a section of an underlying array
	// If an element of a slice is changed, the corresponding element of its array will change
	// slices that share the same array, will see the changes
	sliceFromArr[0] = 70
	fmt.Println("elements of an array after changing an element from slice", arr)
	// Creating a slice with []{1, 2, 3} first creates an array then builds a slice that references it

	// slice defaults
	// [x:y] low-x and high-y bound can be omitted and their values would be 0:len(array)
	defLowValue := arr[:2]
	defHighValue := arr[1:]
	defNoValues := arr[:]
	fmt.Println("default low value", defLowValue)
	fmt.Println("default high value", defHighValue)
	fmt.Println("default no value", defNoValues)

	// copying the slice into an array
	// that array won't be the same as the underlying array of slice sliceIntoArr, but a new array that we can create a new slice from
	sliceIntoArr := []int{1,2,3}
	var arrFromSlice [3]int
	copy(arrFromSlice[:], sliceIntoArr)
	fmt.Println("This is a slice copied into the array", arrFromSlice)
	fmt.Printf("This new array is of type (%T) and the slice is of type (%T)\n", arrFromSlice, sliceIntoArr)

	// NOTE: slices have both the length and the capacity 
	// length is the number of elements it contains, while capacity is number of elements in it's array
	// counting from the first element in the slice
	sliceCapLen := []int{1,2,3,4,5,6}
	printSlice(sliceCapLen)
	// lowering the length while capacity stays the same
	// slice will be empty while the array won't
	sliceCapLen = sliceCapLen[:0]
	printSlice(sliceCapLen)
	// lowering the length
	sliceCapLen = sliceCapLen[:3]
	printSlice(sliceCapLen)
	// reducing both the capacity of the array as well as the length of it's slice
	sliceCapLen = sliceCapLen[1:]
	printSlice(sliceCapLen)

	// NIL SLICE
	// Zero value of a slice is nil
	// nil slice capacity=0 length=0 and no underlying array
	var nilSlice []int
	fmt.Println("nil slice", nilSlice,"length =", len(nilSlice),"capacity =", cap(nilSlice))
	// we can check if slice is nil
	if(nilSlice == nil) {
		fmt.Println("It's nil slice")
	}

	// TODO: this will be required for my future project
	// Slices of slices
	slicesOfSlices := [][]string{
		[]string{"_","_","_"},  // or {"_","_","_"},
		[]string{"_","_","_"},  // or {"_","_","_"},
		[]string{"_","_","_"},  // or {"_","_","_"},
	}

	
	slicesOfSlices[0][0] = "X"
	slicesOfSlices[0][2] = "O"
	slicesOfSlices[1][0] = "X"
	slicesOfSlices[1][1] = "O"
	slicesOfSlices[2][2] = "X"
	fmt.Println(slicesOfSlices)

	for i := 0; i < len(slicesOfSlices); i++ {
		fmt.Printf("%s\n", strings.Join(slicesOfSlices[i], " "))
	}


	///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// for as while
	sum := 1
	for ; sum < 1000; {
		sum += sum
		// fmt.Println(sum)
	}
	fmt.Println(sum)

	// any for loop with no conditions will be an infinite loop
	/*
	for {
		}
	*/

	// regular for loop with conditions
	for i:= 0; i < 10; i++ {
		// adding if statements
		if i % 2 == 0 {
			fmt.Println("Printing i from regular for loop", i)
		}
	}

	// trying to print two things from the fmt.Println() will print them in the same line 
	fmt.Println(
		testIf(3,2,10), 
		testIf(3,2,8),
	)

	/////////////////////////////////////////////////////////////////////////////////////////////
	// switch statement
	// Go will run only one case that is selected (that matches) and won't run the rest
	// switch cases don't need to be ints or constants
	var switchTest int = 1
	switch switchTest{
		case 1:
			fmt.Println(switchTest)
			// break statement is provided automatically in Go
		case 2:
			fmt.Println(switchTest)
		default:
			fmt.Println("Not an int")
	}

	switch os := runtime.GOOS; os {
	case "macos": 
	fmt.Println("its", os)
	case "linux":
		fmt.Println("its", os)
	default:
		fmt.Println("its", os)
	}

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	fmt.Println(today.String())
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
// switch without statement is as switch true
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
	
	///////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// defer will wait for surrounding functions to execute
	deferTest()
	
}

// for printing slices
func printSlice(s []int){
 fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

/////////////////////////////////////////////////////////////////////////////////////////////////////
func deferTest(){
	// fmt.Println("Dobar")
	// defer will wait for surrounding functions to execute
	// defer fmt.Println("dan")
fmt.Println("brojanje")
	for i := 0; i < 10; i++{
		defer fmt.Println(i)
	}
	fmt.Println("zavrseno")
	// defer calls are pushed onto a stack, after a function returns, LIFO (last in first ou)
}

/////////////////////////////////////////////////////////////////////////////////////////////////////
func add(x, y int)int{
	return x + y
}

func testIf(x, n, lim float64)float64{
	// having v declared and assigned with value in if statement, limits it to the scope of that if statement > check the next comment
	if v:= math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n",v,lim)
	}
	// if I was to return v, outside of the if statement's scope, it would be undefined, and would throw an error at compile time
	return lim
}

func split(sum int)(x, y int){
	x = sum * 4 / 9
	y = sum - x
	return
}


////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// BASIC TYPES
/*
> bool

> string

> int  int8  int16  int32  int64
> uint uint8 uint16 uint32 uint64 uintptr

> byte // alias for uint8

> rune // alias for int32
     // represents a Unicode code point

> float32 float64

> complex64 complex128
The example shows variables of several types, and also that variable declarations may be "factored" into blocks, 
as with import statements.
*/
/*
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)
*/
/*
The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems.
 When you need an integer value you should use int unless you have a specific reason to use a sized or unsigned integer type.
*/


/*
package main

import "fmt"

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
*/


/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/* NEWTON'S method for sqrt
package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	if x < 0 {
		return math.NaN() // Return NaN for negative input
	}
	z := x
	tolerance := 1e-10
	for {
		nextZ := (z + x/z) / 2
		if math.Abs(z-nextZ) < tolerance {
			break
		}
		z = nextZ
	}
	return z
}

func main() {
	fmt.Println(Sqrt(9))
}

*/