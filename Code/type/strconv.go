package main

import (
	"fmt"
	"strconv"
)

func main() {

	var (
		intval = 65
		float64val = 2.2
		stringval = "3"

	)

	fmt.Println(intval,float64val,stringval)
	fmt.Printf("%T,%#v\n",float64(intval),float64(intval))

	v,err := strconv.Atoi(stringval)
	fmt.Println(err,v)

	vv,err := strconv.ParseFloat(stringval,64)
	fmt.Println(err,vv)

	fmt.Println(strconv.Itoa(intval))
	fmt.Println(strconv.FormatFloat(float64val,'f',10,64))

}
