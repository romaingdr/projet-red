// FICHIER UTILISE POUR LA CONFIGURATION DU SOCKET TCP AINSI QUE LE COMBAT COOP DES JOUEURS CONNECTES

package utils

import (
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"time"
)

var (
	joueurs       = make(map[string]int)
	nbPlayers     int
	ip            string
	clienConns    []net.Conn
	side          string
	ennemy        = EnnemyObjective{Name: "Ennemi 1", HpCurrent: 200, HpMax: 200, Damages: 40}
	nbDegats      = 0
	degatsEnvoyes = 0
	degatsJoueur  = make(map[string]int)
)

// LocalIP est utilisé afin de récuperer l'adresse ip locale de l'utilisateur
func LocalIP() string {
	// Crée une connexion UDP à une adresse quelconque
	conn, _ := net.Dial("udp", "8.8.8.8:80")
	defer conn.Close()

	// Obtient les informations sur le socket local
	localAddr := conn.LocalAddr().(*net.UDPAddr)

	// Renvoie l'adresse IP locale au format string
	return localAddr.IP.String()
}

// isDead2 sert à vérifier si un des joueurs participants à l'objectif ou l'ennemi est mort
func isDead2() bool {
	for joueur := range joueurs {
		if joueurs[joueur] <= 0 {
			return true
		}
	}
	if ennemy.HpCurrent <= 0 {
		return true
	}
	return false
}

// randomKey sert à renvoyer un joueur aléatoire parmis la liste de joueurs participants à l'objectif
func randomKey(m map[string]int) string {
	// Initialiser le générateur de nombres aléatoires avec une graine
	rand.Seed(time.Now().UnixNano())

	// Créer une liste de clés à partir de la map
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}

	// Générer un index aléatoire
	randomIndex := rand.Intn(len(keys))

	// Retourner la clé correspondante à l'index aléatoire
	return keys[randomIndex]
}

// multiObjectives est la fonction du menu du mode de jeu Multijoueur - Objectif
func multiObjectives(p *Personnage) {
	fmt.Println("[1] - Créer un salon")
	fmt.Println("[2] - Rejoindre un salon")
	fmt.Println("[3] - Retour")
	choice, _ := Inputint()
	switch choice {
	case 1:
		ClearConsole()
		fmt.Print("Nombre de joueurs (2 - 9 Max.): ")
		nbPlayers, _ = Inputint()
		if nbPlayers < 2 || nbPlayers > 9 {
			ClearConsole()
			Red.Println("Le mode objectif se joue entre 2 et 9 joueurs seulement !")
			multiObjectives(p)
		} else {
			createParty(p)
		}
	case 2:
		ClearConsole()
		joinParty(p)
	case 3:
		ClearConsole()
		p.Menu()
	default:
		ClearConsole()
		Red.Println("Veuillez saisir une donnée valide")
		multiObjectives(p)
	}
}

// afficheJoueursHp est utilisé afin d'afficher les points de vie des joueurs participants
func afficheJoueursHp() {
	ClearConsole()
	for joueur := range joueurs {
		Green.Println(joueur+" - ", joueurs[joueur], " hp")
	}
}

// afficheEnnemi est utilisé afin d'afficher les points de vie de l'ennemi
func afficheEnnemi() {
	Red.Println(ennemy.Name + " - " + strconv.Itoa(ennemy.HpCurrent) + "/" + strconv.Itoa(ennemy.HpMax))
}

// afficheSalonAttente sert à afficher le salon d'attente le temps que tout les joueurs soient connectés au serveur
func afficheSalonAttente() {
	ClearConsole()
	Blue.Println("Adresse du salon : ", ip)
	fmt.Println("----- JOUEURS (", len(joueurs), "/", nbPlayers, ") -----")
	for joueur := range joueurs {
		fmt.Println(joueur)
	}
}

