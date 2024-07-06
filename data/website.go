package data

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Define the structures based on the RSS format.
type Rss struct {
	Channel Channel `xml:"channel"`
}

type Channel struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Items       []Item `xml:"item"`
	Link        string `xml:"link"`
}

type Item struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Link        string `xml:"link"`
	PubDate     string `xml:"pubDate"`
}

type Website struct {
	Title       string
	Description string
	Link        string
	BlogPosts   []BlogPost
}

type BlogPost struct {
	Title       string
	Description string
	Link        string
	PubDate     string
}

func parseDate(dateString string) string {
	// Layout to parse the original date string
	layout := "Mon, 02 Jan 2006 15:04:05 MST"

	// Parse the date string
	parsedTime, err := time.Parse(layout, dateString)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return ""
	}
	// Format the parsed time to the new layout
	return parsedTime.Format("2006-01-02")
}

func GetWebsiteData() Website {
	url := "https://www.apl.directory/rss"

	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return Website{}
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return Website{}
	}

	var rss Rss
	if err := xml.Unmarshal(body, &rss); err != nil {
		fmt.Printf("Error parsing the feed: %v\n", err)
		return Website{}
	}

	var blogPosts []BlogPost
	for _, item := range rss.Channel.Items {
		// format pubDate

		// string formattedPubDate := item.PubDate
		blogPosts = append(blogPosts, BlogPost{
			Title:       item.Title,
			Description: item.Description,
			Link:        item.Link,
			PubDate:     parseDate(item.PubDate),
		})
	}

	return Website{
		Title:       rss.Channel.Title,
		Description: rss.Channel.Description,
		Link:        rss.Channel.Link,
		BlogPosts:   blogPosts,
	}
}
