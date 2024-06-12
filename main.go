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
var b float64 = 3.2

// this can be done only when type is present, if type is omitted, the value needs to be assigned
// var test1 = 1
// var test2 > expecting either type or initializer
var test int

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// STRUCTS
type TestStruct struct {
	X int
	Y int
}

type MethodStruct struct {
	X, Y float64
}

// this will be used to point out the fact that the methods are "attached" to a type they are defined for
type TestMethod struct {
	X, Y float64
}

type NonStruct float64

// struct for maps
type MapStruct struct {
	Lat, Long float64
}

// func for func closures
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	// type is implicit > := can only be used inside of functions
	a := 2

	///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// POINTERS
	pointerToa := &a                                                // pointerToa will point to a's memory location
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
	var c, d = 7455555555555548942, "string d"
	test = 12
	// if I want to convert a variable to a dif type, it needs to be done explicitly T(v) type(value)
	var test3 = float64(test)
	// if a variable is declared with no value it wll get it's default values, 0, false, ""
	var ifNoValue int

	fmt.Println("Random number generated: ", rand.Float32())

	fmt.Println(add(1, 2))

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
	fmt.Printf("Printing a regular struct of type: (%T) with value: %v\n", TestStruct{1, 2}, TestStruct{1, 2})
	// accessing structs field
	v := TestStruct{1, 2}
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
		fmt.Printf("Field %d: %v\n", i, field.Interface())
		// using just field instead of field.Interface() in this case would give the same value, but the value from field is of type
		// reflect.Value which isn't the actual value but the objects which holds the metadata and the actual value
		// field.Interface() extracts the actual value from reflect.Value
	}

	///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// ARRAYS
	// [n]T n-items T-type
	// arrays are static in length
	var arr [5]int
	for i := 0; i < len(arr); i++ {
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
	fmt.Println("printing a slice from arr", sliceFromArr)
	arr[2] = 98
	for i := 0; i < len(sliceFromArr); i++ {
		fmt.Println("printing slice elements using for loop", sliceFromArr[i])
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
	sliceIntoArr := []int{1, 2, 3}
	var arrFromSlice [3]int
	copy(arrFromSlice[:], sliceIntoArr)
	fmt.Println("This is a slice copied into the array", arrFromSlice)
	fmt.Printf("This new array is of type (%T) and the slice is of type (%T)\n", arrFromSlice, sliceIntoArr)

	// NOTE: slices have both the length and the capacity
	// length is the number of elements it contains, while capacity is number of elements in it's array
	// counting from the first element in the slice
	sliceCapLen := []int{1, 2, 3, 4, 5, 6}
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
	fmt.Println("nil slice", nilSlice, "length =", len(nilSlice), "capacity =", cap(nilSlice))
	// we can check if slice is nil
	if nilSlice == nil {
		fmt.Println("It's nil slice")
	}

	// creating slices using make
	// make([]T, len, cap)
	// NOTE: this allocates a zeroed and returns a slice that points to that array
	sliceMake := make([]int, 0, 5)
	fmt.Printf("Slice with make ")
	printSlice(sliceMake)

	// TODO: this will be required for my future project
	// Slices of slices
	slicesOfSlices := [][]string{
		[]string{"_", "_", "_"}, // or {"_","_","_"},
		[]string{"_", "_", "_"}, // or {"_","_","_"},
		[]string{"_", "_", "_"}, // or {"_","_","_"},
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

	// Appending to a slice
	// NOTE: append(slice []T, values)
	// you can add multiple values
	// NOTE: if the corresponding array is too small to fit the new values, a new array will be allocated.
	// NOTE: the returned slice (returned from the append()) will point to the newly allocated array
	var sliceAppend []int
	sliceAppend = append(sliceAppend, 1)
	printSlice(sliceAppend)

	// adding more than 1 value using append
	sliceAppend = append(sliceAppend, 2, 3, 5, 5, 6)
	printSlice(sliceAppend)

	// NOTE: Range
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	// i will be the index, value(v) will be the current element of slice value
	for i, value := range pow {
		fmt.Println(value)
		fmt.Printf("2**%d = %d\n", i, value)
	}

	// using only index
	for i := range pow {
		fmt.Println("Using just the index", i)
	}

	// NOTE: index or value can be skipped
	/*
		for _, v := range pow
		for i, _ := range pow
	*/

	///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// NOTE: MAPS
	// similar concept with JS Objects key-value pairs
	// map maps keys to values
	// zero value of a map is nil and it has no keys and keys can't be added

	var m map[string]MapStruct
	if m == nil {
		fmt.Println("Declared map", m)
	}

	m = make(map[string]MapStruct)
	m["Banjaluka"] = MapStruct{
		40.566, 14.965,
	}
	fmt.Println("Map key value", m["Banjaluka"])
	fmt.Printf("Type = %T of map and it's value = %v\n", m, m)

	// NOTE: overwriting the original value of m
	// assignment can be done without using MapStruct

	// to assign a value to a key and not overwrite the whole thing m[key] = MapStruct{value}
	m = map[string]MapStruct{
		"Sarajevo": {15.364, 56.198},
		"Zagreb":   {25.364, -56.198},
		/*
			"Sarajevo": MapStruct{ //
				15.364, 56.198,
			},
			"Zagreb": MapStruct{
				25.364, -56.198,
			},
		*/
	}

	// NOTE: testing maps
	// map[datatype]datatype -> can use structs as structs are just a composite datatype or can use regular datatype as int, string...
	// ... in the place of both the int and string belows
	var testMap = map[string]int{
		"one": 1,
		"two": 2,
	}
	fmt.Println("Test map with ints", testMap)

	fmt.Println("Map after using Map literals", m)
	fmt.Println("map bl", m["Banjaluka"]) // this will print {0, 0} as it's value got overwritten in the block above marked with note

	// NOTE: Mutating maps
	// m[key] = value
	testMap["one"] = 2
	fmt.Println("Reassigning the value of a key", testMap["one"])
	// to get an element > element = m[key]
	elFromMap := testMap["one"]
	fmt.Println("Getting an element from the map", elFromMap)
	// to delete an element > delete(m, key)
	testMap["three"] = 3
	fmt.Println("Adding another element to the map", testMap["three"], "to map >", testMap)
	delete(testMap, "three")
	fmt.Println("Map after one element is deleted", testMap)
	// check if key is present element, ok = m[key]
	el, ok := testMap["three"]
	fmt.Println("Print value", el, "Exist in map?", ok)
	el2, ok := testMap["one"]
	fmt.Println("Print value", el2, "Exist in map?", ok)

	// NOTE: Function values
	// Functions are values too and can be passed as args for other functions
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println("Function values ****************************************************")
	fmt.Println(math.Sqrt(5*5 + 12*12))
	fmt.Println(hypot(5, 12))
	fmt.Println(math.Sqrt(3*3 + 4*4))
	fmt.Println(compute(hypot))    // hypot called with inside compute resulting in fn(3,4) > hypot(3,4)
	fmt.Println(compute(hypot))    // hypot called with inside compute resulting in fn(3,4) > hypot(3,4)
	fmt.Println(compute(math.Pow)) // math.Pow > math.Pow(3*3*3*3) 3 to the power of 4 which ar like default args
	fmt.Println("Function values ****************************************************")

	// NOTE: Function closures
	// Go functions may be closures. A closure is a function value that references variables from outside its body.
	// The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.
	// For example, the adder function returns a closure. Each closure is bound to its own sum variable.

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println("Positive", pos(i))
		fmt.Println("Negative", neg(-2*i))
	}

	// NOTE: For loop
	///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// for as while
	sum := 1
	for sum < 1000 {
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
	for i := 0; i < 10; i++ {
		// adding if statements
		if i%2 == 0 {
			fmt.Println("Printing i from regular for loop", i)
		}
	}

	// trying to print two things from the fmt.Println() will print them in the same line
	fmt.Println(
		testIf(3, 2, 10),
		testIf(3, 2, 8),
	)

	/////////////////////////////////////////////////////////////////////////////////////////////
	// switch statement
	// Go will run only one case that is selected (that matches) and won't run the rest
	// switch cases don't need to be ints or constants
	var switchTest int = 1
	switch switchTest {
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

	// NOTE: Fibonacci
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

	// NOTE: METHODS
	valueForMethods := MethodStruct{3, 4}
	testPointer := &valueForMethods
	testValPoint := *testPointer
	fmt.Println("this is the method section starting **********************************************************************")
	fmt.Println("Method on the MethodStruct", valueForMethods.Abs())
	fmt.Println("Testing pointer **********************************************", testPointer)
	fmt.Println("Testing pointer **********************************************", testValPoint)
	// trying to access the same method from a different type
	valForDifMethod := TestMethod{3, 4}
	// without defining the method again for this struct, Abs would be undefined
	fmt.Println("Testing method on a dif struct", valForDifMethod.Abs())
	// this works only if the method is defined for this struct or any datatype
	// methods on regular types
	valForRegType := NonStruct(-2.2)
	fmt.Println("Non Struct meaning regular type method", valForRegType)

	// NOTE: Pointer receivers
	pValRec := MethodStruct{3, 4}
	pValRec.Scale(10)
	fmt.Println("Pointer receiver", pValRec.Abs()) // prints 5 if the receiver is not a pointer

	// *T receiver of type T and it can't be another receiver
	// Methods with pointer receivers can modify the value to which the receiver points (as Scale does here).
	// Since methods often need to modify their receiver, pointer receivers are more common than value receivers.
	// if a regular value receiver is used on the struct in this example, it would be working on a copy of that struct
	// this is the same as for any functions argument
	// using a pointer receiver, allows us to change the value of the original struct,

	pValRec2 := MethodStruct{3, 4}
	Scale2(&pValRec2, 10) // memory location of pValRec2, and 10
	fmt.Println("Pointer in a regular func", pValRec2.Abs())
	fmt.Println("Memory location", &pValRec2) // not getting the exact memory location needs to be formatted
	// to get the proper memory location
	fmt.Printf("Proper memory location %p\n", &pValRec2)

} // NOTE: end of main func

// NOTE: METHODS
// Go doesn't have classes
// but it's possible to define methods on types (will be using MethodStruct from line 30)
// Abs is the name of the function and the receiver is the part between func keyword and the method name
// Abs with receiver of type MethodStruct with v as a value
// if I had a new struct this method wouldn't be accessible from it
// NOTE: methods are just like functions with a receiver arg
func (v MethodStruct) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// defining the function for a dif struct
func (v TestMethod) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// NOTE: Method can be declared on non-struct types as well
func (v NonStruct) Abs() float64 {
	if v == 0 {
		return float64(-v)
	}
	return float64(v)
}

// NOTE:
/*
You can only declare a method with a receiver whose type is defined in the same package as the method. You cannot declare \
a method with a receiver whose type is defined in another package (which includes the built-in types such as int).
*/

// NOTE: pointer receiver
func (v *MethodStruct) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// scale as regular function
func Scale2(v *MethodStruct, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// NOTE:
/*
Functions with a pointer argument must take a pointer:
var v Vertex
ScaleFunc(v, 5)  // Compile error!
ScaleFunc(&v, 5) // OK

while methods with pointer receivers take either a value or a pointer as the receiver when they are called:
var v Vertex
v.Scale(5)  // OK
p := &v
p.Scale(10) // OK
For the statement v.Scale(5), even though v is a value and not a pointer, the method with the pointer receiver is called automatically.
That is, as a convenience, Go interprets the statement v.Scale(5) as (&v).Scale(5) since the Scale method has a pointer receiver.


The equivalent thing happens in the reverse direction.
Functions that take a value argument must take a value of that specific type:
var v Vertex
fmt.Println(AbsFunc(v))  // OK
fmt.Println(AbsFunc(&v)) // Compile error!
while methods with value receivers take either a value or a pointer as the receiver when they are called:

var v Vertex
fmt.Println(v.Abs()) // OK
p := &v
fmt.Println(p.Abs()) // OK
In this case, the method call p.Abs() is interpreted as (*p).Abs().
*/

// for printing slices
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// ///////////////////////////////////////////////////////////////////////////////////////////////////
func deferTest() {
	// fmt.Println("Dobar")
	// defer will wait for surrounding functions to execute
	// defer fmt.Println("dan")
	fmt.Println("brojanje")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("zavrseno")
	// defer calls are pushed onto a stack, after a function returns, LIFO (last in first ou)
}

// ///////////////////////////////////////////////////////////////////////////////////////////////////
func add(x, y int) int {
	return x + y
}

func testIf(x, n, lim float64) float64 {
	// having v declared and assigned with value in if statement, limits it to the scope of that if statement > check the next comment
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// if I was to return v, outside of the if statement's scope, it would be undefined, and would throw an error at compile time
	return lim
}

func split(sum int) (x, y int) {
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

// NOTE: Exercise: Slices
/*
Implement Pic. It should return a slice of length dy, each element of which is a slice of dx 8-bit unsigned integers. When you run the program, it will display your picture, interpreting the integers as grayscale (well, bluescale) values.

The choice of image is up to you. Interesting functions include (x+y)/2, x*y, and x^y.

(You need to use a loop to allocate each []uint8 inside the [][]uint8.)

(Use uint8(intValue) to convert between types.)

package main

import (
	"golang.org/x/tour/pic"
	//"fmt"
)

func Pic(dx, dy int) [][]uint8 {
	// dy being the length
	pic := make([][]uint8, dy)

	// iterate over pic
	for y := 0; y < dy; y++{
		// at every iteration create a row (slice)
		row := make([]uint8, dx)
		for x := 0; x < dx; x++ {
			// than iterate over each row assigning value for each element in that row(slice)
			row[x] = uint8((x+y)/2)
		}
	pic[y] = row
	//fmt.Println(pic[y])
	}
	// returning the slice of length dy
	return pic
}

func main() {
	pic.Show(Pic)
}
*/

/*
Map exercise
Implement WordCount. It should return a map of the counts of each “word” in the string s.
The wc.Test function runs a test suite against the provided function and prints success or failure.

You might find strings.Fields helpful.


package main

import (
	"golang.org/x/tour/wc"
	"strings"
	"fmt"
)

func WordCount(s string) map[string]int {
	// declare a map
	WordCounts := make(map[string]int)
	// split the string into words by putting it into a slice
	words := strings.Fields(s)
	fmt.Println(words)
	// iterate over the slice and add the key to the map and increment it's value
	// map values will be 0
	// if key doesn't exist in the map, WordCounts[word]++ will add the key and increment it's zero value which is 0
	for _, word := range words{
		fmt.Println(WordCounts[word])
		WordCounts[word]++
	}
	// return the map
	return WordCounts
}

func main() {
	wc.Test(WordCount)
}
*/
// NOTE: Fibonacci function

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	first := 0
	second := 1
	return func() int {
		next := first
		first, second = second, first+second
		return next
	}
}
