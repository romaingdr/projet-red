package utils

import "fmt"

// battleTutorial est le tutoriel de combat.
func (p *Personnage) battleTutorial() {

	// Affichage ligne par ligne du tutoriel
	ClearConsole()
	SpeedMsg("Bienvenue dans le tutoriel de combat !", 30, "blue")
	Input()
	ClearConsole()
	Red.Print("Ennemi 1 - 100 / 100")
	SpeedMsg(" <-- Ici sont affichés les points de vie de l'ennemi", 30, "white")
	Input()
	Green.Print("Vous - 100 / 100")
	SpeedMsg(" <-- Et ici les vôtres", 20, "white")
	Input()
	ClearConsole()
	SpeedMsg("Le combat se joue en tour par tour", 20, "blue")
	fmt.Println("")
	SpeedMsg("A chaque fois que vous jouez, plusieurs options s'offrent à vous : ", 20, "blue")
	fmt.Println("")

	// Appel de BattleMenuTutoriel dans le package utils pour afficher le menu du joueur pendant le tutoriel
	BattleMenuTutorial()

	// Continue d'afficher chaque ligne du tutoriel
	Input()
	ClearConsole()
	SpeedMsg("Lors de chaque attaque, vous verrez le nombre de dégats infligés : ", 20, "blue")
	fmt.Println()
	Green.Print("Vous avez infligé 20 dégats à Ennemi 1")
	Input()
	ClearConsole()
	SpeedMsg("Mais vous pouvez également en recevoir : ", 20, "blue")
	fmt.Println()
	Red.Print("Vous avez reçu 50 dégats (coup critique) de Ennemi 1 !")
	Input()
	ClearConsole()
	SpeedMsg("Lors de votre tour, vous pourrez également utiliser vos abilités : ", 20, "blue")
	fmt.Println("")

	// Choix du spell par le biais de AbilitiesTutorial qui affiche les spells et demande d'en choisir un
	spell := AbilitiesTutorial()

	// Continue d'afficher ligne par ligne le tutoriel
	SpeedMsg(spell+" à infligé 50 dégats à Ennemi 1", 20, "green")
	Input()
	ClearConsole()
	SpeedMsg("Félicitation, vous êtes prêt pour votre premier combat ! Bonne chance", 20, "blue")
	Input()
	ClearConsole()

	// Passage niveau 2 à la fin du tutoriel
	p.niveau = 2
	Blue.Println("Vous avez atteint le niveau 2 !")

	// Retour au menu
	p.Menu()
}

// abilitiesTutorial affiche les abilités disponibles dans le tutoriel de combat et permet à l'utilisateur d'en choisir une.
func AbilitiesTutorial() string {

	fmt.Println("---- Abilités ----")
	fmt.Println("[1] Coup de poing")
	fmt.Println("[2] Frénésie sanguinaire")
	fmt.Println("[3] Lame démoniaque")
	fmt.Println("------------------")
	choice, _ := Inputint()
	switch choice {
	case 1:
		return "Coup de poing"
	case 2:
		return "Frénésie sanguinaire"
	case 3:
		return "Lame démoniaque"
	default:
		ClearConsole()
		Red.Println("Veuillez choisir une option valide")
		return AbilitiesTutorial()
	}
}

// battleMenuTutorial affiche le menu du tutoriel de combat.
func BattleMenuTutorial() {
	fmt.Println("----- A votre tour -----")
	fmt.Print("[1] Attaque auto")
	SpeedMsg("<-- Ceci vous permet d'attaquer l'adversaire avec votre compétence basique", 20, "white")
	Input()
	fmt.Print("[2] Abilités")
	SpeedMsg("<-- Ceci vous permet d'utiliser une abilité sur l'adversaire", 20, "white")
	Input()
	fmt.Print("[3] Inventaire")
	SpeedMsg("<-- Ceci vous permet de consulter votre inventaire pendant le combat", 20, "white")
	fmt.Println("")
	fmt.Println("------------------------")
}
