package handlers

import (
	"net/http"

	templates "github.com/joeyjurjens/peekaboo/pkg/templates/dashboard"
)

func DashboardIndexHandler(w http.ResponseWriter, r *http.Request) {
	component := templates.DashboardIndex()
	component.Render(r.Context(), w)
}
