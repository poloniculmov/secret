package main

import (
	"encoding/csv"
	"github.com/davecgh/go-spew/spew"
	"log"
	"os"
	"secretSanta/santa"
	"strings"
)

func check(error error) {
	if error != nil {
		log.Fatal(error)
	}
}

func readFile() string {
	lines, _ := os.ReadFile("secret_santa.csv")
	return string(lines)
}

func parseCSV(in string) [][]string {
	r := csv.NewReader(strings.NewReader(in))

	records, err := r.ReadAll()
	check(err)

	return records
}

func main() {
	lines := parseCSV(readFile())
	var santas []santa.Santa
	for _, l := range lines[1:] {
		santas = append(santas, santa.NewFromCSV(l))
	}
	drawResult :=santa.Draw(santas)
	spew.Dump(drawResult)

	var em []santa.Email

	var drawText string
	for index, elem := range drawResult {
		drawText += santas[index].FullName() + " drew " + santas[elem].FullName() +"\n"
		em = append(em, santa.NewEmail(santas[index], santas[elem]))
	}
	writeDraw(drawText)
	spew.Dump(em)
	for _, elem := range em {
		elem.Send()
	}
}

func writeDraw(strings string)  {
	f, err := os.Create("draw.txt")
	check(err)

	defer f.Close()
	f.WriteString(strings)

	f.Sync()
}