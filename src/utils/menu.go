package utils

import (
	"fmt"
	"strconv"
)

// Menu affiche le menu principal du jeu.
func (p *Personnage) Menu() {

	// Affichage des choix du menu
	fmt.Println("----- Menu -----")
	fmt.Println("[1] Personnage")
	fmt.Println("[2] Inventaire")
	fmt.Println("[3] Marchand")
	fmt.Println("[4] Abilités")
	if p.niveau == 1 {
		fmt.Println("[5] Combat - tutoriel")
	} else {
		if p.niveau < 8 {
			if p.ennemi == 2 {
				fmt.Print("[5] Combat")
				Red.Println(" - Boss")
			} else {
				fmt.Println("[5] Combat")
			}
		} else {
			fmt.Println("[5] fin ?")
		}
	}
	fmt.Println("[6] Quitter le jeu")
	fmt.Println("----------------")

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
		p.showSkills()
		p.Menu()

	// Combat tutoriel si niveau 1 sinon menu de combat
	case 5:
		ClearConsole()
		if p.niveau == 1 {
			p.battleTutorial()
		} else {
			p.battle()
		}

	// Ferme le jeu complétement
	case 6:
		ClearConsole()
		Red.Println("Fermeture du jeu...")

	// Choix non proposé
	default:
		ClearConsole()
		Red.Println("Veuillez saisir une donnée valide !")
		p.Menu()
	}
}

// accessInventory permet au joueur d'accéder à son inventaire.
func (p *Personnage) accessInventory() {

	// Affichage de l'inventaire
	fmt.Println("Item            Quantité")
	fmt.Println("---------------------------")
	for _, item := range p.inventory {
		// Utilisation de fmt.Printf pour aligner les colonnes
		fmt.Printf("%-20s %-7d\n", item.Name, item.Quantite)
	}
	fmt.Println("---------------------------")
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
	fmt.Println("Ennemis battus : " + strconv.Itoa(p.ennemi) + " /3")
	fmt.Println("Points de vie : ", p.currentHp, "/", p.maxHP)
	fmt.Println("--------------")
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
					Blue.Printf("Vous avez acheté : %s pour %d coro !\n", selectedItem.Name, selectedItem.Price)
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
			Blue.Println("Sortie du marchand")
			p.Menu()
			break
		} else { // Choix non proposé
			ClearConsole()
			Red.Println("Veuillez saisir une donnée valide")
		}
	}
}
