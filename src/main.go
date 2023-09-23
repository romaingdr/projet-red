package main

// Import nécessaires au fonctionnement du jeu
import (
	"fmt"
	"src/utils"
)

// Fonction main
func main() {
	utils.ClearConsole()
	var p1 utils.Personnage
	utils.StartScreen(&p1, 1)
}

func draw() {
	fmt.Println("    /~\\")
	fmt.Println("   |oo )")
	fmt.Println("   _\\=/_")
	fmt.Println("  /  _  \\")
	fmt.Println(" //|/.\\|\\\\")
	fmt.Println("||  \\_/  || ")
	fmt.Println("|| |\\ /| ||")
	fmt.Println(" # \\_ _/ #")
	fmt.Println("   | | |")
	fmt.Println("   | | |")
	fmt.Println("   []|[]")
	fmt.Println("   | | |")
	fmt.Println("  /_]_[_\\")
}
