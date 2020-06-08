package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strings"
)

func main() {
	url := "https://www.cdnbus.in/ATID-389"
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Add("User-Agent", "Mozilla/6.0")
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	defer resp.Body.Close()
	if 200 != resp.StatusCode {
		fmt.Println("status error:", resp.StatusCode, resp.Status)
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println("err:", err)
	}
	bigImage := doc.Find(".bigImage img")
	fmt.Println(bigImage.AttrOr("title", ""))
	fmt.Println(bigImage.AttrOr("src", ""))
	info := doc.Find(".header")
	info.Each(func(i int, selection *goquery.Selection) {
		item := selection.Text()
		value := ""
		if strings.HasPrefix(item, "發行日期:") {
			value = selection.Parent().Text()
			value = strings.Replace(value, "發行日期:", "", 1)
		} else if strings.HasPrefix(item, "長度:") {
			value = selection.Parent().Text()
			value = strings.Replace(value, "長度:", "", 1)
		} else if strings.HasPrefix(item, "演員") {
			stars := doc.Find(".star-name")
			stars.Each(func(i int, selection *goquery.Selection) {
				starName := selection.Text()
				value += strings.TrimSpace(starName)
			})
		} else {
			value = selection.Next().Text()
		}
		fmt.Println(item, value)
	})

}
