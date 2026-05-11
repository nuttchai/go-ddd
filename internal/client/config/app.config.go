package config

type AppConfig struct {
	env  EnvConfig
	rest RESTConfig
	db   DBConfig
}

var DefaultAppConfig = &AppConfig{}

func (c *AppConfig) SetRESTConfig(port string) {
	c.rest.SetPort(port)
}

func (c *AppConfig) SetDBMetaData(dsn, driver string) {
	c.db.SetMetaData(dsn, driver)
}

func (c *AppConfig) SetDBInstance(db any) {
	c.db.SetInstance(db)
}

func (c *AppConfig) SetENV(env string) {
	c.env.Set(env)
}

func (c *AppConfig) GetENV() string {
	return c.env.Get()
}

func (c *AppConfig) GetRESTPort() string {
	return c.rest.GetPort()
}

func (c *AppConfig) GetDBConfig() *DBConfig {
	return &c.db
}
