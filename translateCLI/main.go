package main

import (
	"fmt"
	"flag"
	"os"
	"strings"
	"sync"
	"github.com/r0meroh/goProjects/translateCLI/cli"
)

var sourceLanguage string
var targetLanguage string
var sourceText string

var wg sync.WaitGroup

func init(){
	flag.StringVar(&sourceLanguage, "s", "en", "Source language[en]")
	flag.StringVar(&targetLanguage, "t", "fr", "Target language[fr]")
	flag.StringVar(&sourceText, "st", "", "Test to translate")
}


func main() {
	flag.Parse()

	
	if flag.NFlag() == 0{
		fmt.Println("options: \n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	strChannel := make(chan string)
	wg.Add(1)

	reqBody := &cli.RequestBody{
		SourceLanguage: sourceLanguage,
		TargetLanguage: targetLanguage,
		SourceText: sourceText,
	}
	
	go cli.RequestTranslate(reqBody, strChannel, &wg)

	processedString := strings.ReplaceAll(<-strChannel, "+", " ")
	fmt.Printf("%s\n",processedString)
	close(strChannel)
	wg.Wait()
}