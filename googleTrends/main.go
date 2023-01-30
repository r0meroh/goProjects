package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel *Channel `xml:"channel"`
}

type Channel struct {
	Title    string `xml:"title"`
	ItemList []Item `xml:"item"`
}
type Item struct {
	Title     string `xml:"title"`
	Link      string `xml:"link"`
	Traffic   string `xml:"approx_traffic"`
	NewsItems []News `xml:"news_item"`
}

type News struct {
	Headline     string `xml:"news_item_title"`
	HeadlineLink string `xml:"news_item_url"`
}

func main() {
	fmt.Println("hello")
	var r RSS
	data := readGoogleTrends()
	err := xml.Unmarshal(data, &r)
	if err != nil{
		fmt.Println(err)
	}

	for i := range r.Channel.ItemList {
		rank := (i + 1)
		fmt.Println("#", rank)
		fmt.Println("Search Term: ", r.Channel.ItemList[i].Title)
		fmt.Println("Link to the Trend: ", r.Channel.ItemList[i].Link)
		fmt.Println("Headline: ", r.Channel.ItemList[i].NewsItems[0].Headline)
		fmt.Println("Link to article: ", r.Channel.ItemList[i].NewsItems[0].HeadlineLink)
		fmt.Println("------------------------------------------")
	}

}

func getGoogleTrends() *http.Response{
	resp, err := http.Get("https://trends.google.com/trends/trendingsearches/daily/rss?geo=US")
	if err != nil {
		fmt.Println(err)
	}
	return resp
}

func readGoogleTrends()[]byte{
	resp := getGoogleTrends()
	data, err := ioutil.ReadAll((resp.Body))
	if err != nil {
		fmt.Println(err)
	}
	return data
}


