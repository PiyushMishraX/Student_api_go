package main // package declaration at start of file
import (
	"fmt"
	"log"
	"net/http"

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
		Addr: cfg.Addr,
		Handler: router,
	}

	// fmt.Println("Server started")
	fmt.Printf("Server started %s", cfg.HTTPServer.Addr)



	go func () {
		err := server.ListenAndServe() // start  // blocking
		if err != nil {
			log.Fatalf("failed to start server")
		}
	} ()

	// err := server.ListenAndServe() // start  // blocking
	// if err != nil {
	// 	log.Fatalf("failed to start server")
	// } // prouduction aren't such simple adding graceful stop 
	// to stop ongoing request from not stop while shutdown , it is required in production 
	// creating seperate goroutine and channel
	
	


	


}




// without flag appllication do not run becuase "Must"