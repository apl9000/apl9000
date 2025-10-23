package data

import (
	"fmt"
)

type WebsiteResponseData struct {
	Metadata      Metadata      `json:"metadata"`
	Blog          Blog          `json:"blog"`
	QuoteOfTheDay QuoteOfTheDay `json:"quoteOfTheDay"`
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

type QuoteOfTheDay struct {
	Text   string `json:"text"`
	Author string `json:"author"`
}

type Website struct {
	Title         string
	Description   []string
	Url           string
	Socials       []Social
	BlogPosts     []BlogPost
	QuoteOfTheDay QuoteOfTheDay
}

func GetWebsiteData() Website {
	url := "https://www.apl.directory/api"

	response, err := MakeRequest(url, &WebsiteResponseData{})

	if err != nil {
		fmt.Println("Error fetching website data:", err)
		return Website{}
	}
	data := response.(*WebsiteResponseData)

	posts := data.Blog.Posts

	return Website{
		Title:       data.Metadata.Title,
		Description: data.Metadata.Description,
		Url:         data.Metadata.Url,
		Socials:     data.Metadata.Socials,
		BlogPosts:   posts,
		QuoteOfTheDay: QuoteOfTheDay{
			Text:   data.QuoteOfTheDay.Text,
			Author: data.QuoteOfTheDay.Author,
		},
	}
}
