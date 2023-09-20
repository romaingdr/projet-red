package utils

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

var (
	MaxHpEnnemy     int
	CurrentHpEnnemy int
)

func sliceArgument(phrase string) (string, string) {
	mots := strings.Split(phrase, "|")
	if len(mots) == 2 {
		return mots[0], mots[1]
	} else {
		return "", ""

	}
}

func GetLocalIP() (string, error) {
	// Crée une connexion UDP à une adresse quelconque
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	// Obtient les informations sur le socket local
	localAddr := conn.LocalAddr().(*net.UDPAddr)

	// Renvoie l'adresse IP locale au format string
	return localAddr.IP.String(), nil
}

func MultiStartScreen(p *Personnage) {
	fmt.Println("[1] Créer un serveur")
	fmt.Println("[2] Rejoindre un serveur")
	fmt.Println("[3] Retour")
	choice, _ := Inputint()
	switch choice {
	case 1:
		createServer(p)
	case 2:
		joinServer(p)
	case 3:
		p.Menu()
	default:
		ClearConsole()
		Red.Println("Veuillez saisir une donnée valide")
		MultiStartScreen(p)
	}
}

// Coté client
func joinServer(p *Personnage) {
	ClearConsole()
	fmt.Print("IP du serveur : ")
	serverIP := Input()
	serverPort := 12345

	ClearConsole()
	// Crée une connexion TCP vers le serveur
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIP, serverPort))
	if err != nil {
		fmt.Println("Erreur lors de la connexion au serveur :", err)
		return
	}
	defer conn.Close()

	// On envoie notre pseudo
	message := p.nom + "|" + strconv.Itoa(p.maxHP)
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Erreur lors de l'envoi du pseudo au serveur :", err)
		return
	}

	// Lisez des données du serveur
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Erreur lors de la lecture depuis le serveur :", err)
		return
	}

	// Traitez les données reçues du serveur ici
	message1 := string(buffer[:n])
	pseudo, hp := sliceArgument(message1)
	hpAdversaire, _ := strconv.Atoi(hp)
	CurrentHpEnnemy = hpAdversaire
	MaxHpEnnemy = hpAdversaire
	fmt.Println("Vous avez rejoint :", pseudo)
	afficheHp(p, pseudo, CurrentHpEnnemy, MaxHpEnnemy)

	// Créez une boucle pour maintenir la connexion ouverte et échanger des informations
	for {
		// Lisez des données du serveur
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Erreur lors de la lecture depuis le serveur :", err)
			return
		}

		// Traitez les données reçues du serveur ici
		messageFromServer := string(buffer[:n])
		fmt.Println("Message du serveur :", messageFromServer)

		// Vous pouvez également envoyer des données au serveur si nécessaire
		// messageToServer := "Données à envoyer"
		// _, err = conn.Write([]byte(messageToServer))
		// if err != nil {
		//     fmt.Println("Erreur lors de l'envoi de données au serveur :", err)
		//     return
		// }
	}
}

// Coté serveur
func createServer(p *Personnage) {
	// Crée un socket TCP
	listener, err := net.Listen("tcp", "0.0.0.0:12345")
	if err != nil {
		fmt.Println("Erreur lors de la création du serveur :", err)
		return
	}
	defer listener.Close()

	ip_locale, err := GetLocalIP()
	if err != nil {
		fmt.Println("Erreur lors de la récupération de l'adresse IP locale :", err)
		return
	}

	// Message d'attente
	ClearConsole()
	fmt.Println("Adresse de connexion : ", ip_locale)
	fmt.Println("En attente d'un joueur pour combattre...", ip_locale)

	for {
		// Accepte les connexions entrantes
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Erreur lors de l'acceptation d'une connexion entrante :", err)
			continue
		}

		// Traite la connexion dans une goroutine
		go handleConnection(conn, p)
	}
}

func handleConnection(conn net.Conn, p *Personnage) {
	defer conn.Close()

	// Message de connexion
	ClearConsole()
	fmt.Println("Un joueur a été trouvé !")

	// Envoi de notre pseudo au client
	message := p.nom + "|" + strconv.Itoa(p.maxHP)
	conn.Write([]byte(message))

	// Récupération du pseudo du client
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Erreur lors de la lecture depuis le client :", err)
		return
	}

	message1 := string(buffer[:n])
	pseudo, hp := sliceArgument(message1)
	hpAdversaire, _ := strconv.Atoi(hp)
	CurrentHpEnnemy = hpAdversaire
	MaxHpEnnemy = hpAdversaire
	fmt.Println("Vous jouez contre :", pseudo)
	afficheHp(p, pseudo, CurrentHpEnnemy, MaxHpEnnemy)

	// Créez une boucle pour maintenir la connexion ouverte et échanger des informations
	for {
		// Lisez des données du client
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Erreur lors de la lecture depuis le client :", err)
			return
		}

		// Traitez les données reçues du client ici
		messageFromClient := string(buffer[:n])
		fmt.Println("Message du client :", messageFromClient)

		// Vous pouvez également envoyer des données au client si nécessaire
		// messageToClient := "Données à envoyer"
		// _, err = conn.Write([]byte(messageToClient))
		// if err != nil {
		//     fmt.Println("Erreur lors de l'envoi de données au client :", err)
		//     return
		// }
	}
}

func afficheHp(p *Personnage, pseudo string, current int, max int) {
	SpeedMsg(p.nom+" - "+strconv.Itoa(p.currentHp)+"/"+strconv.Itoa(p.maxHP)+"\n", 20, "green")
	SpeedMsg(pseudo+" - "+strconv.Itoa(current)+"/"+strconv.Itoa(max)+"\n", 20, "red")
}
