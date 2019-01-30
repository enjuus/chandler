package main

import (
  "fmt"
  "log"
  "path"
  "io"
  "regexp"
  flag "github.com/ogier/pflag"
  "github.com/PuerkitoBio/goquery"
  "net/http"
  "os"
)

var help bool
var name bool
var source string
var destination string
var threadName string

func init() {
  flag.BoolVarP(&help, "help", "h", false, "Display this help message")
  flag.BoolVarP(&name, "name", "n", false, "Store images in a subfolder named after the thread")
  flag.StringVarP(&source, "source", "s", "", "The thread to download [required]")
  flag.StringVarP(&destination, "destination", "d", "", "The path to save [required]")
  flag.Parse()
}

func GetImages() {
  resp, err := http.Get(source)
  if err != nil {
    panic(err)
  }

  defer resp.Body.Close()

  document, err := goquery.NewDocumentFromReader(resp.Body)
  if err != nil {
    log.Fatal("Error loading HTTP response body.", err)
  }

    reg, _ := regexp.Compile("[^a-zA-Z0-9 ]+")
    folder := document.Find(".board .subject:first-of-type").Text()
  if name == true {
    destination = destination + "/" + reg.ReplaceAllString(folder, "")
  }

  // find images
  document.Find(".board a.fileThumb").Each(func(index int, element *goquery.Selection) {
    imgSrc, exists := element.Attr("href")
    if exists {
      fmt.Println(imgSrc)
      StoreFile(imgSrc)
    }
  })
}

func StoreFile(imgSrc string) {
  if _, err := os.Stat(destination); os.IsNotExist(err) {
    os.MkdirAll(destination, os.ModePerm);
  }

  response, err := http.Get("https:" + imgSrc)
  if err != nil {
    log.Fatal(err)
  }

  defer response.Body.Close()

  file, _ := os.Create(destination + "/" + path.Base(imgSrc))
  defer file.Close()

  io.Copy(file, response.Body)
}

func main() {
  if help == true {
    PrintHelpMessage()
  }
  if source == "" && destination == "" {
    PrintHelpMessage()
  }

  GetImages()

}

func PrintHelpMessage() {
  fmt.Printf("Usage: %s [options]\n", os.Args[0])
  fmt.Printf("Options:\n")
  flag.PrintDefaults()
  os.Exit(1)
}
