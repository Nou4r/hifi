package middleware

import (
	"api/config"
	"api/types"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"
)

func SigninUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req types.SigninRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	req.Username = strings.TrimSpace(req.Username)
	req.Password = strings.TrimSpace(req.Password)
	if req.Username == "" || req.Password == "" {
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}

	base := fmt.Sprintf("%s://%s", config.SubsonicScheme, config.SubsonicHost)

	jar, _ := cookiejar.New(nil)
	client := &http.Client{Jar: jar}

	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	loginCh := startLoginUser(ctx, client, base+"/admin/rest/ping.view", req.Username, req.Password, startLogin(ctx, client, base+"/admin/login_do", "jack", "123"))

	res := <-loginCh
	if res.Err != nil {
		http.Error(w, res.Err.Error(), http.StatusBadGateway)
		return
	}
	if res.Status >= 400 {
		http.Error(w, string(res.Body), http.StatusBadRequest)
		return
	}

	w.Header().Set(config.HeaderContentType, config.ContentTypeJSON)
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{"message": string(res.Body), "username": req.Username, "password": req.Password, "host": config.HostUrl})

}

func startLoginUser(ctx context.Context, client *http.Client, createURL, newUser, newPass string, loginCh <-chan types.LoginResult) <-chan types.CreateResult {
	out := make(chan types.CreateResult, 1)

	go func() {
		defer close(out)

		select {

		case lr := <-loginCh:
			if lr.Err != nil || !lr.OK {
				out <- types.CreateResult{Err: fmt.Errorf("login failed")}
				return
			}

		case <-ctx.Done():
			out <- types.CreateResult{Status: 0, Body: nil, Err: ctx.Err()}
			return
		}

		u, _ := url.Parse(createURL)
		q := u.Query()
		q.Set("u", newUser)
		q.Set("p", newPass)
		q.Set("c", "gonic")
		q.Set("f", "json")
		u.RawQuery = q.Encode()

		out <- types.CreateResult{
			Status: http.StatusOK,
			Body:   []byte(u.String()),
			Err:    nil,
		}

	}()
	return out
}
