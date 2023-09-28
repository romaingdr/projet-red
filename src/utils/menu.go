// FICHIER UTILISE POUR LA GESTION DU MENU ET DE SES SOUS MENUS

package utils

import (
	"fmt"
	"os"
	"strconv"
)

// Variable globales utiles pour le menu
var (
	db *QuickDB
)

// Menu affiche le menu principal du jeu.
func (p *Personnage) Menu() {
	Cyan.Println("----- Menu -----")
	fmt.Println("[1] Personnage")
	fmt.Println("[2] Inventaire")
	fmt.Println("[3] Marchand")
	fmt.Println("[4] Forgeron")
	fmt.Println("[5] Abilités")
	if p.niveau == 0 {
		fmt.Println("[6] Tutoriel")
	} else if p.niveau == 1 && p.ennemi == 0 {
		fmt.Println("[6] Commencer l'histoire")
	} else {
		if p.niveau < 7 {
			fmt.Println("[6] continuer l'histoire")
		} else if p.niveau == 7 {
			fmt.Println("[6] fin ?")
		} else {
			Green.Println("[6] Jeu terminé | Crédits")
		}
	}
	fmt.Println("[7] Avancement")
	fmt.Println("[8] Multijoueur - PvP")
	fmt.Println("[9] Multijoueur - Objectifs")
	fmt.Println("[10] Sauvegarder la progression")
	fmt.Println("[11] Quitter le jeu")
	Cyan.Println("----------------")
	fmt.Println("[12] Qui sont-ils | Bonus option")

	choice, _ := Inputint()

	switch choice {

	// Affichage des infos du joueur
	case 1:
		ClearConsole()
		p.displayInfo()
		p.Menu()
		break

	// Affichage de l'inventaire
	case 2:
		ClearConsole()
		p.accessInventory()
		p.Menu()
		break

	// Affichage du marchand
	case 3:
		ClearConsole()
		p.Marchand()

	// Affichage des abilités
	case 4:
		ClearConsole()
		p.forgeron()

	// Combat tutoriel si niveau 1 sinon menu de combat
	case 5:
		ClearConsole()
		p.showSkills()
		p.Menu()

	case 6:
		ClearConsole()
		if p.niveau == 0 {
			p.battleTutorial()
		} else {
			p.battle()
		}

	case 7:
		ClearConsole()
		avancement(p)
	case 8:
		ClearConsole()
		MultiStartScreen(p)
	case 9:
		ClearConsole()
		multiObjectives(p)
	case 10:
		save(p)
	case 11:
		ClearConsole()
		os.Exit(0)
	case 12:
		ClearConsole()
		Red.Println("QUI SONT-ILS ? : ABBA | Steven Spielberg | Queen")
		p.Menu()
	// Choix non proposé
	default:
		ClearConsole()
		Red.Println("Veuillez saisir une donnée valide !")
		p.Menu()
	}
}

// avancement affiche la timeline de l'histoire du jeu, les ennemis rencontrés
func avancement(p *Personnage) {
	fmt.Println("  -------------------")
	fmt.Println(" 🔫 -  World War 3")
	fmt.Println("  |    Basic Soldier 1")
	fmt.Println("  |    Basic Soldier 2")
	fmt.Println("  |    RED SOLDIER")
	fmt.Println("  | ")
	fmt.Println(" 🏰  -  Chateau")
	fmt.Println("  |    Bat 1")
	fmt.Println("  |    Bat 2")
	fmt.Println("  |    FALLEN KING")
	fmt.Println("  | ")
	fmt.Println(" 🌱  -  Jardin")
	fmt.Println("  |    Snail")
	fmt.Println("  |    Mosquito")
	fmt.Println("  |    Wasp")
	fmt.Println("  | ")
	fmt.Println(" 💧  -  Ville sous marine")
	fmt.Println("  |    Dolphin")
	fmt.Println("  |    Shark")
	fmt.Println("  |    Mermaid")
	fmt.Println("  | ")
	fmt.Println(" 🛸 -  Starship")
	fmt.Println("  |   ")
	fmt.Println(" 🪐  - Uranus")
	fmt.Println("  -----------------------")
	fmt.Print("appuyez sur entrée pour continuer")
	Input()
	ClearConsole()
	p.Menu()

}

