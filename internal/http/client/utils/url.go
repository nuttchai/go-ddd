package url

const (
	basePath = "/api/v1/"
)

type UrlBuilder struct {
	domain string
}

func (u *UrlBuilder) BuildPath(path ...string) string {
	urlPath := ""
	if len(path) > 0 {
		urlPath = path[0]
	}
	return basePath + u.domain + urlPath
}

func NewUrlBuilder(domain ...string) *UrlBuilder {
	baseDomain := ""
	if len(domain) > 0 {
		baseDomain = domain[0]
	}
	return &UrlBuilder{
		domain: baseDomain,
	}
}
