package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

type GameState struct {
	Name            string
	Score           int
	CurrentQuestion int
	Questions       []Question
}

type Question struct {
	Text    string
	Options []string
	Answer  int
}

func (g *GameState) Initialize() {
	fmt.Println("Welcome to the Quiz Game!")
	fmt.Println("Please enter your name: ")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')

	if err != nil {
		panic(err)
	}

	g.Name = name
	g.Score = 0
	g.CurrentQuestion = 0
	// g.Questions = loadQuestions("quiz.csv")

	fmt.Printf("Hello, %s! Let's start the quiz.\n", g.Name)
}

func (g *GameState) loadCsv() {
	file, err := os.Open("quiz.csv")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := csv.NewReader(file)
	record, err := scanner.ReadAll()

	if err != nil {
		panic(err)
	}

	for index, line := range record[:] {
		if index == 0 {
			continue
		}

		question := Question{
			Text:    line[0],
			Options: line[1:5],
			Answer:  toInt(line[5]),
		}

		g.Questions = append(g.Questions, question)

	}
}

func (g *GameState) Run() {
	for g.CurrentQuestion < len(g.Questions) {
		q := g.Questions[g.CurrentQuestion]
		fmt.Printf("\033[33m- %d: %s \033[0m\n", g.CurrentQuestion+1, q.Text)
		for i, option := range q.Options {
			fmt.Printf("%d. %s\n", i+1, option)
		}
		fmt.Print("Your answer: ")
		var answer int
		fmt.Scanln(&answer)

		if answer == q.Answer {
			fmt.Println("Correct!")
			g.Score += 10
		} else {
			fmt.Printf("Wrong! The correct answer was %d.\n", q.Answer)
		}
		g.CurrentQuestion++
	}
	fmt.Printf("Quiz over! Your final score is %d out of %d.\n", g.Score, len(g.Questions))
}

func (g *GameState) CalculateScore() string {
	if g.Score >= 20 {
		return "Passed"
	}
	return "Failed"
}

func (g *GameState) GameCountDown() {
	seconds := 15
	currentQuestion := 0

	fmt.Println("⏳ Iniciando contagem regressiva:")

	for i := seconds; i > 0; i-- {
		if i == 10 {
			fmt.Println("⚠️  Atenção! Faltam 10 segundos!")
		} else {
			fmt.Printf("⏳ %d segundos restantes...\n", i)
		}

		time.Sleep(1 * time.Second)

		if currentQuestion != g.CurrentQuestion {
			currentQuestion = g.CurrentQuestion
			i = seconds + 1
		}
	}

	fmt.Println("🚀 Tempo esgotado!")
	os.Exit(0)
}

func toInt(s string) int {
	var i int
	fmt.Sscanf(s, "%d", &i)
	return i
}

func main() {
	game1 := &GameState{}
	go game1.GameCountDown()
	go game1.loadCsv()
	game1.Initialize()
	game1.Run()
	fmt.Println(game1.CalculateScore())
}
