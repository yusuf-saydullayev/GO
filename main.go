package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

var (
	userName = "Yusuf"
	userAge  = 23
	userJob  = "Developer"
)

const (
	size   float64 = 123.123
	width  float32 = 1.23123
	lenght float32 = 123123.123123123
	space          = 123123.123123123
)

var (
	i8  int8  = -128 | 127
	i16 int16 = -32768 | 32768
	i32 int32 = -2147483648 | 2147483648
	i64 int64 = -9223372036854775808 | -9223372036854775808
)

var (
	u8  uint8  = 0 | 255
	u16 uint16 = 0 | 65535
	u32 uint32 = 0 | 4294967295
	u64 uint64 = 0 | 18446744073709551615
)

const (
	Monday = 1 + iota
	TuesDay
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

// multiple results
func swap(x, y string) (string, string, string) {
	c := x + y
	return x, y, c
}

func swap2(x, y string) string {
	return x + y
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return x, y
}

var c, python, java bool

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

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z = z - (z*z-x)/(2*z)
		fmt.Printf("Iteration %d: z = %v\n", i+1, z)
	}

	return z
}

func oSystem() string {
	os := runtime.GOOS
	return fmt.Sprintf("Operation system is %v", os)
}

func helloTime() {
	t := time.Now().Hour()

	switch {
	case t < 12:
		fmt.Println("Good morning!")
	case t < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

type Developer struct {
	devType string
}

type Product struct {
	id    uint
	name  string
	price int
}

type User struct {
	id          uint
	name, email string
	phones      []uint
	cart        []Product
	Developer
}

func (uAuth *User) getUserToken() (token string, ok bool) {

	userEmail, userName := uAuth.email, uAuth.name

	if userEmail == "" || userName == "" {
		return "", false
	}

	token = fmt.Sprintf("Bearer: %v:%v", uAuth.email, uAuth.name)
	ok = true

	return token, ok
}

func bubbleSort(arr []int) {
	l := len(arr)
	fmt.Println("Array length", l)

	for i := 0; i < l-1; i++ {
		for j := 0; j < l-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func binarySearch(arr []int, target int) int {
	f := 0
	l := len(arr) - 1

	for i := 0; i < len(arr); i++ {
		mid := (f + l) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			f = arr[mid] + 1
		} else if arr[mid] > target {
			l = arr[mid] - 1
		}
	}

	return -1
}

func square(x int) int {
	return x * x
}

func squarePoint(x *int) {
	*x = *x * *x
}

var returnUsername func(string) string
var returnAge func(birthYear uint) (age int, currentYear int)

func calculate(x, y int, action func(int, int) int) int {
	return action(x, y)
}

func add(x, y int) int {
	return x + y
}

func main() {

	// =============================================================

	var Sum = calculate(1, 2, add)
	fmt.Println("Sum = ", Sum)

	// returnUsername = func(name string) string { return name }

	// fmt.Println(returnUsername("Yusuf"))

	// returnAge = func(birthYear uint) (age int, currentYear int) {
	// 	var cyear = time.Now().Year()
	// 	age = cyear - int(birthYear)
	// 	return age, cyear
	// }

	// age, currentYear := returnAge(2000)
	// fmt.Printf("Age is %v and current year is %v\n", age, currentYear)

	// Pointers ==========================================

	// var intPointer *int
	// var uIntPointer *uint
	// default value of pointer is nil
	// fmt.Printf(" %T, %v\n", intPointer, intPointer)   // *int, <nil>
	// fmt.Printf(" %T, %v\n", uIntPointer, uIntPointer) // *uint, <nil>
	// fmt.Printf("check pointer has value or not %v\n", intPointer == nil) // true
	// fmt.Printf("%T, %v\n", intPointer, *intPointer)   // don't do this, because intPointer is nil and it will be panic
	// intPointer = &userAge    // set address of userAge to intPointer
	// fmt.Println(intPointer)  // get address of userAge 0xc0000b6010
	// fmt.Println(*intPointer) // get value of userAge 23 use * before pointer

	// var b int = 255
	// bPointer := &b                           // bPointer is a pointer to b
	// fmt.Println("Value of b is", *bPointer)  // value 255
	// fmt.Println("Address of b is", bPointer) // address 0xc0000b6010

	// var newPointer = new(int) // new() function creates a pointer to the type and initializes it to default value
	// fmt.Printf("newPointer type: %T, newPointer value: %v\n newPointer address: %v\n", newPointer, *newPointer, newPointer)
	// *newPointer = 10 // to set value to pointer use * before pointer
	// fmt.Println("newPointer value after set", *newPointer)

	// var x int = 10
	// square(x) // square doesn't have side effects because it doesn't change the value of x (pass by value)
	// fmt.Println("x after square", x)
	// squarePoint(&x) // squarePoint has side effects because it changes the value of x (pass by reference)
	// fmt.Println("x after squarePoint", x)

	//=========================================================
	// arr := []int{64, 34, 25, 12, 22, 11, 90}
	// fmt.Println("Unsorted array:", arr)
	// bubbleSort(arr)
	// fmt.Println("Sorted array:", arr)
	// Struct ===============================================
	// Yusufboy := User{
	// 	id:    1,
	// 	name:  "Yusufboy",
	// 	email: "asas@gmail.com",
	// 	phones: []uint{
	// 		99889359184,
	// 		141414124124123,
	// 	},
	// 	cart: []Product{
	// 		{
	// 			id:    1,
	// 			name:  "Samsung",
	// 			price: 222123,
	// 		},
	// 		{
	// 			id:    2,
	// 			name:  "Xiaomi",
	// 			price: 123123,
	// 		},
	// 	},
	// 	Developer: Developer{
	// 		devType: "Frontend",
	// 	},
	// }

	// userToken, ok := Yusufboy.getUserToken()
	// if ok {
	// 	fmt.Println(userToken)
	// }

	// fmt.Println(Yusufboy)
	// m := Yusufboy.email
	// if m == "" {
	// 	fmt.Println("Don't have email")
	// }
	// fmt.Println(Yusufboy.cart)

	// Anonim struct
	// car := struct {
	// 	Name, Model, Time, Certificate string
	// }{
	// 	Name:        "BMW",
	// 	Model:       "X5",
	// 	Time:        fmt.Sprint(time.Now().ISOWeek()),
	// 	Certificate: time.Now().String(),
	// }

	// fmt.Println(car)
	// fmt.Printf("Time: %v \n", car.Time)
	// fmt.Printf("Certificate %v \n", car.Certificate)
	// car.Name = "Ferrari"
	// fmt.Printf("New name: %v", car.Name)

	// ===========================================================
	// m := make(map[string]int)
	// m = map[string]int{
	// 	"Tashkent": 01,
	// 	"Namangan": 02,
	// 	"Andijon":  03,
	// }
	// fmt.Println(m)
	// m["Farg'ona"] = 04
	// fmt.Println("Andijon", m["Andijon"])
	// fmt.Println("len", len(m))
	// delete(m, "Andijon")
	// fmt.Println(m["Andijon"]) // result is 0
	// _, ok := m["Andij"]
	// if ok {
	// 	fmt.Println("Andij is in m")
	// }
	// fmt.Println("Andij is not in m")
	// fmt.Println(m)

	// make ================================================
	// m := make([]int, 3,50)
	// fmt.Println(m)
	// fmt.Printf("Length %v\n",len(m))
	// fmt.Printf("Capacity %v\n",cap(m))

	// ========================================================================
	// const (
	// 	isAdmin = 1 << iota
	// 	isHeadquartes
	// 	isCTO

	// 	canSeeEurope
	// 	canSeeAsia
	// 	canSeeAfrica
	// )

	// userRoles := canSeeAfrica | canSeeAsia | canSeeEurope
	// user1 := canSeeAfrica | canSeeAsia
	// user2 := canSeeEurope

	// fmt.Println("user1 can see Europe", userRoles&user1 == canSeeEurope)
	// fmt.Println("user2 see Europe", userRoles&user2 == canSeeEurope)

	// user 1
	// 32 | 16 (canSeeAfrica | canSeeAsia) & 8(canSeeEurope)
	// 100000 = 32  // 10000 = 16
	// 1000 = 8		// 1000 = 8
	// -------		// ------
	// 100000 = 32	// 10000 = 16

	// 32 | 16 == 8(canSeeEurope) can not see canSeeEurope

	// Slices ===================================================================
	// as := []string{"yellow", "anna", "zick", "gordon", "light"}
	// slices.Sort(as) // mutating current as
	// fmt.Println(sort.SearchStrings(as, "v"))
	// fmt.Println("length:", len(as))
	// fmt.Println("array:", as)

	// Arrays ===================================================
	// q := [...]string{"yellow", "anna", "zick", "gordon", "light"}
	// sortedNumbers := q[:] // NOT COPY, JUST REFERENCE
	// slices.Sort(sortedNumbers) // SORTING (MUTATING ORIGIANL ARRAY(q) TOO !!!!!!)
	// fmt.Println("sorted:", sortedNumbers)
	// fmt.Println("original array:", q) // SAME RESULT WITH sortedNumbers

	// ar := [...]int{4, 5, 6, 3, 7, 8, 0, 2, 7, 23, 457547, 56756}
	// g := ar  // COPY FROM ORIGINAL ARRAY, IF ADD & (g := &ar) OR [:] (g := ar[:])  NOT COPY
	// slices.Sort(g[:]) // MUTATING ONLY g, because we add [:]. [:] is a slice
	// fmt.Println("Original array", ar)
	// fmt.Println("Mutated array", g)

	// var twoD [2][3]int // empty array with [[0 0 0] [0 0 0]]
	// var twoD2 = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	// fmt.Println(twoD)
	// fmt.Println(twoD2)

	// strings operations ==========================================
	// a := "Hello! This is string in go"
	// b := "Hello! This is string in go"

	// fmt.Println(strings.Compare(a, b)) // use this: == instead this strings.Compare
	// fmt.Println(strings.Contains(a, "in go"))
	// fmt.Println(strings.Count(a, "s")) // retruned how many are there
	// fmt.Println(strings.Index(a, "H")) // returns 0, because Go uses a zero-based indexing system
	// fmt.Println(strings.Split(a," "))
	// fmt.Println(strings.ReplaceAll(a,"go","Go"))
	// fmt.Println(strings.ToUpper(b))

	// array,defer, len, cap ======================================================
	// names := [3]string{"Yusuf", "Arnold", "John"}
	// famylies := []string{"Yusuf", "Arnold", "John"}
	// names[0] = "Alex"
	// famylies = append(famylies, "Clark")

	// r:= famylies[:2] // get Yusuf and Aronld
	// fmt.Println(r)
	// fmt.Println(famylies)
	// fmt.Println(cap(nNames))
	// fmt.Println(len(nNames))
	// fmt.Println(nNames)
	// fmt.Println(names)
	// defer fmt.Println(famylies[1])
	// defer helloTime()

	// for loop============================================================
	// for x := 1.0; x <= 5.0; x++ {
	// 	result := Sqrt(x)
	// 	fmt.Printf("Square root of %v is %v\n", x, result)
	// 	fmt.Printf("Difference from math.Sqrt: %v\n", math.Sqrt(x)-result)
	// 	fmt.Println()
	// }

	// Switch Case ================================================================
	// switch os := runtime.GOOS; os {
	// case "darwin":
	// 	fmt.Printf("Operation system is %v", os)
	// case "freebsd":
	// 	fmt.Printf("Operation system is %v", os)
	// case "linux":
	// 	fmt.Printf("Operation system is %v", os)
	// default:
	// 	fmt.Printf("Operation system is %v\n", os)
	// }

	// whatAmI := func(i interface{}) {
	// 	switch t := i.(type) {
	// 	case bool:
	// 		fmt.Println("I'm a bool")
	// 	case int:
	// 		fmt.Println("I'm an int")
	// 	default:
	// 		fmt.Printf("Don't know type %T\n", t)
	// 	}
	// }
	// whatAmI(true)
	// whatAmI(1)
	// whatAmI("hey")

	//----------------------------------------------------
	// fmt.Println("When's Saturday?")
	// fmt.Println(time.Monday)

	// m:= time.Monday
	// fmt.Println(m)

	// today := time.Now().Weekday()

	// switch time.Sunday {
	// case today + 0:
	// 	fmt.Println("Today.")
	// case today + 1:
	// 	fmt.Println("Tomorrow.")
	// case today + 2:
	// 	fmt.Println("In two days.")
	// default:
	// 	fmt.Println("Too far away.")
	// }

	//====================================================================
	// var String string = "admin"
	// fmt.Println(String)
	// String = "asdasdasdasd"
	// fmt.Println(String)
	// var Number int = 12312312312312
	// var Boolean bool = true

	// fmt.Println(String,Number,Boolean)
	// var n bool

	// a, b := swap("aaa", "bbbb")
	// var a string = aaa
	// var b string = bbbb
	// fmt.Println(a, b)
	// fmt.Println(10 / 3)

	// Strig to bytes =============================
	// s:="this is string"
	// b:= []byte(s)
	// s2 := string(b)
	// fmt.Println(b)
	// fmt.Println(s2)

	// Runes =============================
	// r:= 'A'
	// fmt.Printf("%T, %v", r,r)

	// =============================
	// var a, b, c = swap("a", "b")
	// var g = swap2("a", "b")
	// fmt.Println(a, b, c)
	// fmt.Println(g)

	// funcitons ==========================
	// var x, y int = 3, 5
	// var d = x*x + y*y -> 34
	// var h = float64(d) -> float64 type 34
	// var g = math.Sqrt(h)-> math.Sqrt(x float64) float64 = 5.830951894845301
	// var e = uint(g) -> 5
	// var f = uint(math.Sqrt(float64(x*x + y*y)))
	// fmt.Println(x, y, f)

	// loop ====================================
	// sum := 0
	// for i := 0; i < 10; i++ {
	// 	sum += 10
	// }
	// fmt.Println(sum)

	// sum := 1
	// for sum < 1000 {
	// 	sum += sum
	// }
	// fmt.Println(sum)

	//========================================
	// sum := 234
	// fmt.Println(sqrt(float64(sum)), sqrt(-4))

	//===========================================
	// fmt.Println(
	// 	pow(3, 2, 10),
	// 	pow(3, 3, 20),
	// )

	//===========================================
	// a := 10 // 1010
	// b := 3  // 0011
	// fmt.Println(10 & 3)
	// fmt.Println(10 | 3)
	// fmt.Println(10 ^ 3)
	// fmt.Println(10 &^ 3)

}

func getUserToken() {
	panic("unimplemented")
}

//========================================
// 10 & 3 --> "and" operator
// if 1 & 0 = 0
// if 0 & 0 = 0
// if 1 & 1 = 1
// if 0 & 1 = 0
//-----------------
// 1010
// 0011
// 0010 = 2 -> binary 0010 to decimal 2

//============================================
// 10 | 3 --> "or" operator
// if  1 | 0 = 1
// if  0 | 0 = 0
// if  1 | 1 = 1
// if  0 | 1 = 1
//-----------------
// 1010
// 0011
// 1011 = 11 -> binary 1011 to decimal 11

//=====================================
// 10 ^ 3 --> "exclusive or" operator. If two 1 result 0
// if 1 ^ 0 =  1
// if 0 ^ 0 =  0
// if 1 ^ 1 =  0
// if 0 ^ 1 =  1
//----------------------
// 1010
// 0011
// 1001 = 9 -> binary 1001 to decimal 9

//============================================
// 10 &^ 3 --> "and non" operator. If two 0 result 1
// if  1 &^ 0 = 0
// if  0 &^ 0 = 1
// if  1 &^ 1 = 0
// if  0 &^ 1 = 0
//-------------------------------
// 1010
// 0011
// 0100 = 4 -> binary 0100 to decimal 4
