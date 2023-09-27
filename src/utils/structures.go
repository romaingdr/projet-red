// FICHIER UTILISE POUR STOCKER LES STRUCTURES DU JEU

package utils

// Ennemy Structure
type Ennemy struct {
	Name           string
	HpCurrent      int
	HpMax          int
	DamagesMin     int
	DamgesMax      int
	CriticalChance int
	IsBoss         bool
}

// Spell Structure
type Spell struct {
	Name        string
	Description string
	Damages     int
	StillUse    int
	MaxUse      int
}

// Personnage structure
type Personnage struct {
	nom       string `json:"nom"`
	classe    string `json:"classe"`
	niveau    int    `json:"niveau"`
	ennemi    int    `json:"ennemi"`
	currentHp int    `json:"currentHp"`
	maxHP     int    `json:"maxHP"`
	inventory []Item `json:"inventory"`
	skill     []Spell
}

// Item structure
type Item struct {
	Name     string `json:"name"`
	Quantite int    `json:"quantite"`
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

type EnnemyObjective struct {
	Name      string
	HpCurrent int
	HpMax     int
	Damages   int
}
