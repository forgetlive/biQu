package main

import "fmt"
import "log"
import "net/http"
import "time"
import "io/ioutil"
import "github.com/PuerkitoBio/goquery"
import "fmt"
import "strings"


func main(){
  url := "https://www.biqukan.com/38_38101/"
  page, err :=Fetch(url)
  if err != nil {
    fmt.Println(err)
  }
 Parse(page)
}



func Fetch(url string) (page string, err error) {
  client := &http.Client{
    Timeout: 5 * time.Second,
  }

  requests, err := http.NewRequest("GET", url, nil)
  if err != nil {
    log.Fatal(err)
  }
  requests.Header.Set("User-Agent", "Not Firefox")

  response, err := client.Do(requests)
  if err != nil {
    log.Fatal(err)
  }
  defer response.Body.Close()
  if response.StatusCode != 200 {
    log.Fatal("web not access")
  }
  bodyBytes, err := ioutil.ReadAll(response.Body)
  if err != nil {
    log.Fatal(err)
  }
  _ = bodyBytes
  return  string(bodyBytes), nil
}



func Parse(page string){
  document,_ := goquery. NewDocumentFromReader(strings.NewReader(page))
  document.Find("dd >a ").Each(ProcessElement)
}

func ProcessElement(index int ,element *goquery.Selection) {
  text := element.Text()
  link, ok := element.Attr("href")
  if ok {
    fmt.Printf("%s--%s\n", text, "https://www.xbiquge6.com" + link)
  }

}

