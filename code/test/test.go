package test

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Personnage struct {
	Nom         string `json:"nom"`
	Description string `json:"description"`
	ImageAcc    string `json:"image"`
	ImagePage   string `json:"image-2"`
	Role        string `json:"role"`
}

type INDEX struct {
	Personnages []Personnage `json:"personnages"`
}

func RajoutDonneesJson() {
	// Exemple de tableau de personnages avec description et image
	personnages := []Personnage{
		{"Brimstone", `Tout droit venu des États-Unis d'Amérique, Brimstone possède un arsenal orbital qui permet à son escouade de toujours avoir l'avantage. La précision et la fiabilité de ses compétences utilitaires font de lui un commandant sans égal sur le terrain.`, "1-Brimstone.png", "1_Brimstone.png", "Controleur"},
		{"Pheonix", `En provenance du Royaume-Uni, Phoenix illumine le champ de bataille avec ses pouvoirs astraux et son style de combat flamboyant. Peu importe que les renforts arrivent ou non, il fonce au combat quand il le décide.`, "2-Phoenix.png", "2_Phoenix.png", "Duelliste"},
		{"Sage", "Véritable pilier originaire de Chine, Sage assure sa sécurité et celle de son équipe où qu'elle aille. Elle peut réanimer ses alliés tombés au combat et repousser les assauts ennemis pour offrir des oasis de tranquillité sur un champ de bataille infernal.", "3-Sage.png", "3_Sage.png", "Sentinelle"},
		{"Sova", "Né dans l'hiver éternel de la toundra russe, Sova traque, trouve et élimine ses ennemis avec une efficacité et une précision redoutables. Ses incroyables talents d'éclaireur et son arc personnalisé lui garantissent que sa cible ne fuira jamais très longtemps.", "4-Sova.png", "4_Sova.png", "Initiateur"},
		{"Viper", "Viper est une chimiste américaine qui déploie un arsenal d'appareils toxiques pour contrôler le champ de bataille et entraver la vision des ennemis. Si les toxines ne suffisent pas à abattre sa proie, ses machinations finiront le travail.", "5-Viper.png", "5-Viper.png", "Controleur"},
		{"Cypher", "Informateur originaire du Maroc, Cypher est un véritable réseau de surveillance à lui tout seul. Il révèle tous les secrets. Il détecte toutes les manœuvres. Rien n'échappe à Cypher.", "6-Cypher.png", "6_Cypher.png", "Sentinelle"},
		{"Reyna", "Originaire du coeur du Mexique, Reyna est une experte des combats singuliers qui se renforce à chaque élimination qu'elle réussit. Son efficacité n'est limitée que par son habileté, ce qui la rend très dépendante de ses propres performances.", "7-Reyna.png", "7_Reyna.png", "Duelliste"},
		{"Killjoy", "Véritable génie originaire d'Allemagne, Killjoy sécurise et défend les positions clés sans effort grâce à son arsenal d'inventions. Si son équipement ne suffit pas à arrêter l'ennemi, ce sont les entraves de ses robots qui en feront du menu fretin.", "8-Killjoy.png", "8_Killjoy.png", "Sentinelle"},
		{"Breach", "Breach, le Suédois bionique, tire de puissantes décharges cinétiques pour ouvrir un chemin en territoire ennemi. Grâce aux dégâts et aux diversions ainsi provoqués, aucun combat n'est jamais en sa défaveur.", "9-Breach.png", "9_Breach.png", "Initiateur"},
		{"Omen", "Véritable fantôme d'un souvenir, Omen chasse dans les ténèbres. Il aveugle les ennemis, se téléporte d'un bout à l'autre du champ de bataille et laisse la peur se répandre parmi ses adversaires qui se demandent qui sera sa prochaine victime.", "10-Omen.png", "10_Omen.png", "Controleur"},
		{"Jett", "Représentante de sa patrie, la Corée du Sud, Jett dispose d'un style de combat basé sur l'agilité et l'esquive, qui lui permet de prendre des risques qu'elle seule peut se permettre de prendre. Elle tourne autour des affrontements et découpe ses ennemis avant même qu'ils ne s'en rendent compte.", "11-Jett.png", "11_Jett.png", "Dueliste"},
		{"Raze", "Armée de sa personnalité et de sa grosse artillerie, Raze fait une entrée explosive depuis le Brésil. Grâce à sa force brute, elle excelle à débusquer les ennemis retranchés et à faire le ménage dans les espaces étroits, le tout avec une bonne dose de « boum ».", "12-Raze.png", "12_Raze.png", "Dueliste"},
		{"Skye", "Originaire d'Australie, Skye et sa bande de bêtes sauvages ouvrent la voie à travers les territoires hostiles. Grâce à ses créations qui entravent l'ennemi et à sa faculté à soigner les autres, l'équipe est plus forte et plus en sécurité quand elle compte Skye dans ses rangs.", "13-Skye.png", "13_Skye.png", "Initiateur"},
		{"Yoru", "Le Japonais Yoru perce des trous dans la réalité pour s'infiltrer derrière les lignes ennemies sans se faire repérer. En faisant preuve d'autant de ruse que d'agressivité, il prend ses cibles par surprise avant qu'elles n'aient le temps de se retourner.", "14-Yoru.png", "14_Yoru.png", "Duelliste"},
		{"Astra", "L'agent ghanéen Astra canalise les énergies du cosmos pour façonner le champ de bataille à sa convenance. Avec une maîtrise complète de sa forme astrale et un talent pour la planification stratégique, elle a toujours une large avance sur ses ennemis.", "15-Astra.png", "15_Astra.png", "Controleur"},
		{"Kayo", "KAY/O est une machine de guerre conçue dans un but précis : neutraliser les radiants. La neutralisation des compétences ennemies réduit les possibilités de riposte des adversaires, ce qui confère un avantage décisif à son équipe.", "16-Kayo.png", "16_Kayo.png", "Initiateur"},
		{"Chamber", "Aussi classe que bien équipé, le concepteur d'armes français Chamber repousse les assaillants avec une précision mortelle. Il met à profit son arsenal bien particulier pour tenir sa position et éliminer les ennemis de loin en prévoyant une solution aux défis posés par chaque stratégie.", "17-Chamber.png", "17_Chamber.png", "Sentinelles"},
		{"Neon", "L'agent philippin, Neon, s'élance vers l'avant à une vitesse fulgurante, libérant de grosses décharges de radiance biomagnétique générées frénétiquement par son corps. Elle se lance à la poursuite des ennemis qui n'ont pas le temps de s'y préparer et les élimine aussi vite que l'éclair.", "18-Neon.png", "18_Neon.png", "Dueliste"},
		{"Fade", "Originaire de Turquie, la chasseuse de primes Fade utilise le pouvoir des cauchemars pour s'emparer des secrets ennemis. Elle traque ses cibles et révèle leurs plus grandes peurs pour mieux les briser dans l'obscurité.", "19-Fade.png", "19_Fade.png", "Initiateur"},
		{"Harbor", "Venu de la côte indienne, Harbor déferle sur le terrain grâce à une technologie antique qui lui permet de contrôler l'eau. Il déchaîne des torrents bouillonnants et de terribles lames d'eau pour protéger ses alliés et noyer ses adversaires.", "20-Harbor.png", "20_Harbor.png", "Controleur"},
		{"Gekko", "Originaire de Los Angeles, Gekko dirige une bande de créatures chaotiques, mais très soudées. Ses amis s'occupent de disperser les ennemis, puis Gekko rassemble sa petite troupe pour recommencer.", "21-Gekko.png", "21_Gekko.png", "Initiateur"},
		{"Deadlock", "Deadlock, l'agent spécial norvégien, déploie un éventail de nanocâbles ultra-modernes pour défendre sa position contre le plus violent des assauts. Nul n'échappe à sa vigilance, ni ne survit à sa cruelle ténacité.", "22-Deadlock.png", "22_Deadlock.png", "Sentinelle"},
		{"Iso", "Venu de Chine et spécialiste de missions délicates, Iso se fond dans le flux de Radianite pour démanteler les réseaux ennemis. Capable de restructurer l'énergie ambiante en protection anti-balles, il ne dévie pas de la route vers son prochain duel à mort.", "23-Iso.png", "23_Iso.png", "Dueliste"},
		{"Clove", "Clove, l'intenable Écossais·e, déstabilise l'ennemi dans le feu de l'action comme le froid de la mort. Iel sème le trouble dans les rangs ennemis, même depuis sa tombe, car sa mort ne dure jamais très longtemps.", "24-Clove.png", "24_Clove.png", "Controleur"},
		{"Vyse", "La génie du métal Vyse utilise le métal liquide pour isoler, piéger et désarmer ses ennemis. Par la ruse et la manipulation, elle force tous ses adversaires à craindre le champ de bataille lui-même.", "25-Vyse.png", "25_Vyse.png", "Sentinelle"},
	}

	// Création ou ouverture du fichier "personnages.json" en mode écriture
	file, err := os.Create("static/personnages.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // Assurez-vous de fermer le fichier une fois terminé

	// Conversion des données en JSON
	jsonData, err := json.MarshalIndent(personnages, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	// Écriture des données JSON dans le fichier
	_, err = file.Write(jsonData)
	if err != nil {
		log.Fatal(err)
	}

}

// Fonction pour charger les données JSON et retourner l'index des personnages
func ChargerPersonnages() *INDEX {
	// Ouvrir le fichier "personnages.json" en mode lecture
	file, err := os.Open("static/personnages.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // Assurez-vous de fermer le fichier une fois terminé

	// Déclarer une variable pour stocker les personnages
	var personnages []Personnage

	// Lire et décoder les données JSON du fichier dans la variable 'personnages'
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&personnages)
	if err != nil {
		log.Fatal(err)
	}

	// Retourner l'index des personnages
	return &INDEX{
		Personnages: personnages,
	}
}

// Fonction pour afficher les personnages dans la console (à des fins de débogage)
func AfficherPersonnages() {
	// Charger les personnages
	indexData := ChargerPersonnages()

	// Affichage des personnages depuis l'index
	for i, perso := range indexData.Personnages {
		fmt.Println("/////////////////////////////////////////////////////////////////////////////////////////")
		fmt.Printf("Numéro du personnage : %d\n", i+1)
		fmt.Printf("Nom : %s\n", perso.Nom)
		fmt.Printf("Description : %s\n", perso.Description)
		fmt.Println("Image :  ", "static/img/personnages/", perso.ImageAcc)
		fmt.Println("Image : ", "static/img/personnages/", perso.ImagePage)
		fmt.Printf("Rôle : %s\n\n", perso.Role)
	}
}

func main() {
	AfficherPersonnages() // Affiche les personnages pour déboguer
}
