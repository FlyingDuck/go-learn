package main

import "fmt"

func main()  {
	// The Println versions also insert a blank between arguments and append a newline to the output while the Print versions add blanks only if the operand on neither side is a string.
	fmt.Println("Hello", "Go", "I am", "Bennett", 100)
	
	// %d %x
	var x uint64 = 1<<64 - 1
	fmt.Printf("%d %x; %d %x\n", x, x, int64(x), int64(x))
	
	// %v %+v %#v
	t := &T{ 7, -2.35, "abc\tdef" }
	fmt.Printf("%v\n", t)
	fmt.Printf("%+v\n", t)
	fmt.Printf("%#v\n", t)
	//fmt.Printf("%#v\n", timeZone)
	fmt.Println(t)
}

type T struct {
	a int
	b float64
	c string
}

func (t *T) String() string  {
	return fmt.Sprintf("%d, %g, %q", t.a, t.b, t.c)
}


