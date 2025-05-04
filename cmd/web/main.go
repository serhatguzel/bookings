package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/serhatguzel/bookings/internal/config"
	"github.com/serhatguzel/bookings/internal/handlers"
	"github.com/serhatguzel/bookings/internal/render"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager // SCS oturum yöneticisi

func main() {

	// change this to true when in production
	app.InProduction = false

	session = scs.New()                            // SCS oturum yöneticisi oluşturuyoruz
	session.Lifetime = 24 * time.Hour              // oturumun süresi 24 saat
	session.Cookie.Persist = true                  // tarayıcı kapatıldığında oturum devam etsin
	session.Cookie.SameSite = http.SameSiteLaxMode // tarayıcıdan gelen isteklerde oturum bilgisi gönderilsin
	session.Cookie.Secure = app.InProduction       // sadece HTTPS üzerinden gönderilsin

	app.Session = session // app'e oturum yöneticisini ekliyoruz

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache", err)
	}

	app.TemplateCache = tc
	app.UseCache = false // true ise şablonu disten okur false ise diskten okumaz

	repo := handlers.NewRepo(&app) // Repository'i başlatıyoruz ve app'i Repository'e bağlıyoruz

	handlers.NewHandlers(repo) // Repository'i Handlers'e bağlıyoruz

	render.NewTemplates(&app)

	// Favicon isteğini engelle
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		// Tarayıcılar genellikle bir web sayfasını yüklerken /favicon.ico isteği yapar.
		// Bu istek, Home veya About handler'larınızın çağrılmasına neden olabilir.
		//Bu handler hiçbir şey yapmaz ve boş bir yanıt döner.
	})

	fmt.Println("Starting the application on port", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)

}
