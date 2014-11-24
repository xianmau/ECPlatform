package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"utils/global"
)

type GoodsCategoryInfo struct {
	CategoryName  string `json:"categoryname"`
	CategoryInfo  string `json:"categoryinfo"`
	CategoryImage string `json:"categoryimage"`
}

func GetGoodsCategoryInfo(CategoryName string) (*GoodsCategoryInfo, error) {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("select * from `tb_goods_category_info` where `CategoryName`=?", CategoryName)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		var goodsCategoryInfo GoodsCategoryInfo
		var CategoryName string
		var CategoryInfo string
		var CategoryImage string
		if err := rows.Scan(&CategoryName, &CategoryInfo, &CategoryImage); err != nil {
			return nil, err
		}
		goodsCategoryInfo.CategoryName = CategoryName
		goodsCategoryInfo.CategoryInfo = CategoryInfo
		goodsCategoryInfo.CategoryImage = CategoryImage
		return &goodsCategoryInfo, nil
	}
	return nil, nil
}

func GetGoodsCategoryInfoList() ([]GoodsCategoryInfo, error) {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("select * from `tb_goods_category_info` where true")
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	categoryInfoList := []GoodsCategoryInfo{}
	for rows.Next() {
		var goodsCategoryInfo GoodsCategoryInfo
		var CategoryName string
		var CategoryInfo string
		var CategoryImage string
		if err := rows.Scan(&CategoryName, &CategoryInfo, &CategoryImage); err != nil {
			return nil, err
		}
		goodsCategoryInfo.CategoryName = CategoryName
		goodsCategoryInfo.CategoryInfo = CategoryInfo
		goodsCategoryInfo.CategoryImage = CategoryImage
		categoryInfoList = append(categoryInfoList, goodsCategoryInfo)
	}
	return categoryInfoList, nil
}

func CreateGoodsCategoryInfo(CategoryName string, CategoryInfo string, CategoryImage string) error {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return err
	}
	_, err = db.Exec("insert into `tb_goods_category_info`(`CategoryName`,`CategoryInfo`,`CategoryImage`) values(?,?,?)", CategoryName, CategoryInfo, CategoryImage)
	if err != nil {
		return err
	}
	return nil
}

func EditGoodsCategoryInfo(CategoryName string, CategoryInfo string, CategoryImage string) error {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return err
	}
	_, err = db.Exec("update `tb_goods_category_info` set `CategoryInfo`=?,`CategoryImage`=? where `CategoryName`=?", CategoryInfo, CategoryImage, CategoryName)
	if err != nil {
		return err
	}
	return nil
}

func DeleteGoodsCategoryInfo(CategoryName string) error {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return err
	}
	_, err = db.Exec("delete from `tb_goods_category_info` where `CategoryName`=?", CategoryName)
	if err != nil {
		return err
	}
	return nil
}
