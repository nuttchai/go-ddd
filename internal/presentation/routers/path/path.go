package path

const (
	basePath = "/api/v1/"
)

func CreatePath(path string) string {
	return basePath + path
}
