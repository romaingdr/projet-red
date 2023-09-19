package main

// Import nécessaires au fonctionnement du jeu
import (
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"src/utils"
	"strconv"
	"time"
)

// Fonction main
func main() {
	utils.ClearConsole()
	var p1 Personnage
	p1.createCharacter()
	p1.Menu()
}

// Variables gloables pour le jeu
var (
	damage_reduce = 0

	// Couleurs
	Red   = color.New(color.FgRed)
	Blue  = color.New(color.FgBlue)
	Green = color.New(color.FgGreen)

	// Listes des ennemis rencontrés durant le jeu
	Monstres = [][]Ennemy{
		{
			{"Red soldier", 10, 200, 20, 50, 20, false},
			{"Red soldier", 10, 200, 20, 50, 20, false},
			{"Red soldier", 10, 200, 20, 50, 20, true}}}
)

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

// createCharacter initialise un nouveau personnage.
func (p *Personnage) createCharacter() {

	// Affichage des choix
	fmt.Println("---- Création du personnage ----")
	fmt.Println("[1] Créer un personnage")
	fmt.Println("[2] Personnage par défaut")
	fmt.Println("--------------------------------")
	choice, _ := utils.Inputint()
	switch choice {

	// Création d'un personnage (nom + classe)
	case 1:
		hpMax := 0
		spells := []Spell{}
		nom := "0"
		utils.ClearConsole()
		print("Nom de votre personnage >> ")
		nom = utils.Input()
		for !(utils.OnlyLetters(nom)) {
			utils.ClearConsole()
			Red.Println("Veuillez n'utiliser que des lettres | 10 caractères maximum")
			print("Nom de votre personnage >> ")
			nom = utils.Input()
		}
		nom = utils.CapitalizeString(nom)
		utils.ClearConsole()
		classe := utils.ChooseClass()
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
		utils.ClearConsole()
		p.Initialize(nom, classe, 2, hpMax, hpMax, []Item{{"Argent", 10000}, {"Potions", 3}}, spells)
		utils.SpeedMsg("Bienvenue, "+nom+" ! \n", 60, "blue")
		fmt.Println()
		fmt.Print("Appuyez pour entrer dans la partie")
		utils.Input()
		utils.ClearConsole()

	// Personnage par défaut (nom: Romain + classe: Chasseur)
	case 2:
		spells := []Spell{{"Auto", "Attaque automatique du champion", 15, -1, -1},
			{"Lame Sanglante", "Inflige un coup de lame empoisonnée", 25, 3, 3},
			{"Maitrise du terrain", "Le chasseur se concentre pour infliger une violente attaque", 0, 2, 2},
			{"Attaque rapide", "Inflige une attaque automatique double", 30, 1, 1},
			{"(%) Critical chance", "inflige le double des dégats", 10, -1, -1}}

		p.Initialize("Romain", "Chasseur", 2, 125, 125, []Item{{"Argent", 10000}, {"Potions", 3}}, spells)
		utils.ClearConsole()
		utils.SpeedMsg("Bienvenue, Romain !\n", 60, "blue")
		fmt.Println()
		fmt.Print("Appuyez pour entrer dans la partie")
		utils.Input()
		utils.ClearConsole()

	// Choix non proposé
	default:
		utils.ClearConsole()
		Red.Println("Veuillez saisir une donnée valide")
		p.createCharacter()
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

		choice, _ := utils.Inputint()
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
				utils.ClearConsole()
				Red.Println("Vous ne pouvez plus utiliser cette compétence !")
				continue
			}
		case 4:

			// On quitte le menu des spells
			utils.ClearConsole()
			p.playerRound(e)
			return

		default:
			// Choix pas proposé
			utils.ClearConsole()
			Red.Println("Veuillez saisir une donnée valide !")
			continue // Donc on repart la boucle a 0 jusqu'à qu'il mette un bon numéro
		}

		// Affichage une fois que la compétence est utilisée
		utils.ClearConsole()

		Red.Println(e.Name + " - " + strconv.Itoa(e.HpCurrent) + "/" + strconv.Itoa(e.HpMax))
		Green.Println(p.nom + " - " + strconv.Itoa(p.currentHp) + "/" + strconv.Itoa(p.maxHP))
		fmt.Println("----- A votre tour -----")
		utils.SpeedMsg(p.nom+" attaque avec "+nom+"...\n", 20, "default")

		// Petite attente avant le résultat
		time.Sleep(2 * time.Second)

		if critBool {
			degats *= 2
			utils.SpeedMsg("[COUP CRITIQUE] Vous infligez "+strconv.Itoa(degats)+" de dégâts à l'ennemi\n", 20, "green")
		} else {
			utils.SpeedMsg("Vous infligez "+strconv.Itoa(degats)+" de dégâts à l'ennemi\n", 20, "green")
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
				e.HpCurrent -= 10
				Green.Println("[POISON] Vous infligez 10 dégats supplémentaires")

			} else if p.classe == "Titan" { // Réduction des dégats
				damage_reduce = 65
				Green.Println("[BULLE] Vous obtenez 65% de réductions des dégats pour le prochain tour")
			}

		} else if choice == 2 {

			if p.classe == "Arcaniste" { // Vol de vie (100% des dégats)
				p.currentHp += degats
				Green.Println("[VOL DE VIE] Vous récuperez " + strconv.Itoa(degats) + " points de vie")

			} else if p.classe == "Chasseur" { // Damage reduce (pas complété) + Auto stack
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
		fmt.Println("------------------------")
		fmt.Print("Appuyez sur entrée pour continuer")
		utils.Input()
		return
	}
}

