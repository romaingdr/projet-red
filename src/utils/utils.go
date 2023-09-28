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

// Variables globales des couleurs de police
var (
	Red    = color.New(color.FgRed)
	Blue   = color.New(color.FgBlue)
	Green  = color.New(color.FgGreen)
	Yellow = color.New(color.FgYellow)
	Cyan   = color.New(color.FgCyan)
)

// splitWords sert à séparer les arguments d'une chaine de cactères avec la séparation " " et de les mettres dans une liste
func splitWords(input string) []string {
	words := strings.Split(input, " ")
	return words
}

// ClearConsole efface la console.
func ClearConsole() {
	const clearScreen = "\033[H\033[2J"
	fmt.Print(clearScreen)
}

// convertInfos sert à mettre toutes les informations du personnage séparés par " " dans une chaine de caractère
func convertInfos(p *Personnage) string {
	var ligneSauvegarde string
	ligneSauvegarde += p.nom + " "
	ligneSauvegarde += p.classe + " "
	ligneSauvegarde += strconv.Itoa(p.currentHp) + " "
	ligneSauvegarde += strconv.Itoa(p.maxHP) + " "
	ligneSauvegarde += strconv.Itoa(p.niveau) + " "
	ligneSauvegarde += strconv.Itoa(p.ennemi) + " "
	return ligneSauvegarde
}

// convertInfosItems sert à mettre toutes les informations des items séparés par " " dans une chaine de caractère
func convertInfosItems(p *Personnage) (string, string) {
	var itemSauvegarde string
	var NbSauvegarde string
	for i := 0; i < len(p.inventory); i++ {
		if i > 0 {
			itemSauvegarde += " " // Ajouter un espace entre les éléments
			NbSauvegarde += " "   // Ajouter un espace entre les éléments
		}
		itemSauvegarde += p.inventory[i].Name
		NbSauvegarde += strconv.Itoa(p.inventory[i].Quantite)
	}
	return itemSauvegarde, NbSauvegarde
}

// Inputint lit une entrée de l'utilisateur et renvoie un entier.
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

// Input lit une entrée de l'utilisateur et renvoie une chaîne de caractères.
func Input() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

// save est utilisé pour sauvegarder la progression en cours dans un fichier 'database.json'
func save(p *Personnage) {
	db = NewQuickDB("database.json")
	sauvegardePerso := convertInfos(p)
	db.Set("sauvegarde", sauvegardePerso)
	sauvegardeItems, sauvegardeNb := convertInfosItems(p)
	db.Set("sauvegardeItems", sauvegardeItems)
	db.Set("sauvegardeNb", sauvegardeNb)
	ClearConsole()
	Green.Println("La progression a été sauvegardée avec succés !")
	p.Menu()
}

// OnlyLetters vérifie si une chaîne de caractères ne contient que des lettres.
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

// CapitalizeString met en majuscule la première lettre d'une chaîne de caractères.
func CapitalizeString(input string) string {
	if len(input) > 0 {
		input = strings.ToUpper(string(input[0])) + strings.ToLower(input[1:])
	}
	return input
}

// SpeedMsg permet d'afficher progressivement un message dans le terminal en spécifiant la couleur et la vitesse
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
	case "yellow":
		selectedColor = Yellow
	default:
		selectedColor = defaultColor
	}

	for _, char := range message {
		selectedColor.Print(string(char))
		time.Sleep(time.Duration(speed) * time.Millisecond)
	}
}
