package views

import "net/http"

func RegisterHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	// User Group already set upon registration
	// Active = true upon registration
	// User Group = User upon registration
	return nil
}
