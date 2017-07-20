package main

import (
	"errors"
	"fmt"
	"math"
	"time"
)

func main() {

	var number int
	fmt.Print("Enter a number:")
	fmt.Scanln(&number)

	switch number {
	case 1:
		variable()
	case 2:
		loops()
	case 3:
		arrays()
	case 4:
		slices()
	case 5:
		maps()
	case 6:
		functions()
	case 7:
		closures()
	case 8:
		pointers()
	case 9:
		structPerson()
		structRectange()
	case 10:
		errorsEx()
	case 11:
		goroutines()
	case 12:
		channels()
	case 13:
		channelSync()
	default:
		fmt.Println("Enter a correct number")
	}
}

func variable() {
	//String type infered from initialized vale
	var str = "initial"
	fmt.Println(str)

	var a, b int
	a, b = 1, 2
	fmt.Println(a, b)

	f := "short"
	fmt.Println(f)

	//Constants
	const c string = "constant"
	const n = 5000000
	const d = 3e20 / n

	//cannot reassign to n
	//n = 200
	fmt.Println(d)
	//A numeric constant has no type until
	//it’s given one, such as by an explicit cast.
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}

func loops() {

	fmt.Println("\n ---for---")
	i := 0
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	fmt.Println("\n---If Else---")
	if m := 2; m < 0 {
		fmt.Println(m, " is negative")
	} else if m < 10 {
		fmt.Println(m, " has 1 digit")
	} else {
		fmt.Println(m, " has multiple digits")
	}

	fmt.Println("\n---Switch---")

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		{
			fmt.Println("It's the weekend")
		}
	default:
		fmt.Println("It's weekday")
	}

}

func arrays() {

	//var names [3]string = {"Ir","Ri","Su"} --Does not work
	names := [3]string{"Ir", "Ri", "Su"}
	fmt.Println(names)

	var twoDims [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoDims[i][j] = i
		}
	}
	fmt.Println("2d", twoDims)
	fmt.Println("2d", twoDims[1][:])
}

func slices() {

	s := make([]int, 3)
	fmt.Println(s)
	s[0] = 10
	s[1] = 20
	s[2] = 30
	s1 := append(s, 35)
	fmt.Println("s1:", s1)
	s = append(s, 40)
	s = append(s, 50)
	fmt.Println("s:", s)

	c := make([]int, len(s))
	copy(c, s)
	fmt.Println("c:", c) //10 20 30 40 50

	l1 := c[2:5]    //Note excluding 6th
	fmt.Println(l1) //30 40 50
	l2 := c[1:]
	fmt.Println(l2)

	for i, v := range s {
		fmt.Println("Value at ", i, ":", v)
	}

	//variable length slice
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)
}

