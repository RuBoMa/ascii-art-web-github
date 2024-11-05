package ascii

import (
	"log"
	"os"
	"strings"
)

// reading the banner file and printing ascii art matching the input argument
func PrintAsciiArt(input, banner string) (string, error) {

	var result string
	var httpError error

	bannerFile, err := os.ReadFile("./banners/" + banner + ".txt")
	if err != nil {
		log.Print("Error reading banner file: ", err)
		httpError = err
	} else if len(input) != 0 {
		cleanBannerFile := strings.ReplaceAll(string(bannerFile), "\r\n", "\n")
		bannerFileLines := strings.Split(string(cleanBannerFile), "\n")

		input = strings.ReplaceAll(input, "\\n", "\n")
		input = cleanInput(input)
		inputSlice := strings.Split(input, "\n")

		onlyNewLines := true

		//An empty index generates a newline, else the index is looped to match indexes from bannerFileLines
		for _, inputLine := range inputSlice {
			if inputLine == "" {
				result += "\n"
			} else {
				onlyNewLines = false
				for i := 1; i <= 8; i++ {
					for _, char := range inputLine {
						result += bannerFileLines[i+(int(char-32)*9)]
					}
					result += "\n"
				}
			}
		}
		if onlyNewLines {
			result = result[:1]
		}

	}

	return result, httpError
}

// Removing everything that is not printable ascii character from the input
func cleanInput(input string) string {
	var cleanInput strings.Builder

	for _, char := range input {
		if char >= 32 && char <= 126 || char == '\n' {
			cleanInput.WriteRune(char)
		}
	}

	return cleanInput.String()
}

// lisää infoboxi jossa kerrotaan käyttäjälle mitkä charachterit toimii inputkentässä
// '+"'"+'
