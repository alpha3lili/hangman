package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	LightGreen   = "\033[38;5;10m"  // Vert fluo
	DarkGreen    = "\033[38;5;28m"  // Vert foncé
	LightYellow  = "\033[38;5;226m" // Jaune fluo
	OrangeYellow = "\033[38;5;214m" // Jaune orangé
	LightOrange  = "\033[38;5;208m" // Orange fluo
	PinkRed      = "\033[38;5;197m" // Pink red
	LightRed     = "\033[38;5;196m" // Rouge fluo
	Reset        = "\033[0m"        // Réinitialisation de la couleur
)

var LosePoint int
var Losepointsmax int
var Inconnue []string
var begin bool = true

func Hangman() {
	rand.Seed(time.Now().UnixNano())
	LosePoint = 0
	Losepointsmax = 11
	var Sidewords string
	var lettre string
	fmt.Println(len(List))
	var Word string = List[rand.Intn(len(List))]
	for i := 0; i < len(Word); i++ {
		Inconnue = append(Inconnue, "_")
	}
	for {
		for _, a := range Inconnue {
			Sidewords += a
		}
		if begin {
			Random(Word)
			begin = false
		}
		fmt.Println("============H.A.N.G.M.A.N============")
		fmt.Println("")
		fmt.Println("Word: ", Inconnue)
		fmt.Println("")
		fmt.Println("Losepoints: ", LosePoint)
		Error()
		fmt.Println("")
		fmt.Println("")
		fmt.Print("Enter a letter: ")
		fmt.Scanln(&lettre)
		if len(lettre) > 1 {
			if lettre != Word { //Si je ne renvoie pas le même mot, j'augmente mon nombre d'erreur de 2
				LosePoint += 2
			} else {
				Win()
			}
		} else {
			var error int
			for i, k := range Word { //Je parcours le mot
				if string(k) == lettre { //Si la lettre est égale à la lettre du mot
					Inconnue[i] = lettre
				} else { //Sinon j'augmente mon nombre d'erreur de 1
					error++
				}
			}
			if error == len(Word) { //Si le nombre d'erreur est égale à la longueur du mot, j'augmente mon nombre d'erreur de 1
				LosePoint++
			}
		}
	}
}

