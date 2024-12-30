package server

import (
	"ValorantApi/code/test"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
)

// Générer un fichier HTML pour chaque personnage
func GenererPagesPersonnages() {
	// Charger les personnages depuis le fichier JSON
	indexData := test.ChargerPersonnages()

	// Charger le template de base
	tmplPath := "html/Personnages_template.html"

	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		log.Fatalf("Erreur lors du chargement du template: %v", err)
	}

	// Créer un dossier pour les pages si il n'existe pas déjà
	outputDir := "pages-personnages"
	err = os.MkdirAll(outputDir, 0755)
	if err != nil {
		log.Fatalf("Erreur lors de la création du dossier de sortie: %v", err)
	}

	// Parcourir chaque personnage et créer un fichier HTML pour chacun
	for _, perso := range indexData.Personnages {
		// Créer un fichier HTML pour ce personnage
		fileName := filepath.Join(outputDir, fmt.Sprintf("%s.html", perso.Nom))
		file, err := os.Create(fileName)
		if err != nil {
			log.Printf("Erreur lors de la création du fichier %s: %v", fileName, err)
			continue
		}
		defer file.Close()

		// Exécuter le template et écrire le contenu dans le fichier
		err = tmpl.Execute(file, perso)
		if err != nil {
			log.Printf("Erreur lors de l'exécution du template pour %s: %v", perso.Nom, err)
			continue
		}

		// Confirmation que la page a été générée
		fmt.Printf("Page générée pour %s : %s\n", perso.Nom, fileName)
	}
}
