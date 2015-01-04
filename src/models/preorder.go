package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"utils/global"
)

type Preorder struct {
	Id      int    `json:"id"`
	GoodsId int    `json:"goodsid"`
	UserId  int    `json:"userid"`
	Name    string `json:"name"`
	Tel     string `json:"tel"`
	Addr    string `json:"addr"`
	Num     int    `json:"num"`
	Time    string `json:"time"`
	Status  int    `json:"status"`
}

func GetPreorder(Id string) (*Preorder, error) {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("select * from `tb_preorder` where `Id`=?", Id)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		var preorder Preorder
		var Id int
		var GoodsId int
		var UserId int
		var Name string
		var Tel string
		var Addr string
		var Num int
		var Time string
		var Status int
		if err := rows.Scan(&Id, &GoodsId, &UserId, &Name, &Tel, &Addr, &Num, &Time, &Status); err != nil {
			return nil, err
		}
		preorder.Id = Id
		preorder.GoodsId = GoodsId
		preorder.UserId = UserId
		preorder.Name = Name
		preorder.Tel = Tel
		preorder.Addr = Addr
		preorder.Num = Num
		preorder.Time = Time
		preorder.Status = Status
		return &preorder, nil
	}
	return nil, nil
}

func GetPreorderListByGoods(GoodsId string) ([]Preorder, error) {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("select * from `tb_preorder` where `GoodsId`=?", GoodsId)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	preorderList := []Preorder{}
	for rows.Next() {
		var preorder Preorder
		var Id int
		var GoodsId int
		var UserId int
		var Name string
		var Tel string
		var Addr string
		var Num int
		var Time string
		var Status int
		if err := rows.Scan(&Id, &GoodsId, &UserId, &Name, &Tel, &Addr, &Num, &Time, &Status); err != nil {
			return nil, err
		}
		preorder.Id = Id
		preorder.GoodsId = GoodsId
		preorder.UserId = UserId
		preorder.Name = Name
		preorder.Tel = Tel
		preorder.Addr = Addr
		preorder.Num = Num
		preorder.Time = Time
		preorder.Status = Status
		preorderList = append(preorderList, preorder)
	}
	return preorderList, nil
}

func CreatePreorder(GoodsId string, UserId string, Name string, Tel string, Addr string, Num string, Status string) error {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return err
	}
	_, err = db.Exec("insert into `tb_preorder`(`GoodsId`,`UserId`,`Name`,`Tel`,`Addr`,`Num`,`Status`) values(?,?,?,?,?,?,?)", GoodsId, UserId, Name, Tel, Addr, Num, Status)
	if err != nil {
		return err
	}
	return nil
}

func EditPreorder(Id string, Status string) error {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return err
	}
	_, err = db.Exec("update `tb_preorder` set `Status`=? where `Id`=?", Status, Id)
	if err != nil {
		return err
	}
	return nil
}

func DeletPreorder(Id string) error {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return err
	}
	_, err = db.Exec("delete from `tb_preorder` where `Id`=?", Id)
	if err != nil {
		return err
	}
	return nil
}
