package main

import (
	"bufio"
	"fmt"
	"os"
	_ "strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	dateString := dateString()
	filename := dateString + ".md"

	//create markdown file
	createMarkDown(dateString, filename)

	scrape("", filename)
	scrape("go", filename)
	scrape("javascript", filename)
	scrape("rust", filename)
	scrape("python", filename)

}

func dateString() string {
	y, m, d := time.Now().Date()
	mStr := fmt.Sprintf("%d", m)
	dStr := fmt.Sprintf("%d", d)
	if m < 10 {
		mStr = fmt.Sprintf("0%d", m)
	}
	if d < 10 {
		dStr = fmt.Sprintf("0%d", d)
	}
	return fmt.Sprintf("%d-%s-%s", y, mStr, dStr)

}

func createMarkDown(date string, filename string) {

	// open output file
	fo, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	// make a write buffer
	w := bufio.NewWriter(fo)
	w.WriteString("### " + date + "\n")
	w.Flush()
}

func scrape(language string, filename string) {
	var doc *goquery.Document
	var e error
	// var w *bufio.Writer

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if language == "" {
		if _, err = f.WriteString(fmt.Sprintf("\n#### %s\n", "Top")); err != nil {
			panic(err)
		}

		if doc, e = goquery.NewDocument("https://github.com/trending"); e != nil {
			panic(e.Error())
		}
	} else {
		if _, err = f.WriteString(fmt.Sprintf("\n#### %s\n", language)); err != nil {
			panic(err)
		}

		if doc, e = goquery.NewDocument(fmt.Sprintf("https://github.com/trending?l=%s", language)); e != nil {
			panic(e.Error())
		}
	}

	doc.Find("ol.repo-list li").Each(func(i int, s *goquery.Selection) {
		title := strings.TrimSpace(s.Find("h3 a").Text())
		owner := s.Find("span.prefix").Text()
		description := s.Find("p.col-9").Text()
		description = strings.TrimSpace(description)
		url, _ := s.Find("h3 a").Attr("href")
		url = "https://github.com" + url
		ownerImg, _ := s.Find("p.repo-list-meta a img").Attr("src")
		fmt.Println("title: ", title)
		fmt.Println("owner: ", owner)
		fmt.Println("URL: ", url)
		fmt.Println("Owner Img: ", ownerImg)
		if _, err = f.WriteString("* [" + title + "](" + url + "): " + description + "\n"); err != nil {
			println(err.Error())
			panic(err)
		}
	})
}
