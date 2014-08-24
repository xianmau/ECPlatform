package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"utils/global"
)

type Goods struct {
	Id           int     `json:"id"`
	Title        string  `json:"title"`
	Category     string  `json:"category"`
	Recommend    string  `json:"recommend"`
	Content      string  `json:"content"`
	Origin       string  `json:"origin"`
	Unit         string  `json:"unit"`
	Price        float64 `json:"price"`
	Shop         string  `json:"shop"`
	BuyLink      string  `json:"buylink"`
	Images       string  `json:"images"`
	Certificates string  `json:"certificates"`
	Status       int     `json:"status"`
}

func GetGoods(Id string) (*Goods, error) {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("select * from `tb_goods` where `Id`=?", Id)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		var goods Goods
		var Id int
		var Title string
		var Recommend string
		var Category string
		var Content string
		var Origin string
		var Unit string
		var Price float64
		var Shop string
		var BuyLink string
		var Images string
		var Certificates string
		var Status int
		if err := rows.Scan(&Id, &Title, &Category, &Recommend, &Content, &Origin, &Unit, &Price, &Shop, &BuyLink, &Images, &Certificates, &Status); err != nil {
			return nil, err
		}
		goods.Id = Id
		goods.Title = Title
		goods.Category = Category
		goods.Recommend = Recommend
		goods.Content = Content
		goods.Origin = Origin
		goods.Unit = Unit
		goods.Price = Price
		goods.Shop = Shop
		goods.BuyLink = BuyLink
		goods.Images = Images
		goods.Certificates = Certificates
		goods.Status = Status
		return &goods, nil
	}
	return nil, nil
}

func GetGoodsList() ([]Goods, error) {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("select * from `tb_goods` where true order by `Id` desc")
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	goodsList := []Goods{}
	for rows.Next() {
		var goods Goods
		var Id int
		var Title string
		var Category string
		var Recommend string
		var Content string
		var Origin string
		var Unit string
		var Price float64
		var Shop string
		var BuyLink string
		var Images string
		var Certificates string
		var Status int
		if err := rows.Scan(&Id, &Title, &Category, &Recommend, &Content, &Origin, &Unit, &Price, &Shop, &BuyLink, &Images, &Certificates, &Status); err != nil {
			return nil, err
		}
		goods.Id = Id
		goods.Title = Title
		goods.Category = Category
		goods.Recommend = Recommend
		goods.Content = Content
		goods.Origin = Origin
		goods.Unit = Unit
		goods.Price = Price
		goods.Shop = Shop
		goods.BuyLink = BuyLink
		goods.Images = Images
		goods.Certificates = Certificates
		goods.Status = Status
		goodsList = append(goodsList, goods)
	}
	return goodsList, nil
}

func GetGoodsListByCategory(Category string) ([]Goods, error) {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("select A.* from `tb_goods` as A left join `tb_goods_category` as B on A.`Category`=B.`Name` where A.`Category`=? or B.`Parent`=? order by `Id` desc", Category, Category)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	goodsList := []Goods{}
	for rows.Next() {
		var goods Goods
		var Id int
		var Title string
		var Category string
		var Recommend string
		var Content string
		var Origin string
		var Unit string
		var Price float64
		var Shop string
		var BuyLink string
		var Images string
		var Certificates string
		var Status int
		if err := rows.Scan(&Id, &Title, &Category, &Recommend, &Content, &Origin, &Unit, &Price, &Shop, &BuyLink, &Images, &Certificates, &Status); err != nil {
			return nil, err
		}
		goods.Id = Id
		goods.Title = Title
		goods.Category = Category
		goods.Recommend = Recommend
		goods.Content = Content
		goods.Origin = Origin
		goods.Unit = Unit
		goods.Price = Price
		goods.Shop = Shop
		goods.BuyLink = BuyLink
		goods.Images = Images
		goods.Certificates = Certificates
		goods.Status = Status
		goodsList = append(goodsList, goods)
	}
	return goodsList, nil
}

func CreateGoods(Title string, Category string, Recommend string, Content string, Origin string, Unit string, Price string, Shop string, BuyLink string, Images string, Certificates string, Status string) error {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return err
	}
	_, err = db.Exec("insert into `tb_goods`(`Title`,`Category`,`Recommend`,`Content`,`Origin`,`Unit`,`Price`,`Shop`,`BuyLink`,`Images`,`Certificates`,`Status`) values(?,?,?,?,?,?,?,?,?,?,?,?)", Title, Category, Recommend, Content, Origin, Unit, Price, Shop, BuyLink, Images, Certificates, Status)
	if err != nil {
		return err
	}
	return nil
}

func EditGoods(Id string, Title string, Category string, Recommend string, Content string, Origin string, Unit string, Price string, Shop string, BuyLink string, Images string, Certificates string, Status string) error {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return err
	}
	_, err = db.Exec("update `tb_goods` set `Title`=?,`Category`=?,`Recommend`=?,`Content`=?,`Origin`=?,`Unit`=?,`Price`=?,`Shop`=?,`BuyLink`=?,`Images`=?,`Certificates`=?,`Status`=? where `Id`=?", Title, Category, Recommend, Content, Origin, Unit, Price, Shop, BuyLink, Images, Certificates, Status, Id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteGoods(Id string) error {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return err
	}
	_, err = db.Exec("delete from `tb_goods` where `Id`=?", Id)
	if err != nil {
		return err
	}
	return nil
}
