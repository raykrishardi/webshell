package handler

import (
	"net/http"

	"github.com/raykrishardi/webshell-front/internal/entity"
	"github.com/raykrishardi/webshell-front/internal/pkg/config"
	"github.com/raykrishardi/webshell-front/internal/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "webshell.page.gohtml", &entity.TemplateData{})
}
