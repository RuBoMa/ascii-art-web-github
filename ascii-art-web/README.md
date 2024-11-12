# ascii-art-web-github

ASCII Art Generator

Description:
ASCII Art Generator is a Go-based tool that converts text input into ASCII art. This project is designed for fun and learning, allowing users to create customized text art for web displays.

Authors: Fanni & Roope

Features:
Text-to-ASCII art conversion
Supports multiple ASCII styles (standard, shadow, thinkertoy)
web-based usage

Usage:

# Start the Server: by running the following command in your terminal: go run main.go

# Open in Browser: In your browser, navigate to http://localhost:8080

After page is live, just simply type text to the input field and press "generate"

# Generate ASCII Art:

# Select a font style from the options on the webpage.

# Enter your desired text in the input field. Note: max 500 characters, only printable ascii-characters

# Click the Generate button to see your text transformed into ASCII art!

Implemention details:

Input Handling: The algorithm first reads user input and validates it to ensure it’s compatible with the ASCII art style. (standard, shadow, thinkertoy)

Text-to-ASCII Mapping: Each character in the input is matched with pre-defined ASCII patterns. The program retrieves the correct ASCII representation for each character.

Line Assembly: Characters are assembled line by line to create the final ASCII image, ensuring each row aligns properly.

Output Rendering: Finally, the assembled ASCII text is formatted and displayed on the webpage.
