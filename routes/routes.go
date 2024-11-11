package routes

import (
	"auth-service/controllers"
	"net/http"
	"auth-service/middleware"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/signup", controllers.SignUp)
	mux.HandleFunc("/signin", controllers.SignIn)
	mux.HandleFunc("/refresh-token", controllers.RefreshToken)
	
	mux.Handle("/protected", middleware.AuthMiddleware(http.HandlerFunc(controllers.Protected)))
	
	return mux
}
