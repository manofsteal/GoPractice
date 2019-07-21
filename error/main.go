package main

import (
	"fmt"
)

/*

 built-in interface
 type error interface {
	Error() string
 }
*/

type DevideZero struct {
	What string
}

func (e *DevideZero) Error() string {
	return e.What
}

func Devide(d1 float64, d2 int) (float64, error) {
	if d2 == 0 {
		return 0, &DevideZero{
			What: "Cannot devide zero",
		}
	}

	return d1 / float64(d2), nil
}
func main() {

	r, e := Devide(1, 1)

	if e != nil {
		fmt.Println(e.Error())
	} else {
		fmt.Println(r)
	}

	r, e = Devide(1, 0)

	if e != nil {
		fmt.Println(e.Error())
	} else {
		fmt.Println(r)
	}

}
