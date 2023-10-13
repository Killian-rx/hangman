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
	nindex := len(rand)/2 - 1
	nlettre := string(rand[nindex])
	dash := make([]string, len(rand))
	for i := 0; i < len(rand); i++ {
		if i != nindex {
			dash[i] = "_"
		} else {
			dash[i] = nlettre
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
