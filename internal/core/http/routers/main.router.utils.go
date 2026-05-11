package router

type urlBuilder struct {
	basePath string
}

func (u *urlBuilder) buildPath(urlPath ...string) string {
	path := u.basePath
	if len(urlPath) > 0 {
		path = path + urlPath[0]
	}
	return path
}

func newUrlBuilder(basePath string, baseDomain ...string) *urlBuilder {
	if len(baseDomain) > 0 {
		basePath = basePath + baseDomain[0]
	}
	return &urlBuilder{basePath: basePath}
}