// playerRound configure le round côté joueur
func (p *Personnage) playerRound(e *Ennemy) {

	// Affichage du menu lors du tour du joueur
	Red.Println(e.Name + " - " + strconv.Itoa(e.HpCurrent) + "/" + strconv.Itoa(e.HpMax))
	Green.Println(p.nom + " - " + strconv.Itoa(p.currentHp) + "/" + strconv.Itoa(p.maxHP))
	fmt.Println("----- A votre tour -----")
	fmt.Println("[1] Attaque auto")
	fmt.Println("[2] Abilités")
	fmt.Println("[3] Inventaire")
	fmt.Println("------------------------")
	choice, _ := utils.Inputint()

	switch choice {

	// Envoie d'une attaque auto
	case 1:

		utils.ClearConsole()

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
		utils.SpeedMsg(p.nom+" utilise une attaque automatique\n", 20, "default")
		if critBool {
			degats *= 2
			utils.SpeedMsg("[COUP CRITIQUE] Vous infligez "+strconv.Itoa(degats)+" de dégats à l'ennemi\n", 20, "green")
			e.HpCurrent -= degats
		} else {
			utils.SpeedMsg("Vous infligez "+strconv.Itoa(degats)+" de dégats à l'ennemi\n", 20, "green")
			e.HpCurrent -= degats
		}
		fmt.Println("------------------------")
		fmt.Println()
		fmt.Print("Appuyez sur entrée pour continuer")
		utils.Input()

	// Affichage des spells
	case 2:
		utils.ClearConsole()
		p.abilitiesBattle(e)

	// Affichage de l'inventaire
	case 3:
		utils.ClearConsole()
		fmt.Println("Inventaire")
		p.playerRound(e)

	// Choix non proposé
	default:
		utils.ClearConsole()
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
	utils.SpeedMsg(e.Name+" attaque...\n", 20, "default")
	degats := rand.Intn(e.DamgesMax-e.DamagesMin+1) + e.DamagesMin
	critic := rand.Intn(100) + 1
	if critic <= e.CriticalChance {
		degats *= 2
	}
	time.Sleep(2 * time.Second)
	fmt.Println("ancien degats : ", degats)
	new_degats := degats * (1 - (damage_reduce / 100))
	fmt.Println("nouveaux degats : ", new_degats)
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
	utils.Input()
	damage_reduce = 0
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

// Menu affiche le menu principal du jeu.
func (p *Personnage) Menu() {

	// Affichage des choix du menu
	fmt.Println("----- Menu -----")
	fmt.Println("[1] Personnage")
	fmt.Println("[2] Inventaire")
	fmt.Println("[3] Marchand")
	fmt.Println("[4] Abilités")
	if p.niveau == 1 {
		fmt.Println("[5] Combat - tutoriel")
	} else {
		fmt.Println("[5] Combat")
	}
	fmt.Println("[6] Quitter le jeu")
	fmt.Println("----------------")

	choice, _ := utils.Inputint()

	switch choice {

	// Affichage des infos du joueur
	case 1:
		utils.ClearConsole()
		p.displayInfo()
		p.Menu()
		break

	// Affichage de l'inventaire
	case 2:
		utils.ClearConsole()
		p.accessInventory()
		p.Menu()
		break

	// Affichage du marchand
	case 3:
		utils.ClearConsole()
		p.Marchand()

	// Affichage des abilités
	case 4:
		utils.ClearConsole()
		p.showSkills()
		p.Menu()

	// Combat tutoriel si niveau 1 sinon menu de combat
	case 5:
		utils.ClearConsole()
		if p.niveau == 1 {
			p.battleTutorial()
		} else {
			p.battle()
		}

	// Ferme le jeu complétement
	case 6:
		utils.ClearConsole()
		Red.Println("Fermeture du jeu...")

	// Choix non proposé
	default:
		utils.ClearConsole()
		Red.Println("Veuillez saisir une donnée valide !")
		p.Menu()
	}
}

// battleTutorial est le tutoriel de combat.
func (p *Personnage) battleTutorial() {

	// Affichage ligne par ligne du tutoriel
	utils.ClearConsole()
	utils.SpeedMsg("Bienvenue dans le tutoriel de combat !", 30, "blue")
	utils.Input()
	utils.ClearConsole()
	Red.Print("Ennemi 1 - 100 / 100")
	utils.SpeedMsg(" <-- Ici sont affichés les points de vie de l'ennemi", 30, "white")
	utils.Input()
	Green.Print("Vous - 100 / 100")
	utils.SpeedMsg(" <-- Et ici les vôtres", 20, "white")
	utils.Input()
	utils.ClearConsole()
	utils.SpeedMsg("Le combat se joue en tour par tour", 20, "blue")
	fmt.Println("")
	utils.SpeedMsg("A chaque fois que vous jouez, plusieurs options s'offrent à vous : ", 20, "blue")
	fmt.Println("")

	// Appel de BattleMenuTutoriel dans le package utils pour afficher le menu du joueur pendant le tutoriel
	utils.BattleMenuTutorial()

	// Continue d'afficher chaque ligne du tutoriel
	utils.Input()
	utils.ClearConsole()
	utils.SpeedMsg("Lors de chaque attaque, vous verrez le nombre de dégats infligés : ", 20, "blue")
	fmt.Println()
	Green.Print("Vous avez infligé 20 dégats à Ennemi 1")
	utils.Input()
	utils.ClearConsole()
	utils.SpeedMsg("Mais vous pouvez également en recevoir : ", 20, "blue")
	fmt.Println()
	Red.Print("Vous avez reçu 50 dégats (coup critique) de Ennemi 1 !")
	utils.Input()
	utils.ClearConsole()
	utils.SpeedMsg("Lors de votre tour, vous pourrez également utiliser vos abilités : ", 20, "blue")
	fmt.Println("")

	// Choix du spell par le biais de AbilitiesTutorial qui affiche les spells et demande d'en choisir un
	spell := utils.AbilitiesTutorial()

	// Continue d'afficher ligne par ligne le tutoriel
	utils.SpeedMsg(spell+" à infligé 50 dégats à Ennemi 1", 20, "green")
	utils.Input()
	utils.ClearConsole()
	utils.SpeedMsg("Félicitation, vous êtes prêt pour votre premier combat ! Bonne chance", 20, "blue")
	utils.Input()
	utils.ClearConsole()

	// Passage niveau 2 à la fin du tutoriel
	p.niveau = 2
	Blue.Println("Vous avez atteint le niveau 2 !")

	// Retour au menu
	p.Menu()
}

// isDead vérifie la mort d'au moins un des deux combattant d'un duel
func (p *Personnage) isDead(e *Ennemy) bool {
	if p.currentHp <= 0 || e.HpCurrent <= 0 {
		return true
	} else {
		return false
	}
}

// battle est la fonction de combat
func (p *Personnage) battle() {
	utils.ClearConsole()
	// Sauvegarde des spells qui changent

	actualAuto := p.skill[0].Damages
	spell2 := p.skill[2].Damages

	// Configuration de l'ennemi
	ennemi1 := Monstres[p.niveau-2][p.ennemi]

	// Affichage du duel et des statistiques de l'ennemi
	utils.SpeedMsg(p.nom+" VS "+ennemi1.Name+"\n", 20, "red")
	fmt.Println("-----------------------------")
	utils.SpeedMsg("Statistiques de l'ennemi : \n", 20, "default")
	fmt.Println("Points de vie : " + strconv.Itoa(ennemi1.HpMax))
	fmt.Println("Dégats : " + strconv.Itoa(ennemi1.DamagesMin) + " - " + strconv.Itoa(ennemi1.DamgesMax))
	fmt.Println("Chance de coup critique : " + strconv.Itoa(ennemi1.CriticalChance))
	fmt.Println("-----------------------------")
	utils.SpeedMsg("Bonne chance combattant !", 20, "green")
	fmt.Println()
	fmt.Print("Appuyez sur entrée pour continuer")
	utils.Input()
	utils.ClearConsole()

	// Boucle infini jusqu'à qu'un des deux combattants soit mort
	for !(p.isDead(&ennemi1)) {
		// Tour de l'ennemi
		p.ennemyRound(&ennemi1)
		utils.ClearConsole()

		if !(p.isDead(&ennemi1)) {
			// Tour du joueur si il n'est pas mort
			p.playerRound(&ennemi1)
			utils.ClearConsole()
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
		utils.SpeedMsg("Récompenses : \n", 30, "default")
		time.Sleep(1 * time.Second)
		utils.SpeedMsg("+ 300 coro\n", 30, "blue")
		time.Sleep(1 * time.Second)
		utils.SpeedMsg("+ 3 Potions\n", 30, "blue")
		time.Sleep(1 * time.Second)
		utils.SpeedMsg("+ 1 Niveau\n", 30, "blue")

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
		utils.Input()
		utils.ClearConsole()

		if p.ennemi == 0 {
			Blue.Println("Vous avez atteint le niveau " + strconv.Itoa(p.niveau))
		}

		// Re attribution des utilisations de chaque spell
		for i := 1; i < 4; i++ {
			p.skill[i].StillUse = p.skill[i].MaxUse
		}
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
		utils.SpeedMsg("Préparez vous et revenez plus fort ! \n", 30, "default")
		fmt.Println()
		fmt.Print("Appuyez sur entrée pour continuer")
		utils.Input()

		// Ré-attribution des utilisations de chaque spell ainsi que remise à 100% des points de vie du joueur
		p.currentHp = p.maxHP
		for i := 1; i < 4; i++ {
			p.skill[i].StillUse = p.skill[i].MaxUse
		}
		// Réèattribution des dégats a l'auto
		p.skill[0].Damages = actualAuto
		p.skill[2].Damages = spell2
		utils.ClearConsole()
		p.Menu()
	}
}

// accessInventory permet au joueur d'accéder à son inventaire.
func (p *Personnage) accessInventory() {

	// Affichage de l'inventaire
	fmt.Println("Item            Quantité")
	fmt.Println("---------------------------")
	for _, item := range p.inventory {
		// Utilisation de fmt.Printf pour aligner les colonnes
		fmt.Printf("%-20s %-7d\n", item.Name, item.Quantite)
	}
	fmt.Println("---------------------------")
	fmt.Println("[1] Utiliser une potion")
	fmt.Println("[2] Quitter l'inventaire")
	choice, _ := utils.Inputint()
	switch choice {

	// Le joueur prend une potion
	case 1:
		utils.ClearConsole()
		p.takePotion()
		p.accessInventory()

	// Le joueur quitte l'inventaire
	case 2:
		utils.ClearConsole()

	// Choix non proposé
	default:
		utils.ClearConsole()
		Red.Println("Veuillez saisir une donnée valide")
		p.accessInventory()
	}
}

// showSkills affiche les abilités du personnage.
func (p *Personnage) showSkills() {
	// Affichage des sorts du personnage
	fmt.Printf("%-20s%-10s%-60s\n", "Nom", "Dégâts", "Description")
	fmt.Println("------------------------------------------------------------------")

	for _, spell := range p.skill {
		fmt.Printf("%-20s%-10d%-60s\n", spell.Name, spell.Damages, spell.Description)
	}
	fmt.Println("-------------------------------------------------------------------------------------------")
}

// displayInfo affiche les informations du personnage.
func (p *Personnage) displayInfo() {
	fmt.Println("--- ", p.nom, " ---")
	fmt.Println("Classe : ", p.classe)
	fmt.Println("Niveau : ", p.niveau)
	fmt.Println("Ennemis battus : " + strconv.Itoa(p.ennemi) + " /3")
	fmt.Println("Points de vie : ", p.currentHp, "/", p.maxHP)
	fmt.Println("--------------")
}

// Fonction pour vérifier si un item avec un nom donné se trouve dans une liste d'items
func (p *Personnage) alreadyBuy(itemName string) bool {
	for _, item := range p.inventory {
		if item.Name == itemName {
			return true
		}
	}
	return false
}

// Marchand permet au joueur d'interagir avec le marchand.
func (p *Personnage) Marchand() {

	// Configuration de la liste de vente du marchand ainsi que toutes les spécifités des items
	var itemMarchand = []Article{
		{"Potions", 10, "Potion pour récupèrer 40 points de vie", 0, 0, false},
		{"Guinzoo", 100, "+20% hp | +15 ad", 15, p.maxHP / 5, true},
		{"Masque grenouille", 1000, "+600% hp", 0, p.maxHP * 6, true},
		{"Avrilvert", 350, "+100% ad", p.skill[0].Damages, 0, true},
		{"Rook", 450, "+80% hp", 0, p.maxHP * 80 / 100, true},
		{"Anneau de gel", 500, "+60%hp | +20 ad ", 20, p.maxHP * 60 / 100, true},
	}

	// Affichage des items que vend le marchand
	for {
		Green.Println("Argent :  ", p.nbItem("Argent"))
		fmt.Println("    Article             Prix                Description")
		fmt.Println("-----------------------------------------------------")
		for i, item := range itemMarchand {
			if p.alreadyBuy(item.Name) && item.Unique == true {
				Blue.Printf("%s - acheté !\n", item.Name)
			} else {
				fmt.Printf("[%d] %-20s %-7d %-s\n", i+1, item.Name, item.Price, item.Description)
			}
		}
		fmt.Println("-----------------------------------------------------")
		fmt.Printf("[%d] Sortir \n", len(itemMarchand)+1)
		choice, _ := utils.Inputint()

		// Le choix est un item
		if choice > 0 && choice <= len(itemMarchand) {
			// Selection de l'item
			selectedItem := itemMarchand[choice-1]

			if p.enoughMoney(selectedItem.Price) { // Vérification si le joueur a assez d'argent
				if !(p.alreadyBuy(selectedItem.Name)) || selectedItem.Unique == false { // Vérification si le joueur ne l'a pas déjà acheté si il est unique
					p.addInventory(selectedItem.Name, 1)                              // Ajout de l'item a l'inventaire
					p.inventory[p.findIndex("Argent")].Quantite -= selectedItem.Price // On retire l'argent de l'inventaire
					utils.ClearConsole()
					p.skill[0].Damages += selectedItem.Ad // On ajoute les dégats que donne l'item
					p.maxHP += selectedItem.Health        // On ajoute la vie que donne l'item sur les hp max
					p.currentHp += selectedItem.Health    // On ajoute la vie que donne l'item sur les hp actuels
					Blue.Printf("Vous avez acheté : %s pour %d coro !\n", selectedItem.Name, selectedItem.Price)
				} else { // Item déjà acheté
					utils.ClearConsole()
					Red.Println("Vous avez déjà acheté cet item")
				}
			} else { // Pas assez d'argent pour l'acheter
				utils.ClearConsole()
				Red.Println("Vous n'avez pas assez d'argent !")
			}

		} else if choice == len(itemMarchand)+1 { // Choix de sortie du marchand
			utils.ClearConsole()
			Blue.Println("Sortie du marchand")
			p.Menu()
			break
		} else { // Choix non proposé
			utils.ClearConsole()
			Red.Println("Veuillez saisir une donnée valide")
		}
	}
}

// takePotion permet au joueur de prendre une potion si il en a une.
func (p *Personnage) takePotion() {

	if p.nbItem("Potions") > 0 { // Vérifie qu'il a une potion
		if p.currentHp == p.maxHP { // Vérifie qu'il est pas déjà full hp
			Red.Println("Vous êtes déjà au maximum de points de vie !")
		} else {
			if p.maxHP-p.currentHp < 50 { // Si il lui manque une seule potion pour etre full
				p.currentHp = p.maxHP // On le met full pour éviter le débordement d'hp
			} else {
				p.currentHp += 50 // Sinon on ajoute simplement 50
			}
			p.removeInventory("Potions", 1) // On supprime 1 potion de son inventaire
			Blue.Println("- 1 potion, points de vie : ", p.currentHp)
			Blue.Println("Potions restantes : ", p.nbItem("Potions"))
		}
	}
}

// enoughMoney vérifie si le joueur a suffisamment d'argent pour acheter un objet.
func (p *Personnage) enoughMoney(cost int) bool {
	if p.nbItem("Argent")-cost < 0 {
		return false
	} else {
		return true
	}
}

// addInventory ajoute un objet à l'inventaire du joueur.
func (p *Personnage) addInventory(itemName string, quantity int) {
	for i, item := range p.inventory { // Parcourt l'inventaire
		if item.Name == itemName { // On trouve l'item
			p.inventory[i].Quantite += quantity // On augmente sa quantité
			return
		}
	} // Si on le trouve pas
	newItem := Item{Name: itemName, Quantite: quantity} // On le crée
	p.inventory = append(p.inventory, newItem)          // Et on l'ajoute
}

// removeInventory retire un objet de l'inventaire du joueur.
func (p *Personnage) removeInventory(itemName string, quantity int) {
	for i, item := range p.inventory {
		if item.Name == itemName {
			// L'élément existe dans la liste.
			if item.Quantite <= quantity {
				// Retire complètement l'élément si la quantité est suffisante.
				p.inventory = append(p.inventory[:i], p.inventory[i+1:]...)
			} else {
				// Met à jour la quantité de l'élément si la quantité est insuffisante.
				p.inventory[i].Quantite -= quantity
			}
		}
	}
}

// nbItem renvoie la quantité d'un objet dans l'inventaire du joueur.
func (p *Personnage) nbItem(itemName string) int {
	for _, item := range p.inventory {
		if item.Name == itemName {
			return item.Quantite
		}
	}
	return -1
}

// findIndex renvoie l'index d'un objet dans l'inventaire du joueur.
func (p *Personnage) findIndex(itemName string) int {
	index := -1
	for i, item := range p.inventory { // On parcourt l'inventaire
		if item.Name == itemName { // On trouve le bon item
			index = i // On configure l'index sur celui auquel on vient de le trouver
			break
		}
	}
	return index // On retourne l'index ou on l'a trouvé (si on trouve pas index = -1)
}
