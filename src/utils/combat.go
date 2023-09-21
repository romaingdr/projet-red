// FICHIER UTILISE POUR LA GESTION DU COMBAT ET DES TOURS ENTRE L'ENNEMI ET LE JOUEUR

package utils

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// Variables gloables pour le jeu
var (
	damage_reduce = 0
	poison        = 0
	// Listes des ennemis rencontrés durant le jeu
	Monstres = [][]Ennemy{
		{
			{"Basic soldier 1", 200, 200, 20, 50, 25, false},
			{"Basic soldier 2", 200, 200, 20, 50, 25, false},
			{"Red soldier", 300, 300, 20, 50, 25, true},
		},

		{
			{"Bat 1", 350, 350, 20, 50, 25, false},
			{"Bat 2", 350, 350, 20, 50, 25, false},
			{"Donuts man", 550, 550, 20, 50, 40, true},
		},

		{
			{"Mosquito", 200, 200, 20, 50, 80, false},
			{"Snail", 700, 700, 20, 50, 15, false},
			{"Wasp", 600, 600, 20, 50, 60, true},
		},

		{
			{"Eveel 1", 750, 750, 20, 50, 60, false},
			{"Eveel 2", 750, 750, 20, 50, 60, false},
			{"Mermaid duo", 900, 900, 20, 50, 70, true},
		},

		{
			{"The key master 1", 1300, 1300, 20, 50, 80, false},
			{"The key master 2", 1300, 1300, 20, 50, 80, false},
			{"The general", 1700, 1700, 20, 50, 80, true},
		},

		{
			{"Alien 1", 700, 700, 60, 120, 70, false},
			{"Alien 2", 700, 700, 60, 120, 70, false},
			{"Fallen garden", 2500, 2500, 60, 120, 100, true},
		},
	}
)

