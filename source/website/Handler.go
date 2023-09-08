package website

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"EvoGuard/source/website/api"
	"EvoGuard/source/website/pages"
)

/*
	Handler.go will introduce the required concept for
	making the HTTPS server requestable and reachable
	from remote clients.
*/

// NewHandler will try to spawn a new server HTTP handler
func NewHandler() error {
	server, err := NewConfig()
	if err != nil {
		return err
	}

	handler := mux.NewRouter()
	handler.PathPrefix(server.Assets.ServeUnder).Handler(http.FileServer(http.Dir(server.Assets.Directory)))
	service := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", server.Address, server.Port),
		Handler: handler,
	}

	// Main handling for each page
	handler.Use(pages.MiddlewareFunc)
	handler.HandleFunc("/", pages.Login)
	handler.HandleFunc("/login", pages.Login)
	handler.HandleFunc("/register", pages.RegisterHandler)
	handler.HandleFunc("/activate/{id}", pages.ActivateHandler)

	/* Dashboard handling registering */
	dashboard := handler.PathPrefix("/{token}/dashboard").Subrouter()
	dashboard.Use(pages.DashboardMiddleware)
	dashboard.HandleFunc("/", pages.Dashboard)
	dashboard.HandleFunc("/apps", pages.DashboardApps)
	dashboard.HandleFunc("/clients", pages.DashboardClients)
	dashboard.HandleFunc("/products", pages.DashboardProducts)
	dashboard.HandleFunc("/settings", pages.DashboardSettings)
	dashboard.HandleFunc("/messages", pages.DashboardMessages)
	dashboard.HandleFunc("/analytics", pages.DashboardAnalytics)
	dashboard.HandleFunc("/administration", pages.DashboardAdministration)

	/* API handling registering */
	apiRouter := handler.PathPrefix("/{token}/api/").Subrouter()
	apiRouter.Use(api.MiddlewareFunc)
	apiRouter.HandleFunc("/me", api.Myself)
	apiRouter.HandleFunc("/signout", api.Signout)
	apiRouter.HandleFunc("/users/getusers", api.GetUsers)
	apiRouter.HandleFunc("/pages/dashboard", api.BundledDashboardAdministrator)
	apiRouter.HandleFunc("/users/createuser", api.CreateUser)

	cancel := make(chan error)
	go NewNotify(cancel, server)
	go pages.Cleaner(60)
	defer close(cancel)

	/* ListenAndServeTLS will actually boot the server */
	//if err := service.ListenAndServeTLS(server.Cert, server.Key); err != nil {
	//	cancel <- err
	//	time.Sleep(500 * time.Millisecond)
	//	return err
	//}

	if err := service.ListenAndServe(); err != nil {
		cancel <- err
		time.Sleep(500 * time.Millisecond)
		return err
	}

	return nil
}
