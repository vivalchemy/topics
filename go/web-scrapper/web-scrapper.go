package webscrapper

import (
	"errors"
	"fmt"
	// "html"
	"golang.org/x/net/html"
	"net/http"
)

func Scrapper(url string) {
	// http.ListenAndServe(":8080", nil)
	web_response, err := http.Get(url)
	if err != nil {
		errors.New("Error: " + err.Error())
	}
	// fmt.Println(web_response)
	htmlDoc, _ := html.Parse(web_response.Body)
	// fmt.Println(htmlDoc)
	printNode(htmlDoc)
	// fmt.Println(htmlDoc.Data)

}

func printNode(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("<%s>\n", n.Data)
	} else if n.Type == html.TextNode {
		fmt.Print(n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		printNode(c)
	}
	if n.Type == html.ElementNode {
		fmt.Printf("</%s>\n", n.Data)
	}
}
