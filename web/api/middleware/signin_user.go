package middleware

import (
	"api/config"
	"api/types"
	"fmt"
	"net/http"
	"strings"
)

func SigninUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req types.SigninRequest

	req.Username = strings.TrimSpace(req.Username)
	req.Password = strings.TrimSpace(req.Password)
	if req.Username == "" || req.Password == "" {
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}

	base := fmt.Sprintf("%s://%s", config.SubsonicScheme, config.SubsonicHost)

	fmt.Println(base)

}
