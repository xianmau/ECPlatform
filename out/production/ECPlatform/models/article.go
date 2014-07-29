package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"utils/global"
)

type Article struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Category string `json:"category"`
	Content  string `json:"content"`
	Time     string `json:"time"`
	Click    int    `json:"click"`
	Status   int    `json:"status"`
}

func GetArticle(Id string) (*Article, error) {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("select * from `tb_article` where `Id`=?", Id)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		var article Article
		var Id int
		var Title string
		var Category string
		var Content string
		var Time string
		var Click int
		var Status int
		if err := rows.Scan(&Id, &Title, &Category, &Content, &Time, &Click, &Status); err != nil {
			return nil, err
		}
		article.Id = Id
		article.Title = Title
		article.Category = Category
		article.Content = Content
		article.Time = Time
		article.Click = Click
		article.Status = Status
		return &article, nil
	}
	return nil, nil
}

func GetArticleList() ([]Article, error) {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("select * from `tb_article` where true order by `Id` desc")
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	articleList := []Article{}
	for rows.Next() {
		var article Article
		var Id int
		var Title string
		var Category string
		var Content string
		var Time string
		var Click int
		var Status int
		if err := rows.Scan(&Id, &Title, &Category, &Content, &Time, &Click, &Status); err != nil {
			return nil, err
		}
		article.Id = Id
		article.Title = Title
		article.Category = Category
		article.Content = Content
		article.Time = Time
		article.Click = Click
		article.Status = Status
		articleList = append(articleList, article)
	}
	return articleList, nil
}

func CreateArticle(Title string, Category string, Content string, Status string) error {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return err
	}
	_, err = db.Exec("insert into `tb_article`(`Title`,`Category`,`Content`,`Status`) values(?,?,?,?)", Title, Category, Content, Status)
	if err != nil {
		return err
	}
	return nil
}

func EditArticle(Id string, Title string, Category string, Content string, Status string) error {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return err
	}
	_, err = db.Exec("update `tb_article` set `Title`=?,`Category`=?,`Content`=?,`Status`=? where `Id`=?", Title, Category, Content, Status, Id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteArticle(Id string) error {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return err
	}
	_, err = db.Exec("delete from `tb_article` where `Id`=?", Id)
	if err != nil {
		return err
	}
	return nil
}