// battle est la fonction de combat
func (p *Personnage) battle() {
	ClearConsole()
	Script(p)
	ClearConsole()
	if p.niveau == 7 { // LE JEU EST TERMINE
		SpeedMsg("Félicitation "+p.nom+"\n", 30, "blue")
		time.Sleep(2 * time.Second)
		SpeedMsg("Vous avez battu tout les ennemis\n", 20, "blue")
		time.Sleep(1 * time.Second)
		SpeedMsg("Merci d'avoir joué à notre jeu !\n", 20, "blue")
		time.Sleep(1 * time.Second)
		SpeedMsg("En esperant vous revoir bientôt !", 30, "blue")
		Input()
		os.Exit(0)
	}
	// Sauvegarde des spells qui changent

	actualAuto := p.skill[0].Damages
	actualCrit := p.skill[4].Damages
	spell2 := p.skill[2].Damages
	spell3 := p.skill[3].Damages

	// Configuration de l'ennemi
	ennemi1 := Monstres[p.niveau-1][p.ennemi]

	// Affichage du duel et des statistiques de l'ennemi
	SpeedMsg(p.nom+" VS "+ennemi1.Name+"\n", 20, "red")
	fmt.Println("-----------------------------")
	SpeedMsg("Statistiques de l'ennemi : \n", 20, "default")
	fmt.Println("Points de vie : " + strconv.Itoa(ennemi1.HpMax))
	fmt.Println("Dégats : " + strconv.Itoa(ennemi1.DamagesMin) + " - " + strconv.Itoa(ennemi1.DamgesMax))
	fmt.Println("Chance de coup critique : " + strconv.Itoa(ennemi1.CriticalChance))
	fmt.Println("-----------------------------")
	SpeedMsg("Bonne chance combattant !", 20, "green")
	fmt.Println()
	fmt.Print("Appuyez sur entrée pour continuer")
	Input()
	ClearConsole()

	// Boucle infini jusqu'à qu'un des deux combattants soit mort
	for !(p.isDead(&ennemi1)) {
		// Tour de l'ennemi
		p.ennemyRound(&ennemi1)
		ClearConsole()

		if !(p.isDead(&ennemi1)) {
			// Tour du joueur si il n'est pas mort
			p.playerRound(&ennemi1)
			ClearConsole()
		}
	}

	// Victoire du joueur
	if ennemi1.HpCurrent <= 0 {

		// Décompte jusqu'à 0 des hp de l'ennemi
		hpEnnemy := strconv.Itoa(ennemi1.HpMax)

		for i := ennemi1.HpMax; i >= 0; i-- {
			Red.Print("\r", ennemi1.Name+" - "+fmt.Sprintf("%3d", i)+"/"+hpEnnemy)
			time.Sleep(time.Millisecond * 10)
		}

		// Affichage du résultat
		fmt.Println()
		fmt.Println()
		Green.Println("VOUS AVEZ GAGNÉ !")
		fmt.Println()
		SpeedMsg("Récompenses : \n", 30, "default")
		time.Sleep(1 * time.Second)
		SpeedMsg("+ 300 coro\n", 30, "blue")
		time.Sleep(1 * time.Second)
		SpeedMsg("+ 3 Potions\n", 30, "blue")
		time.Sleep(1 * time.Second)
		SpeedMsg("+ 1 Niveau\n", 30, "blue")

		// Attribution des récompenses
		p.addInventory("Potions", 3)
		p.addInventory("Argent", 300)

		if ennemi1.IsBoss {
			p.niveau += 1
			p.ennemi = 0
		} else {
			p.ennemi += 1
		}

		fmt.Println()
		fmt.Print("Appuyez sur entrée pour continuer")
		Input()
		ClearConsole()

		if p.ennemi == 0 {
			Blue.Println("Vous avez atteint le niveau " + strconv.Itoa(p.niveau))
		}

		// Re attribution des utilisations de chaque spell
		for i := 1; i < 4; i++ {
			p.skill[i].StillUse = p.skill[i].MaxUse
		}
		// Ré-attribution des dégats a l'auto
		p.skill[0].Damages = actualAuto
		p.skill[2].Damages = spell2
		p.skill[3].Damages = spell3
		p.skill[4].Damages = actualCrit
		ClearConsole()
		p.Menu()
	} else {
		// Défaite du joueur
		// Décompte jusqu'à 0 des hp du joueur
		hpPlayer := strconv.Itoa(p.maxHP)
		for i := p.maxHP; i >= 0; i-- {
			Green.Print("\r", p.nom+" - "+fmt.Sprintf("%3d", i)+"/"+hpPlayer)
			time.Sleep(time.Millisecond * 10)
		}

		// Affichage du résultat du duel
		fmt.Println()
		fmt.Println()
		Red.Println("VOUS AVEZ PERDU !")
		fmt.Println()
		SpeedMsg("Préparez vous et revenez plus fort ! \n", 30, "default")
		fmt.Println()
		fmt.Print("Appuyez sur entrée pour continuer")
		Input()

		// Ré-attribution des utilisations de chaque spell ainsi que remise à 100% des points de vie du joueur
		p.currentHp = p.maxHP
		for i := 1; i < 4; i++ {
			p.skill[i].StillUse = p.skill[i].MaxUse
		}
		// Réèattribution des dégats a l'auto
		p.skill[0].Damages = actualAuto
		p.skill[2].Damages = spell2
		p.skill[3].Damages = spell3
		p.skill[4].Damages = actualCrit
		ClearConsole()
		p.Menu()
	}
}

