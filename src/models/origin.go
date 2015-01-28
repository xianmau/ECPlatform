package models

import (
	"database/sql"
	_ "utils/mysql"
	"utils/global"
)

type Origin struct {
	Name string `json:"name"`
}

//func GetOrigin(Name string) (*Origin, error) {
//	db, err := sql.Open("mysql", global.Config["conn_str"])
//	defer db.Close()
//	if err != nil {
//		return nil, err
//	}
//	rows, err := db.Query("select * from `tb_origin` where `Name`=?", Name)
//	defer rows.Close()
//	if err != nil {
//		return nil, err
//	}
//	if rows.Next() {
//		var origin Origin
//		var Name string
//		if err := rows.Scan(&Name); err != nil {
//			return nil, err
//		}
//		origin.Name = Name
//		return &origin, nil
//	}
//	return nil, nil
//}

func GetOriginList() ([]Origin, error) {
	db, err := sql.Open("mysql", global.Config["conn_str"])
	defer db.Close()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("select * from `tb_origin` where true")
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	originList := []Origin{}
	for rows.Next() {
		var origin Origin
		var Name string
		if err := rows.Scan(&Name); err != nil {
			return nil, err
		}
		origin.Name = Name
		originList = append(originList, origin)
	}
	return originList, nil
}

func CreateOrigin(Name string) error {
	db, err := sql.Open("mysql", global.Config["conn_str"])
	defer db.Close()
	if err != nil {
		return err
	}
	_, err = db.Exec("insert into `tb_origin`(`Name`) values(?)", Name)
	if err != nil {
		return err
	}
	return nil
}

//func EditOrigin(Name string, NewName string) error {
//	db, err := sql.Open("mysql", global.Config["conn_str"])
//	defer db.Close()
//	if err != nil {
//		return err
//	}
//	_, err = db.Exec("update `tb_origin` set `Name`=? where `Name`=?", NewName, Name)
//	if err != nil {
//		return err
//	}
//	return nil
//}

func DeleteOrigin(Name string) error {
	db, err := sql.Open("mysql", global.Config["conn_str"])
	defer db.Close()
	if err != nil {
		return err
	}
	_, err = db.Exec("delete from `tb_origin` where `Name`=?", Name)
	if err != nil {
		return err
	}
	return nil
}
