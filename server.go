package server

type Server struct {
	Addr string `toml:"addr"`
	Port string `toml:"port"`
}

type PostgresCFG struct {
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	DBName   string `toml:"dbname"`
	SSLMode  string `toml:"sslmode"`
}

type Config struct {
	PGcfg  PostgresCFG `toml:"pg_db"`
	SRVcfg Server      `toml:"server"`
}
