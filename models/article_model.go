package models

import "phosphophyllite/database"

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

//AddArticle ...
func AddArticle(article Article) (int64, error) {
	i, err := InsertArticle(article)
	return i, err
}

//InsertArticle ...
func InsertArticle(article Article) (int64, error) {
	return database.ModifyDB("insert into article(title,tags,brief,content,author,createtime) values(?,?,?,?,?,?)",
		article.Title, article.Tags, article.Brief, article.Content, article.Author, article.Createtime)
}
