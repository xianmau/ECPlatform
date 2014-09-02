package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"utils/global"
)

type WebHomeInfo struct {
	Id  int `json:"id"`
	Floor1  string `json:"floor1"`
	Floor2  string `json:"floor2"`
	Floor3  string `json:"floor3"`
	Floor4  string `json:"floor4"`
	Floor5  string `json:"floor5"`
	Floor6  string `json:"floor6"`
	Floor7  string `json:"floor7"`
	Floor8  string `json:"floor8"`
}

func GetWebHomeInfo(Id string) (*WebHomeInfo, error) {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("select * from `tb_web_home_info` where `Id`=?", Id)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		var webHomeInfo WebHomeInfo
		var Id int
		var Floor1 string
		var Floor2 string
		var Floor3 string
		var Floor4 string
		var Floor5 string
		var Floor6 string
		var Floor7 string
		var Floor8 string
		if err := rows.Scan(&Id, &Floor1, &Floor2, &Floor3, &Floor4, &Floor5, &Floor6, &Floor7, &Floor8); err != nil {
			return nil, err
		}
		webHomeInfo.Id = Id
		webHomeInfo.Floor1 = Floor1
		webHomeInfo.Floor2 = Floor2
		webHomeInfo.Floor3 = Floor3
		webHomeInfo.Floor4 = Floor4
		webHomeInfo.Floor5 = Floor5
		webHomeInfo.Floor6 = Floor6
		webHomeInfo.Floor7 = Floor7
		webHomeInfo.Floor8 = Floor8
		return &webHomeInfo, nil
	}
	return nil, nil
}

func GetWebHomeInfoList() ([]WebHomeInfo, error) {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("select * from `tb_web_home_info` where true")
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	webHomeInfoList := []WebHomeInfo{}
	for rows.Next() {
		var webHomeInfo WebHomeInfo
		var Id int
		var Floor1 string
		var Floor2 string
		var Floor3 string
		var Floor4 string
		var Floor5 string
		var Floor6 string
		var Floor7 string
		var Floor8 string
		if err := rows.Scan(&Id, &Floor1, &Floor2, &Floor3, &Floor4, &Floor5, &Floor6, &Floor7, &Floor8); err != nil {
			return nil, err
		}
		webHomeInfo.Id = Id
		webHomeInfo.Floor1 = Floor1
		webHomeInfo.Floor2 = Floor2
		webHomeInfo.Floor3 = Floor3
		webHomeInfo.Floor4 = Floor4
		webHomeInfo.Floor5 = Floor5
		webHomeInfo.Floor6 = Floor6
		webHomeInfo.Floor7 = Floor7
		webHomeInfo.Floor8 = Floor8
		webHomeInfoList = append(webHomeInfoList, webHomeInfo)
	}
	return webHomeInfoList, nil
}

func CreateWebHomeInfo(Floor1, Floor2, Floor3, Floor4, Floor5, Floor6, Floor7, Floor8 string) error {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return err
	}
	_, err = db.Exec("insert into `tb_web_home_info`(`Floor1`,`Floor2`,`Floor3`,`Floor4`,`Floor5`,`Floor6`,`Floor7`,`Floor8`) values(?,?,?,?,?,?,?,?)", Floor1, Floor2, Floor3, Floor4, Floor5, Floor6, Floor7, Floor8)
	if err != nil {
		return err
	}
	return nil
}

func EditWebHomeInfo(Id string, Floor1, Floor2, Floor3, Floor4, Floor5, Floor6, Floor7, Floor8 string) error {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return err
	}
	_, err = db.Exec("update `tb_web_home_info` set `Floor1`=?,`Floor2`=?,`Floor3`=?,`Floor4`=?,`Floor5`=?,`Floor6`=?,`Floor7`=?,`Floor8`=? where `Id`=?", Floor1, Floor2, Floor3, Floor4, Floor5, Floor6, Floor7, Floor8, Id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteWebHomeInfo(Id string) error {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return err
	}
	_, err = db.Exec("delete from `tb_web_home_info` where `Id`=?", Id)
	if err != nil {
		return err
	}
	return nil
}
