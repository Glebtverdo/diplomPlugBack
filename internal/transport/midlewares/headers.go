package midlewares

import "net/http"

func GlobalHeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set your global headers here
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Requested-With", "XMLHttpRequest")

		// Call the next handler, which can be another middleware in the chain or the final handler.
		next.ServeHTTP(w, r)
	})
}
