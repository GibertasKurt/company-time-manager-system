package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gibertaskurt/company-time-manager-system/models"
	"github.com/uadmin/uadmin"
)

func RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("RegistrationHandler called")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		uadmin.ReturnJSON(w, r, map[string]any{
			"status":  "error",
			"err_msg": "Invalid HTTP method",
		})
		return
	}

	// Define a struct matching the expected JSON payload
	var reqBody struct {
		FirstName  string `json:"first_name"`
		LastName   string `json:"last_name"`
		Username   string `json:"username"`
		Password   string `json:"password"`
		Department string `json:"department_id"`

		// Add more fields here if needed (last_name, email, etc.)
	}

	// Decode the JSON request
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		uadmin.ReturnJSON(w, r, map[string]any{
			"status":  "error",
			"err_msg": "Invalid JSON",
		})
		return
	}

	fmt.Println("username: ", reqBody.Username)
	fmt.Println("password: ", reqBody.Password)
	fmt.Println("firstname: ", reqBody.FirstName)
	fmt.Println("lastname: ", reqBody.LastName)
	fmt.Println("department: ", reqBody.Department)
	// Check if the username already exists
	var existingUser uadmin.User
	if err := uadmin.Get(&existingUser, "username = ?", reqBody.Username); err == nil {
		w.WriteHeader(http.StatusConflict)
		uadmin.ReturnJSON(w, r, map[string]any{
			"status":  "error",
			"err_msg": "Username already exists",
		})
		return
	}

	// Save user
	user := uadmin.User{
		Username:  reqBody.Username,
		Password:  reqBody.Password,
		Email:     reqBody.Username,
		FirstName: reqBody.FirstName,
		LastName:  reqBody.LastName,
	}

	if err := uadmin.Save(&user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		uadmin.ReturnJSON(w, r, map[string]any{
			"status":  "error",
			"err_msg": "Failed to save user",
		})
		return
	}

	// Save employee
	employee := models.Employee{
		UserID:        user.ID,
		FirstName:     reqBody.FirstName,
		LastName:      reqBody.LastName,
		Email:         reqBody.Username,
		DepartmentsID: 1, // Default department ID, change as needed
		IsActive:      true,
	}
	uadmin.Save(&employee)

	uadmin.ReturnJSON(w, r, map[string]interface{}{
		"status":   "ok",
		"response": "User registered successfully",
		"user":     user,
	})
}
