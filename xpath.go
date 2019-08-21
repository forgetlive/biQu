package main

import (
  "fmt"
  "github.com/antchfx/htmlquery"
)

func main(){
  //contentUrl:="https://www.xbiquge6.com/9_9208/5095645.html"
  //contentXpath :=  "//*[@id='content']"

  linkUrl :="https://www.xbiquge6.com/9_9208/"
  linkXpath :="//dd/a"
  ParseLink(linkUrl, linkXpath)
}



func ParseContent(url, xpath string ) {
  doc, err := htmlquery.LoadURL(url)
  if err != nil {
    fmt.Println(err)
  }
   list := htmlquery.Find(doc, xpath)
  for _, n := range list {
   // fmt.Println(htmlquery.OutputHTML(n, false))
     fmt.Println(htmlquery.InnerText(n))
  }
}

func ParseLink(url, xpath string ) {
  doc, err := htmlquery.LoadURL(url)
  if err != nil {
    fmt.Println(err)
  }
   list := htmlquery.Find(doc, xpath)
  for _, n := range list {
     //fmt.Println(htmlquery.InnerText(n))
     text :=htmlquery.InnerText(n)
     link :=htmlquery.SelectAttr(n, "href")
     fmt.Printf("%s %s\n", text, "https://xbiquge6.com"+ link )
  }
}
