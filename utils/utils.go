package utils

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"html/template"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/russross/blackfriday"
	"github.com/sourcegraph/syntaxhighlight"
)

//MD5 ...
func MD5(str string) string {
	md5str := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return md5str
}

//SwitchTimeStampToData ...
func SwitchTimeStampToData(timeStamp int64) string {
	t := time.Unix(timeStamp, 0)
	return t.Format("2006-01-02 15:04:05")
}

//SwitchMarkdownToHTML ...
func SwitchMarkdownToHTML(content string) template.HTML {
	markdown := blackfriday.MarkdownCommon([]byte(content))

	doc, _ := goquery.NewDocumentFromReader(bytes.NewReader(markdown))

	doc.Find("code").Each(func(i int, selection *goquery.Selection) {
		light, _ := syntaxhighlight.AsHTML([]byte(selection.Text()))
		selection.SetHtml(string(light))
		fmt.Println(selection.Html())
		fmt.Println("light:", string(light))
	})
	htmlString, _ := doc.Html()
	return template.HTML(htmlString)
}
