package utils

import "fmt"

// Script affiche le texte du lore en fonction du niveau
func Script(p *Personnage) {
	if p.niveau == 1 {
		ScriptNiv1(p)
	} else if p.niveau == 2 {
		ScriptNiv2(p)
	} else if p.niveau == 3 {
		ScriptNiv3(p)
	} else if p.niveau == 4 {
		ScriptNiv4(p)
	} else if p.niveau == 5 {
		ScriptNiv5(p)
	} else if p.niveau == 6 {
		ScriptNiv6(p)
	}
}

// ScriptNiv1 affiche le texte du lore sur tout le niveau 1
func ScriptNiv1(p *Personnage) {
	if p.ennemi == 0 {
		SpeedMsg("Romain", 20, "green")
		Input()
	} else if p.ennemi == 1 {
		fmt.Println("Premier Niveau , 2eme ennemi")
		Input()
	} else {
		fmt.Println("Dernier ennemi du niveau 1")
		Input()
	}
}

// ScriptNiv2 affiche le texte du lore sur tout le niveau 1
func ScriptNiv2(p *Personnage) {
	if p.ennemi == 0 {
		fmt.Println("Deuxieme Niveau , 1er ennemi")
		Input()
	} else if p.ennemi == 1 {
		fmt.Println("Deuxieme Niveau , 2eme ennemi")
		Input()
	} else {
		fmt.Println("Dernier ennemi du niveau 2")
		Input()
	}
}

// ScriptNiv3 affiche le texte du lore sur tout le niveau 1
func ScriptNiv3(p *Personnage) {
	if p.ennemi == 0 {
		fmt.Println("Troisieme Niveau , 1er ennemi")
		Input()
	} else if p.ennemi == 1 {
		fmt.Println("Troisieme Niveau , 2eme ennemi")
		Input()
	} else {
		fmt.Println("Dernier ennemi du niveau 3")
		Input()
	}
}

// ScriptNiv4 affiche le texte du lore sur tout le niveau 1
func ScriptNiv4(p *Personnage) {
	if p.ennemi == 0 {
		fmt.Println("Quatrieme Niveau , 1er ennemi")
		Input()
	} else if p.ennemi == 1 {
		fmt.Println("Quatrieme Niveau , 2eme ennemi")
		Input()
	} else {
		fmt.Println("Dernier ennemi du niveau 4")
		Input()
	}
}

// ScriptNiv5 affiche le texte du lore sur tout le niveau 1
func ScriptNiv5(p *Personnage) {
	if p.ennemi == 0 {
		fmt.Println("Cinquieme Niveau , 1er ennemi")
		Input()
	} else if p.ennemi == 1 {
		fmt.Println("Cinquieme Niveau , 2eme ennemi")
		Input()
	} else {
		fmt.Println("Dernier ennemi du niveau 5")
		Input()
	}
}

// ScriptNiv6 affiche le texte du lore sur tout le niveau 1
func ScriptNiv6(p *Personnage) {
	if p.ennemi == 0 {
		fmt.Println("Sixieme Niveau , 1er ennemi")
		Input()
	} else if p.ennemi == 1 {
		fmt.Println("Sixieme Niveau , 2eme ennemi")
		Input()
	} else {
		fmt.Println("Dernier ennemi du niveau 6")
		Input()
	}
}
