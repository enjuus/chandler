package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	flag "github.com/ogier/pflag"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"
	"time"
)

var help bool
var source string
var destination string
var watcher bool
var verbose bool
var interval int

func init() {
	flag.BoolVarP(&help, "help", "h", false, "Display this help message")
	flag.StringVarP(&source, "source", "s", "", "The thread to download [required]")
	flag.StringVarP(&destination, "destination", "d", "", "The path to save to. See README for more options. [required]")
	flag.BoolVarP(&watcher, "watcher", "w", false, "Watch the thread for new files")
	flag.BoolVarP(&verbose, "verbose", "v", false, "Enable output")
	flag.IntVarP(&interval, "interval", "i", 20, "The times to check per second")
	flag.Parse()
}

func GetImages() {
	resp, err := http.Get(source)
	if err != nil {
		log.Fatal("Thread cannot be found")
	}

	defer resp.Body.Close()

	document, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body.", err)
	}

	RenameFilePath(document)
	// find images
	document.Find(".board a.fileThumb").Each(func(index int, element *goquery.Selection) {
		imgSrc, exists := element.Attr("href")
		if exists {
			StoreFile(imgSrc)
		}
	})
}

func RenameFilePath(document *goquery.Document) {
	reg, _ := regexp.Compile("[^a-zA-Z0-9 ]+")
	parts := strings.Split(source, "/")
	thread := reg.ReplaceAllString(document.Find(".board .subject:first-of-type").Text(), "")
	threadId := parts[5]
	board := parts[3]

	destination = strings.Replace(destination, "{BOARD}", board, -1)
	if thread != "" {
		destination = strings.Replace(destination, "{THREAD}", thread, -1)
	} else {
		destination = strings.Replace(destination, "{THREAD}", threadId, -1)
	}
	destination = strings.Replace(destination, "{THREADID}", threadId, -1)
}

func StoreFile(imgSrc string) {
	if _, err := os.Stat(destination); os.IsNotExist(err) {
		os.MkdirAll(destination, os.ModePerm)
	}

	response, err := http.Get("https:" + imgSrc)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	path := destination + "/" + path.Base(imgSrc)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		Output("Storing new file")
		file, _ := os.Create(path)
		defer file.Close()
		io.Copy(file, response.Body)
	}
}

func doEvery(d time.Duration, f func()) {
	for range time.Tick(d) {
		GetImages()
	}
}

func Output(s string) {
	if verbose == true {
		fmt.Println(s)
	}
}

func main() {
	if help == true {
		PrintHelpMessage()
	}

	if source == "" && destination == "" {
		PrintHelpMessage()
	}
	if watcher == true {
		Output("Starting watcher")
		GetImages()
		doEvery(time.Duration(interval)*time.Second, GetImages)
	} else {
		GetImages()
	}

}

func PrintHelpMessage() {
	fmt.Printf("Usage: %s [options]\n", os.Args[0])
	fmt.Printf("Options:\n")
	flag.PrintDefaults()
	os.Exit(1)
}
