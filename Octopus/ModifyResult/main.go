package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"os"
)

type mysqlConfig struct {
	Username string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
}

type conclusions struct {
	Id     int    `json:"id"`
	Ctx    string `json:"ctx"`
	Mark   string `json:"mark"`
	Advice string `json:"advice"`
}

type modify struct {
	Measure     int           `json:"measure"`
	Conclusions []conclusions `json:"conclusions"`
}

type config struct {
	Mysql  mysqlConfig `json:"mysql"`
	Modify modify      `json:"modify"`
}

func main() {
	// 连接数据库
	// 读取修改的配置
	// 备份
	// 修改
	cnf, err := cnf()
	if err != nil {
		panic(err)
	}

	db, err := conn(cnf.Mysql)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	sql1 := "SELECT c.id,c.conclusion  FROM qy_user_evaluate_log l left join qy_user_meaconclusion c on l.id=c.ulogid WHERE l.measure_id = ?"
	res, err := db.Query(sql1, cnf.Modify.Measure)
	if err != nil {
		panic(err)
	}

	type result struct {
		ID         int
		Conclusion string
	}

	type sqlCon struct {
		Name         string `json:"name"`
		Score        any    `json:"score"`
		FactorId     int    `json:"factor_id"`
		ConclusionId int    `json:"conclusion_id"`
		Mark         string `json:"mark"`
		Reference    string `json:"reference"`
		Comment      string `json:"comment"`
		Advice       string `json:"advice"`
		GraphShow    int    `json:"graph_show"`
	}

	var resx []result

	for res.Next() {
		var u result
		err := res.Scan(&u.ID, &u.Conclusion)
		if err != nil {
			panic(err)
		}

		resx = append(resx, u)
	}

	cmap := make(map[int]struct {
		Mark   string
		Ctx    string
		Advice string
	})
	updatex := make(map[int]string)
	for _, conclusion := range cnf.Modify.Conclusions {
		cmap[conclusion.Id] = struct {
			Mark   string
			Ctx    string
			Advice string
		}{Mark: conclusion.Mark, Ctx: conclusion.Ctx, Advice: conclusion.Advice}
	}

	for _, x := range resx {
		var s []*sqlCon
		needModify := false

		err := json.Unmarshal([]byte(x.Conclusion), &s)
		if err != nil {
			panic(err)
		}
		for key, xx := range s {
			_, ok := cmap[xx.ConclusionId]
			if !ok {
				continue
			}

			needModify = true
			s[key].Mark = cmap[xx.ConclusionId].Mark
			s[key].Comment = cmap[xx.ConclusionId].Ctx
			s[key].Advice = cmap[xx.ConclusionId].Advice
		}

		if needModify {
			marshal, err := json.Marshal(s)
			if err != nil {
				panic(err)
			}
			updatex[x.ID] = string(marshal)
		}

	}

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	stmt, err := tx.Prepare("UPDATE qy_user_meaconclusion SET conclusion = ? WHERE id = ?")
	if err != nil {
		tx.Rollback()
		panic(err)
	}

	for id, conclusion := range updatex {
		if _, err := stmt.Exec(conclusion, id); err != nil {
			tx.Rollback()
			panic(err)
		}
	}

	if err := tx.Commit(); err != nil {
		panic(err)
	}

}

// 目前就只有mysql
func conn(cfg mysqlConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	return sql.Open("mysql", dsn)
}

// 读取配置
func cnf() (config, error) {
	// "用户名:密码@tcp(数据库地址:端口)/数据库名?charset=utf8mb4&parseTime=True&loc=Local"
	file, err := os.OpenFile("config.json", os.O_RDONLY, os.ModePerm)
	defer file.Close()
	if err != nil {
		return config{}, err
	}

	byteFile, err := io.ReadAll(file)
	if err != nil {
		return config{}, err
	}
	var cfg config

	err = json.Unmarshal(byteFile, &cfg)

	if err != nil {
		return config{}, err
	}

	return cfg, err
}
