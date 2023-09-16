package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// Fonction main du programme et initialisation du personnage

func main() {
	clearConsole()
	var p1 Personnage
	p1.createCaracter()
	p1.Menu()

}

func choixClasse() string {
	clearConsole()
	green := color.New(color.FgGreen)
	blue := color.New(color.FgBlue)
	green.Println("Nom du personnage validé !")
	blue.Println("Choisissez votre classe : ")
	classes := []string{"Titan", "Arcaniste", "Chasseur"}
	println("")
	println("[1] Titan : blablabla")
	println("[2] Arcaniste : blablabla")
	println("[3] Chasseur : blablabla")
	println("")
	print("")
	choice := inputint()
	if choice > 0 && choice < 4 {
		return classes[choice-1]
	} else {
		choixClasse()
	}
	return ""
}

func seulementLettres(input string) bool {
	for _, char := range input {
		if !unicode.IsLetter(char) {
			return false
		}
	}
	return true
}

func majString(input string) string {
	// Convertir la première lettre en majuscule
	if len(input) > 0 {
		input = strings.ToUpper(string(input[0])) + strings.ToLower(input[1:])
	}
	return input
}

func welcomeMsg(message string) {
	green := color.New(color.FgGreen)
	for _, char := range message {
		green.Print(string(char))
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println("")
	fmt.Println("Appuyez sur entrée pour continuer")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	clearConsole()
}

func (p *Personnage) createCaracter() {
	red := color.New(color.FgRed)
	fmt.Println("---- Création du personnage ----")
	fmt.Println("[1] Créer un personnage")
	fmt.Println("[2] Personnage par défault")
	fmt.Println("--------------------------------")
	choice := inputint()
	switch choice {
	case 1:
		hp_max := 0
		nom := "0"
		clearConsole()
		print("Nom de votre personnage ")
		nom = input()
		for !(seulementLettres(nom)) {
			clearConsole()
			red.Println("Veuillez n'utiliser que des lettres")
			print("Nom de votre personnage ")
			nom = input()
		}
		nom = majString(nom)
		clearConsole()
		classe := choixClasse()
		switch classe {
		case "Titan":
			hp_max = 150
		case "Chasseur":
			hp_max = 125
		case "Aracaniste":
			hp_max = 100
		}
		clearConsole()
		p.Initialize(nom, classe, 1, hp_max, hp_max, map[string]int{"Potions": 3, "Argent": 10}, []string{"Coup de poing"})
		welcomeMsg("Bienvenue, " + nom + " ! ")

	case 2:
		p.Initialize("Romain", "Chasseur", 1, 67, 125, map[string]int{"Potions": 3, "Argent": 10}, []string{"Coup de poing"})
		clearConsole()
		welcomeMsg("Bienvenue, Romain !")

	default:
		clearConsole()
		red.Println("Veuillez saisir une donnée valide")
		p.createCaracter()
	}

}

func (p *Personnage) Initialize(nom string, classe string, niveau int, hp int, hp_max int, inventaire map[string]int, skill []string) {
	p.nom = nom
	p.classe = classe
	p.niveau = niveau
	p.current_hp = hp
	p.max_hp = hp_max
	p.inventory = inventaire
	p.skill = skill
}

// Structure du personnage

type Personnage struct {
	nom        string
	classe     string
	niveau     int
	current_hp int
	max_hp     int
	inventory  map[string]int
	skill      []string
}

// Interaction avec la console

func clearConsole() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

func inputint() int {

	fmt.Print(">> ")

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	input := scanner.Text()

	chiffre, _ := strconv.Atoi(input)

	return chiffre
}

func input() string {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print(">> ")

	// Utilisez Scan() pour lire une ligne de texte
	scanner.Scan()

	return scanner.Text()
}

// Menu principal du jeu

func (p *Personnage) Menu() {
	red := color.New(color.FgRed)
	fmt.Println("----- Menu -----")
	fmt.Println("[1] Personnage")
	fmt.Println("[2] Inventaire")
	fmt.Println("[3] Marchand")
	fmt.Println("[4] Abilités")
	fmt.Println("[5] Quitter le jeu")
	fmt.Println("----------------")

	choice := inputint()

	switch choice {
	case 1:
		clearConsole()
		p.displayInfo()
		p.Menu()
		break
	case 2:
		clearConsole()
		p.accessInvetory()
		p.Menu()
		break
	case 3:
		clearConsole()
		p.marchand()
	case 4:
		clearConsole()
		p.showSkills()
		p.Menu()
	case 5:
		clearConsole()
		red.Println("Fermeture du jeu...")
	default:
		clearConsole()
		red.Println("Veuillez saisir une donnée valide !")
		p.Menu()
	}
}

// Interaction directe-joueur

func (p *Personnage) dead() {
	health := p.current_hp
	if health <= 0 {
		p.current_hp = p.max_hp / 2
	}
}

func (p *Personnage) poisonPot() {
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		p.current_hp -= 10
		fmt.Println("[poison] -10hp, ", p.current_hp, "/", p.max_hp)
	}
}

