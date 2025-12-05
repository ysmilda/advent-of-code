package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/ysmilda/Advent-of-code/foundation/aocstrconv"
)

var sessionToken string

func main() {
	flag.StringVar(&sessionToken, "session-token", "", "the session token to use for downloading the input")
	flag.Parse()

	if sessionToken == "" {
		log.Fatal("session-token should be set")
	}

	years, err := os.ReadDir("./solutions")
	if err != nil {
		log.Fatal("unable to read directory:", err)
	}

	for _, year := range years {
		if !year.IsDir() {
			continue
		}

		yearPath := filepath.Join("./solutions", year.Name())
		days, err := os.ReadDir(yearPath)
		if err != nil {
			log.Fatal("unable to read directory:", err)
		}

		for _, day := range days {
			dayPath := filepath.Join(yearPath, day.Name())
			dayFolderContents, err := os.ReadDir(dayPath)
			if err != nil {
				log.Fatal("unable to read directory:", err)
			}

			skip := false
			for _, file := range dayFolderContents {
				if file.Name() == "input.txt" {
					skip = true
				}
			}
			if skip {
				continue
			}

			yearInt := aocstrconv.MustAtoi(year.Name())
			dayInt := aocstrconv.MustAtoi(day.Name()[3:])

			log.Printf("Fetching input for %d day %d\n", yearInt, dayInt)

			req, err := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", yearInt, dayInt), nil)
			if err != nil {
				log.Fatal("unable to create request", err)

			}

			req.AddCookie(&http.Cookie{
				Name:  "session",
				Value: sessionToken,
			})

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				log.Fatal("unable to get input:", err)

			}

			InputFile, err := os.Create(dayPath + "/input.txt")
			if err != nil {
				log.Fatal("unable to write input to file:", err)

			}

			input, err := io.ReadAll(resp.Body)
			if err != nil {
				InputFile.Close()
				log.Fatal("unable to read input from response:", err)
			}

			if input[len(input)-1] == '\n' {
				input = input[:len(input)-1]
			}

			InputFile.Write(input)
			InputFile.Close()
		}
	}
}
