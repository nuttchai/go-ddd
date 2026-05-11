package config

// DBConfig holds database connection metadata and the live driver instance.
type DBConfig struct {
	dsn      string
	driver   string
	instance any
}

func (db *DBConfig) GetDSN() string {
	return db.dsn
}

func (db *DBConfig) SetMetaData(dsn, driver string) {
	db.dsn = dsn
	db.driver = driver
}

func (db *DBConfig) SetInstance(instance any) {
	db.instance = instance
}
