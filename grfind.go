package grfind

import (
	"encoding/xml"
	"net/http"
	"net/url"
)

const (
	uri = "https://goodreads.com/"
)

type GRfindInfo struct {
	Client http.Client
	Key    string
	Secret string
}

type GoodreadsResponse struct {
	Request grfindRequest `xml:"Request" json:"request"`
	Search  grfindSearch  `xml:"search" json:"search"`
}

type grfindRequest struct {
	Authentication string `xml:"authentication" json:"authentication"`
	Key            string `xml:"key" json:"key"`
}

type grfindSearch struct {
	Query   string        `xml:"query" json:"query"`
	Results grfindResults `xml:"results" json:"results"`
}

type grfindResults struct {
	Work []grfindWork `xml:"work" json:"work"`
}

type grfindWork struct {
	ID       int            `xml:"id" json:"id"`
	BestBook grfindBestBook `xml:"best_book" json:"best_book"`
}

type grfindBestBook struct {
	ID            int          `xml:"id" json:"id"`
	Title         string       `xml:"title" json:"title"`
	Author        grfindAuthor `xml:"author" json:"author"`
	ImageURL      string       `xml:"image_url" json:"image_url"`
	SmallImageURL string       `xml:"small_image_url" json:"small_image_url"`
}

type grfindAuthor struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

func (g *GRfindInfo) generateURL(endpoint string, params map[string]string) string {
	v := url.Values{}
	v.Set("key", g.Key)
	for key, value := range params {
		v.Set(key, value)
	}
	query := v.Encode()

	return uri + endpoint + "?" + query
}

func (g *GRfindInfo) GetBooks(q string) ([]grfindWork, error) {
	url := g.generateURL("search/index.xml", map[string]string{"q": q})

	res, err := g.Client.Get(url)
	if err != nil {
		return []grfindWork{}, err
	}
	defer res.Body.Close()

	var gr GoodreadsResponse
	d := xml.NewDecoder(res.Body)
	err = d.Decode(&gr)
	if err != nil {
		return []grfindWork{}, err
	}
	return gr.Search.Results.Work, nil
}
