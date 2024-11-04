package ascii

import (
	"log"
	"os"
	"strings"
)

// reading the banner file and printing ascii art matching the input argument
func PrintAsciiArt(input, banner string) string {

	var result strings.Builder

	if len(input) == 0 {
		return "Error"
	}
	bannerFile, err := os.ReadFile("./banners/" + banner + ".txt")
	if err != nil {
		log.Print("ERROR: Couldn't read the banner file: ", err)
		result.WriteString("ERROR: banner issues")
	} else if len(input) != 0 {
		cleanBannerFile := cleanBanner(bannerFile)
		bannerFileLines := strings.Split(string(cleanBannerFile), "\n")

		input = strings.ReplaceAll(input, "\\n", "\n")
		inputSlice := strings.Split(input, "\n")
		removedChar := ""
		inputSlice, removedChar = cleanInput(inputSlice)

		if removedChar != "" {
			result.WriteString("These characters were removed: " + removedChar + "\n\n")
		}

		//An empty index generates a newline, else the index is looped to match indexes from bannerFileLines
		for _, inputLine := range inputSlice {
			if inputLine == "" {
				result.WriteString("\n")
			} else {
				for i := 1; i <= 8; i++ {
					for _, char := range inputLine {
						result.WriteString(bannerFileLines[i+(int(char-32)*9)])
					}
					result.WriteString("\n")
				}
			}
		}

	}

	return result.String()
}

// Replacing carriage return with a newline
func cleanBanner(fileContent []byte) string {

	return strings.ReplaceAll(string(fileContent), "\r\n", "\n")

}

// Removing everything that is not printable ascii character nor new line from the input
func cleanInput(input []string) ([]string, string) {
	onlyNewLines := true
	var cleanInput []string
	var removedChar strings.Builder

	for _, lines := range input {
		if lines == "" {
			cleanInput = append(cleanInput, lines)
		} else {
			var printableChars strings.Builder
			onlyNewLines = false
			for _, char := range lines {
				if char >= 32 && char <= 126 {
					printableChars.WriteString(string(char))
				} else {
					removedChar.WriteString(string(char) + " ")
				}
			}
			if printableChars.String() != "" {
				cleanInput = append(cleanInput, printableChars.String())
			}
		}

	}
	if onlyNewLines {
		cleanInput = cleanInput[:1]
	}

	return cleanInput, removedChar.String()
}

// lisää infoboxi jossa kerrotaan käyttäjälle mitkä charachterit toimii inputkentässä
// '+"'"+'
