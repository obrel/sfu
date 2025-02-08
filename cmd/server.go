package cmd

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/obrel/go-lib/pkg/log"
	"github.com/obrel/sfu/internal/handler"
	"github.com/obrel/sfu/internal/sfu"
	"github.com/spf13/cobra"
)

var (
	serverCmd = &cobra.Command{
		Use: "server",
	}
)

func serverPreRun(cmd *cobra.Command, args []string) error {
	return nil
}

func serverRun(cmd *cobra.Command, args []string) error {
	// The HTTP Server
	server := &http.Server{
		Addr:    "0.0.0.0:4000",
		Handler: handler.NewHandler(sfu.NewSFU()).Service(),
	}

	// Server run context
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sig

		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, _ := context.WithTimeout(serverCtx, 30*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.For("server", "run").Fatal(context.DeadlineExceeded)
			}
		}()

		// Trigger graceful shutdown
		err := server.Shutdown(shutdownCtx)
		if err != nil {
			log.For("server", "run").Fatal(err)
		}

		serverStopCtx()
	}()

	// Run the server
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return err
	}

	// Wait for server context to be stopped
	<-serverCtx.Done()

	return nil
}

func init() {
	serverCmd.PreRunE = serverPreRun
	serverCmd.RunE = serverRun
}
