package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"tasker/initializers"
	"tasker/task"
	"tasker/user"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func init() {
	godotenv.Load()
	var err error
	DB, err = initializers.ConnectDB()
	if err != nil {
		panic(err)
	}
	initializers.DB = DB
	fmt.Println(DB)
}

func main() {
	port := "3000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{os.Getenv("API_URL_PREPROD"), os.Getenv("API_URL_PROD")},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "OPTIONS", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-type", "X-CSRF-Token"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	Handle(r, http.MethodGet, "/db/flush", FlushDB)
	Handle(r, http.MethodPost, "/user/connect", user.GetUserConnect)
	Handle(r, http.MethodPost, "/user", user.CreateUser)
	r.Group(func(r chi.Router) {
		r.Use(CheckAuth)
		Handle(r, http.MethodGet, "/user", user.GetUser)
		Handle(r, http.MethodGet, "/tasks", task.GetTasks)
		Handle(r, http.MethodPost, "/tasks", task.CreateTask)
		Handle(r, http.MethodPatch, "/tasks/{id}", task.PatchTask)
		Handle(r, http.MethodDelete, "/tasks/{id}", task.DeleteTask)
	})
	http.ListenAndServe(":"+port, r)
}

func FlushDB(wrapper *initializers.Wrapper) {
	err := initializers.ExecFlushDB(DB)
	if err != nil {
		wrapper.Error(err.Error())
		return
	}
	wrapper.Render(map[string]any{
		"message": "DB is flushed",
	})
}

func CheckAuth(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			NewWrapper(w, r).Error("Not authorized", http.StatusUnauthorized)
			return
		}
		wrapper := NewWrapper(w, r)
		wrapper.Data = make(map[string]any)
		wrapper.Data["token"] = auth
		userID, err := user.GetUserAuth(wrapper)
		if err != nil {
			NewWrapper(w, r).Error("Not authorized : "+err.Error(), http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(wrapper.Request.Context(), "user", userID)
		handler.ServeHTTP(w, r.WithContext(ctx))
	})
}

func Handle(r chi.Router, method string, path string, handler func(w *initializers.Wrapper)) {
	r.MethodFunc(method, path, func(w http.ResponseWriter, r *http.Request) {
		wrapper := NewWrapper(w, r)
		if method == http.MethodPost {
			errorMsg, errorCode := wrapper.HandlePOST(wrapper.Request)
			if errorMsg != "" {
				wrapper.Error(errorMsg, errorCode)
				return
			}
		}
		handler(wrapper)
	})
}

func HandleGET(r *http.Request) (errorMSG string, errorCode int) {
	if r.Method != http.MethodGet {
		return "Not authorized", http.StatusMethodNotAllowed
	}
	return "", 0
}

func NewWrapper(w http.ResponseWriter, r *http.Request) *initializers.Wrapper {
	return &initializers.Wrapper{
		Writer:  w,
		Request: r,
	}
}

func Index(wrapper *initializers.Wrapper) {
	wrapper.Render(map[string]any{
		"message": "Hello world",
	})
}
