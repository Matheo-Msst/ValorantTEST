package server

import (
	"ValorantApi/code/test"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func Server() {
	const port = ":1610"
	// Servir le dossier statique
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Routes du serveur
	http.HandleFunc("/", AccueilHandler)
	// Générer les routes pour chaque personnage
	GenererRoutesPersonnages()

	fmt.Println("Serveur démarré sur http://localhost:1610")
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("Erreur lors du démarrage du serveur : %v\n", err)
	}
}

func AfficherTemplate(w http.ResponseWriter, tmpl string, indexData *test.INDEX) {
	// Obtenir le chemin absolu du répertoire de travail actuel
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Erreur lors de la récupération du répertoire de travail: %v\n", err)
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
		return
	}

	// Construire le chemin absolu vers le fichier du template
	// Remarque : Le dossier "html" est à la racine et non dans "ValorantApi"
	tmplPath := filepath.Join(cwd, "html", fmt.Sprintf("%s.html", tmpl))

	// Charger le template
	t, err := template.ParseFiles(tmplPath)
	if err != nil {
		fmt.Printf("Erreur de chargement du template: %v\n", err)
		http.Error(w, fmt.Sprintf("Erreur lors du chargement du template: %s", err), http.StatusInternalServerError)
		return
	}

	// Exécuter le template
	err = t.Execute(w, indexData)
	if err != nil {
		fmt.Printf("Erreur d'exécution du template: %v\n", err)
		http.Error(w, fmt.Sprintf("Erreur lors de l'exécution du template: %s", err), http.StatusInternalServerError)
	}
}

// Handler pour la page d'accueil
func AccueilHandler(w http.ResponseWriter, r *http.Request) {
	// Appeler la fonction pour récupérer les données des personnages
	indexData := test.ChargerPersonnages()

	// Afficher le template avec les données des personnages
	AfficherTemplate(w, "ValorantAccueil", indexData)
}

// Fonction pour afficher un template pour un personnage individuel
func AfficherPersonnage(w http.ResponseWriter, r *http.Request) {
	// Extraire le nom du personnage depuis l'URL (en enlevant "/personnage/")
	personnageNom := r.URL.Path[len("/personnage/"):]

	// Charger les personnages depuis les données JSON
	indexData := test.ChargerPersonnages()

	// Chercher le personnage correspondant dans la liste
	var personnage *test.Personnage
	for _, perso := range indexData.Personnages {
		if perso.Nom == personnageNom {
			personnage = &perso
			break
		}
	}

	// Si le personnage n'est pas trouvé, retourner une erreur 404
	if personnage == nil {
		http.NotFound(w, r)
		return
	}

	// Charger le template pour afficher le personnage
	tmplPath := "html/Personnages_template.html" // Chemin vers le template de base
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		log.Fatalf("Erreur lors du chargement du template: %v", err)
	}

	// Exécuter le template avec les données du personnage
	err = tmpl.Execute(w, personnage)
	if err != nil {
		log.Printf("Erreur lors de l'exécution du template pour %s: %v", personnageNom, err)
		http.Error(w, "Erreur lors de l'affichage du personnage", http.StatusInternalServerError)
	}
}

// Fonction pour générer les routes pour chaque personnage
func GenererRoutesPersonnages() {
	// Charger les personnages depuis les données JSON
	indexData := test.ChargerPersonnages()

	// Parcourir tous les personnages et générer une route pour chacun
	for _, perso := range indexData.Personnages {
		// Pour chaque personnage, créer une route dynamique /personnage/{Nom}
		// Exemple : /personnage/Jett pour le personnage Jett
		http.HandleFunc("/personnage/"+perso.Nom, AfficherPersonnage)
		fmt.Printf("Route générée pour: /personnage/%s\n", perso.Nom)
	}
}
