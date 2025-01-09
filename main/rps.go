package main

import (
	"fmt"
	"rockpaperscissors/game"
)

func main() {
	var n1 string
	var n2 string
	fmt.Print("Bitte den Namen von Spieler 1 eingeben: ")
	fmt.Scanln(&n1)
	fmt.Print("Bitte den Namen von Spieler 2 eingeben: ")
	fmt.Scanln(&n2)

	g := game.NewBotGame(n1, n2)

	p2_human := "N"
	fmt.Printf("Wird %s von einem Menschen gespielt? Falls ja, antworte mit \"Ja\" oder \"J\"", n2)
	fmt.Scanln(&p2_human)
	if p2_human == "J" || p2_human == "Ja" {
		fmt.Println("Zweispielerspiel")
		g = game.NewHumanGame(n1, n2)
	} else {
		fmt.Println("Einzelspielerspiel")
	}

	for i := 0; i < 10; i++ {
		g.Play()
	}
}
