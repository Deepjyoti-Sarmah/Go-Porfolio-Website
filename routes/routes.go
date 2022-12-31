package routes

import (
	"net/http"

	"github.com/DeepjyotiSarmah/portfolio/controllers"
	"github.com/DeepjyotiSarmah/portfolio/middleware"
)

func Route() {
	http.HandleFunc("/", middleware.ErrorHandling(controllers.Template))
	// enabling and adding css files
	fileServer := http.FileServer(http.Dir("./views/static"))
	http.Handle("/static/", http.StripPrefix("/static", fileServer))
}
