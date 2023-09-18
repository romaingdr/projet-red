package main

import (
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"src/utils"
	"strconv"
	"time"
)

// Fonction main
func main() {
	utils.ClearConsole()
	var p1 Personnage
	p1.createCharacter()
	p1.Menu()
}

var (
	Red      = color.New(color.FgRed)
	Blue     = color.New(color.FgBlue)
	Green    = color.New(color.FgGreen)
	Monstres = []Ennemy{{"Red soldier", 200, 200, 20, 50, 20}}
)

type Ennemy struct {
	Name           string
	HpCurrent      int
	HpMax          int
	DamagesMin     int
	DamgesMax      int
	CriticalChance int
}

type Spell struct {
	Name        string
	Description string
	Damages     int
}

// Personnage structure
type Personnage struct {
	nom       string
	classe    string
	niveau    int
	currentHp int
	maxHP     int
	inventory []Item
	skill     []Spell
}

// Item structure
type Item struct {
	Name     string
	Quantite int
}

// Article structure
type Article struct {
	Name        string
	Price       int
	Description string
	Ad          int
	Health      int
	Unique      bool
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
		spells := []Spell{}
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
		classe := utils.ChooseClass()
		switch classe {
		case "Titan":
			hpMax = 180
			spells = []Spell{{"Auto", "Attaque automatique du titan", 10},
				{"La bulle", "Le titan s'enferme dans une bulle et reduit les dégâts subis", 0},
				{"Frappe ultime", "Le titan inflige une violente attaque", 50},
				{"Dé titanesque", "Le titan a 67% de chance d'infliger 400% de dégats, sinon il perd 70 points de vies", 40},
				{"(%) Critical chance", "inflige le double des dégats", 15}}
		case "Chasseur":
			hpMax = 135
			spells = []Spell{{"Auto", "Attaque automatique du chasseur", 20},
				{"Lame Sanglante", "Inflige un coup de lame empoisonnée", 25},
				{"Maitrise du terrain", "Le chasseur se concentre pour infliger une violente attaque", 0},
				{"Attaque rapide", "Inflige 200% des dégats de l'attaque automatique", 30},
				{"(%) Critical chance", "inflige le double des dégats", 10}}
		case "Arcaniste":
			hpMax = 100
			spells = []Spell{{"Auto", "Attaque automatique de l'arcaniste", 30},
				{"Trou noir", "Execute l'ennemi en dessous de 15% de points de vie", 40},
				{"Alteration de l'ame", "Vol de vie (150% des dégats infligés)", 15},
				{"Foudre", "La foudre s'abat sur l'ennemi et lui inflige des dégats", 70},
				{"(%) Critical chance", "inflige le double des dégats", 10}}
		}
		utils.ClearConsole()
		p.Initialize(nom, classe, 1, hpMax, hpMax, []Item{{"Argent", 10000}, {"Potions", 3}}, spells)
		utils.SpeedMsg("Bienvenue, "+nom+" ! ", 60, "blue")
		utils.Input()
		utils.ClearConsole()

	case 2:
		spells := []Spell{{"Auto", "Attaque automatique du champion", 15},
			{"Lame Sanglante", "Inflige un coup de lame empoisonnée", 25},
			{"Maitrise du terrain", "Le chasseur se concentre pour infliger une violente attaque", 0},
			{"Attaque rapide", "Inflige une attaque automatique double", 30}}
		p.Initialize("Romain", "Chasseur", 2, 125, 125, []Item{{"Argent", 10000}, {"Potions", 3}}, spells)
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

func (p *Personnage) abilitiesBattle(e *Ennemy) {
	fmt.Printf("%-20s%-10s%-60s\n", "Nom", "Dégâts", "Description")
	fmt.Println("------------------------------------------------------------------")
	for i := 1; i < 4; i++ {
		fmt.Printf("[%d] %-20s %-10d %-s\n", i, p.skill[i].Name, p.skill[i].Damages, p.skill[i].Description)
	}
	fmt.Println("------------------------------------------------------------------")
	fmt.Println("[4] Sortir")
	choice, _ := utils.Inputint()
	switch choice {
	case 1:
		degats := p.skill[1].Damages
		Green.Println("Vous avez infligé " + strconv.Itoa(degats) + " dégats avec " + p.skill[1].Name)
		e.HpCurrent -= degats
		utils.Input()
	case 2:
		degats := p.skill[2].Damages
		Green.Println("Vous avez infligé " + strconv.Itoa(degats) + " dégats avec " + p.skill[2].Name)
		e.HpCurrent -= degats
		utils.Input()
	case 3:
		degats := p.skill[3].Damages
		Green.Println("Vous avez infligé " + strconv.Itoa(degats) + " dégats avec " + p.skill[3].Name)
		e.HpCurrent -= degats
		utils.Input()
	case 4:
		utils.ClearConsole()
		p.playerRound(e)
	default:
		utils.ClearConsole()
		Red.Println("Veuillez saisir une donnée valide !")
		p.abilitiesBattle(e)
	}
}

func (p *Personnage) playerRound(e *Ennemy) {
	Red.Println(e.Name + " - " + strconv.Itoa(e.HpCurrent) + "/" + strconv.Itoa(e.HpMax))
	Green.Println(p.nom + " - " + strconv.Itoa(p.currentHp) + "/" + strconv.Itoa(p.maxHP))
	fmt.Println("----- A votre tour -----")
	fmt.Println("[1] Attaque auto")
	fmt.Println("[2] Abilités")
	fmt.Println("[3] Inventaire")
	fmt.Println("------------------------")
	choice, _ := utils.Inputint()
	switch choice {
	case 1:
		utils.ClearConsole()
		fmt.Println("Auto")
		p.playerRound(e)
	case 2:
		utils.ClearConsole()
		p.abilitiesBattle(e)
	case 3:
		utils.ClearConsole()
		fmt.Println("Inventaire")
		p.playerRound(e)
	default:
		utils.ClearConsole()
		Red.Println("Veuillez saisir une donnée valide")
		p.playerRound(e)
	}

}

func (p *Personnage) ennemyRound(e *Ennemy) {
	rand.Seed(time.Now().UnixNano())
	Red.Println(e.Name + " - " + strconv.Itoa(e.HpCurrent) + "/" + strconv.Itoa(e.HpMax))
	Green.Println(p.nom + " - " + strconv.Itoa(p.currentHp) + "/" + strconv.Itoa(p.maxHP))
	fmt.Println("----- Tour de " + e.Name + " -----")
	utils.SpeedMsg(e.Name+" attaque...\n", 20, "default")
	degats := rand.Intn(e.DamgesMax-e.DamagesMin+1) + e.DamagesMin
	critic := rand.Intn(2)
	if critic == 1 {
		degats *= 2
	}
	time.Sleep(2 * time.Second)
	p.currentHp -= degats
	if critic == 1 {
		Red.Println("[COUP CRITIQUE] Vous avez reçu " + strconv.Itoa(degats) + "dégats")
	} else {
		Red.Println("Vous avez reçu " + strconv.Itoa(degats) + "dégats")
	}
	fmt.Println("---------------------------------")
	utils.Input()
}

// Initialize initialise les données du personnage.
func (p *Personnage) Initialize(nom string, classe string, niveau int, hp int, hpMax int, inventaire []Item, skill []Spell) {
	p.nom = nom
	p.classe = classe
	p.niveau = niveau
	p.currentHp = hp
	p.maxHP = hpMax
	p.inventory = inventaire
	p.skill = skill
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
	utils.BattleMenuTutorial()
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
	spell := utils.AbilitiesTutorial()
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

func (p *Personnage) isDead(e *Ennemy) bool {
	if p.currentHp <= 0 || e.HpCurrent <= 0 {
		return true
	} else {
		return false
	}
}

// battle est la fonction de combat
func (p *Personnage) battle() {
	utils.ClearConsole()

	ennemi1 := Monstres[p.niveau-2]
	utils.SpeedMsg(p.nom+" VS "+ennemi1.Name+"\n", 20, "red")
	fmt.Println("-----------------------------")
	utils.SpeedMsg("Statistiques de l'ennemi : \n", 20, "default")
	fmt.Println("Points de vie : " + strconv.Itoa(ennemi1.HpMax))
	fmt.Println("Dégats : " + strconv.Itoa(ennemi1.DamagesMin) + " - " + strconv.Itoa(ennemi1.DamgesMax))
	fmt.Println("Chance de coup critique : " + strconv.Itoa(ennemi1.CriticalChance))
	fmt.Println("-----------------------------")
	utils.SpeedMsg("Bonne chance combattant !", 20, "green")
	utils.Input()
	utils.ClearConsole()
	for !(p.isDead(&ennemi1)) {
		p.ennemyRound(&ennemi1)
		utils.ClearConsole()
		if !(p.isDead(&ennemi1)) {
			p.playerRound(&ennemi1)
			utils.ClearConsole()
		}
	}
}

// accessInventory permet au joueur d'accéder à son inventaire.
func (p *Personnage) accessInventory() {

	fmt.Println("Item            Quantité")
	fmt.Println("---------------------------")
	for _, item := range p.inventory {
		// Utilisation de fmt.Printf pour aligner les colonnes
		fmt.Printf("%-20s %-7d\n", item.Name, item.Quantite)
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
	fmt.Printf("%-20s%-10s%-60s\n", "Nom", "Dégâts", "Description")
	fmt.Println("------------------------------------------------------------------")

	for _, spell := range p.skill {
		fmt.Printf("%-20s%-10d%-60s\n", spell.Name, spell.Damages, spell.Description)
	}
	fmt.Println("-------------------------------------------------------------------------------------------")
}

// displayInfo affiche les informations du personnage.
func (p *Personnage) displayInfo() {
	fmt.Println("--- ", p.nom, " ---")
	fmt.Println("Classe : ", p.classe)
	fmt.Println("Niveau : ", p.niveau)
	fmt.Println("Points de vie : ", p.currentHp, "/", p.maxHP)
	fmt.Println("--------------")
}

// Fonction pour vérifier si un item avec un nom donné se trouve dans une liste d'items
func (p *Personnage) alreadyBuy(itemName string) bool {
	for _, item := range p.inventory {
		if item.Name == itemName {
			return true
		}
	}
	return false
}

// Marchand permet au joueur d'interagir avec le marchand.
func (p *Personnage) Marchand() {
	var itemMarchand = []Article{
		{"Potions", 10, "Potion pour récupèrer 40 points de vie", 0, 0, false},
		{"Guinzoo", 100, "+20% hp | +15 ad", 15, p.maxHP / 5, true},
		{"Masque grenouille", 1000, "+600% hp", 0, p.maxHP * 6, true},
		{"Avrilvert", 350, "+100% ad", p.skill[0].Damages, 0, true},
		{"Rook", 450, "+80% hp", 0, p.maxHP * 80 / 100, true},
		{"Anneau de gel", 500, "+60%hp | +20 ad ", 20, p.maxHP * 60 / 100, true},
	}

	for {
		Green.Println("Argent :  ", p.nbItem("Argent"))
		fmt.Println("    Article             Prix                Description")
		fmt.Println("-----------------------------------------------------")
		for i, item := range itemMarchand {
			if p.alreadyBuy(item.Name) && item.Unique == true {
				Blue.Printf("%s - acheté !\n", item.Name)
			} else {
				fmt.Printf("[%d] %-20s %-7d %-s\n", i+1, item.Name, item.Price, item.Description)
			}
		}
		fmt.Println("-----------------------------------------------------")
		fmt.Printf("[%d] Sortir \n", len(itemMarchand)+1)
		choice, _ := utils.Inputint()
		if choice > 0 && choice <= len(itemMarchand) {
			selectedItem := itemMarchand[choice-1]
			if p.enoughMoney(selectedItem.Price) {
				if !(p.alreadyBuy(selectedItem.Name)) {
					p.addInventory(selectedItem.Name, 1)
					p.inventory[p.findIndex("Argent")].Quantite -= selectedItem.Price
					utils.ClearConsole()
					p.skill[0].Damages += selectedItem.Ad
					p.maxHP += selectedItem.Health
					p.currentHp += selectedItem.Health
					Blue.Printf("Vous avez acheté : %s pour %d coro !\n", selectedItem.Name, selectedItem.Price)
				} else {
					utils.ClearConsole()
					Red.Println("Vous avez déjà acheté cet item")
				}
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
			p.inventory[i].Quantite += quantity
			return
		}
	}
	newItem := Item{Name: itemName, Quantite: quantity}
	p.inventory = append(p.inventory, newItem)
}

// removeInventory retire un objet de l'inventaire du joueur.
func (p *Personnage) removeInventory(itemName string, quantity int) {
	for i, item := range p.inventory {
		if item.Name == itemName {
			// L'élément existe dans la liste.
			if item.Quantite <= quantity {
				// Retire complètement l'élément si la quantité est suffisante.
				p.inventory = append(p.inventory[:i], p.inventory[i+1:]...)
			} else {
				// Met à jour la quantité de l'élément si la quantité est insuffisante.
				p.inventory[i].Quantite -= quantity
			}
		}
	}
}

// nbItem renvoie la quantité d'un objet dans l'inventaire du joueur.
func (p *Personnage) nbItem(itemName string) int {
	for _, item := range p.inventory {
		if item.Name == itemName {
			return item.Quantite
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