// afficheMenu sert à afficher le menu de combat des joueurs, afin qu'ils puissent attaquer
func afficheMenu(p *Personnage) {
	fmt.Println("---- Tour des alliés ----")
	fmt.Println("[1] Attaque automatique ")
	fmt.Println("[2] Abilités ")
	fmt.Println("-------------------------")
	choice, _ := Inputint()
	switch choice {
	case 1:
		if side == "server" {
			degatsEnvoyes += p.skill[0].Damages
			degatsJoueur[p.nom] = degatsEnvoyes
			nbDegats += 1
			ClearConsole()
			Green.Println("Vous avez infligé ", p.skill[0].Damages, " dégats")
			fmt.Println("En attente des autres joueurs...")
			if nbDegats == nbPlayers {
				ClearConsole()
				resultatsDegats(p)
			}
		} else {
			sendToServer(clientConns, p.nom+"|"+strconv.Itoa(p.skill[0].Damages))
			ClearConsole()
			Green.Println("Vous avez infligé ", p.skill[0].Damages, " dégats")
			fmt.Println("En attente des autres joueurs...")
		}
	case 2:
		ClearConsole()
		afficheJoueursHp()
		afficheEnnemi()
		fmt.Println("----- A votre tour -----")

		for i := 1; i < 4; i++ {
			skill := p.skill[i]
			fmt.Printf("[%d] %-20s %-10d %d/%d\n", i, skill.Name, skill.Damages, skill.StillUse, skill.MaxUse)
		}

		fmt.Println("------------------------")
		fmt.Println("[4] Sortir")

		choice2, _ := Inputint()
		var degats int

		switch choice2 {
		case 1, 2, 3:
			skill := p.skill[choice2]
			degats = skill.Damages
			if skill.StillUse > 0 {
				if side == "server" {
					degatsEnvoyes += degats
					degatsJoueur[p.nom] = degatsEnvoyes
					nbDegats += 1
					ClearConsole()
					Green.Println("Vous avez infligé ", degats, " dégats")
					fmt.Println("En attente des autres joueurs...")
					if nbDegats == nbPlayers {
						ClearConsole()
						resultatsDegats(p)
					}
				} else {
					sendToServer(clientConns, p.nom+"|"+strconv.Itoa(degats))
					ClearConsole()
					Green.Println("Vous avez infligé ", degats, " dégats")
					fmt.Println("En attente des autres joueurs...")
				}
			} else {
				ClearConsole()
				Red.Println("Vous ne pouvez plus utiliser cette compétence !")
				afficheJoueursHp()
				afficheEnnemi()
				afficheMenu(p)
			}
		default:
			ClearConsole()
			Red.Println("Veuillez saisir une donnée valide !")
			afficheJoueursHp()
			afficheEnnemi()
			afficheMenu(p)
		}
	default:
		ClearConsole()
		Red.Println("Veuillez saisir une donnée valide !")
		afficheJoueursHp()
		afficheEnnemi()
		afficheMenu(p)
	}
}

// resultatsDegats sert à récuperer tout les dégats envoyés par les joueurs et les infliger à l'ennemi
func resultatsDegats(p *Personnage) {
	ClearConsole()
	for joueur := range degatsJoueur {
		Green.Println(joueur + " a infligé " + strconv.Itoa(degatsJoueur[joueur]))
		ennemy.HpCurrent -= degatsJoueur[joueur]
	}
	joueurChoisi := randomKey(joueurs)
	joueurs[joueurChoisi] -= ennemy.Damages
	Red.Println("L'ennmi a infligé " + strconv.Itoa(ennemy.Damages) + " dégats à " + joueurChoisi)
	sendToClient(clientConns, joueurChoisi+"|"+strconv.Itoa(joueurs[joueurChoisi]))
	time.Sleep(1 * time.Second)
	sendToClient(clientConns, "&ennemy|"+strconv.Itoa(ennemy.HpCurrent))
	time.Sleep(1 * time.Second)
	if !(isDead2()) {
		sendToClient(clientConns, "start")
		degatsEnvoyes = 0
		nbDegats = 0
		afficheJoueursHp()
		afficheEnnemi()
		afficheMenu(p)
	} else {
		if ennemy.HpCurrent <= 0 {
			sendToClient(clientConns, "1")
			ClearConsole()
			SpeedMsg("Vous avez battu l'ennemi\n", 20, "green")
			fmt.Println()
			SpeedMsg("Félicitation , vous remportez cet objectif !\n", 20, "green")
			fmt.Println()
			fmt.Print("retourner au menu")
			Input()
			ClearConsole()
			p.Menu()
		} else {
			var joueurMort string
			for joueur := range joueurs {
				if joueurs[joueur] <= 0 {
					joueurMort = joueur
				}
			}
			sendToClient(clientConns, "0|"+joueurMort)
			ClearConsole()
			SpeedMsg(joueurMort+" est mort !\n", 20, "red")
			fmt.Println()
			SpeedMsg("Vous avez perdu !\n", 20, "red")
			fmt.Println()
			fmt.Print("retourner au menu")
			Input()
			ClearConsole()
			p.Menu()
		}
	}
}

// createParty sert à créer le groupe multijoueur afin de combattre
func createParty(p *Personnage) {

	listener, _ := net.Listen("tcp", "0.0.0.0:12345")

	defer listener.Close()

	ip = LocalIP()
	joueurs[p.nom] = p.maxHP
	side = "server"

	// Message d'attente
	afficheSalonAttente()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Erreur lors de l'acceptation d'une connexion entrante :", err)
			continue
		}
		go gererConnection(conn, p)
	}
}

