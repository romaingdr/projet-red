// FICHIER UTILISE POUR STOCKER TOUTES LES FONCTIONS UTILES PENDANT LE JEU (input, affichage etc)

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
	Red    = color.New(color.FgRed)
	Blue   = color.New(color.FgBlue)
	Green  = color.New(color.FgGreen)
	Yellow = color.New(color.FgYellow)
	Cyan   = color.New(color.FgCyan)
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
	if len(input) > 10 || len(input) < 3 {
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
	case "cyan":
		selectedColor = Cyan
	default:
		selectedColor = defaultColor
	}

	for _, char := range message {
		selectedColor.Print(string(char))
		time.Sleep(time.Duration(speed) * time.Millisecond)
	}
}
