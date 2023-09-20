package utils

import "fmt"

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

func ScriptNiv1(p *Personnage) {
	if p.ennemi == 0 {
		fmt.Println("Premier Niveau , 1er ennemi")
		Input()
	} else if p.ennemi == 1 {
		fmt.Println("Premier Niveau , 2eme ennemi")
		Input()
	} else {
		fmt.Println("Dernier ennemi du niveau 1")
		Input()
	}
}

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
