package ascii

import (
	"log"
	"os"
	"strings"
)

// reading the banner file and printing ascii art matching the input argument
func PrintAsciiArt(input, banner string) string {

	var result string
	input = cleanInput(input)
	if len(input) == 0 {
		return "Error"
	}
	bannerFile, err := os.ReadFile("./banners/" + banner + ".txt")
	if err != nil {
		log.Print("ERROR: Couldn't read the banner file: ", err)
		result = "ERROR: banner issues"
	} else if len(input) != 0 {
		cleanBannerFile := cleanBanner(bannerFile)
		bannerFileLines := strings.Split(string(cleanBannerFile), "\n")

		input = strings.ReplaceAll(input, "\\n", "\n")
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
		//if the input consists only newlines, deducting one newline
		if onlyNewLines {
			result = result[1:]
		}
	}

	return result
}

// Replacing carriage return with a newline
func cleanBanner(fileContent []byte) string {

	return strings.ReplaceAll(string(fileContent), "\r\n", "\n")

}

// Removing everything that is not printable ascii character nor new line from the input
func cleanInput(s string) string {
	cleanInput := ""

	for _, char := range s {
		if char >= 32 && char <= 127 || char == '\n' {
			cleanInput += string(char)
		}
	}
	return cleanInput
}

// lisää infoboxi jossa kerrotaan käyttäjälle mitkä charachterit toimii inputkentässä
// '+"'"+'
