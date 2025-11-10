package middleware

import (
	"api/types"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func startLogin(ctx context.Context, client *http.Client, loginDoURL, user, pass string) <-chan types.LoginResult {
	token := make(chan types.LoginResult, 1)
	go func() {
		defer close(token)

		form := url.Values{}
		form.Set("username", user)
		form.Set("password", pass)

		req, err := http.NewRequestWithContext(ctx, http.MethodPost, loginDoURL, strings.NewReader(form.Encode()))
		if err != nil {
			token <- types.LoginResult{OK: false, Err: err}
			return
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		resp, err := client.Do(req)
		if err != nil {
			token <- types.LoginResult{OK: false, Err: err}
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode >= 400 {
			b, _ := io.ReadAll(resp.Body)
			token <- types.LoginResult{OK: false, Err: fmt.Errorf("login failed: %d: %s", resp.StatusCode, string(b))}
			return
		}
		token <- types.LoginResult{OK: true, Err: nil}
	}()
	return token
}
