/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

var (
	port      int
	serverCmd = &cobra.Command{
		Use:   "server",
		Short: "Start the HTTP server",
		Long:  `Start the HTTP server on the specified port`,
		Run: func(cmd *cobra.Command, args []string) {
			startServer()
		},
	}
)

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().IntVarP(&port, "port", "p", 8080, "Port to run the server on")
}

func startServer() {
	addr := fmt.Sprintf(":%d", port)

	// Create request handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request to %s", r.URL.Path)
		w.Header().Set("Content-Type", "text/plain")
		welcomeMsg := fmt.Sprintf("Welcome to obot server!\nVersion: %s", appVersion)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(welcomeMsg))
	})

	log.Printf("Starting server on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
