package models

import (
	"fmt"
	"phosphophyllite/config"
	"phosphophyllite/database"
)

//Article ...
type Article struct {
	ID         int
	Title      string
	Tags       string
	Brief      string
	Content    string
	Author     string
	Createtime int64
	// Status int
}

var articleRowNum = 0

//GetArticleRowNums ...
func GetArticleRowNums() int {
	if articleRowNum == 0 {
		articleRowNum = QueryArticleRowNum()
	}
	return articleRowNum
}

//QueryArticleRowNum ...
func QueryArticleRowNum() int {
	row := database.QueryRowDB("select count(id) from article")
	num := 0
	row.Scan(&num)
	return num
}

//SetArticleRowNum ...
func SetArticleRowNum() {
	articleRowNum = QueryArticleRowNum()
}

//AddArticle ...
func AddArticle(article Article) (int64, error) {
	i, err := InsertArticle(article)
	SetArticleRowNum()
	return i, err
}

//InsertArticle ...
func InsertArticle(article Article) (int64, error) {
	return database.ModifyDB("insert into article(title,tags,brief,content,author,createtime) values(?,?,?,?,?,?)",
		article.Title, article.Tags, article.Brief, article.Content, article.Author, article.Createtime)
}

//FindArticleWithPage ...
func FindArticleWithPage(page int) ([]Article, error) {
	page--
	fmt.Println("----------------->page", page)
	return QueryArticleWithPage(page, config.NUM)
}

//QueryArticleWithPage ...
func QueryArticleWithPage(page, num int) ([]Article, error) {
	sql := fmt.Sprintf("limit %d,%d", page*num, num)
	return QueryArticleWithCondition(sql)
}

//QueryArticleWithCondition ...
func QueryArticleWithCondition(sql string) ([]Article, error) {
	sql = "select id, title, tags, brief, content, author, createtime from article " + sql
	// fmt.Println(sql)
	rows, err := database.QueryDB(sql)
	if err != nil {
		return nil, err
	}

	var articleList []Article
	for rows.Next() {
		id := 0
		title := ""
		tags := ""
		brief := ""
		content := ""
		author := ""
		var createtime int64
		createtime = 0
		rows.Scan(&id, &title, &tags, &brief, &content, &author, &createtime)
		article := Article{id, title, tags, brief, content, author, createtime}
		// fmt.Println("----------article: ", article)
		articleList = append(articleList, article)
	}
	return articleList, nil
}
