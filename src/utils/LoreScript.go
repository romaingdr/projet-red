package utils

import (
	"fmt"
	"time"
)

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
		ClearConsole()
		fmt.Println("Notre héro est un soldat américain en pleine 3e guerre mondiale")
		fmt.Println("Son long récit débute en pleine tranchée, n'ayant pour compagnie que ses camarades tombés au combat.")
		fmt.Println("Autour de lui les balles résonnaient et les pas se rapprochaient de sa position...")
		Input()
		ClearConsole()
		fmt.Println("Tandis que notré héro peine à se remettre de sa situation, un obstacle vint lui faire face")
		fmt.Println("Un soldat à la face brulée et couverte du sang de ses pairs se dressa face à lui...")
		Input()
	} else if p.ennemi == 1 {
		fmt.Println("Notre héro qui jusqu'à là n'avait vu que le sang versé de ses frères d'armes venait d'abattre son premier ennemi")
		fmt.Println("Il savait dorénavant qu'il possédait les armes et les capacités afin d'affronter plus fort")
		Input()
		ClearConsole()
		fmt.Println("Déterminé à venger ses compagnons et à faire triompher fierement les couleurs de son pays, il avança")
		fmt.Println("Tant et si bien qu'il voyait à présent la lumière qui pourrait le mener hors de cet enfer")
		Input()
		ClearConsole()
		SpeedMsg("- AHHHHHHHH", 20, "red")
		time.Sleep(300 * time.Millisecond)
	} else {
		fmt.Println("Après 2 ennemi achevés, la fin de la tranchée, la lumière qui semblait contenir ce sentiment de victoire était à portée de main")
		Input()
		ClearConsole()
		fmt.Println("Il courut rapidement, très rapidement")
		fmt.Println("Les grondements sourds autour de lui lui faisaient entreprendre des perspectives d'aides alliés")
		Input()
		fmt.Println("Soudain, dans toute cette adrénaline, il trébuchat sur un corps qui trainait là")
		Input()
		ClearConsole()
		fmt.Println("Son corps qui jusqu'à là avait tenu tout ce chemin ,")
		fmt.Println("ne fit qu'un avec les corps inanimés dont on ne pouvait distringuer les visages")
		Input()
		ClearConsole()
		fmt.Println("Ces corps qui semblaient inanimés")
		Input()
		Red.Println("à l'exeption d'un")
		time.Sleep(1 * time.Second)
	}
}

