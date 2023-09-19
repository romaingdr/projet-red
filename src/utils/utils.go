package utils

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
)

var (
	Red   = color.New(color.FgRed)
	Blue  = color.New(color.FgBlue)
	Green = color.New(color.FgGreen)
)

// clearConsole efface la console.
func ClearConsole() {
	const clearScreen = "\033[H\033[2J"
	fmt.Print(clearScreen)
}

// inputint lit une entrée de l'utilisateur et renvoie un entier.
func Inputint() (int, error) {
	fmt.Print(">> ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	chiffre, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}
	return chiffre, nil
}

// input lit une entrée de l'utilisateur et renvoie une chaîne de caractères.
func Input() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

// onlyLetters vérifie si une chaîne de caractères ne contient que des lettres.
func OnlyLetters(input string) bool {
	if len(input) > 10 {
		return false
	}
	for _, char := range input {
		if !unicode.IsLetter(char) {
			return false
		}
	}
	return true
}

// capitalizeString met en majuscule la première lettre d'une chaîne de caractères.
func CapitalizeString(input string) string {
	if len(input) > 0 {
		input = strings.ToUpper(string(input[0])) + strings.ToLower(input[1:])
	}
	return input
}

// speedMsg affiche un message lettre par lettre avec une couleur spécifiée.
func SpeedMsg(message string, speed int, colorName string) {
	defaultColor := color.New(color.FgWhite)

	var selectedColor *color.Color

	switch colorName {
	case "green":
		selectedColor = Green
	case "red":
		selectedColor = Red
	case "blue":
		selectedColor = Blue
	default:
		selectedColor = defaultColor
	}

	for _, char := range message {
		selectedColor.Print(string(char))
		time.Sleep(time.Duration(speed) * time.Millisecond)
	}
}

// abilitiesTutorial affiche les abilités disponibles dans le tutoriel de combat et permet à l'utilisateur d'en choisir une.
func AbilitiesTutorial() string {

	fmt.Println("---- Abilités ----")
	fmt.Println("[1] Coup de poing")
	fmt.Println("[2] Frénésie sanguinaire")
	fmt.Println("[3] Lame démoniaque")
	fmt.Println("------------------")
	choice, _ := Inputint()
	switch choice {
	case 1:
		return "Coup de poing"
	case 2:
		return "Frénésie sanguinaire"
	case 3:
		return "Lame démoniaque"
	default:
		ClearConsole()
		Red.Println("Veuillez choisir une option valide")
		return AbilitiesTutorial()
	}
}

// battleMenuTutorial affiche le menu du tutoriel de combat.
func BattleMenuTutorial() {
	fmt.Println("----- A votre tour -----")
	fmt.Print("[1] Attaque auto")
	SpeedMsg("<-- Ceci vous permet d'attaquer l'adversaire avec votre compétence basique", 20, "white")
	Input()
	fmt.Print("[2] Abilités")
	SpeedMsg("<-- Ceci vous permet d'utiliser une abilité sur l'adversaire", 20, "white")
	Input()
	fmt.Print("[3] Inventaire")
	SpeedMsg("<-- Ceci vous permet de consulter votre inventaire pendant le combat", 20, "white")
	fmt.Println("")
	fmt.Println("------------------------")
}

// choixClasse affiche les classes disponibles et permet à l'utilisateur de choisir une classe.
func ChooseClass() string {
	ClearConsole()

	Green.Println("Nom du personnage validé !")
	Blue.Println("Choisissez votre classe : ")
	classes := []string{"Titan", "Arcaniste", "Chasseur"}
	println("")
	println("[1] Titan : «Une représentation de la force brute» | 180 Hp - 10 Ad")
	println("[2] Arcaniste : «Manipule les lois de l'univers» | 100 Hp - 30 Ad")
	println("[3] Chasseur : «N'apparait que dans l'ombre» | 125 Hp - 20 Ad")
	println("")

	for {
		choice, _ := Inputint()
		if choice > 0 && choice < 4 {
			return classes[choice-1]
		} else {
			ClearConsole()
			Red.Println("Veuillez saisir une option valide")
			Blue.Println("Choisissez votre classe : ")
			println("")
			println("[1] Titan : a_completer")
			println("[2] Arcaniste : a_completer")
			println("[3] Chasseur : a_completer")
			println("")
		}
	}
}
