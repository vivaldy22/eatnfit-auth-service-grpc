package config

import (
	"database/sql"
	"fmt"
	"github.com/vivaldy22/eatnfit-auth-service/master/level"
	authservice "github.com/vivaldy22/eatnfit-auth-service/proto"
	"github.com/vivaldy22/eatnfit-auth-service/tools/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func InitDB() (*sql.DB, error) {
	dbUser := viper.ViperGetEnv("DB_USER", "root")
	dbPass := viper.ViperGetEnv("DB_PASSWORD", "password")
	dbHost := viper.ViperGetEnv("DB_HOST", "localhost")
	dbPort := viper.ViperGetEnv("DB_PORT", "3306")
	schemaName := viper.ViperGetEnv("DB_SCHEMA", "schema")
	driverName := viper.ViperGetEnv("DB_DRIVER", "mysql")

	dbPath := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, schemaName)
	dbConn, err := sql.Open(driverName, dbPath)

	if err != nil {
		return nil, err
	}

	if err = dbConn.Ping(); err != nil {
		return nil, err
	}

	return dbConn, nil
}

func RunServer(db *sql.DB) {
	host := viper.ViperGetEnv("GRPC_HOST", "localhost")
	port := viper.ViperGetEnv("GRPC_PORT", "1010")

	listener, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		log.Fatal(err)
	}

	srv := grpc.NewServer()
	service := level.NewService(db)
	authservice.RegisterLevelCRUDServer(srv, service)
	reflection.Register(srv)

	if err = srv.Serve(listener); err != nil {
		log.Fatal(err)
	}
	log.Printf("Starting GRPC Eat N' Fit Auth Server at %v port: %v\n", host, port)
}