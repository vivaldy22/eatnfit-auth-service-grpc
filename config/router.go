package config

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/vivaldy22/eatnfit-auth-service/master/level"
	"github.com/vivaldy22/eatnfit-auth-service/middleware"
	auth_service "github.com/vivaldy22/eatnfit-auth-service/proto"
	"github.com/vivaldy22/eatnfit-auth-service/tools/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func CreateRouter() *mux.Router {
	return mux.NewRouter()
}

func RunServer(r *mux.Router) {
	host := viper.ViperGetEnv("GRPC_HOST", "localhost")
	port := viper.ViperGetEnv("GRPC_PORT", "1010")

	listener, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		log.Fatal(err)
	}

	srv := grpc.NewServer()
	auth_service.RegisterLevelCRUDServer(srv, &level.Service{})
	reflection.Register(srv)

	if err = srv.Serve(listener); err != nil {
		log.Fatal(err)
	}
	log.Printf("Starting GRPC Eat N' Fit Auth Server at %v port: %v\n", host, port)
}

func InitRouters(db *sql.DB, r *mux.Router) {
	r.Use(middleware.ActivityLogMiddleware)

	levelService := level.NewService(db)
	level.NewHandler(levelService, r)
}
