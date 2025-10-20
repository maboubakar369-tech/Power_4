package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	// Créer un serveur HTTP avec options
	server := &http.Server{
		Addr:    ":8080",
		Handler: setupRouter(), // fonction qui retourne http.Handler avec routes + statiques
	}

	// Démarrer le serveur dans une goroutine pour pouvoir écouter l’arrêt proprement
	go func() {
		log.Printf("✅ Serveur démarré sur http://localhost%s\n", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Erreur serveur : %v", err)
		}
	}()

	// Capturer le signal d’arrêt (Ctrl+C) pour shutdown propre
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop // on attend le signal d’arrêt

	log.Println("🛑 Arrêt du serveur en cours...")

	// Timeout de 5 secondes pour shutdown propre
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Erreur lors de l'arrêt du serveur : %v", err)
	}

	log.Println("✅ Serveur arrêté proprement")
}

// setupRouter configure les routes et la gestion des fichiers statiques
func setupRouter() http.Handler {
	mux := http.NewServeMux()

	// Servir les fichiers statiques
	fileServer := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	// Routes principales
	registerRoutes(mux)

	return mux
}
