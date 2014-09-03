package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"utils/global"
)

type GoodsStatistic struct {
	GoodsId      int `json:"goodsid"`
	ViewTimes    int `json:"viewtimes"`
	BuyTimes     int `json:"buytimes"`
	LinkTimes    int `json:"linktimes"`
	CollectTimes int `json:"collecttimes"`
}

func GetGoodsStatistic(GoodsId string) (*GoodsStatistic, error) {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("select * from `tb_goods_statistic` where `GoodsId`=?", GoodsId)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		var goodsStatistic GoodsStatistic
		var GoodsId int
		var ViewTimes int
		var BuyTimes int
		var LinkTimes int
		var CollectTimes int
		if err := rows.Scan(&GoodsId, &ViewTimes, &BuyTimes, &LinkTimes, &CollectTimes); err != nil {
			return nil, err
		}
		goodsStatistic.GoodsId = GoodsId
		goodsStatistic.ViewTimes = ViewTimes
		goodsStatistic.BuyTimes = BuyTimes
		goodsStatistic.LinkTimes = LinkTimes
		goodsStatistic.CollectTimes = CollectTimes
		return &goodsStatistic, nil
	}
	return nil, nil
}

func GetGoodsStatisticList() ([]GoodsStatistic, error) {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("select * from `tb_goods_statistic` where true order by `GoodsId` asc")
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	goodsStatisticList := []GoodsStatistic{}
	for rows.Next() {
		var goodsStatistic GoodsStatistic
		var GoodsId int
		var ViewTimes int
		var BuyTimes int
		var LinkTimes int
		var CollectTimes int
		if err := rows.Scan(&GoodsId, &ViewTimes, &BuyTimes, &LinkTimes, &CollectTimes); err != nil {
			return nil, err
		}
		goodsStatistic.GoodsId = GoodsId
		goodsStatistic.ViewTimes = ViewTimes
		goodsStatistic.BuyTimes = BuyTimes
		goodsStatistic.LinkTimes = LinkTimes
		goodsStatistic.CollectTimes = CollectTimes
		goodsStatisticList = append(goodsStatisticList, goodsStatistic)
	}
	return goodsStatisticList, nil
}

func CreateGoodsStatistic(GoodsId string) error {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return err
	}
	_, err = db.Exec("insert into `tb_goods_statistic`(`GoodsId`) values(?)", GoodsId)
	if err != nil {
		return err
	}
	return nil
}

func IncreaseGoodsStatisticViewTimes(GoodsId string) error {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return err
	}
	_, err = db.Exec("update `tb_goods_statistic` set `ViewTimes`=`ViewTimes`+1 where `GoodsId`=?", GoodsId)
	if err != nil {
		return err
	}
	return nil
}

func EditGoodsStatistic(GoodsId string, ViewTimes string, BuyTimes string, LinkTimes string, CollectTimes string) error {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return err
	}
	_, err = db.Exec("update `tb_goods_statistic` set `ViewTimes`=?,`BuyTimes`=?,`LinkTimes`=?,`CollectTimes`=? where `GoodsId`=?", ViewTimes, BuyTimes, LinkTimes, CollectTimes, GoodsId)
	if err != nil {
		return err
	}
	return nil
}

func DeleteGoodsStatistic(GoodsId string) error {
	db, err := sql.Open("mysql", global.Config.Get("conn_str"))
	defer db.Close()
	if err != nil {
		return err
	}
	_, err = db.Exec("delete from `tb_goods_statistic` where `GoodsId`=?", GoodsId)
	if err != nil {
		return err
	}
	return nil
}
