// FICHIER UTILISE POUR LA CONFIGURATION DU SOCKET TCP AINSI QUE LE COMBAT ENTRE LES JOUEURS CONNECTES

package utils

import (
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
)

// Variables globales utiles au mode PvP multijoueur
var (
	MaxHpEnnemy     int
	CurrentHpEnnemy int
	pseudoEnnemy    string
	round           string
	clientConns     []net.Conn
	poisonEnnemy    = 0
	poisonUsed      = false
)

// sliceArgument sert à split les arguments reçus par le socket et séparés par des "|"
func sliceArgument(phrase string) (string, string) {
	mots := strings.Split(phrase, "|")
	if len(mots) == 2 {
		return mots[0], mots[1]
	} else {
		return "", ""

	}
}

// GetLocalIP sert à récuperer l'ip locale de la machine host
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

// MultiStartScreen sert à afficher le menu du mode Multijoueur PvP afin de Créer/Rejoindre un serveur
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
		ClearConsole()
		p.Menu()
	default:
		ClearConsole()
		Red.Println("Veuillez saisir une donnée valide")
		MultiStartScreen(p)
	}
}

// COTE CLIENT
// joinServer sert à rejoindre un serveur avec l'ip locale de l'ordinateur host , connexion sur le port 12345
func joinServer(p *Personnage) {
	ClearConsole()
	fmt.Print("IP du serveur : ")
	serverIP := Input()
	serverPort := 12345

	ClearConsole()

	// connexion tcp au serveur
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIP, serverPort))
	if err != nil {
		Red.Println("Impossible de trouver le serveur !")
		MultiStartScreen(p)
	}
	defer conn.Close()

	clientConns = append(clientConns, conn)

	round = "server"
	// On envoie notre pseudo et nos hp
	message := p.nom + "|" + strconv.Itoa(p.maxHP)
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Erreur lors de l'envoi du pseudo au serveur :", err)
		return
	}

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
	pseudoEnnemy = pseudo
	CurrentHpEnnemy = hpAdversaire
	MaxHpEnnemy = hpAdversaire
	fmt.Println("Vous avez rejoint :", pseudo)
	fmt.Println()

	for i := 10; i > 0; i-- {
		fmt.Printf("\rLa partie commence dans %d secondes", i)
		time.Sleep(1 * time.Second)
	}

	ClearConsole()
	afficheMenuClient(p)

	for {
		// Si une attaque est reçue
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Erreur lors de la lecture depuis le serveur :", err)
			return
		}

		// Traiter l'attaque
		attaque := string(buffer[:n])
		degatsRecu, crit := sliceArgument(attaque)
		degats, _ := strconv.Atoi(degatsRecu)
		if crit == "true" {
			SpeedMsg("[COUP CRITIQUE] "+pseudoEnnemy+" vous inflige "+degatsRecu+" dégats\n", 20, "red")
		} else {
			SpeedMsg(pseudoEnnemy+" vous inflige "+degatsRecu+" dégats\n", 20, "red")
		}
		p.currentHp -= degats
		fmt.Println("--------------------")
		round = "client"
		time.Sleep(4 * time.Second)
		ClearConsole()
		afficheMenuClient(p)

	}
}

// COTE SERVEUR
// createServer sert à créer un serveur sur le port 12345 de la machine host
func createServer(p *Personnage) {
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
	fmt.Println("En attente d'un joueur pour combattre...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Erreur lors de l'acceptation d'une connexion entrante :", err)
			continue
		}
		go handleConnection(conn, p)
	}
}

