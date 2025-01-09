package game

import (
	"fmt"
	"rockpaperscissors/player"
)

type Game struct {
	Player1 player.Player
	Player2 player.Player
}

// NewHumanGame erzeugt ein neues Spiel mit zwei menschlichen Spielern mit den Namen n1 und n2.
func NewHumanGame(n1, n2 string) *Game {
	return &Game{
		Player1: player.NewHuman(n1),
		Player2: player.NewHuman(n2),
	}
}

// NewBotGame erzeugt ein neues Spiel mit einem menschlichen und einem Bot-Spieler
// mit den Namen n1 und n2.
func NewBotGame(n1, n2 string) *Game {
	return &Game{
		Player1: player.NewHuman(n1),
		Player2: player.NewBot(n2),
	}
}

// Play spielt ein Spiel mit den beiden Spielern.
// D.h. die Funktion fragt jeden Spieler nach einem Zug und ermittelt den Gewinner.
// Erhöht den Punktestand des Gewinners und zeigt die Punkte an.
func (g *Game) Play() {
	v1 := g.Player1.GetMove()
	v2 := g.Player2.GetMove()

	fmt.Printf("%v hat %v gewählt.\n", g.Player1.GetName(), v1)
	fmt.Printf("%v hat %v gewählt.\n", g.Player2.GetName(), v2)

	// Hinweis: Prüfen Sie mit der Funktion value.Beats() welcher Spieler gewonnen hat.
	//         Aufrufbeispiel: v1.Beats(v2)
	//         Geben Sie dann je nach Ergebnis den Gewinner aus, erhöhen Sie den Punktestand,
	//         geben Sie die Punktestände aus und rufen Sie ggf. die Funktion g.Play() erneut auf.

	// TODO
}

// PrintScores gibt den Punktestand beider Spieler auf die Konsole aus.
func (g Game) PrintScores() {
	// Hinweis: Verwenden Sie fmt.Println() und fmt.Printf() um die Punktestände auszugeben.

	// TODO
}
