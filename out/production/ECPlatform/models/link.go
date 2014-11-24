package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"utils/global"
)

type Link struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Category string `json:"category"`
	LinkUrl  string `json:"linkurl"`
	ImageUrl  string `json:"imageurl"`
	Status   int    `json:"status"`
}

func GetLink(Id string) (*Link, error) {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("select * from `tb_link` where `Id`=?", Id)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		var link Link
		var Id int
		var Title string
		var Category string
		var LinkUrl string
		var ImageUrl string
		var Status int
		if err := rows.Scan(&Id, &Title, &Category, &LinkUrl, &ImageUrl, &Status); err != nil {
			return nil, err
		}
		link.Id = Id
		link.Title = Title
		link.Category = Category
		link.LinkUrl = LinkUrl
		link.ImageUrl = ImageUrl
		link.Status = Status
		return &link, nil
	}
	return nil, nil
}

func GetLinkList() ([]Link, error) {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("select * from `tb_link` where true order by `Id` desc")
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	linkList := []Link{}
	for rows.Next() {
		var link Link
		var Id int
		var Title string
		var Category string
		var LinkUrl string
		var ImageUrl string
		var Status int
		if err := rows.Scan(&Id, &Title, &Category, &LinkUrl, &ImageUrl, &Status); err != nil {
			return nil, err
		}
		link.Id = Id
		link.Title = Title
		link.Category = Category
		link.LinkUrl = LinkUrl
		link.ImageUrl = ImageUrl
		link.Status = Status
		linkList = append(linkList, link)
	}
	return linkList, nil
}

func CreateLink(Title string, Category string, LinkUrl string, ImageUrl string, Status string) error {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return err
	}
	_, err = db.Exec("insert into `tb_link`(`Title`,`Category`,`LinkUrl`,`ImageUrl`,`Status`) values(?,?,?,?,?)", Title, Category, LinkUrl, ImageUrl, Status)
	if err != nil {
		return err
	}
	return nil
}

func EditLink(Id string, Title string, Category string, LinkUrl string, ImageUrl string, Status string) error {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return err
	}
	_, err = db.Exec("update `tb_link` set `Title`=?,`Category`=?,`LinkUrl`=?,`ImageUrl`=?,`Status`=? where `Id`=?", Title, Category, LinkUrl, ImageUrl, Status, Id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteLink(Id string) error {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return err
	}
	_, err = db.Exec("delete from `tb_link` where `Id`=?", Id)
	if err != nil {
		return err
	}
	return nil
}
