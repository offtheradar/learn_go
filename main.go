package main

import(
	"fmt"
)

func	main(){
	name, power := "Yacine", getPower()
	power = add(power, 9000)
	_, bob := do_some(name)
	if !bob {
		fmt.Printf("It's over %s, %d\n", name, power)
	}
}

func	getPower() int {
	return 9001
}

func	add(a int, b int) int {
	return (a + b)
}

func	do_some(name string) (int, bool) {
	if name == "pedro" {
		return 5, true
	} else {
		return 10, false
	}
}
