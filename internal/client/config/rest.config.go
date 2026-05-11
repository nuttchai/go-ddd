package config

// RESTConfig holds HTTP listener settings.
type RESTConfig struct {
	port string
}

func (r *RESTConfig) GetPort() string {
	return r.port
}

func (r *RESTConfig) SetPort(port string) {
	r.port = port
}