// forgeron ouvre le sous-menu du forgeron dans lequel sont créés les buff avec les items drop par les ennemis
func (p *Personnage) forgeron() {
	outilsMarchand := [][]string{
		{"Victoire militaire", "Arme de soldat", "Casque de soldat", "Légion militaire"},
		{"Victoire médiévale", "Ailes de chauve souris", "Dents de chauve souris", "Couronne royale"},
		{"Victoire jardinale", "Coquille d'escargot", "Fiole de venin", "Dar royal"},
		{"Victoire abyssale", "Queue du dauphin", "Dent de la mer", "Trident empoisonné"},
	}

	Cyan.Println("---- Items à forger ----")
	for i := 0; i < len(outilsMarchand); i++ {
		if p.findIndex(outilsMarchand[i][0]) != -1 {
			Cyan.Println(outilsMarchand[i][0] + " - Acheté !")
		} else {
			fmt.Print("[" + strconv.Itoa(i+1) + "] - " + outilsMarchand[i][0] + " : ")
			for y := 1; y < len(outilsMarchand[i]); y++ {
				if p.findIndex(outilsMarchand[i][y]) != -1 {
					Green.Print(outilsMarchand[i][y])
					fmt.Print(", ")
				} else {
					Red.Print(outilsMarchand[i][y])
					fmt.Print(", ")
				}
			}
			fmt.Println("")
		}
	}
	Cyan.Println("------------------------")
	fmt.Println("[" + strconv.Itoa(len(outilsMarchand)+1) + "] - Sortir")
	Cyan.Println("Chaque forge d'item coûte 5 coros !")
	choice, _ := Inputint()
	var itemsRequis []string
	switch {
	case choice > 0 && choice < len(outilsMarchand):
		peutBuild := true
		if p.findIndex(outilsMarchand[choice-1][0]) == -1 {
			if p.enoughMoney(5) {
				itemsRequis = outilsMarchand[choice-1][1:]
				for i := 0; i < len(itemsRequis); i++ {
					requis := itemsRequis[i]
					if p.findIndex(requis) == -1 {
						ClearConsole()
						Red.Println("Vous n'avez pas les composants nécessaires ! (" + requis + ")")
						peutBuild = false
						p.forgeron()
					}
				}
			} else {
				peutBuild = false
				ClearConsole()
				Red.Println("Vous n'avez pas assez d'argent !")
				p.forgeron()
			}
		} else {
			peutBuild = false
			ClearConsole()
			Red.Println("Vous avez déjà acheté cet item !")
			p.forgeron()
		}

		if peutBuild {
			// ajoute l'item
			p.addInventory(outilsMarchand[choice-1][0], 1)

			// supprimer ses composants
			for i := 0; i < len(itemsRequis); i++ {
				requis := itemsRequis[i]
				indice := p.findIndex(requis)
				fmt.Println(indice)
				p.inventory = append(p.inventory[:indice], p.inventory[indice+1:]...)
				fmt.Println(p.inventory)

			}

			// enlever l'argent
			p.removeInventory("Argent", 5)

			ClearConsole()
			Green.Println("Vous avez acheté " + outilsMarchand[choice-1][0])

			// Ajout des buff stats
			switch choice {
			case 1:
				p.skill[4].Damages += 10
				p.skill[0].Damages += 10
				Green.Println("[BUFF] Vous avez gagné 10% de chance de coup critique")
				Green.Println("[BUFF] Vous avez gagné 10 de dégats automatique")
			case 2:
				p.maxHP += 20
				p.currentHp += 20
				p.skill[4].Damages += 10
				Green.Println("[BUFF] Vous avez gagné 10% de chance de coup critique")
				Green.Println("[BUFF] Vous avez gagné 20 points de vie")
			case 3:
				p.skill[0].Damages += 20
				p.maxHP += 20
				p.currentHp += 20
				Green.Println("[BUFF] Vous avez gagné 10 de dégats automatique")
				Green.Println("[BUFF] Vous avez gagné 20 points de vie")
			case 4:
				p.maxHP += 50
				p.currentHp += 50
				p.skill[0].Damages += 50
				p.skill[4].Damages += 20
				Green.Println("[BUFF] Vous avez gagné 20% de chance de coup critique")
				Green.Println("[BUFF] Vous avez gagné 50 points de vie")
				Green.Println("[BUFF] Vous avez gagné 50 de dégats d'attauqe")
			}

			p.forgeron()
		}
	case choice == len(outilsMarchand)+1:
		ClearConsole()
		p.Menu()
	default:
		ClearConsole()
		Red.Println("Veuillez saisir une donnée valide")
		p.forgeron()
	}
}

// accessInventory permet au joueur d'accéder à son inventaire.
func (p *Personnage) accessInventory() {

	// Affichage de l'inventaire
	Cyan.Println("Item            Quantité")
	Cyan.Println("---------------------------")
	for _, item := range p.inventory {
		// Utilisation de fmt.Printf pour aligner les colonnes
		fmt.Printf("%-20s %-7d\n", item.Name, item.Quantite)
	}
	Cyan.Println("---------------------------")
	fmt.Println("[1] Utiliser une potion")
	fmt.Println("[2] Quitter l'inventaire")
	choice, _ := Inputint()
	switch choice {

	// Le joueur prend une potion
	case 1:
		ClearConsole()
		p.takePotion()
		p.accessInventory()

	// Le joueur quitte l'inventaire
	case 2:
		ClearConsole()

	// Choix non proposé
	default:
		ClearConsole()
		Red.Println("Veuillez saisir une donnée valide")
		p.accessInventory()
	}
}

