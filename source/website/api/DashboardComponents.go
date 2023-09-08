package api

import (
	"EvoGuard/source/database"
	"EvoGuard/source/website/functions"
	"encoding/json"
	"net/http"
)

// BundledDashboardAdministratorResponse is what is returned from the BundledDashboardAdministrator
type BundledDashboardAdministratorResponse struct {
	Status                 bool      `json:"status,omitempty"`
	Username               string    `json:"username"`
	TotalMembers           int       `json:"totalMembers"`
	TotalActivatedLicenses int       `json:"totalActivatedLicenses"`
	TotalPurchasesToday    int       `json:"totalPurchasesToday"`
	NumberApplications     int       `json:"numberApplications"`
	WeeklySalesFields      [7]string `json:"weeklySalesFields"`
	WeeklySalesValues      [7]int    `json:"weeklySalesValues"`
	Applications           []string  `json:"applications"`
	ApplicationsValues     []int     `json:"applicationsValues"`
	ApplicationsColours    []string  `json:"applicationsColours"`
}

// BundledDashboardAdministrator will produce and return all the required values
// which we want to display on the dashboard, we sent these elements from the backend
// towards the frontend via an API request from the javascript file using the fetch
// function.
func BundledDashboardAdministrator(writer http.ResponseWriter, request *http.Request) {
	biscuit, ok := functions.ExtractJWTCrums(request)
	if !ok || biscuit == nil {
		WriteJson(map[string]any{"status": false, "error": "missing authorization"}, http.StatusUnauthorized, writer)
		return
	}

	owner, err := biscuit.GetOwnershipUser()
	if err != nil || owner == nil {
		WriteJson(map[string]any{"status": false, "error": "unable to find signed user inside your token"}, http.StatusUnauthorized, writer)
		return
	}

	members, err := database.DB.GetUsers()
	if err != nil {
		WriteJson(map[string]any{"status": false, "error": "error occurred while fetching from database"}, http.StatusInternalServerError, writer)
		return
	}

	/* This is what we attempt to marshal into a json formula */
	var field *BundledDashboardAdministratorResponse = &BundledDashboardAdministratorResponse{
		Status:                 true,
		Username:               owner.Username,
		Applications:           []string{"CP", "Compact", "KIDDIES"},
		TotalMembers:           len(members),
		TotalActivatedLicenses: 0,
		TotalPurchasesToday:    0,
		NumberApplications:     0,
		WeeklySalesFields:      [7]string{"Monday", "Tuesday"},
		WeeklySalesValues:      [7]int{10, 100},
		ApplicationsValues:     []int{90, 10, 1000},
		ApplicationsColours:    []string{"#0f0e17", "#ff8906", "#f25f4c", "#e53170", "#6246ea"},
	}

	content, err := json.Marshal(field)
	if err != nil {
		WriteJson(map[string]any{"status": true, "error": "event occurred which haulted the marshalling of response"}, http.StatusInternalServerError, writer)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write(content)
}
