package url

const (
	basePath = "/api/v1/"
)

func BuildPath(path string) string {
	return basePath + path
}
