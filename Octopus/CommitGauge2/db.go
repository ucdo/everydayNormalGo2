package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
	"strings"
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

func conn() (*sql.DB, error) {
	cnf := (*cnf).Mysql
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		cnf.User, cnf.Password, cnf.Host, cnf.Port, cnf.DbName)
	db, err := sql.Open("mysql", dataSourceName)
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

func getEvaluate() *[]evaluates {
	conn, err := conn()
	if err != nil {
		panic(err)
	}

	sqlx := fmt.Sprintf("SELECT id,measure_id,user_id FROM qy_user_evaluate_log WHERE `status` = 0")

	userTpl := " AND user_id IN ( %s ) "
	measureTpl := " AND measure_id IN ( %s ) "

	// 检查 cnf.User数组的长度，如果非空，添加到SQL语句中
	if len(cnf.User) > 0 {
		placeholders := strings.Repeat("?,", len(cnf.User))
		sqlx += fmt.Sprintf(userTpl, placeholders[:len(placeholders)-1])
	}

	// 检查 cnf.Gauge数组的长度，如果非空，添加到SQL语句中
	if len(cnf.Gauge) > 0 {
		placeholders := strings.Repeat("?,", len(cnf.Gauge))
		sqlx += fmt.Sprintf(measureTpl, placeholders[:len(placeholders)-1])
	}

	sqlx += "  limit 1"

	// 根据条件准备参数
	var args []interface{}
	if len(cnf.User) > 0 {
		for _, user := range cnf.User {
			args = append(args, strconv.Itoa(user))
		}
		//args = append(args,  cnf.User)
	}
	if len(cnf.Gauge) > 0 {
		for _, gauge := range cnf.Gauge {
			args = append(args, strconv.Itoa(gauge))
		}
		//args = append(args,  cnf.Gauge)
	}

	// 执行SQL查询
	var res *sql.Rows
	if len(args) > 0 {
		res, err = conn.Query(sqlx, args...)
	} else {
		res, err = conn.Query(sqlx)
	}
	if err != nil {
		// 处理错误，而不是使用panic
		log.Println(sqlx)
		log.Printf("An error occurred: %v\n", err)
		return nil
	}
	defer res.Close()

	var tmp []evaluates
	for res.Next() {
		x := evaluates{}

		err := res.Scan(&x.Id, &x.MeasureId, &x.UserId)
		if err != nil {
			log.Println("parse result with error: ", err)
			return nil
		}
		tmp = append(tmp, x)
	}

	return &tmp
}

func getMeasures(e *[]evaluates) *measureMap {
	m := measureMap{}
	if len(*e) == 0 {
		return nil
	}

	var args []interface{}
	for _, v := range *e {
		args = append(args, v.MeasureId)
	}

	sqlx := "SELECT q.measure_id, q.id, q.score FROM qy_measure m LEFT JOIN qy_measure_question q ON q.measure_id = m.id WHERE m.id in (%s)"

	x := strings.Repeat("?,", len(args))
	sqlx = fmt.Sprintf(sqlx, x[:len(x)-1])
	log.Println(sqlx, args)

	db, err := conn()
	if err != nil {
		log.Println("conn error on find measures: ", err)
		return nil
	}

	res, err := db.Query(sqlx, args...)
	if err != nil {
		log.Println("conn error on find measures: ", err)
		return nil
	}
	defer res.Close()

	tmp := struct {
		M int `json:"measure_id"`
		q questions
	}{}

	for res.Next() {
		err := res.Scan(&tmp.M, &tmp.q.Id, &tmp.q.Score)
		if err != nil {
			log.Println("measures: parse result with error: ", err)
			return nil
		}

		if _, exists := m[tmp.M]; !exists {
			// 初始化 m[tmp.M] 为一个新的 questions 切片
			m[tmp.M] = make([]questions, 0)
		}

		m[tmp.M] = append(m[tmp.M], questions{
			Id:      tmp.q.Id,
			Score:   tmp.q.Score,
			Options: "",
		})
	}

	return &m
}
