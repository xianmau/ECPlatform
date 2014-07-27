package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"utils/global"
)

type Shop struct {
	Name           string `json:"name"`
	UserName       string `json:"username"`       // 商店所属用户
	Kind           int    `json:"kind"`           // 商店类型：个体户，合作社
	Introduce      string `json:"introduce"`      // 商店介绍
	ApplyTime      string `json:"applytime"`      // 注册时间
	ApplyStatement string `json:"applystatement"` // 申请说明
	Status         int    // 状态
}

func GetShop(Name string) (*Shop, error) {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("select * from `tb_shop` where `Name`=?", Name)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		var shop Shop
		//var Name string
		var UserName string
		var Kind int
		var Introduce string
		var ApplyTime string
		var ApplyStatement string
		var Status int
		if err := rows.Scan(&Name, &UserName, &Kind, &Introduce, &ApplyTime, &ApplyStatement, &Status); err != nil {
			return nil, err
		}
		shop.Name = Name
		shop.UserName = UserName
		shop.Kind = Kind
		shop.Introduce = Introduce
		shop.ApplyTime = ApplyTime
		shop.ApplyStatement = ApplyStatement
		shop.Status = Status

		return &shop, nil
	}
	return nil, nil
}

func GetShopList() ([]Shop, error) {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("select * from `tb_shop` where true order by `RegisterTime` desc")
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	shopList := []Shop{}
	for rows.Next() {
		var shop Shop
		var Name string
		var UserName string
		var Kind int
		var Introduce string
		var ApplyTime string
		var ApplyStatement string
		var Status int
		if err := rows.Scan(&Name, &UserName, &Kind, &Introduce, &ApplyTime, &ApplyStatement, &Status); err != nil {
			return nil, err
		}
		shop.Name = Name
		shop.UserName = UserName
		shop.Kind = Kind
		shop.Introduce = Introduce
		shop.ApplyTime = ApplyTime
		shop.ApplyStatement = ApplyStatement
		shop.Status = Status

		shopList = append(shopList, shop)
	}
	return shopList, nil
}

func ApplyShop(Name string, UserName string, Kind string, Introduce string, ApplyStatement string) error {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return err
	}
	_, err = db.Exec("insert into `tb_shop`(`Name`,`UserName`,`Kind`,`Introduce`,`ApplyStatement`,`Status`) values(?,?,?,?,?,?)", Name, UserName, Kind, Introduce, ApplyStatement, 0)
	if err != nil {
		return err
	}
	return nil
}

func CreateShop(Name string, UserName string, Kind string, Introduce string, ApplyStatement string, Status string) error {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return err
	}
	_, err = db.Exec("insert into `tb_shop`(`Name`,`UserName`,`Kind`,`Introduce`,`ApplyStatement`,`Status`) values(?,?,?,?,?,?)", Name, UserName, Kind, Introduce, ApplyStatement, Status)
	if err != nil {
		return err
	}
	return nil
}

func EditShop(Name string, UserName string, Kind string, Introduce string, ApplyStatement string, Status string) error {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return err
	}
	_, err = db.Exec("update `tb_shop` set `UserName`=?,`Kind`=?,`Introduce`=?,`ApplyStatement`=?,`Status`=? where `Name`=?", UserName, Kind, Introduce, ApplyStatement, Status, Name)
	if err != nil {
		return err
	}
	return nil
}

func DeleteShop(Name string) error {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return err
	}
	_, err = db.Exec("delete from `tb_shop` where `Name`=?", Name)
	if err != nil {
		return err
	}
	return nil
}
