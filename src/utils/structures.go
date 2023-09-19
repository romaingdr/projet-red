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
	nom       string
	classe    string
	niveau    int
	ennemi    int
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
