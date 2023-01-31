package cli

import (
	"fmt"
	"net/http"
	"sync"
	"github.com/Jeffail/gabs"
)

type RequestBody struct {
	SourceLanguage string
	TargetLanguage string
	SourceText string
}

const translateUrl = "https://translate.googleapis.com/translate_a/single"

func RequestTranslate(body *RequestBody, str chan string, wg *sync.WaitGroup){
	client := &http.Client{}
	req, err := http.NewRequest("Get", translateUrl, nil)
	if err != nil {
		fmt.Println(err)
	} 
	query := req.URL.Query()
	query.Add("client", "gtx")
	query.Add("sl", body.SourceLanguage)
	query.Add("tl", body.TargetLanguage)
	query.Add("dt", "t")
	query.Add("q", body.SourceText)

	req.URL.RawQuery = query.Encode()

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusTooManyRequests{
		str <- "You have been rate limited"
		wg.Done()
		return
	}

	parseJson, err := gabs.ParseJSONBuffer(res.Body)
	if err != nil {
		fmt.Println("problem with gabs library")
	}

	nestOne, err := parseJson.ArrayElement(0)
	if err != nil{
		fmt.Println("error in un-nesting elements from response")
	}
	nestTwo, err := nestOne.ArrayElement(0)
	if err != nil{
		fmt.Println("error in un-nesting first element from nested array")
	}

	translatedString, err := nestTwo.ArrayElement(0)
	if err != nil{
		fmt.Println("error in translated final string")
	}
	str <- translatedString.Data().(string)
	wg.Done()
}  
