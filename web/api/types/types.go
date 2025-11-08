package types

import "net/http"

type Routes struct {
	User string
}

func (r *Routes) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	_, _ = w.Write([]byte(r.User))
}