// handleConnection sert à maintenir la connexion sur le serveur et gérer les sockets d'informations entrantes et sortantes
func handleConnection(conn net.Conn, p *Personnage) {
	defer conn.Close()

	// Message de connexion
	ClearConsole()
	fmt.Println("Un joueur a été trouvé !")
	round = "server"
	clientConns = append(clientConns, conn)

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
	pseudoEnnemy = pseudo
	hpAdversaire, _ := strconv.Atoi(hp)
	CurrentHpEnnemy = hpAdversaire
	MaxHpEnnemy = hpAdversaire
	fmt.Println("Vous jouez contre :", pseudo)
	fmt.Println()

	for i := 10; i > 0; i-- {
		fmt.Printf("\rLa partie commence dans %d secondes", i)
		time.Sleep(1 * time.Second)
	}

	ClearConsole()
	afficheMenuServer(p)

	for {

		// Une attaque est reçue
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Erreur lors de la lecture depuis le client :", err)
			return
		}

		// Gérer l'attaque
		attaque := string(buffer[:n])
		degatsRecu, crit := sliceArgument(attaque)
		degats, _ := strconv.Atoi(degatsRecu)
		if crit == "true" {
			SpeedMsg("[COUP CRITIQUE] "+pseudoEnnemy+" vous inflige "+degatsRecu+" dégats\n", 20, "red")
		} else {
			SpeedMsg(pseudoEnnemy+" vous inflige "+degatsRecu+" dégats\n", 20, "red")
		}
		p.currentHp -= degats
		fmt.Println("--------------------")
		round = "server"
		time.Sleep(4 * time.Second)
		ClearConsole()
		afficheMenuServer(p)

	}
}

// afficheHp sert à afficher progressivement les hp du joueur et de l'ennemi
func afficheHp(p *Personnage, pseudo string, current int, max int) {
	SpeedMsg(p.nom+" - "+strconv.Itoa(p.currentHp)+"/"+strconv.Itoa(p.maxHP)+"\n", 20, "green")
	SpeedMsg(pseudo+" - "+strconv.Itoa(current)+"/"+strconv.Itoa(max)+"\n", 20, "red")
}

// afficheHpNoDelay sert à afficher les hp du joueur et de l'ennemi sans délai d'écriture
func afficheHpNoDelay(p *Personnage, pseudo string, current int, max int) {
	Green.Println(p.nom + " - " + strconv.Itoa(p.currentHp) + "/" + strconv.Itoa(p.maxHP))
	Red.Println(pseudo + " - " + strconv.Itoa(current) + "/" + strconv.Itoa(max))
}

