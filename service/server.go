package service

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/patrickmn/go-cache"
)

type Server struct {
	httpServer *http.Server
	cacheStore *cache.Cache
}

func NewServer(port string, cacheStore *cache.Cache) *Server {
	return &Server{
		httpServer: &http.Server{Addr: ":" + port},
		cacheStore: cacheStore,
	}
}

func (s *Server) Start() error {
	http.HandleFunc("/orders/", s.GetById)

	var startingError error = nil
	go func() {
		// the goroutine will be finished when Stop() is called
		if err := s.httpServer.ListenAndServe(); err != nil {
			startingError = err
		}
	}()

	log.Println("HTTP server is listening on ", s.httpServer.Addr)
	return startingError
}

func (s *Server) Stop() error {
	log.Println("Stopping an HTTP server on ", s.httpServer.Addr)
	return s.httpServer.Shutdown(context.TODO())
}

func (s *Server) GetById(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/orders/"):]

	order, exists := s.cacheStore.Get(id)
	if exists {
		t, _ := template.ParseFiles("service/templates/order.html")
		t.Execute(w, order)
		//fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", order.(entity.Order).OrderUid, order.(entity.Order).TrackNumber)
	} else {
		fmt.Fprintf(w, "<h1>404</h1><div>No order with that id.</div>")
	}
}
