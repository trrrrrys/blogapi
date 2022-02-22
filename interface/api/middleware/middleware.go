package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/unrolled/render"

	"github.com/gorilla/handlers"
	"github.com/urfave/negroni"
)

type Request struct {
	Path   string
	Method string
}

var verifyTargetList []*Request = []*Request{
	{
		Path:   "/v1/contents",
		Method: "POST",
	},
}
var rendering = render.New(render.Options{})

func SetMiddleware(h http.Handler) http.Handler {
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD"})
	headersOk := handlers.AllowedHeaders(([]string{"Authorization", "Content-Type"}))
	credentialOk := handlers.AllowCredentials()
	handler := handlers.CORS(originsOk, methodsOk, headersOk, credentialOk)(h)
	n := negroni.New()
	n.Use(negroni.HandlerFunc(VerifyToken))
	n.UseHandler(handler)
	return n
}

func VerifyToken(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// mutationのみトークン認証
	for _, target := range verifyTargetList {
		if r.URL.Path == target.Path && r.Method == target.Method {
			authHeader := r.Header.Get("Authorization")
			token := strings.Replace(authHeader, "Bearer ", "", 1)
			// vc := verify.NewVerifyClient()
			// if !vc.VerifyToken(token) {
			// 	log.Println("Verify Error")
			// 	// エラー処理
			// 	rendering.JSON(w, http.StatusUnauthorized, "Error")
			// 	return
			// }
			log.Printf("Verify Token %v\n", token)
		}
	}
	next.ServeHTTP(w, r)
}
