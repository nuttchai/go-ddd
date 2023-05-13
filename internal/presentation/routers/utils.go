package router

const (
	basePath = "/api/v1/"
)

func buildPath(path string) string {
	return basePath + path
}
