package controllers

import (
	"log"
	"net/http"
	"text/template"

	"github.com/DeepjyotiSarmah/portfolio/database"
	"github.com/DeepjyotiSarmah/portfolio/models"
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

func TemplateList() *template.Template {
	files := []string{
		"views/templates/header.html",
		"views/templates/navbar.html",
		"views/templates/about.html",
		// "views/templates/facts.html",
		"views/templates/skills.html",
		"views/templates/resume.html",
		"views/templates/portfolio.html",
		// "views/templates/services.html",
		"views/templates/testimonials.html",
		"views/templates/contact.html",
		"views/templates/footer.html",
		"views/templates/base.html",
	}
	return template.Must(template.ParseFiles(files...))
}

var Error = template.Must(template.ParseFiles("views/templates/pageerror.html"))

var Tmpl = TemplateList()

func LoadEnvVariable(key string) string {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatal(err)
	}

	return value
}

func Email(value [4]string) [4]string {
	mail := gomail.NewMessage()

	myEmail := LoadEnvVariable("EMAIL")
	appPassword := LoadEnvVariable("APP_PASSWORD")
	mail.SetHeader("From", myEmail)
	mail.SetHeader("To", myEmail)
	mail.SetHeader("Reply To", value[1])
	mail.SetHeader("Subject", value[2])
	mail.SetBody("text/plain", value[0]+" : \n\n"+value[3])

	a := gomail.NewDialer("smtp.gmail.com", 587, myEmail, appPassword)
	if err := a.DialAndSend(mail); err != nil {
		log.Fatal(err)
	}

	return value

}

func Template(w http.ResponseWriter, r *http.Request) error {
	if r.URL.Path != "/" {
		return Error.Execute(w, nil)
	}
	if r.Method != http.MethodPost {
		return Tmpl.Execute(w, nil)
	}

	send := models.Send{
		Name:    r.FormValue("name"),
		Email:   r.FormValue("email"),
		Subject: r.FormValue("subject"),
		Message: r.FormValue("message"),
	}
	_ = send
	value := [4]string{send.Name, send.Email, send.Subject, send.Message}
	_ = Email(value)
	database.Connection(send.Name, send.Email, send.Subject, send.Message)
	sm := struct{ Success bool }{true}
	return Tmpl.Execute(w, sm)

}
