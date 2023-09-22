package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"stock-simulation/pkg/handlers"
	"stock-simulation/pkg/handlers/api_client_handler"
	"stock-simulation/pkg/handlers/user_handler"
	"stock-simulation/pkg/repository"
	service "stock-simulation/pkg/services"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func InitializeLoginHandler(db *sql.DB) http.HandlerFunc {
	userRepository := repository.NewUserRepository(db)
	tokenService := service.NewTokenService("Nojsdfpajsp[jaspjopanopoasd")

	h := user_handler.Handler{
		UserRepository: userRepository,
		TokenService:   tokenService,
	}
	return h.LoginHandler
}

func InitializeRegisterHandler(db *sql.DB) http.HandlerFunc {
	userRepository := repository.NewUserRepository(db)
	tokenService := service.NewTokenService("Nojsdfpajsp[jaspjopanopoasd")

	h := user_handler.Handler{
		UserRepository: userRepository,
		TokenService:   tokenService,
	}
	return h.RegisterHandler
}

func HandleWebsocket(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}

	handler := api_client_handler.Handler{
		Upgrader: upgrader,
	}

	handler.HandleConnection(w, r) // Tu wywołujesz obsługę połączenia
}

// Routes sets up routes for the server
func Routes() {
	db, err := ConnectToDatabase()
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return
	}
	// nie używaj defer db.Close() tutaj, bo połączenie zostanie zamknięte zaraz po zakończeniu funkcji

	http.HandleFunc("/websocket", HandleWebsocket)
	http.HandleFunc("/login", InitializeLoginHandler(db))
	http.HandleFunc("/register", InitializeRegisterHandler(db))
	// Your regular HTTP routes
	http.HandleFunc("/your/route", handlers.YourHandler)
}

func ConnectToDatabase() (*sql.DB, error) {
	db, err := sql.Open("mysql", ":@tcp(:3377)/")
	if err != nil {
		return nil, err
	}

	// Sprawdź połączenie z bazą danych
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected to the database!")
	return db, nil
}
