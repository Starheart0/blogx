package main

import (
	"fmt"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	reader, err := os.Open("uploads/index.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		fmt.Println(err)
		return
	}
	selection := doc.Find("title")
	fmt.Println(selection.Text())
	selection.SetText("Starheart")
	fmt.Println(selection.Text())
}
