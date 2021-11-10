package utils

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"os/signal"
)
// StartServer func for starting a simple server.
func StartServer(a *fiber.App) {
	fiberConnUrl, _ := URIBuilder("fiber")

	if err := a.Listen(fiberConnUrl); err != nil {
		log.Printf("opps!..Server not running! Reason %v", err)
	}
}

// StartServerWithGracefulShutdown function for starting server with a graceful shutdown.
func StartServerWithGracefulShutdown(a *fiber.App) {
	allIdleConns := make(chan struct{})

	// Goroutine to monitor idle conns and close them
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		if err := a.Shutdown(); err != nil {
			log.Printf("Oops... Server is not shutting down! Reason: %v", err)
		}
		close(allIdleConns)
	}()

	StartServer(a)
	<-allIdleConns

}

