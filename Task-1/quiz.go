package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	var w sync.WaitGroup
	w.Add(1)
	questions, answers := readCSV("problems.csv")
	timePtr := flag.Int("limit", -1, "lt")
	points := 0
	flag.Parse()
	go askQuestions(questions, answers, &points)
	var timeGiven int = *timePtr
	if timeGiven != -1 {
		timer1 := time.NewTimer(time.Duration(timeGiven) * time.Second)
		<-timer1.C
	} else {
		w.Wait()
	}

	printPoints(points)
}
func printPoints(points int) {
	fmt.Printf("You got %d points\n", points)
}
func readCSV(filename string) ([]string, []string) {
	var questions []string
	var answers []string
	file, err := os.OpenFile(filename, os.O_RDONLY, 0755)
	if err != nil {
		fmt.Printf("Error in reading file %s \n %s",filename, err.Error())
		return questions, answers
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err.Error())
		return questions, answers
	}
	for _, el := range records {
		questions = append(questions, el[0])
		answers = append(answers, el[1])
	}
	return questions, answers
}

func askQuestions(questions []string, answers []string, points *int) {
	for index, question := range questions {
		var answer string
		fmt.Printf("%d. %s", index + 1, question)
		fmt.Printf("\nAns:\t")
		_, _ = fmt.Scanf("%s", &answer)
		if answer == answers[index] {
			*points++
		}
	}
}