func (p *Personnage) takePot() {
	red := color.New(color.FgRed)
	blue := color.New(color.FgBlue)
	if p.inventory["Potions"] > 0 {
		if p.current_hp == p.max_hp {
			red.Println("Vous êtes déjà au maximum de points de vie !")
		} else {
			if p.max_hp-p.current_hp < 50 {
				p.current_hp = p.max_hp
			} else {
				p.current_hp += 50
			}
			p.removeInventory("Potions", 1)
			blue.Println("- 1 potion, points de vie : ", p.current_hp)
			blue.Println("Potions restantes : ", p.inventory["Potions"])
		}
	}
}

// Fonctions des sous choix du menu principal

func (p *Personnage) accessInvetory() {
	red := color.New(color.FgRed)
	fmt.Println("--- Inventaire ---")
	for objet := range p.inventory {
		fmt.Println(objet, " : ", p.inventory[objet])
	}
	fmt.Println("--------------")
	fmt.Println("[1] Utiliser une potion")
	fmt.Println("[2] Quitter l'inventaire")
	choice := inputint()
	switch choice {
	case 1:
		clearConsole()
		p.takePot()
		p.accessInvetory()
	case 2:
		clearConsole()
	default:
		clearConsole()
		red.Println("Veuillez saisir une donnée valide")
		p.accessInvetory()
	}
} // inventaire

func (p *Personnage) showSkills() {
	fmt.Println("--- Abilités ---")
	for i := 0; i < len(p.skill); i++ {
		fmt.Println("Sort n°", i+1, " : ", p.skill[i])
	}
	fmt.Println("----------------")
}

func (p *Personnage) displayInfo() {
	fmt.Println("--- ", p.nom, " ---")
	fmt.Println("Classe : ", p.classe)
	fmt.Println("Niveau : ", p.niveau)
	fmt.Println("Points de vie : ", p.current_hp, "/", p.max_hp)
	fmt.Println("--------------")
} //personnage

func (p *Personnage) enoughMoney(cout int) bool {
	if p.inventory["Argent"]-cout < 0 {
		return false
	} else {
		return true
	}
}

func (p *Personnage) marchand() {
	red := color.New(color.FgRed)
	blue := color.New(color.FgBlue)
	green := color.New(color.FgGreen)
	green.Println("Argent :  ", p.inventory["Argent"])
	fmt.Println("--- Marchand ---")
	fmt.Println("[1] x1 Potion - [gratuit]")
	fmt.Println("[2] x1 Potion de poison - [gratuit]")
	fmt.Println("[3] x1 Epee - [5]")
	fmt.Println("[4] Sortir")
	fmt.Println("----------------")
	choice := inputint()
	switch choice {
	case 1:
		p.addInvetory("Potions", 1)
		clearConsole()
		blue.Println("Une potion a été achetée !")
		p.marchand()
	case 2:
		p.addInvetory("Potions de poison", 1)
		clearConsole()
		blue.Println("Une potion de poison a été achetée !")
		p.marchand()
	case 3:
		if p.enoughMoney(5) {
			p.addInvetory("Epée", 1)
			p.removeInventory("Argent", 5)
			clearConsole()
			blue.Println("Une épée a été achetée !")
			p.marchand()
		} else {
			clearConsole()
			red.Println("Vous n'avez pas assez d'argent !")
			p.marchand()
		}
	case 4:
		clearConsole()
		p.Menu()
	default:
		clearConsole()
		red.Println("Veuillez saisir une donnée valide")
	}
} //marchand

// Interactions inventaire

func (p *Personnage) addInvetory(item string, nb int) {
	if _, existe := p.inventory[item]; !existe {
		p.inventory[item] = nb
	} else {
		p.inventory[item] += nb
	}
}

func (p *Personnage) removeInventory(item string, quantity int) {
	currentQuantity, exists := p.inventory[item]

	if exists {
		newQuantity := currentQuantity - quantity
		if newQuantity <= 0 {
			if item != "Argent" {
				delete(p.inventory, item)
			} else {
				p.inventory["Argent"] = 0
			}
		} else {
			p.inventory[item] = newQuantity
		}
	}
}
