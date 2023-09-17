package main

// Import des packages
import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// Fonction main
func main() {
	clearConsole()
	var p1 Personnage
	p1.createCaracter()
	p1.Menu()

}

// Affichage des classes et choix parmis celles-ci
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

// Fonctions pour la personnalisation du personnage avant de commencer le jeu
func (p *Personnage) createCaracter() {
	red := color.New(color.FgRed)
	fmt.Println("---- Création du personnage ----")
	fmt.Println("[1] Créer un personnage")
	fmt.Println("[2] Personnage par défault")
	fmt.Println("--------------------------------")
	choice := inputint()
	switch choice {
	case 1:
		hpMax := 0
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
			hpMax = 150
		case "Chasseur":
			hpMax = 125
		case "Aracaniste":
			hpMax = 100
		}
		clearConsole()
		p.Initialize(nom, classe, 1, hpMax, hpMax, []Item{{"Argent", 100}, {"Potions", 3}}, []string{"Coup de poing"})
		affichageMsg("Bienvenue, " + nom + " ! ")

	case 2:
		p.Initialize("Romain", "Chasseur", 1, 67, 125, []Item{{"Argent", 100}, {"Potions", 3}}, []string{"Coup de poing"})
		clearConsole()
		affichageMsg("Bienvenue, Romain !")

	default:
		clearConsole()
		red.Println("Veuillez saisir une donnée valide")
		p.createCaracter()
	}

}

func (p *Personnage) Initialize(nom string, classe string, niveau int, hp int, hpMax int, inventaire []Item, skill []string) {
	p.nom = nom
	p.classe = classe
	p.niveau = niveau
	p.currentHp = hp
	p.maxHP = hpMax
	p.inventory = inventaire
	p.skill = skill
}

// Fonctions pour la configuration du pseudonyme
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

// Afficher un message lettre par lettre
func speedMsg(message string, speed int, colorName string) {
	red := color.New(color.FgRed)
	blue := color.New(color.FgBlue)
	green := color.New(color.FgGreen)
	defaultColor := color.New(color.FgWhite) // Couleur par défaut

	var selectedColor *color.Color

	switch colorName {
	case "green":
		selectedColor = green
	case "red":
		selectedColor = red
	case "blue":
		selectedColor = blue
	default:
		selectedColor = defaultColor
	}

	for _, char := range message {
		selectedColor.Print(string(char))
		time.Sleep(time.Duration(speed) * time.Millisecond)
	}
}

func affichageMsg(message string) {
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

// Personnage structure
type Personnage struct {
	nom       string
	classe    string
	niveau    int
	currentHp int
	maxHP     int
	inventory []Item
	skill     []string
}

// Item structure
type Item struct {
	Name  string
	Price int
}

// Interactions console
func clearConsole() {
	var cmd *exec.Cmd
	cmd = exec.Command("cmd", "/c", "cls")

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
	scanner.Scan()

	return scanner.Text()
}

// Menu du jeu
func (p *Personnage) Menu() {
	red := color.New(color.FgRed)
	fmt.Println("----- Menu -----")
	fmt.Println("[1] Personnage")
	fmt.Println("[2] Inventaire")
	fmt.Println("[3] Marchand")
	fmt.Println("[4] Abilités")
	if p.niveau == 1 {
		fmt.Println("[5] Combat - tutoriel")
	} else {
		fmt.Println("[5] Combat")
	}
	fmt.Println("[6] Quitter le jeu")
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
		if p.niveau == 1 {
			p.battleTutorial()
		} else {
			p.battle()
		}
	case 6:
		clearConsole()
		red.Println("Fermeture du jeu...")
	default:
		clearConsole()
		red.Println("Veuillez saisir une donnée valide !")
		p.Menu()
	}
}