func Error() {
	if LosePoint == 1 {
		fmt.Println("")
		fmt.Println("    ")
		fmt.Println("    ")
		fmt.Println("    ")
		fmt.Println("    ")
		fmt.Println("    ")
		fmt.Println("     ")
		fmt.Println("     ")
		fmt.Println("    ")
		fmt.Println("     ")
		fmt.Println(LightGreen, " _____", Reset)
	} else if LosePoint == 2 {
		fmt.Println(LightGreen, "      []", Reset)
		fmt.Println(LightGreen, "      []    ", Reset)
		fmt.Println(LightGreen, "      []        ", Reset)
		fmt.Println(LightGreen, "      []         ", Reset)
		fmt.Println(LightGreen, "      []              ", Reset)
		fmt.Println(LightGreen, "      []   ", Reset)
		fmt.Println(LightGreen, "      []      ", Reset)
		fmt.Println(LightGreen, "      []        ", Reset)
		fmt.Println(LightGreen, "      []         ", Reset)
		fmt.Println(LightGreen, "      []")
		fmt.Println(LightGreen, "    __[]__", Reset)
	} else if LosePoint == 3 {
		fmt.Println(DarkGreen, "      []", Reset)
		fmt.Println(DarkGreen, "      []            ", Reset)
		fmt.Println(DarkGreen, "      []         ", Reset)
		fmt.Println(DarkGreen, "      []          ", Reset)
		fmt.Println(DarkGreen, "      []              ", Reset)
		fmt.Println(DarkGreen, "      []          ", Reset)
		fmt.Println(DarkGreen, "      []       ", Reset)
		fmt.Println(DarkGreen, "      []         ", Reset)
		fmt.Println(DarkGreen, "      []            ", Reset)
		fmt.Println(DarkGreen, "      []\\")
		fmt.Println(DarkGreen, "    __[]_\\__", Reset)
	} else if LosePoint == 4 {
		fmt.Println(LightYellow, "======[]==============|========", Reset)
		fmt.Println(LightYellow, "      []           ", Reset)
		fmt.Println(LightYellow, "      []         ", Reset)
		fmt.Println(LightYellow, "      []           ", Reset)
		fmt.Println(LightYellow, "      []           ", Reset)
		fmt.Println(LightYellow, "      []          ", Reset)
		fmt.Println(LightYellow, "      []           ", Reset)
		fmt.Println(LightYellow, "      []           ", Reset)
		fmt.Println(LightYellow, "      []           ", Reset)
		fmt.Println(LightYellow, "      []\\", Reset)
		fmt.Println(LightYellow, "    __[]_\\__", Reset)
	} else if LosePoint == 5 {
		fmt.Println(OrangeYellow, "======[]==============|========", Reset)
		fmt.Println(OrangeYellow, "      []              |", Reset)
		fmt.Println(OrangeYellow, "      []            ", Reset)
		fmt.Println(OrangeYellow, "      []          ", Reset)
		fmt.Println(OrangeYellow, "      []              ", Reset)
		fmt.Println(OrangeYellow, "      []         ", Reset)
		fmt.Println(OrangeYellow, "      []          ", Reset)
		fmt.Println(OrangeYellow, "      []           ", Reset)
		fmt.Println(OrangeYellow, "      []        ", Reset)
		fmt.Println(OrangeYellow, "      []\\", Reset)
		fmt.Println(OrangeYellow, "    __[]_\\__", Reset)
	} else if LosePoint == 6 {
		fmt.Println(LightOrange, "======[]==============|========", Reset)
		fmt.Println(LightOrange, "      []            |", Reset)
		fmt.Println(LightOrange, "      []            |   |", Reset)
		fmt.Println(LightOrange, "      []            ||", Reset)
		fmt.Println(LightOrange, "      []              ", Reset)
		fmt.Println(LightOrange, "      []          ", Reset)
		fmt.Println(LightOrange, "      []            ", Reset)
		fmt.Println(LightOrange, "      []             ", Reset)
		fmt.Println(LightOrange, "      []            ", Reset)
		fmt.Println(LightOrange, "      []\\", Reset)
		fmt.Println(LightOrange, "    __[]_\\__", Reset)
	} else if LosePoint == 7 {
		fmt.Println(LightOrange, "======[]==============|========")
		fmt.Println(LightOrange, "      []            |", Reset)
		fmt.Println(LightOrange, "      []            |   |", Reset)
		fmt.Println(LightOrange, "      []            ||", Reset)
		fmt.Println(LightOrange, "      []              |", Reset)
		fmt.Println(LightOrange, "      []              |", Reset)
		fmt.Println(LightOrange, "      []              |", Reset)
		fmt.Println(LightOrange, "      []               ", Reset)
		fmt.Println(LightOrange, "      []           ", Reset)
		fmt.Println(LightOrange, "      []\\", Reset)
		fmt.Println(LightOrange, "    __[]_\\__", Reset)
	} else if LosePoint == 8 {
		fmt.Println(PinkRed, "======[]==============|========", Reset)
		fmt.Println(PinkRed, "      []            |", Reset)
		fmt.Println(PinkRed, "      []            |   |", Reset)
		fmt.Println(PinkRed, "      []            ||", Reset)
		fmt.Println(PinkRed, "      []              |", Reset)
		fmt.Println(PinkRed, "      []              |\\", Reset)
		fmt.Println(PinkRed, "      []              | \\", Reset)
		fmt.Println(PinkRed, "      []              ", Reset)
		fmt.Println(PinkRed, "      []            ", Reset)
		fmt.Println(PinkRed, "      []\\", Reset)
		fmt.Println(PinkRed, "    __[]_\\__", Reset)
	} else if LosePoint == 9 {
		fmt.Println(PinkRed, "======[]==============|========", Reset)
		fmt.Println(PinkRed, "      []            |", Reset)
		fmt.Println(PinkRed, "      []            |   |", Reset)
		fmt.Println(PinkRed, "      []            ||", Reset)
		fmt.Println(PinkRed, "      []              |", Reset)
		fmt.Println(PinkRed, "      []             /|\\", Reset)
		fmt.Println(PinkRed, "      []            / | \\", Reset)
		fmt.Println(PinkRed, "      []             ", Reset)
		fmt.Println(PinkRed, "      []            ", Reset)
		fmt.Println(PinkRed, "      []\\", Reset)
		fmt.Println(PinkRed, "    __[]_\\__", Reset)
	} else if LosePoint == 10 {
		fmt.Println(LightRed, "======[]==============|========", Reset)
		fmt.Println(LightRed, "      []            |", Reset)
		fmt.Println(LightRed, "      []            |   |", Reset)
		fmt.Println(LightRed, "      []            ||", Reset)
		fmt.Println(LightRed, "      []              |", Reset)
		fmt.Println(LightRed, "      []             /|\\", Reset)
		fmt.Println(LightRed, "      []            / | \\", Reset)
		fmt.Println(LightRed, "      []               \\", Reset)
		fmt.Println(LightRed, "      []                \\", Reset)
		fmt.Println(LightRed, "      []\\", Reset)
		fmt.Println(LightRed, "    __[]_\\__", Reset)
	} else if LosePoint >= 11 {
		fmt.Println(LightRed, "======[]==============|========", Reset)
		fmt.Println(LightRed, "      []            |", Reset)
		fmt.Println(LightRed, "      []            |   |", Reset)
		fmt.Println(LightRed, "      []            ||", Reset)
		fmt.Println(LightRed, "      []              |", Reset)
		fmt.Println(LightRed, "      []             /|\\", Reset)
		fmt.Println(LightRed, "      []            / | \\", Reset)
		fmt.Println(LightRed, "      []             / \\", Reset)
		fmt.Println(LightRed, "      []            /   \\", Reset)
		fmt.Println(LightRed, "      []\\", Reset)
		fmt.Println(LightRed, "    __[]_\\__", Reset)
	} else if LosePoint == 0 {
		fmt.Println("")
	}
}

func Win() {
	fmt.Println("winner")
}

func Random(word string) {
	help := word[rand.Intn(len(word))]
	for i, k := range word {
		if string(k) == string(help) {
			Inconnue[i] = string(k)
		}
	}
}
