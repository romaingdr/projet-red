package main

// Import nécessaires au fonctionnement du jeu
import (
	"src/utils"
)

// Fonction main
func main() {
	utils.ClearConsole()
	var p1 utils.Personnage
	p1.CreateCharacter()
	p1.Menu()
}
