package ascii

import (
	"log"
	"os"
	"strings"
)

// reading the banner file and printing ascii art matching the input argument
func PrintAsciiArt(input, banner string) string {

	var result string

	if len(input) > 0 {
		bannerFileLines := strings.Split(banner, "\n")

		inputSlice := strings.Split(input, "\n")

		//An empty index generates a newline, else the index is looped to match indexes from bannerFileLines
		for _, inputLine := range inputSlice {
			if inputLine == "" {
				result += "\n"
			} else {
				for i := 1; i <= 8; i++ {
					for _, char := range inputLine {
						result += bannerFileLines[i+(int(char-32)*9)]
					}
					result += "\n"
				}
			}
		}
	}
	return result
}

// Replacing carriage return with a newline and checking if all characters are printable
func ValidInput(input string) (string, bool) {
	input = strings.ReplaceAll(input, "\r\n", "\n")

	for _, char := range input {
		if (char < 32 || char > 127) && char != '\n' {
			return input, false // Input has non-printable chacarters
		}
	}

	return input, true // Input consists of only printable ascii characters and newlines
}

// Reading the banner file from the directory and removing carriage returns
func ReadBanner(banner string) (string, error) {
	bannerContent, err := os.ReadFile("./banners/" + banner + ".txt")
	if err != nil {
		log.Print("Error reading banner file: ", err)
		return banner, err
	}
	cleanBanner := strings.ReplaceAll(string(bannerContent), "\r\n", "\n") // removing carriage return from the
	return cleanBanner, nil
}
