package main

import (
	"Octopus/CTCore/structures"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// DONE 创建数据连接

type gaugeRes struct {
	Id      int         // 量表id
	Options []gaugeOpts // 量表的题目以及选项
}

type gaugeOpts struct {
	questionId int    // 题目id
	answerStr  string // 选项列表
	answers    []string
}

type userOpts struct {
	Id      int
	Account string
}

func conn() (*sqlx.DB, error) {
	cnf := *cnf
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		cnf.User, cnf.Password, cnf.Host, cnf.Port, cnf.DbName)
	db, err := sqlx.Connect("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	// 测试创建连接是否成功
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func insert(sql string) error {
	return nil
}

func selectX(dbStruct structures.ResInterface, sql string, args ...any) (*structures.ResInterface, error) {
	res := &dbStruct
	err := db.Select(res, sql, args...)
	if err != nil {
		return nil, err
	}
	return res, nil
}
