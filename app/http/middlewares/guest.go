package middlewares

import (
	"net/http"

	"github.com/runningape/goblog/pkg/auth"
	"github.com/runningape/goblog/pkg/flash"
)

func Guest(next HTTPHandlerFunc) HTTPHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if auth.Check() {
			flash.Warning("登录用户无法访问此页面。")
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		next(w, r)
	}
}
