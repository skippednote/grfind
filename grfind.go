package grfind

import (
	"encoding/xml"
	"net/http"
	"net/url"
)

const (
	uri = "https://goodreads.com/"
)

// GRfind struct to setup GoodReads key and secret.
type GRfind struct {
	Client http.Client
	Key    string
	Secret string
}

type response struct {
	Request request `xml:"Request" json:"request"`
	Search  search  `xml:"search" json:"search"`
}

type request struct {
	Authentication string `xml:"authentication" json:"authentication"`
	Key            string `xml:"key" json:"key"`
}

type search struct {
	Query   string  `xml:"query" json:"query"`
	Results results `xml:"results" json:"results"`
}

type results struct {
	Work []Work `xml:"work" json:"work"`
}

// Work struct contains properties of a book.
type Work struct {
	ID       int      `xml:"id" json:"id"`
	BestBook bestbook `xml:"best_book" json:"best_book"`
}

type bestbook struct {
	ID            int    `xml:"id" json:"id"`
	Title         string `xml:"title" json:"title"`
	Author        author `xml:"author" json:"author"`
	ImageURL      string `xml:"image_url" json:"image_url"`
	SmallImageURL string `xml:"small_image_url" json:"small_image_url"`
}

type author struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

func (g *GRfind) generateURL(endpoint string, params map[string]string) string {
	v := url.Values{}
	v.Set("key", g.Key)
	for key, value := range params {
		v.Set(key, value)
	}
	query := v.Encode()

	return uri + endpoint + "?" + query
}

// GetBooks returns an array of books.
func (g *GRfind) GetBooks(q string) ([]Work, error) {
	url := g.generateURL("search/index.xml", map[string]string{"q": q})

	res, err := g.Client.Get(url)
	if err != nil {
		return []Work{}, err
	}
	defer res.Body.Close()

	var gr response
	d := xml.NewDecoder(res.Body)
	err = d.Decode(&gr)
	if err != nil {
		return []Work{}, err
	}
	return gr.Search.Results.Work, nil
}
