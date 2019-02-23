package main

import (
	"fmt"
	"time"
)

// Person defined by name, age, weight and list of heights per year
type Person struct {
	name   string
	age    int
	weight float32
}

func main() {
	person := Person{"Sergi", 17, 48.5}
	fmt.Println(person)

	for i := 0; i < 8; i++ {
		fmt.Printf("Is the number %d even? %t\n", i, i%2 == 0)
	}

	fridge := make(map[string]int)
	fridge["apples"] = 5
	fridge["beers"] = 0

	for item, quantity := range fridge {
		if quantity == 0 {
			fmt.Printf("WE NEED TO BUY MORE %s!!!\n", item)
		}
	}

	i := 3
	fmt.Println(&i)

	inc1(i)
	fmt.Println(i)

	inc2(&i) // uses pointer
	fmt.Println(i)

	//go count("sheeps")
	//go count("cows")
	sum, sub, mult, div := calculate(5, 8)
	fmt.Println(sum, sub, mult, div)

	/*cmd := GetCommand()
	out, err := exec.Command(cmd[0], cmd[1:]...).Output()

	if err != nil {
		panic(err)
	}

	fmt.Println(string(out))
	*/

	//numbers := []int{5, 7, 1, 7, 9, 2, 2, 7, 8, 2, 6, 25, 836, 25, 215}
	//fmt.Println(sumArray(numbers...))

	/*var a [15]int

	for i := 0; i < len(a); i++ {
		a[i] = i
	}

	fmt.Println("numbers | fibonacci")
	for i := range a {
		fmt.Printf(" %d	| %d\n", i, fib(i))
	}*/
	fmt.Println(fib(16))
	RunJSON()

}

func fib(n int) int {
	if n <= 1 {
		return n
	}

	return fib(n-2) + fib(n-1)
}

func sumArray(nums ...int) int {
	var sum int
	for _, n := range nums {
		sum += n
	}
	return sum
}

func inc1(x int) {
	x++
}

func inc2(x *int) {
	*x++
}

func count(what string) {
	for i := 1; i < 10; i++ {
		fmt.Println(i, what)
		time.Sleep(time.Second)
	}
}

func calculate(x, y int) (int, int, int, int) {
	return x + y, x - y, x * y, x / y
}
