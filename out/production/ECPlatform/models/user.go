package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"utils/global"
)

type User struct {
	Name         string `json:"name"` // 邮箱，手机
	Password     string `json:"password"`
	Level        int    `json:"level"`
	BaseInfo     string `json:"baseinfo"`     // 头像，真实姓名，性别，生日，地址，邮箱，手机
	ReceiveInfo  string `json:"receiveinfo"`  // [{收货信息1},{收货信息2},..,{收货信息n}]
	RegisterTime string `json:"registertime"` // 注册时间
	LoginHistory string `json:"loginhistory"` // [时间1,时间2,...,时间n]
	Status       int    // 状态
}

func GetUserForLogin(Name string, Password string) (*User, error) {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("select * from `tb_user` where `Name`=? and `Password`=?", Name, Password)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		var user User
		//var Name string
		//var Password string
		var Level int
		var BaseInfo string
		var ReceiveInfo string
		var RegisterTime string
		var LoginHistory string
		var Status int
		if err := rows.Scan(&Name, &Password, &Level, &BaseInfo, &ReceiveInfo, &RegisterTime, &LoginHistory, &Status); err != nil {
			return nil, err
		}
		user.Name = Name
		user.Password = Password
		user.Level = Level
		user.BaseInfo = BaseInfo
		user.ReceiveInfo = ReceiveInfo
		user.RegisterTime = RegisterTime
		user.LoginHistory = LoginHistory
		user.Status = Status
		return &user, nil
	}
	return nil, nil
}

func GetUser(Name string) (*User, error) {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("select * from `tb_user` where `Name`=?", Name)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		var user User
		//var Name string
		var Password string
		var Level int
		var BaseInfo string
		var ReceiveInfo string
		var RegisterTime string
		var LoginHistory string
		var Status int
		if err := rows.Scan(&Name, &Password, &Level, &BaseInfo, &ReceiveInfo, &RegisterTime, &LoginHistory, &Status); err != nil {
			return nil, err
		}
		user.Name = Name
		user.Password = Password
		user.Level = Level
		user.BaseInfo = BaseInfo
		user.ReceiveInfo = ReceiveInfo
		user.RegisterTime = RegisterTime
		user.LoginHistory = LoginHistory
		user.Status = Status

		return &user, nil
	}
	return nil, nil
}

func GetUserList() ([]User, error) {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("select * from `tb_user` where true order by `RegisterTime` desc")
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	userList := []User{}
	for rows.Next() {
		var user User
		var Name string
		var Password string
		var Level int
		var BaseInfo string
		var ReceiveInfo string
		var RegisterTime string
		var LoginHistory string
		var Status int
		if err := rows.Scan(&Name, &Password, &Level, &BaseInfo, &ReceiveInfo, &RegisterTime, &LoginHistory, &Status); err != nil {
			return nil, err
		}
		user.Name = Name
		user.Password = Password
		user.Level = Level
		user.BaseInfo = BaseInfo
		user.ReceiveInfo = ReceiveInfo
		user.RegisterTime = RegisterTime
		user.LoginHistory = LoginHistory
		user.Status = Status

		userList = append(userList, user)
	}
	return userList, nil
}

func RegisterUser(Name string, Password string) error {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return err
	}
	_, err = db.Exec("insert into `tb_user`(`Name`,`Password`,`BaseInfo`,`ReceiveInfo`,`LoginHistory`,`Status`) values(?,?,?)", Name, Password, "{}", "[]", "[]", 1)
	if err != nil {
		return err
	}
	return nil
}

func CreateUser(Name string, Password string, Level string, Status string) error {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return err
	}
	_, err = db.Exec("insert into `tb_user`(`Name`,`Password`,`Level`,`BaseInfo`,`ReceiveInfo`,`LoginHistory`,`Status`) values(?,?,?,?,?,?,?)", Name, Password, Level, "{}", "[]", "[]", Status)
	if err != nil {
		return err
	}
	return nil
}

func EditUserPassword(Name string, Password string) error {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return err
	}
	_, err = db.Exec("update `tb_user` set `Password`=? where `Name`=?", Password, Name)
	if err != nil {
		return err
	}
	return nil
}

func EditUserByAdmin(Name string, Level string, Status string) error {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return err
	}
	_, err = db.Exec("update `tb_user` set `Level`=?,`Status`=? where `Name`=?", Level, Status, Name)
	if err != nil {
		return err
	}
	return nil
}

func EditUserBaseInfo(Name string, BaseInfo string) error {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return err
	}
	_, err = db.Exec("update `tb_user` set `BaseInfo`=? where `Name`=?", BaseInfo, Name)
	if err != nil {
		return err
	}
	return nil
}

func EditUserReceiveInfo(Name string, ReceiveInfo string) error {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return err
	}
	_, err = db.Exec("update `tb_user` set `ReceiveInfo`=? where `Name`=?", ReceiveInfo, Name)
	if err != nil {
		return err
	}
	return nil
}

func EditUserLoginHistory(Name string, LoginHistory string) error {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return err
	}
	_, err = db.Exec("update `tb_user` set `LoginHistory`=? where `Name`=?", LoginHistory, Name)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(Name string) error {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return err
	}
	_, err = db.Exec("delete from `tb_user` where `Name`=?", Name)
	if err != nil {
		return err
	}
	return nil
}
