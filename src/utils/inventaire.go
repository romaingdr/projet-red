// FICHIER UTILISE POUR LA GESTION DE L'INVENTAIRE (affichage, ajout, suppression)

package utils

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

// Fonction pour vérifier si un item avec un nom donné se trouve dans une liste d'items
func (p *Personnage) alreadyBuy(itemName string) bool {
	for _, item := range p.inventory {
		if item.Name == itemName {
			return true
		}
	}
	return false
}
