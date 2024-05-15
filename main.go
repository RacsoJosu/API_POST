package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/racsoJosu/rest-ws/handler"
	"github.com/racsoJosu/rest-ws/middleware"
	"github.com/racsoJosu/rest-ws/server"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil{
		log.Fatal("Error ala cargar variables de entorno")
	}

	PORT := os.Getenv("PORT")
	JWT_SECRET := os.Getenv("JWT_SECRET")
	DATABASE_URL := os.Getenv("DATABASE_URL")

	s, err := server.NewServer(context.Background(), &server.Config{
		JWTSecret: JWT_SECRET,
		Port: PORT,
		DataBaseUrl: DATABASE_URL,
	})

	if err != nil {
		log.Fatal(err)
	}
	s.Run(BinderRoutes)
}

func BinderRoutes(s server.Server, r *mux.Router){
	r.Use(middleware.CheckAuthMiddleware(s))
	r.HandleFunc("/", handler.HomeHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/signup", handler.SignUpHandler(s)).Methods(http.MethodPost)
	
	r.HandleFunc("/login", handler.LoginHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/me", handler.MeHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/createPost", handler.InsertPostHandler(s)).Methods(http.MethodPost)
	

}