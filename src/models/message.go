package models

import (
	"database/sql"
	_ "utils/mysql"
	"utils/global"
)

type Message struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Content string `json:"content"`
	Time    string `json:"time"`
	Status  int    `json:"status"`
}

func GetMessage(Id string) (*Message, error) {
	db, err := sql.Open("mysql", global.Config["conn_str"])
	defer db.Close()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("select * from `tb_message` where `Id`=?", Id)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		var message Message
		var Id int
		var Name string
		var Email string
		var Content string
		var Time string
		var Status int
		if err := rows.Scan(&Id, &Name, &Email, &Content, &Time, &Status); err != nil {
			return nil, err
		}
		message.Id = Id
		message.Name = Name
		message.Email = Email
		message.Content = Content
		message.Time = Time
		message.Status = Status
		return &message, nil
	}
	return nil, nil
}

func GetMessageListByGoods() ([]Message, error) {
	db, err := sql.Open("mysql", global.Config["conn_str"])
	defer db.Close()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("select * from `tb_message` where true order by `Time` desc")
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	messageList := []Message{}
	for rows.Next() {
		var message Message
		var Id int
		var Name string
		var Email string
		var Content string
		var Time string
		var Status int
		if err := rows.Scan(&Id, &Name, &Email, &Content, &Time, &Status); err != nil {
			return nil, err
		}
		message.Id = Id
		message.Name = Name
		message.Email = Email
		message.Content = Content
		message.Time = Time
		message.Status = Status
		messageList = append(messageList, message)
	}
	return messageList, nil
}

func CreateMessage(Name string, Email string, Content string) error {
	db, err := sql.Open("mysql", global.Config["conn_str"])
	defer db.Close()
	if err != nil {
		return err
	}
	_, err = db.Exec("insert into `tb_message`(`Name`,`Email`,`Content`) values(?,?,?)", Name, Email, Content)
	if err != nil {
		return err
	}
	return nil
}

func EditMessage(Id string, Status string) error {
	db, err := sql.Open("mysql", global.Config["conn_str"])
	defer db.Close()
	if err != nil {
		return err
	}
	_, err = db.Exec("update `tb_message` set `Status`=? where `Id`=?", Status, Id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteMessage(Id string) error {
	db, err := sql.Open("mysql", global.Config["conn_str"])
	defer db.Close()
	if err != nil {
		return err
	}
	_, err = db.Exec("delete from `tb_message` where `Id`=?", Id)
	if err != nil {
		return err
	}
	return nil
}