// afficheMenuServer sert à afficher le menu de combat côté serveur
func afficheMenuServer(p *Personnage) {
	if !(isDead(p)) {
		if round == "server" {
			afficheHp(p, pseudoEnnemy, CurrentHpEnnemy, MaxHpEnnemy)
			fmt.Println("----- A votre tour -----")
			fmt.Println("[1] Attaque auto")
			fmt.Println("[2] Abilités")

			fmt.Println("------------------------")
			choice, _ := Inputint()
			switch choice {

			// AUTO ATTAQUE
			case 1:
				degats := p.skill[0].Damages
				critic := rand.Intn(100) + 1
				if critic <= p.skill[4].Damages {
					degats *= 2
					critic = 1
				}
				if critic == 1 {
					if poisonEnnemy > 0 {
						fmt.Println("[POISON] Vous infligez 10 dégats supplémentaires")
						degats += 10
						poisonEnnemy -= 1
						poisonUsed = true
					} else {
						poisonUsed = false
					}
					sendMessageToClient(clientConns, strconv.Itoa(degats)+"|true")
				} else {
					if poisonEnnemy > 0 {
						fmt.Println("[POISON] Vous infligez 10 dégats supplémentaires")
						degats += 10
						poisonEnnemy -= 1
						poisonUsed = true
					} else {
						poisonUsed = false
					}
					sendMessageToClient(clientConns, strconv.Itoa(degats)+"|false")
				}

				ClearConsole()
				afficheHp(p, pseudoEnnemy, CurrentHpEnnemy, MaxHpEnnemy)
				fmt.Println("----- A votre tour -----")
				SpeedMsg("Vous utilisez une attaque automatique\n", 20, "default")
				if poisonUsed {
					fmt.Println("[POISON] Vous infligez 10 dégats supplémentaires")
				}
				if critic == 1 {
					SpeedMsg("[COUP CRITIQUE] "+strconv.Itoa(degats)+" infligés à "+pseudoEnnemy+"\n", 20, "green")
				} else {
					SpeedMsg(strconv.Itoa(degats)+" infligés à "+pseudoEnnemy+"\n", 20, "green")
				}
				fmt.Println("------------------------")
				time.Sleep(1700 * time.Millisecond)
				CurrentHpEnnemy -= degats
				round = "client"
				afficheMenuServer(p)

			// Spells
			case 2:
				// Affichage des spells
				ClearConsole()
				afficheHpNoDelay(p, pseudoEnnemy, CurrentHpEnnemy, MaxHpEnnemy)
				fmt.Println("----- A votre tour -----")

				for i := 1; i < 4; i++ {
					skill := p.skill[i]
					fmt.Printf("[%d] %-20s %-10d %d/%d\n", i, skill.Name, skill.Damages, skill.StillUse, skill.MaxUse)
				}

				fmt.Println("------------------------")
				fmt.Println("[4] Sortir")

				choice, _ := Inputint()
				var degats int

				// Calcul si le coup va être un crit (en fonction du pourcentage de chance de crit ==> p.skill[4].Damages)
				rand.Seed(time.Now().UnixNano())
				crit := rand.Intn(100) + 1
				critBool := crit <= p.skill[4].Damages

				switch choice {
				case 1, 2, 3:
					skill := p.skill[choice]
					if skill.StillUse > 0 {
						// Affichage du début
						ClearConsole()
						afficheHp(p, pseudoEnnemy, CurrentHpEnnemy, MaxHpEnnemy)
						fmt.Println("----- A votre tour -----")
						SpeedMsg("Vous utilisez "+skill.Name+"\n", 20, "default")

						// La compétence est utilisée
						degats = skill.Damages
						p.skill[choice].StillUse -= 1

						// Effets du spell

						if choice == 1 { // Spell 1

							if p.classe == "Arcaniste" { // Execution en dessous de 10%
								if CurrentHpEnnemy <= MaxHpEnnemy/10 {
									CurrentHpEnnemy = 0
									degats = CurrentHpEnnemy
								}

							} else if p.classe == "Chasseur" { // Poison de 10 dégats
								poisonEnnemy = 3

							} else if p.classe == "Titan" { // Réduction des dégats
								p.skill[4].Damages += 10
								Green.Println("[BULLE] Vous obtenez 10% de chance de coup critique")
							}

						} else if choice == 2 {

							if p.classe == "Chasseur" { // Damage reduce  + 20% auto
								p.skill[0].Damages += p.skill[0].Damages / 5
								Green.Println("[MAITRISE DU TERRAIN] Votre attaque automatique inflige 20% de dégats supplémentaires")
							}

						} else if choice == 3 { // Amélioration des dégats du spell 2

							if p.classe == "Arcaniste" {
								p.skill[2].Damages += p.skill[2].Damages / 2
								Green.Println("[FOUDRE] Votre altération de l'âme inflige 50% de dégats supplémentaires")
							}
						}

						// Envoi des degats
						if critBool {
							degats *= 2
							if poisonEnnemy > 0 {
								fmt.Println("[POISON] Vous infligez 10 dégats supplémentaires")
								degats += 10
								poisonEnnemy -= 1
							}
							sendMessageToClient(clientConns, strconv.Itoa(degats)+"|true")
						} else {
							if poisonEnnemy > 0 {
								fmt.Println("[POISON] Vous infligez 10 dégats supplémentaires")
								degats += 10
								poisonEnnemy -= 1
							}
							sendMessageToClient(clientConns, strconv.Itoa(degats)+"|false")
						}

						// Affichage des degats
						if critBool {
							SpeedMsg("[COUP CRITIQUE] "+strconv.Itoa(degats)+" infligés à "+pseudoEnnemy+"\n", 20, "green")
						} else {
							SpeedMsg(strconv.Itoa(degats)+" infligés à "+pseudoEnnemy+"\n", 20, "green")
						}

						fmt.Println("------------------------")
						time.Sleep(1700 * time.Millisecond)
						CurrentHpEnnemy -= degats
						round = "client"
						afficheMenuServer(p)

					} else {
						ClearConsole()
						Red.Println("Vous ne pouvez plus utiliser cette compétence !")
						afficheMenuServer(p)
					}

				default:
					ClearConsole()
					afficheMenuServer(p)
				}

			// MAUVAIS INPUT
			default:
				ClearConsole()
				Red.Println("Veuillez saisir une donnée valide")
				afficheMenuServer(p)

			}
		} else {
			ClearConsole()
			afficheHp(p, pseudoEnnemy, CurrentHpEnnemy, MaxHpEnnemy)
			fmt.Println("----- Tour de " + pseudoEnnemy + " -----")
			SpeedMsg(pseudoEnnemy+" attaque...\n", 20, "default")
		}
	} else {
		if p.currentHp <= 0 {
			ClearConsole()
			SpeedMsg("Vainqueur : "+pseudoEnnemy+"\n", 20, "default")
			SpeedMsg("Vous avez perdu le combat !\n", 20, "red")
			fmt.Println()
			fmt.Print("retourner au menu")
			Input()
			ClearConsole()
			p.Menu()
		} else {
			ClearConsole()
			SpeedMsg("Vainqueur : "+p.nom+"\n", 20, "default")
			SpeedMsg("Vous avez gagné le combat !\n", 20, "green")
			fmt.Println()
			fmt.Print("retourner au menu")
			Input()
			ClearConsole()
			p.Menu()
		}
	}
}

