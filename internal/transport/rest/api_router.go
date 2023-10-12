package rest

import (
	"anime/internal/services"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"log"
	"net/http"
)

type APIRouter struct {
	router *chi.Mux
}

func New() *chi.Mux {
	r := chi.NewRouter()

	apiRouter := &APIRouter{router: r}
	r.Get("/home", home)

	r.Mount("/v1", apiRouter.V1())

	return r
}

func (r *APIRouter) V1() *chi.Mux {
	v1 := chi.NewRouter()

	v1.Get("/end", end)

	return v1
}

func home(w http.ResponseWriter, req *http.Request) {
	id := uuid.New()
	sub := make([]services.Subtitles, 1)
	sub[0] = services.Subtitles{
		ID:       id.String(),
		Language: "ru",
		Author:   "GrapefruitSama",
	}

	model := services.Anime{
		ID:           id.String(),
		Language:     "ru",
		Title:        "Восхождение в тени.",
		Genre:        "comedy",
		EpisodeCount: 24,
		DateRelease:  "27.08.1999",
		Subtitles:    sub,
	}

	data, err := json.Marshal(model)
	if err != nil {
		log.Printf("json.Marshal error: %v", err)
	}

	fmt.Fprintf(w, "%s", data)
}

func end(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Sasi End")
}
