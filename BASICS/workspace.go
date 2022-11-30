package main

import (
	"fmt"
	"math"
)

func main(){
	// var n bool=true
	var n bool  //default is 0 as false
	fmt.Printf("%v, %T\n",n,n)
	a:=1998==1998
	fmt.Printf("%v,%T\n",a,a)
	// b:="Anjan"=="Arjun"
	// fmt.Printf("%v,%T\n",b,b)
	c := 4 + 6i
	var d complex128=6+4i
	fmt.Printf("%v,%T\n",c,c)
	fmt.Printf("%v,%T\n",real(c),real(c))
	fmt.Printf("%v,%T\n",imag(c),imag(c))
	fmt.Printf("%v,%T\n",c+d,c+d)
	fmt.Printf("%v,%T\n",c-d,c-d)
	fmt.Printf("%v,%T\n",c*d,c*d)
	fmt.Printf("%v,%T\n",c/d,c/d)
	x:=10
	y:=12
	fmt.Println(x & y) 
	s := "This is a String"
	b := []byte(s)
	fmt.Printf("%v,%T\n", s, s)
	fmt.Printf("%v,%T\n", b, b)
	var n1 int=6    
	var n2 int8=9
	// fmt.Printf("%v,%T",n1+n2,n1+n2)    mismatched datatype we can't add 
	fmt.Printf("%v,%T\n",n1+int(n2),n1+int(n2))   //we can do by conversion to a int type
	// const v int = math.Sqrt(16)  const not applicable for math libraries 
	v:=math.Sqrt(16)
	fmt.Printf("%v,%T",v,v)
}