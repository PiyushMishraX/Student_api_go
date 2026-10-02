package main // package declaration at start of file
import (
	"context"
	// "fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/piyushmishrax/Student_api_go/internal/config"
)

// import "fmt"
// main is entry point

func main() {
	// fmt.Println("Welcome to students api")

	// lead config
	//            // can set coustom logger , or use inbuilt one
	// atabase setup
	// setup router
	// setup server

	// load config
	cfg := config.MustLoad()

	// router 
	// // http pacage in go > 1.22 > we use them http // we can metntion method, get , post  and use parameter etc
	// inbuilt pacakage can be used for creating server and route setup

	router := http.NewServeMux()

	router.HandleFunc("GET /", func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to student api"))


	} )


	// setup server

	server := http.Server {
		// Addr: cfg.HTTPServer.Addr,
		Addr: cfg.Addr,
		Handler: router,
	}

	// fmt.Println("Server started")
	// fmt.Printf("Server started %s", cfg.HTTPServer.Addr)
	// fmt.Printf("Server started %s", cfg.Addr)
	slog.Info("Server started", slog.String("address", cfg.Addr) )


	done := make(chan os.Signal, 1) // go routine runs concurrently so we have to stop it ( using channel block )

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM) // when inturrept signal comes notify channel // learn the calls from google

	go func () {
		err := server.ListenAndServe() // start  // blocking
		if err != nil {
			log.Fatalf("failed to start server")
		}
	} ()

	<-done // while process not done we are blocked ( while done did not get signal again )

	// err := server.ListenAndServe() // start  // blocking
	// if err != nil {
	// 	log.Fatalf("failed to start server")
	// } // prouduction aren't such simple adding graceful stop 
	// to stop ongoing request from not stop while shutdown , it is required in production 
	// creating seperate goroutine and channel
	

	slog.Info("shutting down the server ")

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	

	// server.Shutdown() it sometimes infinitly hangs and do not let the aserver shutdown  /// so we notify when it didn't work 
	// err := server.Shutdown(ctx)
	// if err != nil {
	// 	slog.Error("failed to shutdown server", slog.String("error", err.Error()))
	// }

	if err := server.Shutdown(ctx);  err != nil {
		slog.Error("failed to shutdown server", slog.String("error", err.Error())) // sending context to stop infinite hang
	}

	slog.Info("server shutdown successfully") // understand context in go like js 

	


	


}




// without flag appllication do not run becuase "Must"