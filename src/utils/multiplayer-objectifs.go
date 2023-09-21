// jn|j1|j2|j3 j= commande envoi joueur / n = nb Joueurs

package utils

import (
	"fmt"
	"net"
	"strconv"
)

var (
	joueurs    = make(map[string]int)
	nbPlayers  int
	ip         string
	clienConns []net.Conn
)

func LocalIP() string {
	// Crée une connexion UDP à une adresse quelconque
	conn, _ := net.Dial("udp", "8.8.8.8:80")
	defer conn.Close()

	// Obtient les informations sur le socket local
	localAddr := conn.LocalAddr().(*net.UDPAddr)

	// Renvoie l'adresse IP locale au format string
	return localAddr.IP.String()
}

func multiObjectives(p *Personnage) {
	fmt.Println("[1] - Créer un salon")
	fmt.Println("[2] - Rejoindre un salon")
	fmt.Println("[3] - Retour")
	choice, _ := Inputint()
	switch choice {
	case 1:
		ClearConsole()
		fmt.Print("Nombre de joueurs : ")
		nbPlayers, _ = Inputint()
		createParty(p)
	case 2:
		ClearConsole()
		joinParty(p)
	}
}

func afficheJoueursHp() {
	ClearConsole()
	for joueur := range joueurs {
		Green.Println(joueur+" - ", joueurs[joueur], "/", joueurs[joueur])
	}
}
func afficheSalonAttente() {
	ClearConsole()
	Blue.Println("Adresse du salon : ", ip)
	fmt.Println("----- JOUEURS (", len(joueurs), "/", nbPlayers, ") -----")
	for joueur := range joueurs {
		fmt.Println(joueur)
	}
}

func createParty(p *Personnage) {

	listener, err := net.Listen("tcp", "0.0.0.0:12345")
	if err != nil {
		fmt.Println("Erreur lors de la création du serveur :", err)
		return
	}
	defer listener.Close()

	ip = LocalIP()
	joueurs[p.nom] = p.maxHP

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
			afficheJoueursHp()
		}
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
		fmt.Println(message)

	}
}

func joinParty(p *Personnage) {
	ClearConsole()
	fmt.Print("Adresse du salon : ")
	serverIP := Input()
	serverPort := 12345

	ClearConsole()

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
		pseudo, hp := sliceArgument(message)
		hpInt, _ := strconv.Atoi(hp)
		joueurs[pseudo] = hpInt
		afficheJoueursHp()

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
