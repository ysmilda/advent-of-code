package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"os"
	"text/template"
	"time"
)

type definition struct {
	Day  uint
	Year uint
}

//go:embed day.tmpl
var dayTemplateContent string
var dayTemplate = template.Must(template.New("day").Parse(dayTemplateContent))

//go:embed day_test.tmpl
var dayTestTemplateContent string
var dayTestTemplate = template.Must(template.New("day_test").Parse(dayTestTemplateContent))

var (
	year, day uint
)

func main() {
	flag.UintVar(&year, "year", uint(time.Now().Year()), "the year to generate for")
	flag.UintVar(&day, "day", 0, "the day to generate for")
	flag.Parse()

	if day == 0 {
		log.Fatal("day should have been set")
	}

	path := fmt.Sprintf("./solutions/%d/day%d", year, day)
	if _, err := os.ReadDir(path); err == nil {
		log.Fatal(fmt.Errorf("day %d already exists", day))
	}

	err := os.MkdirAll(path, 0o755)
	if err != nil {
		log.Fatal("Unable to create directory:", err)
	}

	solutionFile, err := os.Create(fmt.Sprintf("%s/day%d.go", path, day))
	if err != nil {
		log.Fatal("unable to create file:", err)
	}
	defer solutionFile.Close()

	testFile, err := os.Create(fmt.Sprintf("%s/day%d_test.go", path, day))
	if err != nil {
		log.Fatal("unable to create file:", err)
	}
	defer testFile.Close()

	def := definition{
		Day:  day,
		Year: year,
	}

	err = dayTemplate.Execute(solutionFile, def)
	if err != nil {
		log.Fatal("unable to execute template:", err)
	}

	err = dayTestTemplate.Execute(testFile, def)
	if err != nil {
		log.Fatal("unable to execute template:", err)
	}
}
