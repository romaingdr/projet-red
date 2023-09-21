// FICHIER UTILISE POUR L'AFFICHAGE DE LA FENETRE DE DEPART

package utils

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"os"
)

// StartScreen affiche la page d'accueil en ASCII ART pour choisir entre jouer une nouvelle partie et quitter le jeu
func StartScreen(p *Personnage, state int) {
	ClearConsole()
	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()

	ClearConsole()

	if state == 1 {
		fmt.Println("இஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇ")
		fmt.Println("இ                                     இ")
		fmt.Println("இ     ➥       NOUVELLE PARTIE         இ")
		fmt.Println("இ                                     இ")
		fmt.Println("இ              QUITTER LE JEU         இ")
		fmt.Println("இ                                     இ")
		fmt.Println("இஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇ")
	} else {
		fmt.Println("இஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇ")
		fmt.Println("இ                                     இ")
		fmt.Println("இ              NOUVELLE PARTIE        இ")
		fmt.Println("இ                                     இ")
		fmt.Println("இ     ➥       QUITTER LE JEU          இ")
		fmt.Println("இ                                     இ")
		fmt.Println("இஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇஇ")
	}

	_, key, err := keyboard.GetKey()
	if err != nil {
		panic(err)
	}

	switch key {
	case keyboard.KeyArrowUp:
		ClearConsole()
		StartScreen(p, 1)
		ClearConsole()
	case keyboard.KeyArrowDown:
		state = 0
		ClearConsole()
		StartScreen(p, 0)
		ClearConsole()
	case keyboard.KeyEnter:
		if state == 1 {
			ClearConsole()
			p.CreateCharacter()
		} else {
			ClearConsole()
			os.Exit(0)
		}
		return // Return the current state when the Enter key is pressed
	default:
		StartScreen(p, state)
	}
}
