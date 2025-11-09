package middleware

import (
	"api/config"
	"api/types"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func SignupUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req types.SignupRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	req.Username = strings.TrimSpace(req.Username)
	req.Password = strings.TrimSpace(req.Password)

	if req.Username == "" || req.Password == "" {
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}

	endpoint := fmt.Sprintf("http://%s/admin/create_user_do", config.SubsonicHost)

	form := url.Values{}
	form.Set("username", req.Username)
	form.Set("password_one", req.Password)
	form.Set("password_two", req.Password)

	req2, err := http.NewRequest(config.MethodPost, endpoint, strings.NewReader(form.Encode()))
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}
	req2.Header.Set(config.HeaderContentType, config.ContentTypeForm)
	req2.Header.Add(config.Cookies, "next-auth.session-token=eyJhbGciOiJkaXIiLCJlbmMiOiJBMjU2R0NNIn0..jPtBFQtrrCNga6Y1.NlV8ZWOctkeQmLGcSPu202ycZK7lkkbj2uaBIFEOq_0b5jNgiRVSpKyvTVn_4CFtMRxxUcVt48ncdgdPgln16LCsiU9iCFY15GyohufVxIZhu8i0-Tyg7WnnKNhmA6COzbBP93XLUL3fRZvO5wA589O5vQTlOpg_EKsVDGDj17y957YYK6eh-VA0HJlWJH_ifdC3VmzLH9p9zgDpPHxTuddHH5fuZyqNb252Xn_a-zRjhRLRNUKJPoxJ80QaomI-va1D8P5-dI7jHhpKCrL2CMswhM0XtW1igUWrkrbDRHoOCK8M04D5Xc3yvXPa3G4IBxJNAAaKAatq0s5QzHzAmWFu4Bd2xtjnZryGI0usbQVOYR1tpnC1J90_XIFl.g6_Dx_QdH3Y8H0Q-m3r0Vg; gonic=MTc2MjY5NjQzOXxOd3dBTkZOTlFsQlpURUpXVlZKWFdVbFpUbGxMVUVwSlJFVTJUMHhNUTB4TFZrMVRRMWxDU2pST1Z6TkJRa3RCVTBkQ1NqWlVURkU9fMNa-5ztxiGjPfr3mSFigYQIUigq4oEdP18boELHpE5l")

	resp, err := http.DefaultClient.Do(req2)
	if err != nil {
		http.Error(w, "Failed to communicate with user service", http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)

	switch resp.StatusCode {
	case http.StatusInternalServerError:
		http.Error(w, string(b), http.StatusBadRequest)
		return

	default:
		w.Header().Set(config.HeaderContentType, config.ContentTypeJSON)
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"message": "User created successfully",
		})
		return
	}
}
