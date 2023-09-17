# Readme.md - projet-red

## Description

Ce programme est un jeu de rôle textuel simple dans lequel vous créez un personnage et interagissez avec le monde du jeu en utilisant des menus textuels. Le jeu vous permet de choisir le nom, la classe et d'autres attributs de votre personnage, puis de naviguer à travers différents menus pour gérer votre inventaire, acheter des objets chez un marchand, afficher les compétences de votre personnage, etc.

## Fonctionnalités principales

- Créez un personnage en choisissant son nom et sa classe parmi Titan, Arcaniste et Chasseur.
- Gérez l'inventaire de votre personnage, y compris l'utilisation de potions de guérison.
- Visitez un marchand pour acheter des objets.
- Affichez les compétences de votre personnage.
- Interagissez avec le jeu à travers des menus textuels.

## Dernier update

- Création d'une structure Item pour une implémentation plus simple
- Inventaire : Changement du type map[string]int en type []Item
- Marchand : Création d'une liste d'item (items_marchand -> []item ) à la vente pour ne pas les écrire un par un
- Tutoriel : Ajout du tutoriel complet de combat + passage niveau 2 pour le vrai combat
- Marchand / Inventaire : Changement de l'affichage (+ clair)
- Créations du package "utils" pour importer certaines fonctions

## Auteur(s)

romaingdr && avvrt
