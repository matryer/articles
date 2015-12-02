package rest

import "net/http"

type Router interface {
	Route(pathpattern, method string, handler http.Handler)
}

type RestReader interface {
	Read(w http.ResponseWriter, r *http.Request)
}

type RestManyReader interface {
	ReadMany(w http.ResponseWriter, r *http.Request)
}

type RestCreator interface {
	Create(w http.ResponseWriter, r *http.Request)
}

type RestUpdater interface {
	Update(w http.ResponseWriter, r *http.Request)
}

type RestReplacer interface {
	Replace(w http.ResponseWriter, r *http.Request)
}

type RestDeleter interface {
	Delete(w http.ResponseWriter, r *http.Request)
}

func Route(path string, controller interface{}, router Router) {
	var n int
	if c, ok := controller.(RestReader); ok {
		router.Route(path+"/{id}", "GET", http.HandlerFunc(c.Read))
		n++
	}
	if c, ok := controller.(RestManyReader); ok {
		router.Route(path, "GET", http.HandlerFunc(c.ReadMany))
		n++
	}
	if c, ok := controller.(RestCreator); ok {
		router.Route(path, "POST", http.HandlerFunc(c.Create))
		n++
	}
	if c, ok := controller.(RestDeleter); ok {
		router.Route(path+"/{id}", "DELETE", http.HandlerFunc(c.Delete))
		n++
	}
	if c, ok := controller.(RestReplacer); ok {
		router.Route(path+"/{id}", "PUT", http.HandlerFunc(c.Replace))
		n++
	}
	if c, ok := controller.(RestUpdater); ok {
		router.Route(path+"/{id}", "PATH", http.HandlerFunc(c.Update))
		n++
	}
	if n == 0 {
		panic("rest: controller has no suitable methods")
	}
}

type PresentsController struct{}

func (p PresentsController) ReadMany(w http.ResponseWriter, r *http.Request) {
	// TODO: handle read many request
}

func main() {
	router := &r{}
	Route("/presents", PresentsController{}, router)
}

type r struct{}

func (r *r) Route(pathpattern, method string, handler http.Handler) {

}
