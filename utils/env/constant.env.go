package env

type MetaENV struct {
	Name string
	File string
}

var Production = &MetaENV{
	Name: "production",
	File: ".env",
}

var Local = &MetaENV{
	Name: "local",
	File: ".env.local",
}