// showSkills affiche les abilités du personnage.
func (p *Personnage) showSkills() {
	// Affichage des sorts du personnage
	Cyan.Printf("%-20s%-10s%-60s\n", "Nom", "Dégâts", "Description")
	Cyan.Println("------------------------------------------------------------------")

	for _, spell := range p.skill {
		fmt.Printf("%-20s%-10d%-60s\n", spell.Name, spell.Damages, spell.Description)
	}
	Cyan.Println("-------------------------------------------------------------------------------------------")
}

// displayInfo affiche les informations du personnage.
func (p *Personnage) displayInfo() {
	niveaux := []string{"Tutoriel", "World War 3", "Château", "Jardin", "Ville sous-marine", "Starship", "Uranus"}
	Cyan.Println("--- ", p.nom, " ---")
	fmt.Println("Classe : ", p.classe)
	fmt.Println("Niveau : ", p.niveau, "("+niveaux[p.niveau]+")")
	fmt.Println("Ennemis battus : " + strconv.Itoa(p.ennemi) + " /3")
	fmt.Println("Points de vie : ", p.currentHp, "/", p.maxHP)
	Cyan.Println("--------------")
}

// Marchand permet au joueur d'interagir avec le marchand.
func (p *Personnage) Marchand() {

	// Configuration de la liste de vente du marchand ainsi que toutes les spécifités des items
	var itemMarchand = []Article{
		{"Potions", 10, "Potion pour récupèrer 40 points de vie", 0, 0, false},
		{"Guinzoo", 100, "+20% hp | +15 ad", 15, p.maxHP / 5, true},
		{"Masque grenouille", 1000, "+600% hp", 0, p.maxHP * 6, true},
		{"Avrilvert", 350, "+100% ad", p.skill[0].Damages, 0, true},
		{"Rook", 450, "+80% hp", 0, p.maxHP * 80 / 100, true},
		{"Anneau de gel", 500, "+60%hp | +20 ad ", 20, p.maxHP * 60 / 100, true},
	}

	// Affichage des items que vend le marchand
	for {
		Green.Println("Argent :  ", p.nbItem("Argent"))
		Cyan.Println("    Article             Prix                Description")
		Cyan.Println("-----------------------------------------------------")
		for i, item := range itemMarchand {
			if p.alreadyBuy(item.Name) && item.Unique == true {
				Cyan.Printf("%s - acheté !\n", item.Name)
			} else {
				fmt.Printf("[%d] %-20s %-7d %-s\n", i+1, item.Name, item.Price, item.Description)
			}
		}
		Cyan.Println("-----------------------------------------------------")
		fmt.Printf("[%d] Sortir \n", len(itemMarchand)+1)
		choice, _ := Inputint()

		// Le choix est un item
		if choice > 0 && choice <= len(itemMarchand) {
			// Selection de l'item
			selectedItem := itemMarchand[choice-1]

			if p.enoughMoney(selectedItem.Price) { // Vérification si le joueur a assez d'argent
				if !(p.alreadyBuy(selectedItem.Name)) || selectedItem.Unique == false { // Vérification si le joueur ne l'a pas déjà acheté si il est unique
					p.addInventory(selectedItem.Name, 1)                              // Ajout de l'item a l'inventaire
					p.inventory[p.findIndex("Argent")].Quantite -= selectedItem.Price // On retire l'argent de l'inventaire
					ClearConsole()
					p.skill[0].Damages += selectedItem.Ad // On ajoute les dégats que donne l'item
					if p.classe == "Chasseur" {           // On configure l'attaque rapide du chasseur
						p.skill[3].Damages = p.skill[0].Damages * 2
					}
					p.maxHP += selectedItem.Health     // On ajoute la vie que donne l'item sur les hp max
					p.currentHp += selectedItem.Health // On ajoute la vie que donne l'item sur les hp actuels
					Cyan.Printf("Vous avez acheté : %s pour %d coro !\n", selectedItem.Name, selectedItem.Price)
				} else { // Item déjà acheté
					ClearConsole()
					Red.Println("Vous avez déjà acheté cet item")
				}
			} else { // Pas assez d'argent pour l'acheter
				ClearConsole()
				Red.Println("Vous n'avez pas assez d'argent !")
			}

		} else if choice == len(itemMarchand)+1 { // Choix de sortie du marchand
			ClearConsole()
			p.Menu()
			break
		} else { // Choix non proposé
			ClearConsole()
			Red.Println("Veuillez saisir une donnée valide")
		}
	}
}
