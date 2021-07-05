package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"sync"
	"time"
)
type problem struct {
	question string
	answer string
}
func main() {
	var w sync.WaitGroup
	w.Add(1)
	problems := readCSV("problems.csv")
	timePtr := flag.Int("limit", -1, "lt")
	points := 0
	flag.Parse()
	go askQuestions(problems, &points)
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
func readCSV(filename string) ([]problem) {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0755)
	if err != nil {
		fmt.Printf("Error in reading file %s \n %s",filename, err.Error())
		return []problem{}
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	problems := make([]problem, len(records))
	if err != nil {
		fmt.Println(err.Error())
		return []problem{}
	}
	for index, el := range records {
		problems[index] = problem {
			question: el[0],
			answer: el[1],
		}
	}
	return problems
}

func askQuestions(problems []problem, points *int) {
	for index, problem := range problems {
		var answer string
		fmt.Printf("%d. %s", index + 1, problem.question)
		fmt.Printf("\nAns:\t")
		_, _ = fmt.Scanf("%s", &answer)
		if answer == problem.answer {
			*points++
		}
	}
}