// ScriptNiv2 affiche le texte du lore sur tout le niveau 1
func ScriptNiv2(p *Personnage) {
	if p.ennemi == 0 {
		fmt.Println("Le combat contre ce qui semblait être le soldat le plus puissant")
		fmt.Println("avait affaibli critiquement notre héro qui se réveilla dans un endroit inconnu")
		Input()
		ClearConsole()
		SpeedMsg("- Où suis-je ?", 20, "cyan")
		Input()
		ClearConsole()
		fmt.Println("En face de lui se dressait une immense grotte")
		fmt.Println("Avec sur le sol, un miroir, un étrange miroir.")
		Input()
		ClearConsole()
		fmt.Println("D'un élan de curiosité, il s'approche du miroir")
		fmt.Println("En marchant dessus, il passe à travers et fait une chute de plusieurs mètres.")
		fmt.Println("Pour au final atterir dans ce qui semble être à première vue un chateau abandonné.")
		Input()
		ClearConsole()
		SpeedMsg("- Où ai-je encore atterit ?\n", 20, "cyan")
		Input()
		SpeedMsg("- Viens par ici\n", 20, "green")
		Input()
		SpeedMsg("- Qui a parlé ?\n", 20, "cyan")
		Input()
		SpeedMsg("- C'est moi le miroir, mes congénères ont subis un horrible sortilège lancé par le maitre de l'univers\n", 20, "green")
		Input()
		SpeedMsg("- Mais quel est le rapport avec moi ?\n", 20, "cyan")
		Input()
		SpeedMsg("- Le maitre est devenu trop puissant et veut maintenant tout détruire, tu es l'élu\n", 20, "green")
		Input()
		SpeedMsg("- Mais ? je n'.. \n", 20, "cyan")
		SpeedMsg("- Tu as 48h pour remonter tout les miroirs à leurs sources afin de t'en rapprocher et de l'élim..\n", 20, "green")
		SpeedMsg("- Tu as aussi entendu ce bruit ??\n", 20, "cyan")
		Input()
		ClearConsole()
		SpeedMsg("- Qui à osé pénetrer MON chateau ?\n", 20, "red")
		Input()

	} else if p.ennemi == 1 {
		fmt.Println("Les évènements se sont déroulés trop rapidement")
		fmt.Println("- Qui été ce miroir ?")
		fmt.Println("- De qui parlait-il ?")
		fmt.Println("- Pourquoi moi ?")
		Input()
		ClearConsole()
		fmt.Println("Seul dans cet immense château, notre héro était perdu")
		fmt.Println()
		SpeedMsg("- Que dois-je faire maintenant ?", 20, "cyan")
		Input()
		ClearConsole()
		SpeedMsg("- Tu révasses ?", 20, "red")
		Input()
	} else {
		fmt.Println("Une fois la deuxième chauve-souris achevée, un immense roi apparut devant lui")
		Input()
		SpeedMsg("- Tu ne croyais pas t'en sortir comme ça ?\n", 20, "red")
		Input()
		SpeedMsg("- Qui êtes-vous, pourquoi me voulez vous du mal ?\n", 20, "cyan")
		Input()
		SpeedMsg("- Tu poses trop de questions, viens te battre\n", 20, "red")
		Input()
	}
}

// ScriptNiv3 affiche le texte du lore sur tout le niveau 1
func ScriptNiv3(p *Personnage) {
	if p.ennemi == 0 {
		fmt.Println("C'était fini, le cauchemar chateau et de son roi était derrière lui")
		fmt.Println("Il était enfin de retour dans cette mystérieuse grotte avec encore ce miroir au sol")
		fmt.Println("Cette fois ci le héro réfléchit longuement avant de franchir le pas")
		Input()
		ClearConsole()
		SpeedMsg("- Que dois-je faire ?\n", 20, "cyan")
		Input()
		SpeedMsg("- Le miroir a bien dit 48h ?\n", 20, "cyan")
		Input()
		SpeedMsg("- Je dois en savoir plus\n", 20, "cyan")
		Input()
		ClearConsole()
		fmt.Println("Le joueur franchit le miroir et se retrouva cette fois face à un jardin immense")
		fmt.Println("Aucun ennemi ne semblait en vu ni même le miroir qu'il avait rencontré au chateau")
		Input()
		ClearConsole()
		SpeedMsg("- AHHHHHHH\n", 20, "cyan")
		Input()
		SpeedMsg("[Fourmi] - N'ai pas peur je ne te veux pas de mal moi\n", 20, "green")
		Input()
		SpeedMsg("- Ma..Mais tu es un fourmi géante ?\n", 20, "cyan")
		Input()
		SpeedMsg("[Fourmi] - Je suis envoyé par le miroir pour t'aider dans le jardin\n", 20, "green")
		Input()
		SpeedMsg("- Enfin , merci ! Pourriez vous me dire ce que je suis censé faire maintenant ?\n", 20, "cyan")
		Input()
		SpeedMsg("[Fourmi] - Fonce vers le grand arbre tu y trouveras ce dont tu as besoin\n", 20, "green")
		Input()
		ClearConsole()
		fmt.Println("En route vers l'arbre, le héros pris une pause afin de contempler ce joli paysage autour de lui")
		Input()
		SpeedMsg("- Cet escargot a l'air endormi\n", 20, "cyan")
		Input()
		fmt.Println("Il s'approchat de lescargot...")
		Input()
	} else if p.ennemi == 1 {
		fmt.Println("Notre héro en avait maintenant la confirmation, les gens ici n'étaient pas tous là pour l'aider")
		fmt.Println("Il se remis en route et atteignat sans trop de difficulté l'arbre recommandé par la fourmi géante")
		Input()
		ClearConsole()
		SpeedMsg("- C'est ce miroir que tu veux ?\n", 20, "red")
		Input()
		SpeedMsg("- Oui exactement !\n", 20, "cyan")
		Input()
		SpeedMsg("- Il va falloir me passer sur les ailes !\n", 20, "red")
		Input()
	} else {
		SpeedMsg("- Tu viens de tuer mon ami le moustique, tu vas le payer cher", 20, "red")
		Input()
	}
}

