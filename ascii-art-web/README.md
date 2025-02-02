ascii-art-web-github

# ASCII Art Generator

Description:
ASCII Art Generator is a Go-based tool that converts text input into ASCII art.
This project is designed for fun and learning, allowing users to create customized text art.

Authors: Fanni & Roope

Features:
Text-to-ASCII art conversion
Supports multiple ASCII styles (standard, shadow, thinkertoy)
web-based usage

## Usage:

1. Start the Server: by running the following command in your terminal: go run main.go

2. Open in Browser: In your browser, navigate to http://localhost:8080

3. Generate ASCII Art:

- Select a font style from the options on the webpage.

- Enter your desired text in the input field. Note: max 500 characters, only printable ascii-characters

- Click the Generate button to see your text transformed into ASCII art!

### Alternatively start the server in Docker:

1. Run following commands in the terminal:

- chmod + x build.sh
- ./build.sh

Implemention details:

Input Handling: The algorithm first reads user input and validates it to ensure it’s compatible with the ASCII art style. (standard, shadow, thinkertoy)

Text-to-ASCII Mapping: Each character in the input is matched with pre-defined ASCII patterns. The program retrieves the correct ASCII representation for each character.

Line Assembly: Characters are assembled line by line to create the final ASCII image, ensuring each row aligns properly.

Output Rendering: Finally, the assembled ASCII text is formatted and displayed on the webpage.
