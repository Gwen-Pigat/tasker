package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"tasker/dashboard"
	"tasker/initializers"
	"tasker/note"
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
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "OPTIONS", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	r.Get("/db/flush", func(w http.ResponseWriter, r *http.Request) {
		FlushDB(NewWrapper(w, r))
	})

	r.Post("/user/connect", func(w http.ResponseWriter, r *http.Request) {
		user.GetUserConnect(NewWrapper(w, r))
	})

	r.Post("/user", func(w http.ResponseWriter, r *http.Request) {
		user.CreateUser(NewWrapper(w, r))
	})
	r.Group(func(r chi.Router) {
		r.Use(CheckAuth)
		r.Get("/user", func(w http.ResponseWriter, r *http.Request) {
			user.GetUser(NewWrapper(w, r))
		})
		tasksPaths(r)
		notesPaths(r)
		dashbaordPaths(r)
		commonTasksPaths(r)
	})

	fmt.Printf("Server starting on port %s\n", port)
	http.ListenAndServe(":"+port, r)
}

func tasksPaths(r chi.Router) {
	r.Get("/tasks", func(w http.ResponseWriter, r *http.Request) {
		task.GetTasks(NewWrapper(w, r))
	})
	r.Post("/tasks", func(w http.ResponseWriter, r *http.Request) {
		task.CreateTask(NewWrapper(w, r))
	})
	r.Put("/tasks/{id}", func(w http.ResponseWriter, r *http.Request) {
		task.PutTask(NewWrapper(w, r))
	})
	r.Patch("/tasks/{id}", func(w http.ResponseWriter, r *http.Request) {
		task.PatchTask(NewWrapper(w, r))
	})
	r.Delete("/tasks/{id}", func(w http.ResponseWriter, r *http.Request) {
		task.DeleteTask(NewWrapper(w, r))
	})
}

func notesPaths(r chi.Router) {
	r.Get("/notes", func(w http.ResponseWriter, r *http.Request) {
		note.GetNotes(NewWrapper(w, r))
	})
	r.Post("/notes", func(w http.ResponseWriter, r *http.Request) {
		note.CreateNote(NewWrapper(w, r))
	})
	r.Put("/notes/{id}", func(w http.ResponseWriter, r *http.Request) {
		note.PutNote(NewWrapper(w, r))
	})
	r.Delete("/notes/{id}", func(w http.ResponseWriter, r *http.Request) {
		note.DeleteNote(NewWrapper(w, r))
	})
}

func dashbaordPaths(r chi.Router) {
	r.Get("/dashboard/stats", func(w http.ResponseWriter, r *http.Request) {
		dashboard.GetDashboardStats(NewWrapper(w, r))
	})
}

func commonTasksPaths(r chi.Router) {
	r.Get("/common-tasks", func(w http.ResponseWriter, r *http.Request) {
		task.GetCommonTasks(NewWrapper(w, r))
	})
	r.Post("/common-tasks", func(w http.ResponseWriter, r *http.Request) {
		task.CreateCommonTask(NewWrapper(w, r))
	})
	r.Post("/common-tasks/{id}/validate", func(w http.ResponseWriter, r *http.Request) {
		task.ValidateCommonTask(NewWrapper(w, r))
	})
	r.Delete("/common-tasks/{id}", func(w http.ResponseWriter, r *http.Request) {
		task.DeleteCommonTask(NewWrapper(w, r))
	})
}

func FlushDB(wrapper *initializers.Wrapper) {
	if err := initializers.ExecFlushDB(DB); err != nil {
		wrapper.Error(err.Error(), http.StatusInternalServerError)
		return
	}
	wrapper.Render(map[string]any{"message": "DB is flushed"})
}

func CheckAuth(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			NewWrapper(w, r).Error("Authorization header missing", http.StatusUnauthorized)
			return
		}

		wrapper := NewWrapper(w, r)
		wrapper.Data = map[string]any{"token": auth}
		userID, err := user.GetUserAuth(wrapper)
		if err != nil {
			wrapper.Error("Invalid token: "+err.Error(), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user", userID)
		handler.ServeHTTP(w, r.WithContext(ctx))
	})
}

func Handle(r chi.Router, method string, path string, handler func(w *initializers.Wrapper)) {
	r.MethodFunc(method, path, func(w http.ResponseWriter, r *http.Request) {
		wrapper := NewWrapper(w, r)
		handler(wrapper)
	})
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
