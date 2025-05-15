package main

import (
	"fmt"
	"maps"
	"math"
	"time"
	// 	"slices"
)

// Functions
func plus(a int, b int) int {
  return a+b
}

func plusPlus(a,b,c int) int {
  return a+b+c
}

// Multiple return type
func vals() (int, int){
  return 3,7
}

const s string = "constant"

func main(){
	fmt.Println("Hello world!")
	
// 	DATA TYPES
// Numeric	int, float64, uint32, complex64
// Boolean	bool
// String	string
// Composite	array, slice, map, struct
// Reference	pointer, function, interface, channel

	// VARIABLES
	var a = "initial"
	fmt.Println(a)

	var b, c int = 3, 4
	fmt.Println(b,c)

	f:="apple"
	fmt.Println(f)

	fmt.Println(s)
	const n = 500000
	fmt.Println(math.Sin(n))

	const d = 3e20/n
	fmt.Println(d)
	fmt.Println(int64(d))

	// CONTOL FLOW
	// For Loop
	i:=1
	for i <= 3{
		fmt.Print(i)
		i+=1
	}
	fmt.Println()
	for j:=0; j<=3;j++{
	    fmt.Print(j)
	}
	fmt.Println()

// 	Conditionals
if num:=9; num<0{
    fmt.Println("The number${} is negative")
} else if num<10{
    fmt.Printf("The number %d has only one digit",num)
}else{
    fmt.Println("The number has multiple digits.")
}
// Switches
k:=2
fmt.Println("Write", k, "as")
switch k{
    case 1:
    fmt.Println("One")
    case 2:
    fmt.Println("Two")
}

switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("It is weekend")
    default:
        fmt.Println("It is weekday")
}

whoAmI:= func (i interface{}){
    switch t:=i.(type){
        case bool:
        fmt.Println("I'm bool")
        case int:
        fmt.Println("I'm int")
        default:
        fmt.Printf("Don't know type %T \n",t)
    }
}

whoAmI(true)
whoAmI(2)
whoAmI("true")

// Arrays
arr := [5]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", arr)
    
arr=[...]int{1,2,3,4,5}
fmt.Println("dcl",arr)

arr=[...]int{100,3:400,500}
fmt.Println("dcl",arr)

// var twoD=[2][3]int

// Slices
var s []string
fmt.Println("ununit",s,s==nil, len(s)==0)

  // MAPS
  m:=make(map[string]int)
  m["k1"]=12
  m["k2"]=13
  fmt.Println("map:", m)
  _,prs:=m["k3"]
  fmt.Println("prs:", prs)
  
  myMap := map[string]int{"foo":1, "bar":2}
  n2 := map[string]int{"foo":1, "bar":2}
  
  if maps.Equal(myMap,n2){
    fmt.Println("myMap == n2")
  }
  
  // Function call
  res:=plus(1,2)
  fmt.Println("1+2=",res)
  
  res=plusPlus(1,2,3)
  fmt.Println("1+2+3=",res)
  
  val1, val2 := vals()
  fmt.Println(val1, val2)
  
  _, val3 := vals()
  fmt.Println(val3)

}
