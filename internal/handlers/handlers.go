package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/serhatguzel/bookings/internal/config"
	"github.com/serhatguzel/bookings/internal/models"
	"github.com/serhatguzel/bookings/internal/render"
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
	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
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
	render.RenderTemplate(w, r, "about.page.tmpl", td)
}

func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	log.Println("Generals")
	render.RenderTemplate(w, r, "generals.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	log.Println("Majors")
	render.RenderTemplate(w, r, "majors.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	log.Println("Availability")
	render.RenderTemplate(w, r, "search-availability.page.tmpl", &models.TemplateData{})
}

func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	log.Println("Post Availability")
	// Parse form data
	err := r.ParseForm()
	if err != nil {
		log.Println("Error parsing form:", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Get the form values
	start := r.Form.Get("start")
	end := r.Form.Get("end")

	// Log the values for debugging
	log.Println("Start date:", start)
	log.Println("End date:", end)

	w.Write([]byte(fmt.Sprintf("start date is %s and end date is %s", start, end))) // Placeholder response
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	log.Println("Availability JSON")
	resp := jsonResponse{
		OK:      true,
		Message: "Available",
	}

	out, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		log.Println("Error marshalling JSON:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	log.Println(string(out)) // Log the JSON response for debugging
	w.Write(out)
	w.Header().Set("Content-Type", "application/json") // Set the content type to JSON
}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	log.Println("Contact")
	render.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})
}

func (m *Repository) MakeReservation(w http.ResponseWriter, r *http.Request) {
	log.Println("Make Reservation")
	render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{})
}
