package handlers

import (
	"net/http"
	"github.com/Kdarks3215/gardens/pkg/config"
	"github.com/Kdarks3215/gardens/pkg/models"
	"github.com/Kdarks3215/gardens/pkg/render"
)

//Repo the repository used by handlers
var Repo *Respository

//Repository is the repository type
type Respository struct {
	App *config.AppConfig
}

//NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Respository {
	return &Respository{
		App :a,
	}
}
//NewHandlers sets the repository for handlers
func NewHandlers ( r* Respository) {
	Repo = r
}

//Home is handler for home page
func (m *Respository)Home(w http.ResponseWriter, r *http.Request){
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

//About is handler for about page
func (m * Respository)About(w http.ResponseWriter, r *http.Request){
	//perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] 	= remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{StringMap: stringMap})
}