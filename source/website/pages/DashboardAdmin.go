package pages

import (
	"EvoGuard/source/database"
	"EvoGuard/source/website/functions"
	"net/http"
)

// AdminDashboard is the panel for admin management and moderation of other products
func AdminDashboard(writer http.ResponseWriter, request *http.Request, user *database.User) {
	functions.WriteHypertext(writer, make(map[string]string), "resources", "public", "html", "admin", "dashboard.html")
}
