package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"text/template"
)

type ProfileData struct {
	ForeCast string
}

func main() {
	template, err := template.ParseFiles("template/index.html")

	if err != nil {
		fmt.Printf("Error parsing template: %v\n", err)
		return
	}

    // Create a buffer to hold the HTML output
    var htmlBuffer bytes.Buffer

		data := ProfileData{
			ForeCast: "Sunny",
		}
		// Execute the template with the data object
    err = template.Execute(&htmlBuffer, data)

		if err != nil {
			fmt.Printf("Error executing template: %v\n", err)
			return
		}

    htmlContent := htmlBuffer.String()


		    // Write the markdown content to a file
    filename := "README.md"
    err = os.WriteFile(filename, []byte(htmlContent), 0644)
    if err != nil {
        log.Fatal("Error writing Markdown to file: ", err)
    }

    log.Println("Markdown written successfully to", filename)
}