// ScriptNiv4 affiche le texte du lore sur tout le niveau 1
func ScriptNiv4(p *Personnage) {
	if p.ennemi == 0 {
		fmt.Println("Notre héro avait passer avec aisance le jardin, il commençait à prendre habitude des miroirs")
		fmt.Println("Une fois de retour dans la grotte, il s'empressa de foncer dans le nouveau miroir, il y voyait mainentant un jeu")
		Input()
		ClearConsole()
		SpeedMsg("- Une ville sous-marine ? Et je respire normalement ?\n", 20, "cyan")
		Input()
		fmt.Println("Soudain, à l'ombre d'un corail, un dauphin surgit de nul part")
		Input()
		SpeedMsg("Je commençais à avoir faim !\n", 20, "red")
		Input()
	} else if p.ennemi == 1 {
		fmt.Println("Après cette victoire contre le dauphin, notre héro continue sa quête vers le miroir")
		fmt.Println("La ville parassait étonnement calme, le bruit de l'eau cachait les nuisances sonores de la ville")
		fmt.Println("Pour la première fois, le héro se sentait bien...")
		Input()
		ClearConsole()
		SpeedMsg("- Eh toi là bas !, ton miroir se trouve dans le palai derrière moi, mais tu peux rebrousser chemin\n", 20, "red")
		Input()
		SpeedMsg("- AHAHAHAH, ça serait mal me connaitre !\n", 20, "cyan")
		Input()
		SpeedMsg("- ... je t'aurais prévenu\n", 20, "red")
		Input()
	} else {
		fmt.Println("Une fois ce requin vigile battu, notre héro entra dans le palai des sirènes")
		Input()
		SpeedMsg("- BONJOUR A TOUS, JE CHERCHE LE MIROIR !\n", 20, "cyan")
		Input()
		SpeedMsg("- C'est de ça que tu parles ? pas de soucis, essaye toujours", 20, "red")
		Input()
	}
}

// ScriptNiv5 affiche le texte du lore sur tout le niveau 1
func ScriptNiv5(p *Personnage) {
	if p.ennemi == 0 {
		fmt.Println("Décidement, les niveaux s'enchainaient et notre héro ne semblait reculer devant rien")
		fmt.Println("Il fonça dans le miroir mais cette fois ci tout était différent...")
		Input()
		ClearConsole()
		fmt.Println("Il faisait très sombre et l'atmosphère était pesante")
		fmt.Println("En levant les yeux au ciel, notre héro aperçoit, ou du moins distingue une grande silhouette")
		Input()
		ClearConsole()
		SpeedMsg("[G] - C'est donc toi l'élu ? Je lis dans ton esprit que tu ne connais même pas la prophétie...\n", 20, "red")
		Input()
		SpeedMsg("- De quoi parl..\n", 20, "cyan")
		fmt.Println()
		SpeedMsg("[G] - Ici c'est moi qui parle, si tu savais ce que tu vas devoir encore endurer et l'avenir qui t'attend...\n", 20, "red")
		Input()
		SpeedMsg("- Peut importe ce que je dois traverser je le ferais , viens te battre si tu en es capable !\n", 20, "cyan")
		Input()
		SpeedMsg("[G] - ... comme tu voudras", 20, "red")
		Input()
	}
}

