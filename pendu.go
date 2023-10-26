package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// la fonction random va prendre un mot au hasard dans le fichier words.txt
func random(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var mots []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		mots = append(mots, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	rand := mots[rand.Intn(len(mots))]

	return rand, nil
}

// la fonction pendu va spliter les differentes positions du hangman dans le fichier hangman.txt pour qu'on puisse les appeler une par une
func pendu(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		return nil
	}
	defer file.Close()

	var positions []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		positions = append(positions, scanner.Text())
	}

	pendu := strings.Split(strings.Join(positions, "\n"), "=========")

	return pendu
}

func main() {
	rand, err := random("words.txt")
	if err != nil {
		fmt.Printf("Erreur : %v\n", err)
		return
	}
	pendu := pendu("hangman.txt")
	fmt.Printf("\n")
	// ici on defini les diferents index pour les differents niveaux
	// pour le niveau 1, 1 lettre s'affiche, niveau 2, 2 lettes et comme ça jusqu'au niveau 4
	nindex1 := int(len(rand) - 1)
	nindex2 := 0
	nindex3 := int(len(rand) / 2)
	nlettre1 := string(rand[nindex1])
	nlettre2 := string(rand[nindex2])
	nlettre3 := string(rand[nindex3])
	dash := make([]string, len(rand))
	var n int
	var z string
	p := "p"
	jeu := false
	// ici on a une boucle qui ne s'arrete pas tant que l'on appuie pas sur la lettre p
	fmt.Println("Bienvenue au Jeu du Pendu")
	for !jeu {
		fmt.Print("Pour commencer appuyer sur P puis sur Entrer ")
		fmt.Scan(&z)
		z = strings.ToLower(z)
		if p == z {
			jeu = true
		} else {
			continue
		}
	}
	fmt.Println("1 : facile")
	fmt.Println("2 : moyen")
	fmt.Println("3 : difficile")
	fmt.Println("4 : extreme")
	affich := false
	// cette boucle permet d'eviter les valeurs incorrect, on ne peut passer la boucle que si la valeur entrer est compris entre 0 et 4
	for !affich {
		fmt.Print("Choisissez votre difficulté (1 à 4) :")
		fmt.Scan(&n)
		if n > 0 && n < 5 {
			affich = true
		} else {
			continue
		}
	}
	// cette boucle affiche le nombre de lettres au hasard en fonction du niveau choisi
	fmt.Print("Voici le mot à trouver: ")
	for i := 0; i < len(rand); i++ {
		if n == 1 && len(rand) > 5 {
			if i == nindex1 {
				dash[i] = nlettre1
			} else if i == nindex2 {
				dash[i] = nlettre2
			} else if i == nindex3 {
				dash[i] = nlettre3
			} else {
				dash[i] = "_"
			}
		} else if n == 2 {
			if i == nindex1 {
				dash[i] = nlettre1
			} else if i == nindex2 {
				dash[i] = nlettre2
			} else {
				dash[i] = "_"
			}
		} else if n == 3 {
			if i == nindex3 {
				dash[i] = nlettre3
			} else {
				dash[i] = "_"
			}
		} else if n == 4 {
			dash[i] = "_"
		}
	}
	fmt.Println(strings.Join(dash, " "))
	fmt.Println("Bonne chance, vous avez 10 essais")
	rest := len(pendu) - 1
	var list []string
	for {
		verif := false
		fmt.Println("Lettre déjà entrée:", list)
		fmt.Print("Entrez une lettre : ")
		var mot string
		fmt.Scan(&mot)
		// ici on passe la lettre en minuscule si jamais le joueur entre unen lettre majuscule
		mot = strings.ToLower(mot)
		list = append(list, mot) //liste qui permet de stocker les lettres entrées
		fmt.Print("\n")
		egale := false
		// ici on verifie si la lettre entrée est presente dans le mot
		for i := 0; i < len(rand); i++ {
			if mot[0] == rand[i] {
				dash[i] = mot
				verif = true
				egale = true
			}
		}
		// si la variable verif est false c'est à dire que la lettre n'est pas dans le mot on reduit le compteur
		if !verif {
			rest--
			fmt.Println("--------------------------------------------")
			fmt.Println("Essai restant :", rest)
		}
		// ici on affiche la position du hangman correspondante au nombre d'essai restant
		if rest >= 0 && rest < len(pendu) {
			if !egale {
				pendu[rest] += "========"
				fmt.Println(pendu[rest])
			}
		}
		// ici si le nombre d'essai restant est epuisé on break pour terminer le jeu et on affiche perdu
		if rest == 0 {
			fmt.Println("Vous avez perdu, le mot était :", rand)
			break
		}
		fmt.Println(strings.Join(dash, " "))
		// ici on verifie si toutes les lettres du mot à chercher on bien été trouvé si c'est le cas on print le mot est on break
		if !strings.Contains(strings.Join(dash, ""), "_") {
			fmt.Println("Bravo, le mot était", rand)
			break
		}
	}
}
