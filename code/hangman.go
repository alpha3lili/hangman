package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
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
var win bool
var LetterTry []string

func Hangman() {
	ClearScreen()
	//rand.Seed(time.Now().UnixNano())
	win = false
	begin = true
	LetterTry = []string{}
	LosePoint = 0
	Losepointsmax = 11
	var Sidewords string
	var lettre string
	temp := ""
	for _, element := range AllWords {
		if element != '\r' && element != ' ' {
			temp += string(element)
		}
	}
	all := strings.Split(temp, "\n")
	fmt.Println(len(all))
	var Word string = all[rand.Intn(len(all))]
	Inconnue = []string{}
	for i := 0; i < len(Word); i++ {
		Inconnue = append(Inconnue, "_")
	}
	for {
		ClearScreen()
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
		DislplayLetter()
		Error()
		fmt.Println("")
		fmt.Println("")
		fmt.Print("Enter a letter: ")
		fmt.Scanln(&lettre)
		if len(lettre) > 1 {
			if lettre != Word { //Si je ne renvoie pas le même mot, j'augmente mon nombre d'erreur de 2
				LosePoint += 2
				LetterTry = append(LetterTry, lettre)
			} else {
				win = true
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
			for _, ele := range LetterTry {
				if ele == lettre {
					error--
				}
			}

			if error == len(Word) { //Si le nombre d'erreur est égale à la longueur du mot, j'augmente mon nombre d'erreur de 1
				LosePoint++
				LetterTry = append(LetterTry, lettre)
			}
			c := 0
			for _, ele := range Inconnue {
				if ele != "_" {
					c++
				}

			}
			if c == len(Word) {
				win = true
			}
		}
		if LosePoint >= Losepointsmax {
			win = false
			break
		}
		if win {
			break
		}
	}
	Finish()
}

func DislplayLetter() {
	fmt.Println("")
	fmt.Print("Letter or word tried: ")
	for _, a := range LetterTry {
		fmt.Print(a, " ")
	}
	fmt.Println("")
}

func Finish() {
	ClearScreen()
	if win {
		fmt.Println("You win !")
		fmt.Println("Do you want to play again ? (yes/no)")
		var answer string
		fmt.Scanln(&answer)
		if answer == "yes" {
			Hangman()
		} else {
			return
		}
	} else {
		fmt.Println("You lose !")
		fmt.Println("Do you want to play again ? (yes/no)")
		var answer string
		fmt.Scanln(&answer)
		if answer == "yes" {
			Hangman()
		} else {
			return
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
		fmt.Println(LightYellow, "======[]=========|========", Reset)
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
		fmt.Println(OrangeYellow, "======[]=========|========", Reset)
		fmt.Println(OrangeYellow, "      []         |", Reset)
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
		fmt.Println(LightOrange, "======[]=========|========", Reset)
		fmt.Println(LightOrange, "      []         |", Reset)
		fmt.Println(LightOrange, "      []       |   |", Reset)
		fmt.Println(LightOrange, "      []         ||", Reset)
		fmt.Println(LightOrange, "      []              ", Reset)
		fmt.Println(LightOrange, "      []          ", Reset)
		fmt.Println(LightOrange, "      []            ", Reset)
		fmt.Println(LightOrange, "      []             ", Reset)
		fmt.Println(LightOrange, "      []            ", Reset)
		fmt.Println(LightOrange, "      []\\", Reset)
		fmt.Println(LightOrange, "    __[]_\\__", Reset)
	} else if LosePoint == 7 {
		fmt.Println(LightOrange, "======[]=========|========")
		fmt.Println(LightOrange, "      []         |", Reset)
		fmt.Println(LightOrange, "      []       |   |", Reset)
		fmt.Println(LightOrange, "      []         ||", Reset)
		fmt.Println(LightOrange, "      []         |", Reset)
		fmt.Println(LightOrange, "      []         |", Reset)
		fmt.Println(LightOrange, "      []         |", Reset)
		fmt.Println(LightOrange, "      []               ", Reset)
		fmt.Println(LightOrange, "      []           ", Reset)
		fmt.Println(LightOrange, "      []\\", Reset)
		fmt.Println(LightOrange, "    __[]_\\__", Reset)
	} else if LosePoint == 8 {
		fmt.Println(PinkRed, "======[]=========|========", Reset)
		fmt.Println(PinkRed, "      []         |", Reset)
		fmt.Println(PinkRed, "      []       |   |", Reset)
		fmt.Println(PinkRed, "      []         ||", Reset)
		fmt.Println(PinkRed, "      []         |", Reset)
		fmt.Println(PinkRed, "      []         |\\", Reset)
		fmt.Println(PinkRed, "      []         | \\", Reset)
		fmt.Println(PinkRed, "      []              ", Reset)
		fmt.Println(PinkRed, "      []            ", Reset)
		fmt.Println(PinkRed, "      []\\", Reset)
		fmt.Println(PinkRed, "    __[]_\\__", Reset)
	} else if LosePoint == 9 {
		fmt.Println(PinkRed, "======[]=========|========", Reset)
		fmt.Println(PinkRed, "      []         |", Reset)
		fmt.Println(PinkRed, "      []       |   |", Reset)
		fmt.Println(PinkRed, "      []         ||", Reset)
		fmt.Println(PinkRed, "      []         |", Reset)
		fmt.Println(PinkRed, "      []        /|\\", Reset)
		fmt.Println(PinkRed, "      []       / | \\", Reset)
		fmt.Println(PinkRed, "      []             ", Reset)
		fmt.Println(PinkRed, "      []            ", Reset)
		fmt.Println(PinkRed, "      []\\", Reset)
		fmt.Println(PinkRed, "    __[]_\\__", Reset)
	} else if LosePoint == 10 {
		fmt.Println(LightRed, "======[]=========|========", Reset)
		fmt.Println(LightRed, "      []         |", Reset)
		fmt.Println(LightRed, "      []       |   |", Reset)
		fmt.Println(LightRed, "      []         ||", Reset)
		fmt.Println(LightRed, "      []         |", Reset)
		fmt.Println(LightRed, "      []        /|\\", Reset)
		fmt.Println(LightRed, "      []       / | \\", Reset)
		fmt.Println(LightRed, "      []         | ", Reset)
		fmt.Println(LightRed, "      []\\", Reset)
		fmt.Println(LightRed, "    __[]_\\__", Reset)
	} else if LosePoint >= 11 {
		fmt.Println(LightRed, "======[]=========|========", Reset)
		fmt.Println(LightRed, "      []         |", Reset)
		fmt.Println(LightRed, "      []       |   |", Reset)
		fmt.Println(LightRed, "      []         ||", Reset)
		fmt.Println(LightRed, "      []         |", Reset)
		fmt.Println(LightRed, "      []        /|\\", Reset)
		fmt.Println(LightRed, "      []       / | \\", Reset)
		fmt.Println(LightRed, "      []        / \\", Reset)
		fmt.Println(LightRed, "      []       /   \\", Reset)
		fmt.Println(LightRed, "      []\\", Reset)
		fmt.Println(LightRed, "    __[]_\\__", Reset)
	} else if LosePoint == 0 {
		fmt.Println("")
	}
}

func Random(word string) {
	help := word[rand.Intn(len(word))]
	for i, k := range word {
		if string(k) == string(help) {
			Inconnue[i] = string(k)
		}
	}
}

func ClearScreen() {
	//Function to called when we want to clear the Terminal
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
