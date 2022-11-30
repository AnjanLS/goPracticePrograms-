package main
import "fmt"

func main(){
	var car = map[string]string{"brand":"ford","type":"mustang","year":"1964"}
	values:=map[string]int{"A":1,"B":2,"C":3}

	fmt.Printf("car\t%v\n",car)
	fmt.Printf("values\t%v\n",values)
}