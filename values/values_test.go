package values

import "fmt"

//
func ExampleValue_String() {
	rock := Rock
	paper := Paper
	scissors := Scissors

	fmt.Println(rock)
	fmt.Println(paper)
	fmt.Println(scissors)

	// Output:
	// Stein
	// Papier
	// Schere
}

func ExampleValue_Beats() {
	fmt.Println(Rock.Beats(Rock))
	fmt.Println(Rock.Beats(Paper))
	fmt.Println(Rock.Beats(Scissors))
	fmt.Println()

	fmt.Println(Paper.Beats(Rock))
	fmt.Println(Paper.Beats(Paper))
	fmt.Println(Paper.Beats(Scissors))
	fmt.Println()

	fmt.Println(Scissors.Beats(Rock))
	fmt.Println(Scissors.Beats(Paper))
	fmt.Println(Scissors.Beats(Scissors))

	// Output:
	// false
	// false
	// true
	//
	// true
	// false
	// false
	//
	// false
	// true
	// false
}
