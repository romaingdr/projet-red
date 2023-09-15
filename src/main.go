package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

// Fonction main du programme et initialisation du personnage

func main() {
	var p1 Personnage
	p1.Initialize("Romain", "elfe", 1, 40, 100, map[string]int{"Potions": 3, "Argent": 10})
	p1.Menu()

}

func (p *Personnage) Initialize(nom string, classe string, niveau int, hp int, hp_max int, inventaire map[string]int) {
	p.nom = nom
	p.classe = classe
	p.niveau = niveau
	p.current_hp = hp
	p.max_hp = hp_max
	p.inventory = inventaire
}

// Structure du personnage

type Personnage struct {
	nom        string
	classe     string
	niveau     int
	current_hp int
	max_hp     int
	inventory  map[string]int
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
	fmt.Println("[1] Personnage")
	fmt.Println("[2] Inventaire")
	fmt.Println("[3] Marchand")
	fmt.Println("[4] Quitter le jeu")

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
		fmt.Println("Fermeture du jeu...")
	default:
		clearConsole()
		fmt.Println("Veuillez saisir une donnée valide !")
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
	if p.inventory["Potions"] > 0 {
		if p.current_hp == p.max_hp {
			fmt.Println("Vous êtes déjà au maximum de points de vie !")
		} else {
			if p.max_hp-p.current_hp < 50 {
				p.current_hp = p.max_hp
			} else {
				p.current_hp += 50
			}
			p.removeInventory("Potions", 1)
			fmt.Println("- 1 potion, points de vie : ", p.current_hp)
			fmt.Println("Potions restantes : ", p.inventory["Potions"])
		}
	}
}

// Fonctions des sous choix du menu principal

func (p *Personnage) accessInvetory() {
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
		fmt.Println("Sortie de l'inventaire")
	}
} // inventaire

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
		fmt.Println("Une potion a été achetée !")
		p.marchand()
	case 2:
		p.addInvetory("Potions de poison", 1)
		clearConsole()
		fmt.Println("Une potion de poison a été achetée !")
		p.marchand()
	case 3:
		if p.enoughMoney(5) {
			p.addInvetory("Epée", 1)
			p.removeInventory("Argent", 5)
			clearConsole()
			fmt.Println("Une épée a été achetée !")
			p.marchand()
		} else {
			clearConsole()
			fmt.Println("Vous n'avez pas assez d'argent !")
			p.marchand()
		}
	case 4:
		clearConsole()
		fmt.Println("Sortie du marchand...")
		p.Menu()
	default:
		clearConsole()
		fmt.Println("Veuillez saisir une donnée valide")
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
