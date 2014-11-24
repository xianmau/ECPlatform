package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"utils/global"
)

type GoodsCategory struct {
	Name     string `json:"name"`
	Parent   string `json:"parent"`
	Ordering int    `json:"ordering"`
}

func GetGoodsCategory(Name string) (*GoodsCategory, error) {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("select * from `tb_goods_category` where `Name`=?", Name)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		var goodsCategory GoodsCategory
		var Name string
		var Parent string
		var Ordering int
		if err := rows.Scan(&Name, &Parent, &Ordering); err != nil {
			return nil, err
		}
		goodsCategory.Name = Name
		goodsCategory.Parent = Parent
		goodsCategory.Ordering = Ordering
		return &goodsCategory, nil
	}
	return nil, nil
}

func GetGoodsCategoryList() ([]GoodsCategory, error) {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("select * from `tb_goods_category` where true order by `Ordering` desc")
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	categoryList := []GoodsCategory{}
	for rows.Next() {
		var goodsCategory GoodsCategory
		var Name string
		var Parent string
		var Ordering int
		if err := rows.Scan(&Name, &Parent, &Ordering); err != nil {
			return nil, err
		}
		goodsCategory.Name = Name
		goodsCategory.Parent = Parent
		goodsCategory.Ordering = Ordering
		categoryList = append(categoryList, goodsCategory)
	}
	return categoryList, nil
}

func GetSubGoodsCategoryList(root string) ([]GoodsCategory, error) {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("select * from `tb_goods_category` where `Parent`=? order by `Ordering` desc", root)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	categoryList := []GoodsCategory{}
	for rows.Next() {
		var goodsCategory GoodsCategory
		var Name string
		var Parent string
		var Ordering int
		if err := rows.Scan(&Name, &Parent, &Ordering); err != nil {
			return nil, err
		}
		goodsCategory.Name = Name
		goodsCategory.Parent = Parent
		goodsCategory.Ordering = Ordering
		categoryList = append(categoryList, goodsCategory)
	}
	return categoryList, nil
}

func CreateGoodsCategory(Name string, Parent string, Ordering string) error {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return err
	}
	_, err = db.Exec("insert into `tb_goods_category`(`Name`,`Parent`,`Ordering`) values(?,?,?)", Name, Parent, Ordering)
	if err != nil {
		return err
	}
	return nil
}

func EditGoodsCategory(Name string, Parent string, Ordering string) error {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return err
	}
	_, err = db.Exec("update `tb_goods_category` set `Parent`=?,`Ordering`=? where `Name`=?", Parent, Ordering, Name)
	if err != nil {
		return err
	}
	return nil
}

func DeleteGoodsCategory(Name string) error {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return err
	}
	_, err = db.Exec("delete from `tb_goods_category` where `Name`=?", Name)
	if err != nil {
		return err
	}
	return nil
}
