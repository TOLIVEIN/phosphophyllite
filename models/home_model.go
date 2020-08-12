package models

import (
	"bytes"
	"fmt"
	"html/template"
	"phosphophyllite/config"
	"phosphophyllite/utils"
	"strconv"
	"strings"
)

//HomeBlockParam ...
type HomeBlockParam struct {
	ID         int
	Title      string
	Tags       []TagLink
	Brief      string
	Content    string
	Author     string
	Createtime string

	Link string

	UpdateLink string
	DeleteLink string

	IsLogin bool
}

//TagLink ...
type TagLink struct {
	TagName string
	TagURL  string
}

//HomeFooterPage ...
type HomeFooterPage struct {
	HasPre   bool
	HasNext  bool
	Page     string
	PreLink  string
	NextLink string
}

//MakeHomeBlocks ...
func MakeHomeBlocks(articles []Article, isLogin bool) template.HTML {
	htmlHome := ""
	for _, article := range articles {
		homeParam := HomeBlockParam{}
		homeParam.ID = article.ID
		homeParam.Title = article.Title
		homeParam.Tags = createTagsLinks(article.Tags)
		fmt.Println("tag-->", article.Tags)
		homeParam.Brief = article.Brief
		homeParam.Content = article.Content
		homeParam.Author = article.Author
		homeParam.Createtime = utils.SwitchTimeStampToData(article.Createtime)
		homeParam.Link = "/show/" + strconv.Itoa(article.ID)
		homeParam.UpdateLink = "/article/update?id=" + strconv.Itoa(article.ID)
		homeParam.DeleteLink = "/article/delete?id=" + strconv.Itoa(article.ID)
		homeParam.IsLogin = isLogin

		fmt.Println("-----------", homeParam)
		t, _ := template.ParseFiles("views/home_block.html")
		buffer := bytes.Buffer{}

		t.Execute(&buffer, homeParam)
		htmlHome += buffer.String()
	}

	fmt.Println("htmlHome-->", htmlHome)
	return template.HTML(htmlHome)
}

//createTagsLinks ...
func createTagsLinks(tags string) []TagLink {
	var tagLink []TagLink
	tagsParam := strings.Split(tags, "&")
	for _, tag := range tagsParam {
		tagLink = append(tagLink, TagLink{tag, "/?tag=" + tag})
	}
	return tagLink
}

//SetHomeFooterPage ...
func SetHomeFooterPage(p int) HomeFooterPage {
	page := HomeFooterPage{}

	num := GetArticleRowNums()
	allPageNum := (num-1)/config.NUM + 1
	page.Page = fmt.Sprintf("%d/%d", p, allPageNum)

	if p <= 1 {
		page.HasPre = false
	} else {
		page.HasPre = true
	}

	if p >= allPageNum {
		page.HasNext = false
	} else {
		page.HasNext = true
	}

	page.PreLink = "/?page=" + strconv.Itoa(p-1)
	page.NextLink = "/?page=" + strconv.Itoa(p+1)
	return page
}
