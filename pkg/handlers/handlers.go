package handlers

import (
	"log"
	"net/http"

	"github.com/serhatguzel/bookings/pkg/config"
	"github.com/serhatguzel/bookings/pkg/models"
	"github.com/serhatguzel/bookings/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	log.Println("Home")
	remoteIP := r.RemoteAddr

	m.App.Session.Put(r.Context(), "remote_ip", remoteIP) // IP adresini oturuma kaydet
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	log.Println("About")

	stringMap := make(map[string]string)

	stringMap["test"] = "Hello, again!"

	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip") // IP adresini oturumdan al
	stringMap["remote_ip"] = remoteIp                             // IP adresini stringMap'e ekle

	// TemplateData nesnesini oluştur ve StringMap'i ekle
	td := &models.TemplateData{
		StringMap: stringMap,
	}

	// RenderTemplate fonksiyonunu çağır
	render.RenderTemplate(w, "about.page.tmpl", td)
}
