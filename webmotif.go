package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

var motChoisi string
var motAffiche []string
var lettresDevinees []string
var essaisRestants int
var motifPenduASCII = []string{
	``,
	`
=========`,
	`
       |
       |
       |
       |
       |
=========`,
	`
   +---+
       |
       |
       |
       |
       |
=========`,
	`
  +---+
  |   |
      |
      |
      |
      |
=========`,
	`
  +---+
  |   |
  O   |
      |
      |
      |
=========`,
	`
  +---+
  |   |
  O   |
  |   |
      |
      |
=========`,
	`
  +---+
  |   |
  O   |
 /|   |
      |
      |
=========`,
	`
  +---+
  |   |
  O   |
 /|\  |
      |
      |
=========`,
	`
  +---+
  |   |
  O   |
 /|\  |
 /    |
      |
=========`,
	`
  +---+
  |   |
  O   |
 /|\  |
 / \  |
      |
=========`,
}

func choisirMot() string {
	rand.Seed(time.Now().UnixNano())
	mot, erreur := extraire()
	if erreur != nil {
		log.Fatal(erreur)
	}
	return mot[rand.Intn(len(mot))]
}

func afficherMot() string {
	affichage := ""
	for _, char := range motChoisi {
		if contientLettre(lettresDevinees, string(char)) {
			affichage += string(char) + " "
		} else {
			affichage += "_ "
		}
	}
	return strings.TrimSpace(affichage)
}

func contientLettre(slice []string, element string) bool {
	for _, i := range slice {
		if i == element {
			return true
		}
	}
	return false
}

func resetGame() {
	motChoisi = choisirMot()
	lettresDevinees = []string{}
	essaisRestants = len(motifPenduASCII) - 1
	reveler := make([]bool, len(motChoisi))
	// Calculer le nombre random de lettre a révélé
	numAreveler := len(motChoisi)/2 - 1
	if numAreveler < 0 {
		numAreveler = 0
	}
	// révéler x lettre aléatoire dans le mot
	for i := 0; i < numAreveler; i++ {
		randomPosition := rand.Intn(len(motChoisi))
		if !reveler[randomPosition] {
			reveler[randomPosition] = true
		} else {
			i--
		}
	}
	motAffiche = make([]string, len(motChoisi))
	for i, char := range motChoisi {
		if reveler[i] {
			motAffiche[i] = string(char)
		} else {
			motAffiche[i] = "_"
		}
	}
}

func jouerHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		devinette := strings.ToLower(r.Form.Get("lettre"))

		if len(devinette) == 1 && 'a' <= devinette[0] && devinette[0] <= 'z' {
			if !contientLettre(lettresDevinees, devinette) {
				lettresDevinees = append(lettresDevinees, devinette)

				if !strings.Contains(motChoisi, devinette) {
					essaisRestants--
				}
			}
		}
	}

	data := struct {
		Mot     string
		Lettres string
		Essais  int
		Pendu   string
	}{
		Mot:     afficherMot(),
		Lettres: strings.Join(lettresDevinees, ", "),
		Essais:  essaisRestants,
		Pendu:   motifPenduASCII[len(motifPenduASCII)-essaisRestants-1],
	}

	tmpl, err := template.ParseFiles("hangman.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func resetHandler(w http.ResponseWriter, r *http.Request) {
	resetGame()
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func main() {
	resetGame()

	http.HandleFunc("/", jouerHandler)
	http.HandleFunc("/hangman", handler)
	http.HandleFunc("/reset", resetHandler)

	fs := http.FileServer(http.Dir("statics"))
	http.Handle("/statics/", http.StripPrefix("/statics/", fs))

	port := 8080
	fmt.Printf("Serveur en cours d'exécution sur le port %d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println(err)
	}
}

func extraire() ([]string, error) {
	args := os.Args[1:]
	if len(args) == 0 {
		return nil, fmt.Errorf("fichier vide")
	}
	n := args[0]
	content, erreur := os.ReadFile(n)
	if erreur != nil {
		//si il y'a une erreur
		return nil, erreur
	}
	mots := strings.Fields(string(content))
	return mots, nil
}
