package main

import "fmt"

func main() {

	// create connection.
	InitDB()

	// do basic test.
	// if err := basic_test(); err != nil {
	// 	fmt.Printf("error : %s\n", err)
	// }

	// do practical test.
	if err := practical_test(); err != nil {
		fmt.Printf("error : %s\n", err)
	}

}
