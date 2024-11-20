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
	if input == "" {
		return "Your didn't enter any input", false
	}

	input = strings.ReplaceAll(input, "\r\n", "\n")
	nonPrintable := make(map[rune]bool)

	for _, char := range input {
		if (char < 32 || char > 127) && char != '\n' {
			nonPrintable[char] = true // mapping all unique invalid characters
		}
	}

	// If there are any invalid characters, returning the error message with the input and invalid characters
	if len(nonPrintable) > 0 {
		var charSlice []string
		for char := range nonPrintable {
			charSlice = append(charSlice, string(char))
		}
		invalidChars := strings.Join(charSlice, ", ")
		log.Println("Invalid character in input: ", invalidChars)
		errormsg := "\n\nCouldn't create Ascii-Art due to invalid characters.\n\nYou entered: " + input + "\n\n" + "These characters are not allowed: " + invalidChars + "\n\n"
		return errormsg, false
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
