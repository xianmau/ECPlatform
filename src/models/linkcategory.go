package models

import (
	"database/sql"
	_ "utils/mysql"
	"utils/global"
)

type LinkCategory struct {
	Name     string `json:"name"`
	Parent   string `json:"parent"`
	Ordering int    `json:"ordering"`
}

func GetLinkCategory(Name string) (*LinkCategory, error) {
	db, err := sql.Open("mysql", global.Config["conn_str"])
	defer db.Close()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("select * from `tb_link_category` where `Name`=?", Name)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		var linkCategory LinkCategory
		var Name string
		var Parent string
		var Ordering int
		if err := rows.Scan(&Name, &Parent, &Ordering); err != nil {
			return nil, err
		}
		linkCategory.Name = Name
		linkCategory.Parent = Parent
		linkCategory.Ordering = Ordering
		return &linkCategory, nil
	}
	return nil, nil
}

func GetLinkCategoryList() ([]LinkCategory, error) {
	db, err := sql.Open("mysql", global.Config["conn_str"])
	defer db.Close()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("select * from `tb_link_category` where true order by `Ordering` desc")
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	linkList := []LinkCategory{}
	for rows.Next() {
		var linkCategory LinkCategory
		var Name string
		var Parent string
		var Ordering int
		if err := rows.Scan(&Name, &Parent, &Ordering); err != nil {
			return nil, err
		}
		linkCategory.Name = Name
		linkCategory.Parent = Parent
		linkCategory.Ordering = Ordering
		linkList = append(linkList, linkCategory)
	}
	return linkList, nil
}

func CreateLinkCategory(Name string, Parent string, Ordering string) error {
	db, err := sql.Open("mysql", global.Config["conn_str"])
	defer db.Close()
	if err != nil {
		return err
	}
	_, err = db.Exec("insert into `tb_link_category`(`Name`,`Parent`,`Ordering`) values(?,?,?)", Name, Parent, Ordering)
	if err != nil {
		return err
	}
	return nil
}

func EditLinkCategory(Name string, Parent string, Ordering string) error {
	db, err := sql.Open("mysql", global.Config["conn_str"])
	defer db.Close()
	if err != nil {
		return err
	}
	_, err = db.Exec("update `tb_link_category` set `Parent`=?,`Ordering`=? where `Name`=?", Parent, Ordering, Name)
	if err != nil {
		return err
	}
	return nil
}

func DeleteLinkCategory(Name string) error {
	db, err := sql.Open("mysql", global.Config["conn_str"])
	defer db.Close()
	if err != nil {
		return err
	}
	_, err = db.Exec("delete from `tb_link_category` where `Name`=?", Name)
	if err != nil {
		return err
	}
	return nil
}
