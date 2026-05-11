package router

type Router struct {
	userUrlBuilder *urlBuilder
}

func NewRouter() *Router {
	return &Router{
		userUrlBuilder: newUrlBuilder("/api/v1", "/user"),
	}
}
