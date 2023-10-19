package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

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
	fmt.Println("Bienvenue au Jeu du Pendu")
	for !jeu {
		fmt.Print("Pour Commencer Appuyer sur P puis sur Entrer ")
		fmt.Scan(&z)
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
	fmt.Print("Choisissez votre difficulté (1 à 4) :")
	fmt.Scan(&n)
	if n == 1 && len(rand) < 5 {
		fmt.Print("Vous ne pouvez pas choisir le niveau 1 car le mot fait moins de 5 lettres")
		fmt.Print("Difficulté (1 à 4) :")
		fmt.Scan(&n)
	}
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
	for {
		verif := false
		fmt.Print("Entrez une lettre : ")
		var mot string
		fmt.Scan(&mot)
		egale := false
		for i := 0; i < len(rand); i++ {
			if mot[0] == rand[i] {
				dash[i] = mot
				verif = true
				egale = true
			}
		}
		if !verif {
			rest--
			fmt.Println("Essai restant :", rest)
		}
		if rest >= 0 && rest < len(pendu) {
			if !egale {
				pendu[rest] += "========"
				fmt.Println(pendu[rest])
			}
		}
		if rest == 0 {
			fmt.Println("Vous avez perdu, le mot était :", rand)
			break
		}
		fmt.Println(strings.Join(dash, " "))

		if !strings.Contains(strings.Join(dash, ""), "_") {
			fmt.Println("Bravo !!")
			break
		}
	}
}
