package types

type RESTConfig struct {
	port string
}

type DBConfig struct {
	dsn    string
	driver string
}

type AppConfig struct {
	rest RESTConfig
	db   DBConfig
	env  string
}

func (c *AppConfig) GetRESTPort() string {
	return c.rest.port
}

func (c *AppConfig) GetDBConfig() *DBConfig {
	return &c.db
}

func (db *DBConfig) GetDSN() string {
	return db.dsn
}

func (db *DBConfig) GetDriver() string {
	return db.driver
}

func (c *AppConfig) GetENV() string {
	return c.env
}

func (c *AppConfig) SetRESTConfig(port string) {
	c.rest.port = port
}

func (c *AppConfig) SetDBConfig(dsn, driver string) {
	c.db.dsn = dsn
	c.db.driver = driver
}

func (c *AppConfig) SetENV(env string) {
	c.env = env
}