// Fonctions du tutoriel
func abilitiesTutorial() string {
	red := color.New(color.FgRed)
	fmt.Println("---- Abilités ----")
	fmt.Println("[1] Coup de poing")
	fmt.Println("[2] Frénésie sanguinaire")
	fmt.Println("[3] Lame démoniaque")
	fmt.Println("------------------")
	choice := inputint()
	switch choice {
	case 1:
		return "Coup de poing"
	case 2:
		return "Frénésie sanguinaire"
	case 3:
		return "Lame démoniaque"
	default:
		clearConsole()
		red.Println("Veuillez choisir une option valide")
		abilitiesTutorial()
	}
	return ""
}

func battleMenuTutorial() {
	fmt.Println("----- A votre tour -----")
	fmt.Print("[1] Attaque auto")
	speedMsg("<-- Ceci vous permet d'attaquer l'adversaire avec votre compétence basique", 20, "white")
	input()
	fmt.Print("[2] Abilités")
	speedMsg("<-- Ceci vous permet d'utiliser une abilité sur l'adversaire", 20, "white")
	input()
	fmt.Print("[3] Inventaire")
	speedMsg("<-- Ceci vous permet de consulter votre inventaire pendant le combat", 20, "white")
	fmt.Println("")
	fmt.Println("------------------------")
}

