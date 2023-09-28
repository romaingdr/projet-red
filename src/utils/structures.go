// FICHIER UTILISE POUR STOCKER LES STRUCTURES DU JEU

package utils

// Ennemy Structure des ennemis du jeu
type Ennemy struct {
	Name           string
	HpCurrent      int
	HpMax          int
	DamagesMin     int
	DamgesMax      int
	CriticalChance int
	IsBoss         bool
}

// Spell Structure des capacités du personnage principal
type Spell struct {
	Name        string
	Description string
	Damages     int
	StillUse    int
	MaxUse      int
}

// Personnage structure du personnage principal
type Personnage struct {
	nom       string
	classe    string
	niveau    int
	ennemi    int
	currentHp int
	maxHP     int
	inventory []Item
	skill     []Spell
}

// Item structure d'un item dans l'inventaire
type Item struct {
	Name     string
	Quantite int
}

// Article structure d'un article dans le marchand
type Article struct {
	Name        string
	Price       int
	Description string
	Ad          int
	Health      int
	Unique      bool
}

// EnnemyObjective structure de l'ennemi dans l'objectif
type EnnemyObjective struct {
	Name      string
	HpCurrent int
	HpMax     int
	Damages   int
}
