package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"utils/global"
	"encoding/json"
)

type Role struct {
	Name      string   `json:"name"`
	Authority []string `json:"authority"`
}

func GetRole(Name string) (*Role, error) {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("select * from `tb_role` where `Name`=?", Name)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		var role Role
		var Name string
		var Authority string
		if err := rows.Scan(&Name, &Authority); err != nil {
			return nil, err
		}
		role.Name = Name
		err := json.Unmarshal([]byte(Authority), &(role.Authority))
		if err != nil{
			//
		}
		return &role, nil
	}
	return nil, nil
}

func GetRoleList() ([]Role, error) {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("select * from `tb_role` where true")
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	roleList := []Role{}
	for rows.Next() {
		var role Role
		var Name string
		var Authority string
		if err := rows.Scan(&Name, &Authority); err != nil {
			return nil, err
		}
		role.Name = Name
		err := json.Unmarshal([]byte(Authority), &(role.Authority))
		if err != nil{
			//
		}
		roleList = append(roleList, role)
	}
	return roleList, nil
}

func CreateRole(Name string, Authority string) error {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return err
	}
	_, err = db.Exec("insert into `tb_role`(`Name`,`Authority`) values(?,?)", Name, Authority)
	if err != nil {
		return err
	}
	return nil
}

func EditRole(Name string, Authority string) error {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return err
	}
	_, err = db.Exec("update `tb_role` set `Authority`=? where `Name`=?", Authority, Name)
	if err != nil {
		return err
	}
	return nil
}

func DeleteRole(Name string) error {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return err
	}
	_, err = db.Exec("delete from `tb_role` where `Name`=?", Name)
	if err != nil {
		return err
	}
	return nil
}
