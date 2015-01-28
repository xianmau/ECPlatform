package models

import (
	"database/sql"
	_ "utils/mysql"
	"utils/global"
)

type Article struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Category string `json:"category"`
	Abstract string `json:"abstract"`
	Image string `json:"image"`
	Content  string `json:"content"`
	Time     string `json:"time"`
	Click    int    `json:"click"`
	Status   int    `json:"status"`
	Remark   string `json:"remark"`
}

func GetArticle(Id string) (*Article, error) {
	db, err := sql.Open("mysql", global.Config["conn_str"])
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
		var Abstract string
		var Image string
		var Content string
		var Time string
		var Click int
		var Status int
		var Remark string
		if err := rows.Scan(&Id, &Title, &Category, &Abstract, &Image, &Content, &Time, &Click, &Status, &Remark); err != nil {
			return nil, err
		}
		article.Id = Id
		article.Title = Title
		article.Category = Category
		article.Abstract = Abstract
		article.Image = Image
		article.Content = Content
		article.Time = Time
		article.Click = Click
		article.Status = Status
		article.Remark = Remark
		return &article, nil
	}
	return nil, nil
}

func GetArticleByTitle(Title string) (*Article, error) {
	db, err := sql.Open("mysql", global.Config["conn_str"])
	defer db.Close()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("select * from `tb_article` where `Title`=?", Title)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		var article Article
		var Id int
		var Title string
		var Category string
		var Abstract string
		var Image string
		var Content string
		var Time string
		var Click int
		var Status int
		var Remark string
		if err := rows.Scan(&Id, &Title, &Category, &Abstract, &Image, &Content, &Time, &Click, &Status, &Remark); err != nil {
			return nil, err
		}
		article.Id = Id
		article.Title = Title
		article.Category = Category
		article.Abstract = Abstract
		article.Image = Image
		article.Content = Content
		article.Time = Time
		article.Click = Click
		article.Status = Status
		article.Remark = Remark
		return &article, nil
	}
	return nil, nil
}

func GetArticleListByCategory(Category string) ([]Article, error) {
	db, err := sql.Open("mysql", global.Config["conn_str"])
	defer db.Close()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("select * from `tb_article` where `Status`>0 and `Category`=? order by `Id` desc", Category)
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
		var Abstract string
		var Image string
		var Content string
		var Time string
		var Click int
		var Status int
		var Remark string
		if err := rows.Scan(&Id, &Title, &Category, &Abstract, &Image, &Content, &Time, &Click, &Status, &Remark); err != nil {
			return nil, err
		}
		article.Id = Id
		article.Title = Title
		article.Category = Category
		article.Abstract = Abstract
		article.Image = Image
		article.Content = Content
		article.Time = Time
		article.Click = Click
		article.Status = Status
		article.Remark = Remark
		articleList = append(articleList, article)
	}
	return articleList, nil
}

func GetHotArticles(Category string, Top string) ([]Article, error) {
	db, err := sql.Open("mysql", global.Config["conn_str"])
	defer db.Close()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("select * from `tb_article` where `Status`>0 and `Category`=? order by `Click` desc limit 0, ?", Category, Top)
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
		var Abstract string
		var Image string
		var Content string
		var Time string
		var Click int
		var Status int
		var Remark string
		if err := rows.Scan(&Id, &Title, &Category, &Abstract, &Image, &Content, &Time, &Click, &Status, &Remark); err != nil {
			return nil, err
		}
		article.Id = Id
		article.Title = Title
		article.Category = Category
		article.Abstract = Abstract
		article.Image = Image
		article.Content = Content
		article.Time = Time
		article.Click = Click
		article.Status = Status
		article.Remark = Remark
		articleList = append(articleList, article)
	}
	return articleList, nil
}

func GetArticleList() ([]Article, error) {
	db, err := sql.Open("mysql", global.Config["conn_str"])
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
		var Abstract string
		var Image string
		var Content string
		var Time string
		var Click int
		var Status int
		var Remark string
		if err := rows.Scan(&Id, &Title, &Category, &Abstract, &Image, &Content, &Time, &Click, &Status, &Remark); err != nil {
			return nil, err
		}
		article.Id = Id
		article.Title = Title
		article.Category = Category
		article.Abstract = Abstract
		article.Image = Image
		article.Content = Content
		article.Time = Time
		article.Click = Click
		article.Status = Status
		article.Remark = Remark
		articleList = append(articleList, article)
	}
	return articleList, nil
}

func CreateArticle(Title string, Category string, Abstract string, Image string, Content string, Status string, Remark string) error {
	db, err := sql.Open("mysql", global.Config["conn_str"])
	defer db.Close()
	if err != nil {
		return err
	}
	_, err = db.Exec("insert into `tb_article`(`Title`,`Category`,`Abstract`,`Image`,`Content`,`Status`,`Remark`) values(?,?,?,?,?,?,?)", Title, Category, Abstract, Image, Content, Status, Remark)
	if err != nil {
		return err
	}
	return nil
}

func EditArticle(Id string, Title string, Category string, Abstract string, Image string, Content string, Status string, Remark string) error {
	db, err := sql.Open("mysql", global.Config["conn_str"])
	defer db.Close()
	if err != nil {
		return err
	}
	_, err = db.Exec("update `tb_article` set `Title`=?,`Category`=?,`Abstract`=?,`Image`=?,`Content`=?,`Status`=?,`Remark`=? where `Id`=?", Title, Category, Abstract, Image, Content, Status, Remark, Id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteArticle(Id string) error {
	db, err := sql.Open("mysql", global.Config["conn_str"])
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

func IncreaseArticleReadTimes(Id string) error {
	db, err := sql.Open("mysql", global.Config["conn_str"])
	defer db.Close()
	if err != nil {
		return err
	}
	_, err = db.Exec("update `tb_article` set `Click`=`Click`+1 where `Id`=?", Id)
	if err != nil {
		return err
	}
	return nil
}
