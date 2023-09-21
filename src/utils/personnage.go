// FICHIER UTILISE POUR LA CREATION DU PERSONNAGE ET L'ATTRIBUTION DE SES SPELLS

package utils

import (
	"fmt"
)

// createCharacter initialise un nouveau personnage.
func (p *Personnage) CreateCharacter() {
	ClearConsole()

	hpMax := 0
	spells := []Spell{}
	nom := "0"
	print("Nom de votre personnage ☛ ")
	nom = Input()
	for !(OnlyLetters(nom)) {
		ClearConsole()
		Red.Println("Veuillez n'utiliser que des lettres | 10 caractères maximum")
		print("Nom de votre personnage ☛ ")
		nom = Input()
	}
	nom = CapitalizeString(nom)
	ClearConsole()
	classe := ChooseClass()
	switch classe {
	case "Titan":
		hpMax = 180
		spells = []Spell{{"Auto", "Attaque automatique du titan", 10, -1, -1},
			{"La bulle", "Réduit de 65% la prochaine attaque et gagne 10% de chance de coup critique", 0, 5, 5},
			{"Frappe ultime", "Le titan inflige une violente attaque", 50, 3, 3},
			{"Dé titanesque", "Le titan a 67% de chance d'infliger 400% de dégats, sinon il perd 70 points de vies", 30, 100, 100},
			{"(%) Critical chance", "Chance d'obtenir un coup critique (inflige le double des dégats)", 15, -1, -1}}
	case "Chasseur":
		hpMax = 135
		spells = []Spell{{"Auto", "Attaque automatique du chasseur", 20, -1, -1},
			{"Lame Sanglante", "Inflige un poison de 10 dégats par tour pendant 3 tours", 25, 3, 3},
			{"Maitrise du terrain", "Réduit de 50% la prochaine attaque | + 20% de dégats d'attaque automatique", 0, 100, 100},
			{"Attaque rapide", "Inflige 200% des dégats de l'attaque automatique", 40, 3, 3},
			{"(%) Critical chance", "Chance d'obtenir un coup critique (inflige le double des dégats)", 10, -1, -1}}
	case "Arcaniste":
		hpMax = 100
		spells = []Spell{{"Auto", "Attaque automatique de l'arcaniste", 30, -1, -1},
			{"Trou noir", "Execute l'ennemi en dessous de 15% de points de vie", 40, 3, 3},
			{"Alteration de l'ame", "Vol les dégats infligés et les transforment en point de vie", 15, 100, 100},
			{"Foudre", "Améliore les dégâts de l'altération de l'âme de 50%", 70, 2, 2},
			{"(%) Critical chance", "Chance d'obtenir un coup critique (inflige le double des dégats)", 10, -1, -1}}
	}
	ClearConsole()
	p.Initialize(nom, classe, 7, hpMax, hpMax, []Item{{"Argent", 300}, {"Potions", 3}}, spells)
	SpeedMsg("Bienvenue, "+nom+" ! \n", 60, "blue")
	fmt.Println()
	fmt.Print("Appuyez pour entrer dans la partie")
	Input()
	ClearConsole()
	p.Menu()
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

// choixClasse affiche les classes disponibles et permet à l'utilisateur de choisir une classe.
func ChooseClass() string {
	ClearConsole()

	Green.Println("Nom du personnage validé !")
	Blue.Println("Choisissez votre classe : ")
	classes := []string{"Titan", "Arcaniste", "Chasseur"}
	println("")
	println("[1] Titan : «Une représentation de la force brute» | 180 Hp - 10 Ad")
	println("[2] Arcaniste : «Manipule les lois de l'univers» | 100 Hp - 30 Ad")
	println("[3] Chasseur : «N'apparait que dans l'ombre» | 125 Hp - 20 Ad")
	println("")

	for {
		choice, _ := Inputint()
		if choice > 0 && choice < 4 {
			return classes[choice-1]
		} else {
			ClearConsole()
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
