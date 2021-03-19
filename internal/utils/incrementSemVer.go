package utils

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func IncrementCurrentRelease(current string, label SemVer) string {
	s := strings.Split(current, ".")

	sMajor, err := strconv.Atoi(s[0])
	if err != nil {
		log.Fatal(err)
	}

	sMinor, err := strconv.Atoi(s[1])
	if err != nil {
		log.Fatal(err)
	}

	sPatch, err := strconv.Atoi(s[2])
	if err != nil {
		log.Fatal(err)
	}

	switch label {
	case Major:
		sMajor++
		sMinor = 0
		sPatch = 0
	case Minor:
		sMinor++
		sPatch = 0
	case Patch:
		sPatch++
	}

	currentRelease := fmt.Sprintf("%d.%d.%d", sMajor, sMinor, sPatch)

	return currentRelease
}