// Sous fonctions du menu
func (p *Personnage) accessInvetory() {
	red := color.New(color.FgRed)
	fmt.Println("--- Inventaire ---")
	for _, item := range p.inventory {
		fmt.Printf("%s : %d\n", item.Name, item.Price)
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
} //inventaire

func (p *Personnage) showSkills() {
	fmt.Println("--- Abilités ---")
	for i := 0; i < len(p.skill); i++ {
		fmt.Println("Sort n°", i+1, " : ", p.skill[i])
	}
	fmt.Println("----------------")
} //abilités

func (p *Personnage) displayInfo() {
	fmt.Println("--- ", p.nom, " ---")
	fmt.Println("Classe : ", p.classe)
	fmt.Println("Niveau : ", p.niveau)
	fmt.Println("Points de vie : ", p.currentHp, "/", p.maxHP)
	fmt.Println("--------------")
} //personnage

func (p *Personnage) marchand() {
	// couleurs textes
	red := color.New(color.FgRed)
	blue := color.New(color.FgBlue)
	green := color.New(color.FgGreen)

	var itemMarchand = []Item{
		{"Potions", 0},
		{"Potions de poison", 0},
		{"Epée", 5},
	}

	for {
		green.Println("Argent :  ", p.nbItem("Argent"))
		fmt.Println("--- Marchand ---")

		// Affichage des items
		for i, item := range itemMarchand {
			fmt.Printf("[%d] %s - (%d coro)\n", i+1, item.Name, item.Price)
		}

		// Option de sortie
		fmt.Printf("[%d] Sortir\n", len(itemMarchand)+1)
		fmt.Println("----------------")
		choice := inputint()
		if choice > 0 && choice <= len(itemMarchand) {
			selectedItem := itemMarchand[choice-1]
			if p.enoughMoney(selectedItem.Price) {
				p.addInventory(selectedItem.Name, 1)
				p.inventory[p.trouveIndex("Argent")].Price -= selectedItem.Price
				clearConsole()
				blue.Printf("Vous avez acheté : %s pour %d coro !\n", selectedItem.Name, selectedItem.Price)
			} else {
				clearConsole()
				red.Println("Vous n'avez pas assez d'argent !")
			}
		} else if choice == len(itemMarchand)+1 {
			clearConsole()
			blue.Println("Sortie du marchand")
			p.Menu()
			break
		} else {
			clearConsole()
			red.Println("Veuillez saisir une donnée valide")
		}
	}
} //marchand

func (p *Personnage) battleTutorial() {
	red := color.New(color.FgRed)
	blue := color.New(color.FgBlue)
	green := color.New(color.FgGreen)
	clearConsole()
	//blue := color.New(color.FgBlue)
	speedMsg("Bienvenue dans le tutoriel de combat !", 30, "blue")
	input()
	clearConsole()
	red.Print("Ennemi 1 - 100 / 100")
	speedMsg(" <-- Ici sont affichés les points de vie de l'ennemi", 30, "white")
	input()
	green.Print("Vous - 100 / 100")
	speedMsg(" <-- Et ici les vôtres", 20, "white")
	input()
	clearConsole()
	speedMsg("Le combat se joue en tour par tour", 20, "blue")
	fmt.Println("")
	speedMsg("A chaque fois que vous jouez, plusieurs options s'offrent à vous : ", 20, "blue")
	fmt.Println("")
	battleMenuTutorial()
	input()
	clearConsole()
	speedMsg("Lors de chaque attaque, vous verrez le nombre de dégats infligés : ", 20, "blue")
	fmt.Println()
	green.Print("Vous avez infligé 20 dégats à Ennemi 1")
	input()
	clearConsole()
	speedMsg("Mais vous pouvez également en recevoir : ", 20, "blue")
	fmt.Println()
	red.Print("Vous avez reçu 50 dégats (coup critique) de Ennemi 1 !")
	input()
	clearConsole()
	speedMsg("Lors de votre tour, vous pourrez également utiliser vos abilités : ", 20, "blue")
	fmt.Println("")
	spell := abilitiesTutorial()
	speedMsg(spell+" à infligé 50 dégats à Ennemi 1", 20, "green")
	input()
	clearConsole()
	speedMsg("Félicitation, vous êtes prét pour votre premier combat ! Bonne chance", 20, "blue")
	input()
	clearConsole()
	p.niveau = 2
	blue.Println("Vous avez atteint le niveau 2 !")
	p.Menu()
} // Tutoriel de combat

func (p *Personnage) battle() {
	clearConsole()
	blue := color.New(color.FgBlue)
	blue.Println("Prochain update 🤞")
	p.Menu()
}

// Interactions avec l'inventaire
func (p *Personnage) takePot() {
	red := color.New(color.FgRed)
	blue := color.New(color.FgBlue)
	if p.nbItem("Potions") > 0 {
		if p.currentHp == p.maxHP {
			red.Println("Vous êtes déjà au maximum de points de vie !")
		} else {
			if p.maxHP-p.currentHp < 50 {
				p.currentHp = p.maxHP
			} else {
				p.currentHp += 50
			}
			p.removeInventory("Potions", 1)
			blue.Println("- 1 potion, points de vie : ", p.currentHp)
			blue.Println("Potions restantes : ", p.nbItem("Potions"))
		}
	}
}

func (p *Personnage) enoughMoney(cout int) bool {
	if p.nbItem("Argent")-cout < 0 {
		return false
	} else {
		return true
	}
}

func (p *Personnage) addInventory(itemName string, nb int) {
	for i, item := range p.inventory {
		if item.Name == itemName {
			p.inventory[i].Price += nb
			return
		}
	}
	newItem := Item{Name: itemName, Price: nb}
	p.inventory = append(p.inventory, newItem)
}

func (p *Personnage) removeInventory(itemName string, quantity int) {
	for i, item := range p.inventory {
		if item.Name == itemName {
			// L'élément existe dans la liste.
			if item.Price <= quantity {
				// Retirez complètement l'élément si la quantité est suffisante.
				p.inventory = append(p.inventory[:i], p.inventory[i+1:]...)
			} else {
				// Mettez à jour la quantité de l'élément si la quantité est insuffisante.
				p.inventory[i].Price -= quantity
			}
		}
	}
}

func (p *Personnage) nbItem(nomItem string) int {
	for _, item := range p.inventory {
		if item.Name == nomItem {
			return item.Price
		}
	}
	return -1
}

func (p *Personnage) trouveIndex(nomItem string) int {
	index := -1
	for i, item := range p.inventory {
		if item.Name == nomItem {
			index = i
			break
		}
	}
	return index
}
