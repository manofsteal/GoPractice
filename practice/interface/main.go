package main


import "fmt"


// interface
type Vehicle interface {
	Name() string
}


type Intf2 interface {
	Price() string
}


// implement 1

type Motor struct {
	name string
	price string
}

func (v *Motor) Name() string {
	return v.name
}

func (v *Motor) Price() string {
	return v.price
}


// // user

func showName(v Vehicle) string {
	return v.Name()
}

func showPrice(v Intf2) string {
	return v.Price()
}

func main() {

	// const (
	// 	N = 10
	// )

	// motors := make([]Motor, N)

	// for i := 0; i < N; i++ {
	// 	motors[i] = Motor {
	// 		name: fmt.Sprintf("motor %d", i),
	// 	}
	// }
	// fmt.Println(motors)

	motor := Motor {
		name: "motor",
		price: "10000",
	}

	var v Vehicle
	var p Intf2	

	v = &motor
	p = &motor


	fmt.Println(showName(v))
	fmt.Println(showPrice(p))

	// fmt.Println(motor.Name())


}