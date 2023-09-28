// FICHIER UTILISER POUR GERER LA BASE DE DONNEES

package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// QuickDB est la structure de la base de données
type QuickDB struct {
	filePath string
	data     map[string]interface{}
}

// NewQuickDB sert à créer une nouvelle base de données utilisable à partir d'un emplacement de fichier .json
func NewQuickDB(filePath string) *QuickDB {
	db := &QuickDB{
		filePath: filePath,
		data:     make(map[string]interface{}),
	}
	db.loadData()
	return db
}

// loadData charge les données depuis un fichier JSON dans la structure de données interne de QuickDB
func (db *QuickDB) loadData() {
	file, err := ioutil.ReadFile(db.filePath)
	if err != nil {
		return
	}

	if err := json.Unmarshal(file, &db.data); err != nil {
		fmt.Println("Error loading data:", err)
	}
}

// saveData sauvegarde les données actuelles de QuickDB dans un fichier JSON
func (db *QuickDB) saveData() {
	dataJSON, err := json.MarshalIndent(db.data, "", "  ")
	if err != nil {
		fmt.Println("Error saving data:", err)
		return
	}

	err = ioutil.WriteFile(db.filePath, dataJSON, 0644)
	if err != nil {
		fmt.Println("Error saving data to file:", err)
	}
}

// Set permet de définir une valeur associée à une clé dans QuickDB et ensuite enregistre les données mises à jour.
func (db *QuickDB) Set(key string, value interface{}) {
	db.data[key] = value
	db.saveData()
}

// Get récupère la valeur associée à une clé donnée depuis QuickDB.
func (db *QuickDB) Get(key string) interface{} {
	return db.data[key]
}
