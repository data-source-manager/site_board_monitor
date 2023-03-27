package config

type Config struct {
	Name  string      `json:"name,omitempty"`
	Mysql MysqlConfig `json:"Mysql"`
	Log   LogConfig   `json:"Log"`
	Redis RedisConfig `json:"Redis"`
}
type MysqlConfig struct {
	Host     string `json:"host" `
	Port     int    `json:"port" `
	User     string `json:"user" `
	Password string `json:"password" `
	Db       string `json:"db" `
}

type LogConfig struct {
	Model      string `json:"model,omitempty" default:"dev"`
	Level      string ` json:"level" default:"warn" `
	FileName   string ` json:"file_name"`
	MaxSize    int    ` json:"max_size"`
	MaxAge     int    ` json:"max_age"`
	MaxBackups int    ` json:"max_backups"`
}

type RedisConfig struct {
	Host          string `json:"host"`
	Port          int    `json:"port"`
	Pwd           string `json:"pwd"`
	Db            int    `json:"db"`
	SiteBoardMsg  string `json:"site_board_msg"`
	NewsKey       string `json:"news_key"`
	BoardErrorKey string `json:"error_board_key"`
}
