package main

import (
	"bench/controllers"
	"bench/logger"
	"bench/utils"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

func main() {
	var wg sync.WaitGroup
	log := logger.New("bench")
	port := utils.GetEnv("PORT", "3000")
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(cors.AllowAll().Handler)
	r.Mount("/", controllers.New())

	wg.Add(1)

	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%s", port), r)

		if err != nil {
			log.Error(err)
		}

		wg.Done()
	}()

	log.Infof("listening on port %s\n", port)
	wg.Wait()
}