// abilitiesBattle affiche le menu des spells dans un combat
func (p *Personnage) abilitiesBattle(e *Ennemy) {
	for {

		// Affichage des spells
		Red.Println(e.Name + " - " + strconv.Itoa(e.HpCurrent) + "/" + strconv.Itoa(e.HpMax))
		Green.Println(p.nom + " - " + strconv.Itoa(p.currentHp) + "/" + strconv.Itoa(p.maxHP))
		fmt.Println("----- A votre tour -----")

		for i := 1; i < 4; i++ {
			skill := p.skill[i]
			fmt.Printf("[%d] %-20s %-10d %d/%d\n", i, skill.Name, skill.Damages, skill.StillUse, skill.MaxUse)
		}

		fmt.Println("------------------------")
		fmt.Println("[4] Sortir")

		choice, _ := Inputint()
		var degats int
		var nom string

		// Calcul si le coup va être un crit (en fonction du pourcentage de chance de crit ==> p.skill[4].Damages)
		rand.Seed(time.Now().UnixNano())
		crit := rand.Intn(100) + 1
		critBool := crit <= p.skill[4].Damages

		switch choice {
		case 1, 2, 3:
			skill := p.skill[choice]
			if skill.StillUse > 0 {

				// La compétence est utilisée
				degats = skill.Damages
				nom = skill.Name
				p.skill[choice].StillUse -= 1

			} else {
				// La compétence a été trop utilisée
				ClearConsole()
				Red.Println("Vous ne pouvez plus utiliser cette compétence !")
				continue
			}
		case 4:

			// On quitte le menu des spells
			ClearConsole()
			p.playerRound(e)
			return

		default:
			// Choix pas proposé
			ClearConsole()
			Red.Println("Veuillez saisir une donnée valide !")
			continue // Donc on repart la boucle a 0 jusqu'à qu'il mette un bon numéro
		}

		// Affichage une fois que la compétence est utilisée
		ClearConsole()

		Red.Println(e.Name + " - " + strconv.Itoa(e.HpCurrent) + "/" + strconv.Itoa(e.HpMax))
		Green.Println(p.nom + " - " + strconv.Itoa(p.currentHp) + "/" + strconv.Itoa(p.maxHP))
		fmt.Println("----- A votre tour -----")
		SpeedMsg(p.nom+" attaque avec "+nom+"...\n", 20, "default")

		// Petite attente avant le résultat
		time.Sleep(2 * time.Second)

		if critBool {
			degats *= 2
			SpeedMsg("[COUP CRITIQUE] Vous infligez "+strconv.Itoa(degats)+" de dégâts à l'ennemi\n", 20, "green")
		} else {
			SpeedMsg("Vous infligez "+strconv.Itoa(degats)+" de dégâts à l'ennemi\n", 20, "green")
		}

		// On enlève les hp à l'ennemi
		e.HpCurrent -= degats

		// Effets post-spells
		if choice == 1 { // Spell 1

			if p.classe == "Arcaniste" { // Execution en dessous de 10%
				if e.HpCurrent <= e.HpMax/10 {
					e.HpCurrent = 0
					Green.Println("[EXECUTION] Votre trou noir a éxecuté l'ennemi")
				}

			} else if p.classe == "Chasseur" { // Poison de 10 dégats
				poison = 3

			} else if p.classe == "Titan" { // Réduction des dégats
				damage_reduce = 65
				Green.Println("[BULLE] Vous obtenez 65% de réductions des dégats pour le prochain tour")
				p.skill[4].Damages += 10
				Green.Println("[BULLE] Vous obtenez 10% de chance de coup critique")
			}

		} else if choice == 2 {

			if p.classe == "Arcaniste" { // Vol de vie (100% des dégats)
				p.currentHp += degats
				Green.Println("[VOL DE VIE] Vous récuperez " + strconv.Itoa(degats) + " points de vie")

			} else if p.classe == "Chasseur" { // Damage reduce  + 20% auto
				damage_reduce = 50
				Green.Println("[MAITRISE DU TERRAIN] Vous obtenez 50% de réduction des dégats pour le prochain tour")
				p.skill[0].Damages += p.skill[0].Damages / 5
				Green.Println("[MAITRISE DU TERRAIN] Votre attaque automatique inflige 20% de dégats supplémentaires")
			}

		} else if choice == 3 { // Amélioration des dégats du spell 2
			if p.classe == "Arcaniste" {
				p.skill[2].Damages += p.skill[2].Damages / 2
				Green.Println("[FOUDRE] Votre altération de l'âme inflige 50% de dégats supplémentaires")

			} else if p.classe == "Titan" { // Dé titanesque (67% --> +400% damages | 33% --> -70hp)
				rand.Seed(time.Now().UnixNano())
				aleatoire := rand.Intn(3) + 1
				// On inflige
				if aleatoire <= 2 {
					degats *= 4
					Green.Println("[DE TITANESQUE] Vous infligez " + strconv.Itoa(degats) + " dégats supplémentaires")
					e.HpCurrent -= degats
				} else {
					Red.Println("[DE TITANESQUE] Vous perdez 70 points de vie")
					p.currentHp -= 70
				}
			}
		}
		if poison > 0 {
			e.HpCurrent -= 10
			Green.Println("[POISON] Vous infligez 10 dégats supplémentaires")
			poison--
		}
		fmt.Println("------------------------")
		fmt.Print("Appuyez sur entrée pour continuer")
		Input()
		return
	}
}

