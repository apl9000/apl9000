package data

import (
	"fmt"
	"time"
)

type WebsiteResponseData struct {
	Metadata Metadata `json:"metadata"`
	Blog     Blog     `json:"blog"`
}

type Metadata struct {
	Url         string   `json:"url"`
	Title       string   `json:"title"`
	Description []string `json:"description"`
	Socials     []Social `json:"socials"`
}

type Social struct {
	Name string `json:"name"`
	Href string `json:"href"`
}

type Blog struct {
	Posts       []BlogPost `json:"posts"`
	Description string     `json:"description"`
}

type BlogPost struct {
	Title   string `json:"title"`
	Summary string `json:"summary"`
	Slug    string `json:"slug"`
	PubDate string `json:"publishedAt"`
}

type Website struct {
	Title       string
	Description []string
	Url         string
	Socials     []Social
	BlogPosts   []BlogPost
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
	url := "https://www.apl.directory/api"

	response, err := MakeRequest(url, &WebsiteResponseData{})

	if err != nil {
		fmt.Println("Error fetching website data:", err)
		return Website{}
	}
	data := response.(*WebsiteResponseData)

	// Reverse the order of blog posts
	posts := data.Blog.Posts
	for i, j := 0, len(posts)-1; i < j; i, j = i+1, j-1 {
		posts[i], posts[j] = posts[j], posts[i]
	}

	return Website{
		Title:       data.Metadata.Title,
		Description: data.Metadata.Description,
		Url:         data.Metadata.Url,
		Socials:     data.Metadata.Socials,
		BlogPosts:   posts,
	}
}
