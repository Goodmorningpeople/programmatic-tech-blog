package handlers

import (
	"net/http"

	"github.com/Goodmorningpeople/booking_go/pkg/config"
	"github.com/Goodmorningpeople/booking_go/pkg/models"
	"github.com/Goodmorningpeople/booking_go/pkg/render"
)

// repository pattern (very, very cool)
var Repo *Respository

type Respository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Respository {
	return &Respository{
		App: a,
	}
}

func NewHandler(r *Respository) {
	Repo = r
}

// handler funcs that handle the rendered templates and write them to a server with oppotional logic passed to template (response writer w is passed when used in HandleFunc() for this)
func (m *Respository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}
