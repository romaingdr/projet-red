// FICHIER UTILISE POUR RECUPERER LA SAUVEGARDE DE LA PARTIE PRECEDENTE

package utils

import (
	"strconv"
)

// recupSauvegarde récupère les données du fichier 'database.json' et crée un personnage avec
func recupSauvegarde() {
	var p1 Personnage
	var items []Item
	var spells []Spell
	db = NewQuickDB("database.json")
	infosJoueurs := splitWords(db.Get("sauvegarde").(string))
	nomItems := splitWords(db.Get("sauvegardeItems").(string))
	nbItems := splitWords(db.Get("sauvegardeNb").(string))
	nbItemsInt := []int{}

	for i := 0; i < len(nbItems); i++ {
		entier, _ := strconv.Atoi(nbItems[i])
		nbItemsInt = append(nbItemsInt, entier)
	}

	nom := infosJoueurs[0]
	classe := infosJoueurs[1]
	currentHP, _ := strconv.Atoi(infosJoueurs[2])
	maxHP, _ := strconv.Atoi(infosJoueurs[3])
	niveau, _ := strconv.Atoi(infosJoueurs[4])
	ennemi, _ := strconv.Atoi(infosJoueurs[5])

	for i := 0; i < len(nomItems); i++ {
		items = append(items, Item{Name: nomItems[i], Quantite: nbItemsInt[i]})
	}

	switch classe {
	case "Titan":
		spells = []Spell{{"Auto", "Attaque automatique du titan", 10, -1, -1},
			{"La bulle", "Réduit de 65% la prochaine attaque et gagne 10% de chance de coup critique", 0, 5, 5},
			{"Frappe ultime", "Le titan inflige une violente attaque", 50, 3, 3},
			{"Dé titanesque", "Le titan a 67% de chance d'infliger 400% de dégats, sinon il perd 70 points de vies", 30, 100, 100},
			{"(%) Critical chance", "Chance d'obtenir un coup critique (inflige le double des dégats)", 15, -1, -1}}
	case "Chasseur":
		spells = []Spell{{"Auto", "Attaque automatique du chasseur", 20, -1, -1},
			{"Lame Sanglante", "Inflige un poison de 10 dégats par tour pendant 3 tours", 25, 3, 3},
			{"Maitrise du terrain", "Réduit de 50% la prochaine attaque | + 20% de dégats d'attaque automatique", 0, 100, 100},
			{"Attaque rapide", "Inflige 200% des dégats de l'attaque automatique", 40, 3, 3},
			{"(%) Critical chance", "Chance d'obtenir un coup critique (inflige le double des dégats)", 10, -1, -1}}
	case "Arcaniste":
		spells = []Spell{{"Auto", "Attaque automatique de l'arcaniste", 30, -1, -1},
			{"Trou noir", "Execute l'ennemi en dessous de 15% de points de vie", 40, 3, 3},
			{"Alteration de l'ame", "Vol les dégats infligés et les transforment en point de vie", 15, 100, 100},
			{"Foudre", "Améliore les dégâts de l'altération de l'âme de 50%", 70, 2, 2},
			{"(%) Critical chance", "Chance d'obtenir un coup critique (inflige le double des dégats)", 10, -1, -1}}
	}

	p1.Initialize(nom, classe, niveau, ennemi, currentHP, maxHP, items, spells)
	p1.Menu()
}
