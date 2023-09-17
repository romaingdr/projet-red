package main

import (
	"fmt"
	"github.com/fatih/color"
	"src/utils"
)

// Fonction main
func main() {
	utils.ClearConsole()
	var p1 Personnage
	p1.createCharacter()
	p1.Menu()
}

var (
	Red   = color.New(color.FgRed)
	Blue  = color.New(color.FgBlue)
	Green = color.New(color.FgGreen)
)

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

// createCharacter initialise un nouveau personnage.
func (p *Personnage) createCharacter() {
	fmt.Println("---- Création du personnage ----")
	fmt.Println("[1] Créer un personnage")
	fmt.Println("[2] Personnage par défaut")
	fmt.Println("--------------------------------")
	choice, _ := utils.Inputint()
	switch choice {
	case 1:
		hpMax := 0
		nom := "0"
		utils.ClearConsole()
		print("Nom de votre personnage >> ")
		nom = utils.Input()
		for !(utils.OnlyLetters(nom)) {
			utils.ClearConsole()
			Red.Println("Veuillez n'utiliser que des lettres | 10 caractères maximum")
			print("Nom de votre personnage >> ")
			nom = utils.Input()
		}
		nom = utils.CapitalizeString(nom)
		utils.ClearConsole()
		classe := chooseClass()
		switch classe {
		case "Titan":
			hpMax = 150
		case "Chasseur":
			hpMax = 125
		case "Arcaniste":
			hpMax = 100
		}
		utils.ClearConsole()
		p.Initialize(nom, classe, 1, hpMax, hpMax, []Item{{"Argent", 100}, {"Potions", 3}}, []string{"Coup de poing"})
		utils.SpeedMsg("Bienvenue, "+nom+" ! ", 60, "blue")
		utils.Input()
		utils.ClearConsole()

	case 2:
		p.Initialize("Romain", "Chasseur", 1, 125, 125, []Item{{"Argent", 100}, {"Potions", 3}}, []string{"Coup de poing"})
		utils.ClearConsole()
		utils.SpeedMsg("Bienvenue, Romain !", 60, "blue")
		utils.Input()
		utils.ClearConsole()

	default:
		utils.ClearConsole()
		Red.Println("Veuillez saisir une donnée valide")
		p.createCharacter()
	}

}

// Initialize initialise les données du personnage.
func (p *Personnage) Initialize(nom string, classe string, niveau int, hp int, hpMax int, inventaire []Item, skill []string) {
	p.nom = nom
	p.classe = classe
	p.niveau = niveau
	p.currentHp = hp
	p.maxHP = hpMax
	p.inventory = inventaire
	p.skill = skill
}

