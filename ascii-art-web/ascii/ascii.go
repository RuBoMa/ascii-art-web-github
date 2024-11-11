package ascii

import (
	"log"
	"os"
	"strings"
)

// reading the banner file and printing ascii art matching the input argument
func PrintAsciiArt(input, banner string) (string, error) {

	var result string

	if len(input) > 0 {
		bannerFile, err := os.ReadFile("./banners/" + banner + ".txt")
		if err != nil {
			log.Print("Error reading banner file: ", err)
			return result, err
		} else {
			cleanBannerFile := strings.ReplaceAll(string(bannerFile), "\r\n", "\n")
			bannerFileLines := strings.Split(string(cleanBannerFile), "\n")
			log.Println("Banner handled")

			input = replaseNewlines(input)
			input = strings.ReplaceAll(input, "\\n", "\n")
			inputSlice := strings.Split(input, "\n")
			log.Println("Input handled")

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
	}
	return result, nil
}

// Removing everything that is not printable ascii character from the input
func ValidInput(input string) bool {
	nInput := replaseNewlines(input)

	for _, char := range nInput {
		if (char < 32 || char > 127) && char != '\n' {
			log.Println("Not ascii", char)
			return false
		}
	}

	return true
}

func replaseNewlines(input string) string {
	log.Println("replacing newlines")
	newStr := strings.ReplaceAll(input, "\r\n", "\n")
	log.Println("replaced \\r\\n")
	// secStr := strings.ReplaceAll(newStr, "\r", "\n")
	// log.Panicln("replaced \\r")
	return newStr
}
