package main

import (
	"awesomeProject/internal/pkg/handler/handler"
	repo "awesomeProject/internal/pkg/repository"
	srv "awesomeProject/internal/pkg/service"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	if err := run(); err != nil {
		os.Exit(1)
	}
}

func run() (err error) {
	// ============================Log============================ //
	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)
	file, err := os.OpenFile("app-auth.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Failed to open log file: %s\n", err)
		return err
	}
	defer file.Close()
	log.SetOutput(file)
	// ============================DataBase============================ //
	db, err := pgxpool.Connect(context.Background(), fmt.Sprintf(
		"user=%v password=%v host=%v port=%v dbname=%v",
		"admin",
		"admin",
		"127.0.0.1",
		"5432",
		"agona"))
	if err != nil {
		log.Printf("err in db connection: %v", err)
		return err
	}
	log.Println("connected to postgres")
	defer db.Close()
	if err = db.Ping(context.Background()); err != nil {
		return err
	}
	// ----------------------------Database---------------------------- //
	//
	//
	// ============================Init layers============================ //

	catRepo := repo.NewCatRepository(db, log)
	jwtRepo := repo.NewTokensRepository(db, log)
	jwtService := srv.NewJwtService(catRepo, jwtRepo, log)
	authHandler := handler.NewAuthHandler(jwtService, log)
	// ----------------------------Init layers---------------------------- //
	//
	//
	// ============================Create router============================ //
	r := mux.NewRouter().PathPrefix("/api/v1").Subrouter()
	r.HandleFunc("/sign-in", authHandler.SignIn).Methods(http.MethodPost)
	r.HandleFunc("/refresh", authHandler.RefreshToken).Methods(http.MethodPost)
	r.HandleFunc("/logout", authHandler.Logout).Methods(http.MethodPost)
	server := http.Server{
		Addr:    ":8082",
		Handler: r,
	}
	quit := make(chan os.Signal, 1)   // ToDo: Google
	signal.Notify(quit, os.Interrupt) // ToDo: Google
	go func() {
		if err := server.ListenAndServe(); err != nil {
			fmt.Printf("http server listen err: %v\n", err)
		}
	}()
	sig := <-quit
	print(sig)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err = server.Shutdown(ctx); err != nil {
		err = fmt.Errorf("error on server shutdown: %w", err)
		return err
	}
	print("server shutdown")
	return nil
}