// choixClasse affiche les classes disponibles et permet à l'utilisateur de choisir une classe.
func chooseClass() string {
	utils.ClearConsole()

	Green.Println("Nom du personnage validé !")
	Blue.Println("Choisissez votre classe : ")
	classes := []string{"Titan", "Arcaniste", "Chasseur"}
	println("")
	println("[1] Titan : a_completer")
	println("[2] Arcaniste : a_completer")
	println("[3] Chasseur : a_completer")
	println("")

	for {
		choice, _ := utils.Inputint()
		if choice > 0 && choice < 4 {
			return classes[choice-1]
		} else {
			utils.ClearConsole()
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

// Menu affiche le menu principal du jeu.
func (p *Personnage) Menu() {

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

	choice, _ := utils.Inputint()

	switch choice {
	case 1:
		utils.ClearConsole()
		p.displayInfo()
		p.Menu()
		break
	case 2:
		utils.ClearConsole()
		p.accessInventory()
		p.Menu()
		break
	case 3:
		utils.ClearConsole()
		p.Marchand()
	case 4:
		utils.ClearConsole()
		p.showSkills()
		p.Menu()
	case 5:
		utils.ClearConsole()
		if p.niveau == 1 {
			p.battleTutorial()
		} else {
			p.battle()
		}
	case 6:
		utils.ClearConsole()
		Red.Println("Fermeture du jeu...")
	default:
		utils.ClearConsole()
		Red.Println("Veuillez saisir une donnée valide !")
		p.Menu()
	}
}

// abilitiesTutorial affiche les abilités disponibles dans le tutoriel de combat et permet à l'utilisateur d'en choisir une.
func abilitiesTutorial() string {

	fmt.Println("---- Abilités ----")
	fmt.Println("[1] Coup de poing")
	fmt.Println("[2] Frénésie sanguinaire")
	fmt.Println("[3] Lame démoniaque")
	fmt.Println("------------------")
	choice, _ := utils.Inputint()
	switch choice {
	case 1:
		return "Coup de poing"
	case 2:
		return "Frénésie sanguinaire"
	case 3:
		return "Lame démoniaque"
	default:
		utils.ClearConsole()
		Red.Println("Veuillez choisir une option valide")
		return abilitiesTutorial()
	}
}

// battleMenuTutorial affiche le menu du tutoriel de combat.
func battleMenuTutorial() {
	fmt.Println("----- A votre tour -----")
	fmt.Print("[1] Attaque auto")
	utils.SpeedMsg("<-- Ceci vous permet d'attaquer l'adversaire avec votre compétence basique", 20, "white")
	utils.Input()
	fmt.Print("[2] Abilités")
	utils.SpeedMsg("<-- Ceci vous permet d'utiliser une abilité sur l'adversaire", 20, "white")
	utils.Input()
	fmt.Print("[3] Inventaire")
	utils.SpeedMsg("<-- Ceci vous permet de consulter votre inventaire pendant le combat", 20, "white")
	fmt.Println("")
	fmt.Println("------------------------")
}

// battleTutorial est le tutoriel de combat.
func (p *Personnage) battleTutorial() {

	utils.ClearConsole()
	utils.SpeedMsg("Bienvenue dans le tutoriel de combat !", 30, "blue")
	utils.Input()
	utils.ClearConsole()
	Red.Print("Ennemi 1 - 100 / 100")
	utils.SpeedMsg(" <-- Ici sont affichés les points de vie de l'ennemi", 30, "white")
	utils.Input()
	Green.Print("Vous - 100 / 100")
	utils.SpeedMsg(" <-- Et ici les vôtres", 20, "white")
	utils.Input()
	utils.ClearConsole()
	utils.SpeedMsg("Le combat se joue en tour par tour", 20, "blue")
	fmt.Println("")
	utils.SpeedMsg("A chaque fois que vous jouez, plusieurs options s'offrent à vous : ", 20, "blue")
	fmt.Println("")
	battleMenuTutorial()
	utils.Input()
	utils.ClearConsole()
	utils.SpeedMsg("Lors de chaque attaque, vous verrez le nombre de dégats infligés : ", 20, "blue")
	fmt.Println()
	Green.Print("Vous avez infligé 20 dégats à Ennemi 1")
	utils.Input()
	utils.ClearConsole()
	utils.SpeedMsg("Mais vous pouvez également en recevoir : ", 20, "blue")
	fmt.Println()
	Red.Print("Vous avez reçu 50 dégats (coup critique) de Ennemi 1 !")
	utils.Input()
	utils.ClearConsole()
	utils.SpeedMsg("Lors de votre tour, vous pourrez également utiliser vos abilités : ", 20, "blue")
	fmt.Println("")
	spell := abilitiesTutorial()
	utils.SpeedMsg(spell+" à infligé 50 dégats à Ennemi 1", 20, "green")
	utils.Input()
	utils.ClearConsole()
	utils.SpeedMsg("Félicitation, vous êtes prêt pour votre premier combat ! Bonne chance", 20, "blue")
	utils.Input()
	utils.ClearConsole()
	p.niveau = 2
	Blue.Println("Vous avez atteint le niveau 2 !")
	p.Menu()
}

// battle est la fonction de combat (non implémentée).
func (p *Personnage) battle() {
	utils.ClearConsole()

	Blue.Println("Prochain update 🤞")
	p.Menu()
}

// accessInventory permet au joueur d'accéder à son inventaire.
func (p *Personnage) accessInventory() {

	fmt.Println("Item            Quantité")
	fmt.Println("---------------------------")
	for _, item := range p.inventory {
		// Utilisation de fmt.Printf pour aligner les colonnes
		fmt.Printf("%-20s %-7d\n", item.Name, item.Price)
	}
	fmt.Println("---------------------------")
	fmt.Println("[1] Utiliser une potion")
	fmt.Println("[2] Quitter l'inventaire")
	choice, _ := utils.Inputint()
	switch choice {
	case 1:
		utils.ClearConsole()
		p.takePotion()
		p.accessInventory()
	case 2:
		utils.ClearConsole()
	default:
		utils.ClearConsole()
		Red.Println("Veuillez saisir une donnée valide")
		p.accessInventory()
	}
}

// showSkills affiche les abilités du personnage.
func (p *Personnage) showSkills() {
	fmt.Println("--- Abilités ---")
	for i := 0; i < len(p.skill); i++ {
		fmt.Println("Sort n°", i+1, " : ", p.skill[i])
	}
	fmt.Println("----------------")
}

// displayInfo affiche les informations du personnage.
func (p *Personnage) displayInfo() {
	fmt.Println("--- ", p.nom, " ---")
	fmt.Println("Classe : ", p.classe)
	fmt.Println("Niveau : ", p.niveau)
	fmt.Println("Points de vie : ", p.currentHp, "/", p.maxHP)
	fmt.Println("--------------")
}

// Marchand permet au joueur d'interagir avec le marchand.
func (p *Personnage) Marchand() {

	var itemMarchand = []Item{
		{"Potions", 0},
		{"Potions de poison", 0},
		{"Epée", 5},
		{"Casque en fer", 20},
	}

	for {
		Green.Println("Argent :  ", p.nbItem("Argent"))
		fmt.Println("    Article             Prix")
		fmt.Println("---------------------------")
		for i, item := range itemMarchand {
			// Utilisation de fmt.Printf pour aligner les colonnes
			fmt.Printf("[%d] %-20s %-7d\n", i+1, item.Name, item.Price)
		}
		fmt.Println("---------------------------")
		fmt.Printf("[%d] Sortir \n", len(itemMarchand)+1)
		choice, _ := utils.Inputint()
		if choice > 0 && choice <= len(itemMarchand) {
			selectedItem := itemMarchand[choice-1]
			if p.enoughMoney(selectedItem.Price) {
				p.addInventory(selectedItem.Name, 1)
				p.inventory[p.findIndex("Argent")].Price -= selectedItem.Price
				utils.ClearConsole()
				Blue.Printf("Vous avez acheté : %s pour %d coro !\n", selectedItem.Name, selectedItem.Price)
			} else {
				utils.ClearConsole()
				Red.Println("Vous n'avez pas assez d'argent !")
			}
		} else if choice == len(itemMarchand)+1 {
			utils.ClearConsole()
			Blue.Println("Sortie du marchand")
			p.Menu()
			break
		} else {
			utils.ClearConsole()
			Red.Println("Veuillez saisir une donnée valide")
		}
	}
}

// takePotion permet au joueur de prendre une potion si il en a une.
func (p *Personnage) takePotion() {

	if p.nbItem("Potions") > 0 {
		if p.currentHp == p.maxHP {
			Red.Println("Vous êtes déjà au maximum de points de vie !")
		} else {
			if p.maxHP-p.currentHp < 50 {
				p.currentHp = p.maxHP
			} else {
				p.currentHp += 50
			}
			p.removeInventory("Potions", 1)
			Blue.Println("- 1 potion, points de vie : ", p.currentHp)
			Blue.Println("Potions restantes : ", p.nbItem("Potions"))
		}
	}
}

// enoughMoney vérifie si le joueur a suffisamment d'argent pour acheter un objet.
func (p *Personnage) enoughMoney(cost int) bool {
	if p.nbItem("Argent")-cost < 0 {
		return false
	} else {
		return true
	}
}

// addInventory ajoute un objet à l'inventaire du joueur.
func (p *Personnage) addInventory(itemName string, quantity int) {
	for i, item := range p.inventory {
		if item.Name == itemName {
			p.inventory[i].Price += quantity
			return
		}
	}
	newItem := Item{Name: itemName, Price: quantity}
	p.inventory = append(p.inventory, newItem)
}

// removeInventory retire un objet de l'inventaire du joueur.
func (p *Personnage) removeInventory(itemName string, quantity int) {
	for i, item := range p.inventory {
		if item.Name == itemName {
			// L'élément existe dans la liste.
			if item.Price <= quantity {
				// Retire complètement l'élément si la quantité est suffisante.
				p.inventory = append(p.inventory[:i], p.inventory[i+1:]...)
			} else {
				// Met à jour la quantité de l'élément si la quantité est insuffisante.
				p.inventory[i].Price -= quantity
			}
		}
	}
}

// nbItem renvoie la quantité d'un objet dans l'inventaire du joueur.
func (p *Personnage) nbItem(itemName string) int {
	for _, item := range p.inventory {
		if item.Name == itemName {
			return item.Price
		}
	}
	return -1
}

// findIndex renvoie l'index d'un objet dans l'inventaire du joueur.
func (p *Personnage) findIndex(itemName string) int {
	index := -1
	for i, item := range p.inventory {
		if item.Name == itemName {
			index = i
			break
		}
	}
	return index
}
