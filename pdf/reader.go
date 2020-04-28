package reader

import (
	"bufio"
	"log"
	"os"
)

//Saving pieces of the txt(Not the most optimal way :P)
var (
	MainTitle       string
	MainDescription []string
	InstallTitle    string
	InstallDesc     []string
	CreateTitle     string
	CreateDesc      []string
	GitInit         string
	GitClone        string
)

//ReadPdf func from
func ReadPdf() error {
	f, err := os.Open("./github-git-cheat-sheet.txt")
	if err != nil {
		log.Fatalf("Error opening github-git-cheat-sheet: %s", err)
	}
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	f.Close()

	return nil
}
