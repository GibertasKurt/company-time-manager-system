package api

import (
	"net/http"
	"strings"
)

func APIHandler(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/api")
	r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")

	if strings.HasPrefix(r.URL.Path, "/registration") {
		RegistrationHandler(w, r)
		return
	}
	if strings.HasPrefix(r.URL.Path, "/update_break") {
		UpsBreakHandler(w, r)
		return
	}
}

/*
fetch("/api/registration", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            // "X-CSRFToken": csrfToken
        },
        body: JSON.stringify(logData),
    })
        .then(response => {
            if (!response.ok) {
                throw new Error(`Server responded with status ${response.status}`);
            }
            return response.json();
        })
        .then(data => {
            console.log("Server response:", data);
            // Handle success (e.g., redirect or show message)
        })
        .catch(error => {
            console.error("Fetch error:", error);
            alert("Failed to register: " + error.message + ". Please check your inputs.");
        });
*/
