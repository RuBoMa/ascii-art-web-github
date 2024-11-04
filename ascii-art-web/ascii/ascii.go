package ascii

import (
	"log"
	"os"
	"strings"
)

// reading the banner file and printing ascii art matching the input argument
func PrintAsciiArt(input, banner string) string {

	var result string
	input = cleanStr(input)
	if len(input) == 0 {
		return "Error"
	}
	bannerFile, err := os.ReadFile("./banners/" + banner + ".txt")
	if err != nil {
		log.Print("ERROR: Couldn't read the banner file: ", err)
		result = "ERROR: banner issues"
		return result

	} else if len(input) != 0 {
		cleaned := cleanInput(bannerFile)
		//splitting the banner file into a slice by rows (Index 0 = row 1)
		bannerFileLines := strings.Split(string(cleaned), "\n")

		//splitting the input into a slice by \n
		input = strings.ReplaceAll(input, "\\n", "\n")
		words := strings.Split(input, "\n")

		onlyNewLines := true

		//An empty index generates a newline, else the index is looped to match indexes from bannerFileLines
		for _, word := range words {
			if word == "" {
				result += "\n"
			} else {
				onlyNewLines = false
				for i := 1; i <= 8; i++ {
					for _, char := range word {
						result += bannerFileLines[i+(int(char-32)*9)]
					}
					result += "\n"
				}
			}
		}
		//if the input consists only newlines, deducting one newline
		if onlyNewLines && len(result) > 0 {
			result = result[1:]
		}
	}

	return result
}

func cleanInput(fileContent []byte) string {

	return strings.ReplaceAll(string(fileContent), "\r\n", "\n")

}

func cleanStr(s string) string {
	cleanedStr := ""

	for _, char := range s {
		if char >= 32 && char <= 127 || char == '\n' {
			cleanedStr += string(char)
		}
	}
	return cleanedStr
}

// lisää infoboxi jossa kerrotaan käyttäjälle mitkä charachterit toimii inputkentässä
// '+"'"+'