// ScriptNiv6 affiche le texte du lore sur tout le niveau 1
func ScriptNiv6(p *Personnage) {
	if p.ennemi == 0 {
		fmt.Println("Comment ? Comment avait-il fait pour battre le général ?")
		fmt.Println(p.nom + " n'était pas n'importe qui, il était l'élu qui ferait tout changer")
		Input()
		ClearConsole()
		fmt.Println("Une fois de nouveau franchit le portail, il s'arrêta")
		fmt.Println("Posé sur le haut d'une colline rocheuse, il pensa")
		fmt.Println("A tout ce qui venait de se passer en si peu de temps")
		fmt.Println("Tout cela était-il réel ? ça n'avait plus d'importance")
		Input()
		ClearConsole()
		fmt.Println("Derrière la fumée qui se dégageait tout autour de lui, une silhouette se dessinait")
		fmt.Println("Elle apparut de plus en plus grande, comme si elle avançait dans sa direction")
		fmt.Println(p.nom + " quant à lui se sentait de moins en moins à l'aise")
		Input()
		ClearConsole()
		fmt.Println("Quand soudain")
		Input()
		ClearConsole()
		fmt.Println("Il se figea sur place, tétanisé, le regard fixe, il comprit")
		Input()
		ClearConsole()
		fmt.Println("Cette silouhette, c'était la sienne")
		Input()
		ClearConsole()
		SpeedMsg("[?] - Enfin je te rencontre\n", 30, "yellow")
		Input()
		SpeedMsg("- Qu..Qui es tu ?\n", 20, "cyan")
		Input()
		SpeedMsg("[?] - Je suis toi et tu es moi. Je suis l'élu et j'ai pris tes résponsabilités il y a 1 milliard d'années déjà\n", 30, "yellow")
		SpeedMsg("J'ai tout prévu afin de que tu viennes à moi et que je puisse t'éliminer comme toutes les autres fois\n", 30, "yellow")
		SpeedMsg("Je suis le maitre de l'univers", 30, "yellow")
		Input()
	}
}

func finalChoice(p *Personnage) {
	SpeedMsg(p.nom+" ,\n", 30, "default")
	SpeedMsg("Félicitation, vous avez battu le maitre de l'univers\n", 30, "default")
	SpeedMsg("Vous êtes maintenant le combattant le plus puissant que l'on ai connu\n", 30, "default")
	SpeedMsg("Il vous reste dorénavant à faire votre dernier choix\n", 30, "cyan")
	fmt.Println()
	fmt.Println("[1] Rester ici et devenir le maitre de l'univers à votre tour")
	fmt.Println("[2] Rentrer chez vous")
	fmt.Println()
	choixFinal, _ := Inputint()
	switch choixFinal {
	case 1:
		ClearConsole()
		fmt.Println("Vous êtes maintenant le maitre de l'univers")
		fmt.Println("Vous poursuivez une vie sans fin exemptée de toutes émotions")
		fmt.Println("le repos est eternel ici")
		Input()
		ClearConsole()
		fmt.Println("Merci d'avoir joué")
		p.niveau += 1
		Input()
		ClearConsole()
		p.Menu()
	case 2:
		ClearConsole()
		SpeedMsg("[Infirmière] - Reveillez-vous ! Reveillez-vous !\n", 20, "default")
		Input()
		SpeedMsg("- Ou suis-je, pourquoi suis-je sur un lit d'hopital ?\n", 30, "cyan")
		Input()
		SpeedMsg("[Infirmière] - Vous êtes surement en état de choc, on ne pensait plus vous revoir parmis nous !\n", 20, "default")
		Input()
		SpeedMsg("[Infirmière 2] - Vous venez de sortir du coma, une mine vous à exploser sous le pied pendant que vous combattiez\n", 20, "default")
		Input()
		SpeedMsg("[Infirmière] - Bon retour parmis nous !\n", 30, "default")
		Input()
		ClearConsole()
		fmt.Println("Merci d'avoir joué")
		p.niveau += 1
		Input()
		ClearConsole()
		p.Menu()
	default:
		ClearConsole()
		Red.Println("Veuillez saisir une donnée valide")
		finalChoice(p)
	}
}
