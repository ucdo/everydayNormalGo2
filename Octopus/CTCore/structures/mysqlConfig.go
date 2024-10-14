package structures

// MysqlCnf mysql 数据库配置
type MysqlCnf struct {
	User      string `json:"user"`      // JSON键 "user" 映射到此字段
	Password  string `json:"password"`  // JSON键 "password" 映射到此字段
	Host      string `json:"host"`      // JSON键 "host" 映射到此字段
	Port      string `json:"port"`      // JSON键 "port" 映射到此字段
	DbName    string `json:"dbname"`    // JSON键 "dbname" 映射到此字段
	CycleTime int    `json:"cycleTime"` // JSON键 "cycleTime" 映射到此字段
}
