package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"utils/global"
)

type ArticleCategory struct {
	Name     string `json:"name"`
	Parent   string `json:"parent"`
	Ordering int    `json:"ordering"`
}

func GetArticleCategory(Name string) (*ArticleCategory, error) {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("select * from `tb_article_category` where `Name`=?", Name)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		var articleCategory ArticleCategory
		var Name string
		var Parent string
		var Ordering int
		if err := rows.Scan(&Name, &Parent, &Ordering); err != nil {
			return nil, err
		}
		articleCategory.Name = Name
		articleCategory.Parent = Parent
		articleCategory.Ordering = Ordering
		return &articleCategory, nil
	}
	return nil, nil
}

func GetArticleCategoryList() ([]ArticleCategory, error) {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("select * from `tb_article_category` where true order by `Ordering` desc")
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	articleList := []ArticleCategory{}
	for rows.Next() {
		var articleCategory ArticleCategory
		var Name string
		var Parent string
		var Ordering int
		if err := rows.Scan(&Name, &Parent, &Ordering); err != nil {
			return nil, err
		}
		articleCategory.Name = Name
		articleCategory.Parent = Parent
		articleCategory.Ordering = Ordering
		articleList = append(articleList, articleCategory)
	}
	return articleList, nil
}

func CreateArticleCategory(Name string, Parent string, Ordering string) error {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return err
	}
	_, err = db.Exec("insert into `tb_article_category`(`Name`,`Parent`,`Ordering`) values(?,?,?)", Name, Parent, Ordering)
	if err != nil {
		return err
	}
	return nil
}

func EditArticleCategory(Name string, Parent string, Ordering string) error {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return err
	}
	_, err = db.Exec("update `tb_article_category` set `Parent`=?,`Ordering`=? where `Name`=?", Parent, Ordering, Name)
	if err != nil {
		return err
	}
	return nil
}

func DeleteArticleCategory(Name string) error {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return err
	}
	_, err = db.Exec("delete from `tb_article_category` where `Name`=?", Name)
	if err != nil {
		return err
	}
	return nil
}
