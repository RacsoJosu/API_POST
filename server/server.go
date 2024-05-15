package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/racsoJosu/rest-ws/database"
	"github.com/racsoJosu/rest-ws/repository"
)

type Config struct {
	Port        string
	JWTSecret   string
	DataBaseUrl string
}

type Server interface {
	Config() *Config
}

type Broker struct {
	config *Config
	router *mux.Router
}

func (b *Broker) Config() *Config {
	return b.config
}

func NewServer(ctx context.Context, config *Config)(*Broker, error){
	if config.Port == "" {
		return nil, errors.New("El puerto es requerido")
		
	}
	if config.JWTSecret == "" {
		return nil, errors.New("El jwt es requerido")
		
	}
	if config.DataBaseUrl == "" {
		return nil, errors.New("La base de datos es requerida")
	}

	broker := &Broker{
		config: config,
		router: mux.NewRouter(),
	}

	return broker, nil
}

func (b *Broker) Run(binder func(s Server, r *mux.Router)){
	b.router = mux.NewRouter()
	binder(b, b.router)
	repo, err := database.NewPostgresRepository(b.config.DataBaseUrl)
	if err != nil{
		log.Fatal(err)
	}

	repository.SetRepository(repo);
	log.Printf("El servidor ha iniciado en el puerto %s\n ", b.config.Port);
	if err := http.ListenAndServe(b.config.Port, b.router ); err != nil{
		log.Fatal("Error: \n", err)
	}

}