func maps() {
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 10
	m["k3"] = 20

	v1 := m["k3"]
	fmt.Println("v1:", v1)

	delete(m, "k3")
	//https://medium.com/go-walkthrough/go-walkthrough-fmt-55a14bbbfc53
	fmt.Printf("map:%v v1:%d \n", m, v1)

	val, isPresent := m["k3"]
	if isPresent {
		fmt.Println("Key present with value as ", val)
	} else {
		fmt.Println("Key not present")
	}

	n := map[string]int{"foo": 2, "bar": 5}
	fmt.Println(n)

	for k, v := range n {
		fmt.Printf("%s -> %d\n", k, v)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}

func functions() {

	a, b := vals()
	fmt.Println(a, b)

	a, b = variadicFuncSum(10, 15, 20, 25, 30)
	fmt.Println(a, b)
}

func vals() (int, int) {
	return 3, 7
}

func variadicFuncSum(nums ...int) (int, int) {
	total := 0
	counter := 0
	for _, num := range nums {
		total += num
		counter++
	}

	return total, counter
}

func closures() {

	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newSeq := intSeq()
	fmt.Println(newSeq())
	fmt.Println(newSeq())

}

func intSeq() func() int {

	i := 0
	return func() int {
		i++
		return i
	}
}

func pointers() {
	i := 1
	zeroval(i)
	fmt.Println(i)
	zerovalPtr(&i)
	fmt.Println(i)
}

func zeroval(i int) {
	i = 0
}

func zerovalPtr(i *int) {
	*i = 0
}

type person struct {
	name string
	age  int
}

func structPerson() {

	//note the curly braces
	s1 := person{"Bob", 20}
	fmt.Printf("%v type is %#v %T \n", s1, s1, s1)

	s2 := person{name: "Fred", age: 25}
	fmt.Println(s2)

	sp := &s2
	sp.age = 30
	fmt.Printf("Pointer %T has value %v Original s2 has value: %d", sp, *sp, s2.age)

}

type rect struct {
	width, height float32
}

//https://stackoverflow.com/questions/30403642/why-cant-the-interface-be-implemented-with-pointer-receivers
func (r *rect) area() float32 {
	r.width = 1.0
	r.height = 1.0
	return r.width * r.height
}

func (r rect) perim() float32 {
	r.width = 3.0
	r.height = 3.0
	return 2*r.width + 2*r.height
}
func structRectange() {
	r := rect{width: 10.0, height: 20.0}
	fmt.Printf("\n\nArea of rectange %v: %f     \nPerimeter of rectange:%f", r, r.area(), r.perim())
	fmt.Printf("\nMemory address: %p", &r)
	fmt.Printf("\nIs rectangle width: %f height: %f changed? Yes, only in pointer receiver", r.width, r.height)
	//Go automatically handles conversion between values and pointers for method calls. You may want to use
	//a pointer receiver type to avoid copying on method calls or to allow the method to mutate the receiving struct.
	rp := &r
	rp2 := &r
	fmt.Printf("\n\nArea of rectange %v: %f     \nPerimeter of rectange:%f", &rp, rp.area(), rp.perim())
	fmt.Printf("\nMemory address %p points to %p and %p points to %p", &rp, rp, &rp2, rp2)
	//Interface
	c := circle{radius: 5}
	measure(c)
	//measure(r) //r is not a pointer as expected in rect's area function.
	//http://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go
	//That may sound cryptic, but it makes sense when you remember the following: everything in Go is passed by value.
	//Every time you call a function, the data you’re passing into it is copied.
	//In the case of a method with a value receiver, the value is copied when calling the method.
	measure(rp)
}

type circle struct {
	radius float32
}

func (c circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float32 {
	return 2 * math.Pi * c.radius
}

type geometry interface {
	//area(), perim() float32 ERROR
	area() float32
	perim() float32
}

func measure(g geometry) {

	fmt.Printf("\nGeometry type: %T Pointer:%p \n", g, &g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func errorsEx() {

	//build in error interface
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	//custom error type
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	//type assertion
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {

		fmt.Printf("e:%v ok:%v ae type: %T arg:%d prob:%s", e, ok, ae, ae.arg, ae.prob)

	}
}

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}

	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with 42"}
		//return -1, argError.New{arg, "can't work with 42"} argError.New undefined (type argError has no method New)
	}

	return arg + 3, nil
}

func f(str string) {
	for i := 0; i < 3; i++ {
		fmt.Println(str, ":", i)
		//amt: = time.Duration(rand.Intn(250))
		//time.Sleep(time.Millisecond * amt)
	}
}

func goroutines() {
	f("direct")

	go f("goroutines")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	go f("goroutines 2")
	var input string
	fmt.Println("Enter to exit")
	fmt.Scanln(&input)
	fmt.Println("done")
}

func channels() {
	messages := make(chan string)
	fmt.Printf("Channel Type: %#v \n", messages)
	go func() { messages <- "ping" }()
	msg := <-messages
	fmt.Printf("Msg Type: %#v Value: %s", &msg, msg)

}

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func channelSync() {
	done := make(chan bool)
	go worker(done)

	<-done

}
