package models

// Settings app settings
type Settings struct {
	App      App      `json:"app"`
	Postgres Postgres `json:"postgres"`
	Redis    Redis    `json:"redis"`
}
type Redis struct {
	RedisUri string `json:"url"`
}

// Params contains params of server meta data
type App struct {
	ServerName string `json:"serverName"`
	PortRun    string `json:"portRun"`
	LogFile    string `json:"logFile"`
	Token      string `json:"token"`
}

type Postgres struct {
	PG_host     string `json:"pg_host"`
	PG_port     string `json:"pg_port"`
	PG_user     string `json:"pg_user"`
	PG_password string `json:"pg_password"`
	PG_name     string `json:"pg_name"`
}
