package utils

import (
	"fmt"
)

// createCharacter initialise un nouveau personnage.
func (p *Personnage) CreateCharacter() {

	// Affichage des choix
	fmt.Println("---- Création du personnage ----")
	fmt.Println("[1] Créer un personnage")
	fmt.Println("[2] Personnage par défaut")
	fmt.Println("--------------------------------")
	choice, _ := Inputint()
	switch choice {

	// Création d'un personnage (nom + classe)
	case 1:
		hpMax := 0
		spells := []Spell{}
		nom := "0"
		ClearConsole()
		print("Nom de votre personnage >> ")
		nom = Input()
		for !(OnlyLetters(nom)) {
			ClearConsole()
			Red.Println("Veuillez n'utiliser que des lettres | 10 caractères maximum")
			print("Nom de votre personnage >> ")
			nom = Input()
		}
		nom = CapitalizeString(nom)
		ClearConsole()
		classe := ChooseClass()
		switch classe {
		case "Titan":
			hpMax = 180
			spells = []Spell{{"Auto", "Attaque automatique du titan", 10, -1, -1},
				{"La bulle", "Le titan s'enferme dans une bulle et reduit les dégâts subis", 0, 5, 5},
				{"Frappe ultime", "Le titan inflige une violente attaque", 50, 3, 3},
				{"Dé titanesque", "Le titan a 67% de chance d'infliger 400% de dégats, sinon il perd 70 points de vies", 40, 100, 100},
				{"(%) Critical chance", "inflige le double des dégats", 15, -1, -1}}
		case "Chasseur":
			hpMax = 135
			spells = []Spell{{"Auto", "Attaque automatique du chasseur", 20, -1, -1},
				{"Lame Sanglante", "Inflige un coup de lame empoisonnée", 25, 3, 3},
				{"Maitrise du terrain", "Le chasseur se concentre pour infliger une violente attaque", 0, 100, 100},
				{"Attaque rapide", "Inflige 200% des dégats de l'attaque automatique", 30, 3, 3},
				{"(%) Critical chance", "inflige le double des dégats", 10, -1, -1}}
		case "Arcaniste":
			hpMax = 100
			spells = []Spell{{"Auto", "Attaque automatique de l'arcaniste", 30, -1, -1},
				{"Trou noir", "Execute l'ennemi en dessous de 15% de points de vie", 40, 3, 3},
				{"Alteration de l'ame", "Vol de vie (150% des dégats infligés)", 15, 100, 100},
				{"Foudre", "La foudre s'abat sur l'ennemi et lui inflige des dégats", 70, 2, 2},
				{"(%) Critical chance", "inflige le double des dégats", 10, -1, -1}}
		}
		ClearConsole()
		p.Initialize(nom, classe, 2, hpMax, hpMax, []Item{{"Argent", 10000}, {"Potions", 3}}, spells)
		SpeedMsg("Bienvenue, "+nom+" ! \n", 60, "blue")
		fmt.Println()
		fmt.Print("Appuyez pour entrer dans la partie")
		Input()
		ClearConsole()

	// Personnage par défaut (nom: Romain + classe: Chasseur)
	case 2:
		spells := []Spell{{"Auto", "Attaque automatique du champion", 15, -1, -1},
			{"Lame Sanglante", "Inflige un coup de lame empoisonnée", 25, 3, 3},
			{"Maitrise du terrain", "Le chasseur se concentre pour infliger une violente attaque", 0, 2, 2},
			{"Attaque rapide", "Inflige une attaque automatique double", 30, 1, 1},
			{"(%) Critical chance", "inflige le double des dégats", 10, -1, -1}}

		p.Initialize("Romain", "Chasseur", 2, 125, 125, []Item{{"Argent", 10000}, {"Potions", 3}}, spells)
		ClearConsole()
		SpeedMsg("Bienvenue, Romain !\n", 60, "blue")
		fmt.Println()
		fmt.Print("Appuyez pour entrer dans la partie")
		Input()
		ClearConsole()

	// Choix non proposé
	default:
		ClearConsole()
		Red.Println("Veuillez saisir une donnée valide")
		p.CreateCharacter()
	}

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
