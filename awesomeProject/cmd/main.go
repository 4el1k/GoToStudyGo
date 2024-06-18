package main

import (
	"awesomeProject/internal/pkg/handler/handler"
	repo "awesomeProject/internal/pkg/repository"
	srv "awesomeProject/internal/pkg/service"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
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
	// ============================DataBase============================ //
	db, err := pgxpool.Connect(context.Background(), fmt.Sprintf(
		"user=%v password=%v host=%v port=%v dbname=%v",
		"admin",
		"admin",
		"127.0.0.1",
		"5432",
		"agona"))
	if err != nil {
		fmt.Printf("err in db connection: %v", err)
		return err
	}
	fmt.Sprintln("connected to postgres")
	defer db.Close()
	if err = db.Ping(context.Background()); err != nil {
		return err
	}
	// ----------------------------Database---------------------------- //
	//
	//
	// ============================Init layers============================ //
	catRepo := repo.NewCatRepository(db)
	catSrv := srv.NewCatService(*catRepo) // да почему нельзя без *
	catHandler := handler.NewCatHandler(*catSrv)
	// ----------------------------Init layers---------------------------- //
	//
	//
	// ============================Create router============================ //
	r := mux.NewRouter().PathPrefix("/api/v1").Subrouter()
	r.HandleFunc("/cat", catHandler.Create).Methods(http.MethodPost)
	server := http.Server{
		Addr:    ":8081",
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
