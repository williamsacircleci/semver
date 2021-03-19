package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"regexp"
	"strings"
)

func CleanReleaseNotes(releaseNotes string) string {

	strToByte := []byte(releaseNotes)
	noteBytes := bytes.NewReader(strToByte)

	scanner := bufio.NewScanner(noteBytes)

	markdown := "### "
	formattedNotes := ""

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) > 0 {
			parts := strings.Split(line, "\n")
			for i := range parts {
				str := parts[i]
				if strings.Contains(str, markdown) {
					str = strings.Replace(str, markdown, "*", -1)
					str += "*"
				}
				formattedNotes += fmt.Sprintf("%s\n", str)
			}
		}
	}

	regex, err := regexp.Compile("\n\n")
	if err != nil {
		log.Fatalln(err)
	}

	formattedNotes = regex.ReplaceAllString(formattedNotes, "\n")

	return formattedNotes
}
