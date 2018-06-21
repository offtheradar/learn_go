package main

import(
	"fmt"
)

func	main(){
	name, power := "Yacine", getPower()
	power = 9000
	fmt.Printf("It's over %s, %d\n", name, power)
}

func	getPower() int {
	return 9001
}
