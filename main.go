package main

import (
	"apl9000/data"
	"bytes"
	"fmt"
	"log"
	"os"
	"text/template"
)

type ProfileData struct {
	Timestamp       string
	Website         data.Website
	WeatherForecast data.WeatherForecast
	Rates           data.Rates
}

func main() {
	template, err := template.ParseFiles("template/index.html")

	if err != nil {
		fmt.Printf("Error parsing template: %v\n", err)
		return
	}

	data := ProfileData{
		Timestamp:       data.GetCurrentTime(),
		WeatherForecast: data.GetWeatherForecast(),
		Rates:           data.GetRates(),
		Website:         data.GetWebsiteData(),
	}

	var htmlBuffer bytes.Buffer
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
