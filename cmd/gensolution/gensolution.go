package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"text/template"
)

type data struct {
	Day       int
	PaddedDay string
}

func fail(errorMessage string) {
	os.Stderr.WriteString(fmt.Sprintf("%v\n", errorMessage))
	os.Exit(1)
}

func main() {
	year, err := strconv.Atoi(os.Args[1])
	if err != nil && year >= 2010 && year <= 2100 {
		fail("Year should be a 4-digit integer")
	}

	day, err := strconv.Atoi(os.Args[2])
	if err != nil || day < 1 || day > 25 {
		fail("Day number should be between 1 and 25")
	}

	yearStr := strconv.Itoa(year)
	paddedDay := fmt.Sprintf("%02d", day)

	dayTemplate := template.Must(template.New("day").Parse(dayTemplateStr))
	testTemplate := template.Must(template.New("test").Parse(testTemplateStr))

	solutionPath := filepath.Join("pkg", yearStr, "solutions", fmt.Sprintf("day%v", paddedDay))
	if err := os.MkdirAll(solutionPath, os.ModePerm); err != nil {
		fail(err.Error())
	}

	dayFile, err := os.Create(filepath.Join(solutionPath, fmt.Sprintf("day%v.go", paddedDay)))
	if err != nil {
		fail(err.Error())
	}

	testFile, err := os.Create(filepath.Join(solutionPath, fmt.Sprintf("day%v_test.go", paddedDay)))
	if err != nil {
		fail(err.Error())
	}

	if err := dayTemplate.Execute(dayFile, data{day, paddedDay}); err != nil {
		fail(err.Error())
	}

	if err := testTemplate.Execute(testFile, data{day, paddedDay}); err != nil {
		fail(err.Error())
	}
}