// playerRound configure le round côté joueur
func (p *Personnage) playerRound(e *Ennemy) {
	// Confiuration de l'attaque rapide du chasseur
	if p.classe == "Chasseur" {
		p.skill[3].Damages = p.skill[0].Damages * 2
	}

	// Affichage du menu lors du tour du joueur
	Red.Println(e.Name + " - " + strconv.Itoa(e.HpCurrent) + "/" + strconv.Itoa(e.HpMax))
	Green.Println(p.nom + " - " + strconv.Itoa(p.currentHp) + "/" + strconv.Itoa(p.maxHP))
	fmt.Println("----- A votre tour -----")
	fmt.Println("[1] Attaque auto")
	fmt.Println("[2] Abilités")

	fmt.Println("------------------------")
	choice, _ := Inputint()

	switch choice {

	// Envoie d'une attaque auto
	case 1:

		ClearConsole()

		// Degats de l'auto
		degats := p.skill[0].Damages

		// Calcul si c'est un coup critique avec le % (% ==> p.skill[4].Damges)
		rand.Seed(time.Now().UnixNano())
		crit := rand.Intn(100) + 1
		critBool := crit <= p.skill[4].Damages

		// Affichage de l'envoi de l'auto
		Red.Println(e.Name + " - " + strconv.Itoa(e.HpCurrent) + "/" + strconv.Itoa(e.HpMax))
		Green.Println(p.nom + " - " + strconv.Itoa(p.currentHp) + "/" + strconv.Itoa(p.maxHP))
		fmt.Println("----- A votre tour -----")
		SpeedMsg(p.nom+" utilise une attaque automatique\n", 20, "default")
		if critBool {
			degats *= 2
			SpeedMsg("[COUP CRITIQUE] Vous infligez "+strconv.Itoa(degats)+" de dégats à l'ennemi\n", 20, "green")
			e.HpCurrent -= degats
		} else {
			SpeedMsg("Vous infligez "+strconv.Itoa(degats)+" de dégats à l'ennemi\n", 20, "green")
			e.HpCurrent -= degats
		}
		if poison > 0 {
			e.HpCurrent -= 10
			Green.Println("[POISON] Vous infligez 10 dégats supplémentaires")
			poison--
		}
		fmt.Println("------------------------")
		fmt.Println()
		fmt.Print("Appuyez sur entrée pour continuer")
		Input()

	// Affichage des spells
	case 2:
		ClearConsole()
		p.abilitiesBattle(e)

	// Choix non proposé
	default:
		ClearConsole()
		Red.Println("Veuillez saisir une donnée valide")
		p.playerRound(e)
	}

}

// ennemyRound configure le round côté ennemi-ordinateur
func (p *Personnage) ennemyRound(e *Ennemy) {
	rand.Seed(time.Now().UnixNano())
	Red.Println(e.Name + " - " + strconv.Itoa(e.HpCurrent) + "/" + strconv.Itoa(e.HpMax))
	Green.Println(p.nom + " - " + strconv.Itoa(p.currentHp) + "/" + strconv.Itoa(p.maxHP))
	fmt.Println("----- Tour de " + e.Name + " -----")
	SpeedMsg(e.Name+" attaque...\n", 20, "default")
	degats := rand.Intn(e.DamgesMax-e.DamagesMin+1) + e.DamagesMin
	critic := rand.Intn(100) + 1
	if critic <= e.CriticalChance {
		degats *= 2
		critic = 1
	}
	time.Sleep(2 * time.Second)
	degats = int(float64(degats) * (1 - float64(damage_reduce)/100))

	p.currentHp -= degats
	if critic == 1 {
		Red.Println("[COUP CRITIQUE] Vous avez reçu " + strconv.Itoa(degats) + " dégats")
	} else {
		if damage_reduce > 0 {
			Green.Println("[REDUCTION DES DEGATS] Les dégats ont été reduits de " + strconv.Itoa(damage_reduce) + "%")
			Red.Println("Vous avez reçu " + strconv.Itoa(degats) + " dégats")
		} else {
			Red.Println("Vous avez reçu " + strconv.Itoa(degats) + " dégats")
		}

	}
	fmt.Println("---------------------------------")
	fmt.Println()
	fmt.Print("Appuyez sur entrée pour continuer")
	Input()
	damage_reduce = 0
}

// isDead vérifie la mort d'au moins un des deux combattant d'un duel
func (p *Personnage) isDead(e *Ennemy) bool {
	if p.currentHp <= 0 || e.HpCurrent <= 0 {
		return true
	} else {
		return false
	}
}