// afficheMenuServer sert à afficher le menu de combat côté client
func afficheMenuClient(p *Personnage) {
	if !(isDead(p)) {

		if round == "server" {
			ClearConsole()
			afficheHp(p, pseudoEnnemy, CurrentHpEnnemy, MaxHpEnnemy)
			fmt.Println("----- Tour de " + pseudoEnnemy + " -----")
			SpeedMsg(pseudoEnnemy+" attaque...\n", 20, "default")
		} else {
			ClearConsole()
			afficheHp(p, pseudoEnnemy, CurrentHpEnnemy, MaxHpEnnemy)
			fmt.Println("----- A votre tour -----")
			fmt.Println("[1] Attaque auto")
			fmt.Println("[2] Abilités")

			fmt.Println("------------------------")
			choice, _ := Inputint()
			switch choice {
			case 1:
				degats := p.skill[0].Damages
				critic := rand.Intn(100) + 1
				if critic <= p.skill[4].Damages {
					degats *= 2
					critic = 1
				}
				if critic == 1 {
					if poisonEnnemy > 0 {
						fmt.Println("[POISON] Vous infligez 10 dégats supplémentaires")
						degats += 10
						poisonEnnemy -= 1
						poisonUsed = true
					} else {
						poisonUsed = false
					}
					sendMessageToServer(clientConns, strconv.Itoa(degats)+"|true")
				} else {
					if poisonEnnemy > 0 {
						fmt.Println("[POISON] Vous infligez 10 dégats supplémentaires")
						degats += 10
						poisonEnnemy -= 1
						poisonUsed = true
					} else {
						poisonUsed = false
					}
					sendMessageToServer(clientConns, strconv.Itoa(degats)+"|false")
				}
				ClearConsole()
				afficheHp(p, pseudoEnnemy, CurrentHpEnnemy, MaxHpEnnemy)
				fmt.Println("----- A votre tour -----")
				SpeedMsg("Vous utilisez une attaque automatique\n", 20, "default")
				if poisonUsed {
					fmt.Println("[POISON] Vouq infligez 10 dégats supplémentaires")
				}
				if critic == 1 {
					SpeedMsg("[COUP CRITIQUE] "+strconv.Itoa(degats)+" infligés à "+pseudoEnnemy+"\n", 20, "green")
				} else {
					SpeedMsg(strconv.Itoa(degats)+" infligés à "+pseudoEnnemy+"\n", 20, "green")
				}
				fmt.Println("------------------------")
				time.Sleep(1700 * time.Millisecond)
				CurrentHpEnnemy -= degats
				round = "server"
				ClearConsole()
				afficheMenuClient(p)

			// Spells
			case 2:
				// Affichage des spells
				ClearConsole()
				afficheHpNoDelay(p, pseudoEnnemy, CurrentHpEnnemy, MaxHpEnnemy)
				fmt.Println("----- A votre tour -----")

				for i := 1; i < 4; i++ {
					skill := p.skill[i]
					fmt.Printf("[%d] %-20s %-10d %d/%d\n", i, skill.Name, skill.Damages, skill.StillUse, skill.MaxUse)
				}

				fmt.Println("------------------------")
				fmt.Println("[4] Sortir")

				choice, _ := Inputint()
				var degats int

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
						p.skill[choice].StillUse -= 1
						ClearConsole()
						afficheHp(p, pseudoEnnemy, CurrentHpEnnemy, MaxHpEnnemy)
						fmt.Println("----- A votre tour -----")
						SpeedMsg("Vous utilisez "+skill.Name+"\n", 20, "default")

						// Effets du spell

						if choice == 1 { // Spell 1

							if p.classe == "Arcaniste" { // Execution en dessous de 10%
								if CurrentHpEnnemy <= MaxHpEnnemy/10 {
									CurrentHpEnnemy = 0
									degats = CurrentHpEnnemy
								}

							} else if p.classe == "Chasseur" { // Poison de 10 dégats
								poisonEnnemy = 3

							} else if p.classe == "Titan" { // Réduction des dégats
								p.skill[4].Damages += 10
								Green.Println("[BULLE] Vous obtenez 10% de chance de coup critique")
							}

						} else if choice == 2 {

							if p.classe == "Chasseur" { // Damage reduce  + 20% auto
								p.skill[0].Damages += p.skill[0].Damages / 5
								Green.Println("[MAITRISE DU TERRAIN] Votre attaque automatique inflige 20% de dégats supplémentaires")
							}

						} else if choice == 3 { // Amélioration des dégats du spell 2

							if p.classe == "Arcaniste" {
								p.skill[2].Damages += p.skill[2].Damages / 2
								Green.Println("[FOUDRE] Votre altération de l'âme inflige 50% de dégats supplémentaires")
							}
						}

						// Envoi des degats
						if critBool {
							degats *= 2
							if poisonEnnemy > 0 {
								fmt.Println("[POISON] Vous infligez 10 dégats supplémentaires")
								degats += 10
								poisonEnnemy -= 1
							}
							sendMessageToServer(clientConns, strconv.Itoa(degats)+"|true")
						} else {
							if poisonEnnemy > 0 {
								fmt.Println("[POISON] Vous infligez 10 dégats supplémentaires")
								degats += 10
								poisonEnnemy -= 1
							}
							sendMessageToServer(clientConns, strconv.Itoa(degats)+"|false")
						}

						// Affichage des degats

						if critBool {
							SpeedMsg("[COUP CRITIQUE] "+strconv.Itoa(degats)+" infligés à "+pseudoEnnemy+"\n", 20, "green")
						} else {
							SpeedMsg(strconv.Itoa(degats)+" infligés à "+pseudoEnnemy+"\n", 20, "green")
						}
						fmt.Println("------------------------")
						time.Sleep(1700 * time.Millisecond)
						CurrentHpEnnemy -= degats
						round = "server"
						afficheMenuClient(p)

					} else {
						ClearConsole()
						Red.Println("Vous ne pouvez plus utiliser cette compétence !")
						afficheMenuClient(p)
					}
				default:
					ClearConsole()
					afficheMenuServer(p)
				}
			default:
				ClearConsole()
				Red.Println("Veuillez saisir une donnée valide")
				afficheMenuClient(p)
			}
		}
	} else {
		if p.currentHp <= 0 {
			ClearConsole()
			SpeedMsg("Vainqueur : "+pseudoEnnemy+"\n", 20, "default")
			SpeedMsg("Vous avez perdu le combat !\n", 20, "red")
			fmt.Println()
			fmt.Print("retourner au menu")
			Input()
			ClearConsole()
			p.Menu()
		} else {
			ClearConsole()
			SpeedMsg("Vainqueur : "+p.nom+"\n", 20, "default")
			SpeedMsg("Vous avez gagné le combat !\n", 20, "green")
			fmt.Println()
			fmt.Print("retourner au menu")
			Input()
			ClearConsole()
			p.Menu()
		}
	}
}

// sendMessageToServer sert à envoyer un message au serveur depuis le client
func sendMessageToServer(clientConns []net.Conn, message string) {
	for _, clientConn := range clientConns {
		_, err := clientConn.Write([]byte(message))
		if err != nil {
			fmt.Println("Erreur lors de l'envoi de données au client :", err)
			// Gérez l'erreur de manière appropriée
		}
	}
}

// sendMessageToServer sert à envoyer un message au client depuis le serveur
func sendMessageToClient(clientConns []net.Conn, message string) {
	for _, clientConn := range clientConns {
		_, err := clientConn.Write([]byte(message))
		if err != nil {
			fmt.Println("Erreur lors de l'envoi de données au client :", err)
			// Gérez l'erreur de manière appropriée
		}
	}
}

// isDead sert à vérifier si un des deux joueurs est mort
func isDead(p *Personnage) bool {
	if p.currentHp <= 0 || CurrentHpEnnemy <= 0 {
		return true
	} else {
		return false
	}
}
