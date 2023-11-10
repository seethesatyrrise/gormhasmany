package orm

type pgConfig struct {
	host     string
	user     string
	password string
	dbname   string
	port     int
}

func getDefaultConfig() *pgConfig {
	return &pgConfig{
		host:     "localhost",
		user:     "postgres",
		password: "postgres",
		dbname:   "gorm",
		port:     5432,
	}
}