// gererConnection sert à entretenir la connexion au serveur et à gérer individuellement les messages recu par les participants
func gererConnection(conn net.Conn, p *Personnage) {
	defer conn.Close()

	// Envoi de notre pseudo au client
	messageEnvoye := p.nom + "|" + strconv.Itoa(p.maxHP)
	conn.Write([]byte(messageEnvoye))

	// Récupération du pseudo du client
	buffer := make([]byte, 1024)
	n, _ := conn.Read(buffer)

	message := string(buffer[:n])
	pseudo, hp := sliceArgument(message)
	clientConns = append(clientConns, conn)
	hpInt, _ := strconv.Atoi(hp)
	joueurs[pseudo] = hpInt

	afficheSalonAttente()

	if len(joueurs) == nbPlayers {
		for joueur := range joueurs {
			sendMessageToClient(clientConns, joueur+"|"+strconv.Itoa(joueurs[joueur]))
		}
		time.Sleep(2 * time.Second)
		sendMessageToClient(clientConns, "start")
		afficheJoueursHp()
		afficheEnnemi()
		afficheMenu(p)
	}
	for {

		// Un message est recu
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Erreur lors de la lecture depuis le client :", err)
			return
		}

		// Gérer un message recu
		message := string(buffer[:n])
		pseudo, degats := sliceArgument(message)
		nbDegats += 1
		degatsInt, _ := strconv.Atoi(degats)
		degatsEnvoyes += degatsInt
		degatsJoueur[pseudo] = degatsInt
		if nbDegats == nbPlayers {
			ClearConsole()
			resultatsDegats(p)
		}

	}
}

// joinParty sert à rejoindre un groupe via une adress ip et un port afin de participer
func joinParty(p *Personnage) {
	ClearConsole()
	fmt.Print("Adresse du salon : ")
	serverIP := Input()
	serverPort := 12345

	ClearConsole()
	side = "client"
	// connexion tcp au serveur
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIP, serverPort))
	if err != nil {
		Red.Println("Impossible de trouver le serveur !")
		multiObjectives(p)
	}
	defer conn.Close()
	clientConns = append(clientConns, conn)

	// On envoie notre pseudo
	message := p.nom + "|" + strconv.Itoa(p.maxHP)
	conn.Write([]byte(message))

	// On lit le pseudo du host
	// Lis le pseudo / hp du joueur host
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Erreur lors de la lecture depuis le serveur :", err)
		return
	}

	message1 := string(buffer[:n])
	pseudo, hp := sliceArgument(message1)
	hpAdversaire, _ := strconv.Atoi(hp)
	joueurs[pseudo] = hpAdversaire

	fmt.Println("En attente du début de la partie...")
	for {
		// Si un message est recu
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Erreur lors de la lecture depuis le serveur :", err)
			return
		}

		// Traiter le message
		message := string(buffer[:n])
		if message == "start" {
			afficheJoueursHp()
			afficheEnnemi()
			afficheMenu(p)
		} else if message[0] == '0' {
			ClearConsole()
			_, joueurMort := sliceArgument(message)
			SpeedMsg(joueurMort+" est mort !\n", 20, "red")
			fmt.Println()
			SpeedMsg("Vous avez perdu !\n", 20, "red")
			fmt.Println()
			fmt.Print("retourner au menu")
			Input()
			ClearConsole()
			p.Menu()
		} else if message == "1" {
			ClearConsole()
			SpeedMsg("Vous avez battu l'ennemi\n", 20, "green")
			fmt.Println()
			SpeedMsg("Félicitation , vous remportez cet objectif !\n", 20, "green")
			fmt.Println()
			fmt.Print("retourner au menu")
			Input()
			ClearConsole()
			p.Menu()
		} else {
			pseudo, hp := sliceArgument(message)
			hpInt, _ := strconv.Atoi(hp)
			if pseudo == "&ennemy" {
				ennemy.HpCurrent = hpInt
			} else {
				joueurs[pseudo] = hpInt
			}
		}
	}
}

// sendMessageToServer sert à envoyer un message au serveur depuis le client
func sendToServer(clientConns []net.Conn, message string) {
	for _, clientConn := range clientConns {
		_, err := clientConn.Write([]byte(message))
		if err != nil {
			fmt.Println("Erreur lors de l'envoi de données au client :", err)
			// Gérez l'erreur de manière appropriée
		}
	}
}

// sendMessageToServer sert à envoyer un message au client depuis le serveur
func sendToClient(clientConns []net.Conn, message string) {
	for _, clientConn := range clientConns {
		_, err := clientConn.Write([]byte(message))
		if err != nil {
			fmt.Println("Erreur lors de l'envoi de données au client :", err)
			// Gérez l'erreur de manière appropriée
		}
	}
}
