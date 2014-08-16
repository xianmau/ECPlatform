package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"utils/global"
)

type Admin struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Role     Role   `json:"role"`
}

func GetAdminForLogin(Name string, Password string) (*Admin, error) {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("select `tb_admin`.*,`tb_role`.`Authority` from `tb_admin` left join `tb_role` on `tb_admin`.`Role`=`tb_role`.`Name` where `tb_admin`.`Name`=? and `tb_admin`.`Password`=?", Name, Password)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		var admin Admin
		//var Name string
		//var Password string
		var Department string
		var RoleName string
		var RoleAuthority sql.NullString
		if err := rows.Scan(&Name, &Password, &Department, &RoleName, &RoleAuthority); err != nil {
			return nil, err
		}
		admin.Name = Name
		admin.Password = Password
		admin.Role.Name = RoleName
		if RoleAuthority.Valid {
			if err := json.Unmarshal([]byte(RoleAuthority.String), &admin.Role.Authority); err != nil {
				return nil, err
			}
		}
		return &admin, nil
	}
	return nil, nil
}

func GetAdmin(Name string) (*Admin, error) {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("select `tb_admin`.*,`tb_role`.`Authority` from `tb_admin` left join `tb_role` on `tb_admin`.`Role`=`tb_role`.`Name` where `tb_admin`.`Name`=?", Name)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		var admin Admin
		//var Name string
		var Password string
		var RoleName string
		var RoleAuthority sql.NullString
		if err := rows.Scan(&Name, &Password, &RoleName, &RoleAuthority); err != nil {
			return nil, err
		}
		admin.Name = Name
		admin.Password = Password
		admin.Role.Name = RoleName
		if RoleAuthority.Valid {
			if err := json.Unmarshal([]byte(RoleAuthority.String), &admin.Role.Authority); err != nil {
				return nil, err
			}
		}
		return &admin, nil
	}
	return nil, nil
}

func GetAdminList() ([]Admin, error) {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("select `tb_admin`.*,`tb_role`.`Authority` from `tb_admin` left join `tb_role` on `tb_admin`.`Role`=`tb_role`.`Name` where true")
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	adminList := []Admin{}
	for rows.Next() {
		var admin Admin
		var Name string
		var Password string
		var RoleName string
		var RoleAuthority sql.NullString
		if err := rows.Scan(&Name, &Password, &RoleName, &RoleAuthority); err != nil {
			return nil, err
		}
		admin.Name = Name
		admin.Password = Password
		admin.Role.Name = RoleName
		if RoleAuthority.Valid {
			if err := json.Unmarshal([]byte(RoleAuthority.String), &admin.Role.Authority); err != nil {
				return nil, err
			}
		}
		adminList = append(adminList, admin)
	}
	return adminList, nil
}

func CreateAdmin(Name string, Password string, RoleName string) error {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return err
	}
	_, err = db.Exec("insert into `tb_admin`(`Name`,`Password`,`Role`) values(?,?,?)", Name, Password, RoleName)
	if err != nil {
		return err
	}
	return nil
}

func EditAdmin(Name string, Password string, RoleName string) error {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return err
	}
	if Password == "" {
		_, err = db.Exec("update `tb_admin` set `Role`=? where `Name`=?", Password, RoleName, Name)
		if err != nil {
			return err
		}
		return nil
	}
	_, err = db.Exec("update `tb_admin` set `Password`=?,`Role`=? where `Name`=?", Password, RoleName, Name)
	if err != nil {
		return err
	}
	return nil
}

func DeleteAdmin(Name string) error {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return err
	}
	_, err = db.Exec("delete from `tb_admin` where `Name`=?", Name)
	if err != nil {
		return err
	}
	return nil
}
