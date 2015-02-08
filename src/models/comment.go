package models

import (
	"database/sql"
	_ "utils/mysql"
	"utils/global"
)

type Comment struct {
	Id      int    `json:"id"`
	GoodsId int    `json:"goodsid"`
	Name    string `json:"name"`
	Content string `json:"content"`
	Time    string `json:"time"`
	Status  int    `json:"status"`
}

func GetComment(Id string) (*Comment, error) {
	db, err := sql.Open("mysql", global.Config["conn_str"])
	defer db.Close()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("select * from `tb_comment` where `Id`=?", Id)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		var comment Comment
		var Id int
		var GoodsId int
		var Name string
		var Content string
		var Time string
		var Status int
		if err := rows.Scan(&Id, &GoodsId, &Name, &Content, &Time, &Status); err != nil {
			return nil, err
		}
		comment.Id = Id
		comment.GoodsId = GoodsId
		comment.Name = Name
		comment.Content = Content
		comment.Time = Time
		comment.Status = Status
		return &comment, nil
	}
	return nil, nil
}

func GetCommentListByGoods(GoodsId string) ([]Comment, error) {
	db, err := sql.Open("mysql", global.Config["conn_str"])
	defer db.Close()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("select * from `tb_comment` where `GoodsId`=? order by `Time` desc", GoodsId)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	commentList := []Comment{}
	for rows.Next() {
		var comment Comment
		var Id int
		var GoodsId int
		var Name string
		var Content string
		var Time string
		var Status int
		if err := rows.Scan(&Id, &GoodsId, &Name, &Content, &Time, &Status); err != nil {
			return nil, err
		}
		comment.Id = Id
		comment.GoodsId = GoodsId
		comment.Name = Name
		comment.Content = Content
		comment.Time = Time
		comment.Status = Status
		commentList = append(commentList, comment)
	}
	return commentList, nil
}

func CreateComment(GoodsId string, Name string, Content string, Status string) error {
	db, err := sql.Open("mysql", global.Config["conn_str"])
	defer db.Close()
	if err != nil {
		return err
	}
	_, err = db.Exec("insert into `tb_comment`(`GoodsId`,`Name`,`Content`,`Status`) values(?,?,?,?)", GoodsId, Name, Content, Status)
	if err != nil {
		return err
	}
	return nil
}

func EditComment(Id string, Status string) error {
	db, err := sql.Open("mysql", global.Config["conn_str"])
	defer db.Close()
	if err != nil {
		return err
	}
	_, err = db.Exec("update `tb_comment` set `Status`=? where `Id`=?", Status, Id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteComment(Id string) error {
	db, err := sql.Open("mysql", global.Config["conn_str"])
	defer db.Close()
	if err != nil {
		return err
	}
	_, err = db.Exec("delete from `tb_comment` where `Id`=?", Id)
	if err != nil {
		return err
	}
	return nil
}
