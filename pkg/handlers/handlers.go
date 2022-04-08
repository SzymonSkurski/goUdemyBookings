package handlers

import (
	"net/http"

	"github.com/SzymonSkursrki/goUdemyBookings/pkg/config"
	"github.com/SzymonSkursrki/goUdemyBookings/pkg/models"
	"github.com/SzymonSkursrki/goUdemyBookings/pkg/render"
)

// Repo ther repository used by handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// New Handlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	//how to send data to template ?
	//inject by reference &TemplateData{}
	//some fake data to inject
	stringMap := make(map[string]string)
	stringMap["test"] = "test string"
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
