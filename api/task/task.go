package task

import (
	"net/http"
	"strings"
	"tasker/initializers"
	"time"

	"github.com/go-chi/chi/v5"
)

type Task struct {
	ID       int     `json:"id"`
	Title    string  `json:"title"`
	DateAdd  *string `json:"dateAdd"`
	DateTo   *string `json:"dateTo"`
	Content  *string `json:"content"`
	IsDone   bool    `json:"isDone"`
	IsCommon bool    `json:"isCommon"`
	RefUser  int     `json:"refUser"`
}

type CommonTask struct {
	ID      int     `json:"id"`
	Title   string  `json:"title"`
	DateAdd *string `json:"dateAdd"`
	RefUser int     `json:"refUser"`
}

var taskSetup = map[string]string{
	"payload": "id,date_add,date_to,title,content,is_done,is_common,ref_user",
	"table":   "task",
}

func stringPtr(s string) *string {
	return &s
}

func scanTask(row interface{ Scan(...any) error }) (Task, error) {
	var t Task
	err := row.Scan(&t.ID, &t.DateAdd, &t.DateTo, &t.Title, &t.Content, &t.IsDone, &t.IsCommon, &t.RefUser)
	return t, err
}

func CreateTask(wrapper *initializers.Wrapper) {
	var payload struct {
		Title string `json:"title"`
	}
	if err := wrapper.ParseJSON(&payload); err != nil {
		wrapper.Error(err.Error(), http.StatusBadRequest)
		return
	}

	task := Task{
		Title:   payload.Title,
		IsDone:  false,
		RefUser: wrapper.ReturnUser(),
		DateAdd: stringPtr(time.Now().UTC().Format(initializers.Format)),
	}

	_, err := initializers.DB.Exec(
		"INSERT INTO "+taskSetup["table"]+" (title, date_add, is_done, ref_user) VALUES (?,?,?,?)",
		task.Title, task.DateAdd, task.IsDone, task.RefUser,
	)
	if err != nil {
		wrapper.Error(err.Error(), http.StatusInternalServerError)
		return
	}
	GetTasks(wrapper)
}

func PutTask(wrapper *initializers.Wrapper) {
	var payload struct {
		Title   string  `json:"title"`
		DateAdd string  `json:"dateAdd"`
		DateTo  *string `json:"dateTo"`
	}
	if err := wrapper.ParseJSON(&payload); err != nil {
		wrapper.Error(err.Error(), http.StatusBadRequest)
		return
	}

	id := chi.URLParam(wrapper.Request, "id")
	userID := wrapper.ReturnUser()

	// Verify ownership
	var exists bool
	err := initializers.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM "+taskSetup["table"]+" WHERE id=? AND ref_user=?)", id, userID).Scan(&exists)
	if err != nil || !exists {
		wrapper.Error("Task not found or access denied", http.StatusNotFound)
		return
	}

	isDone := payload.DateTo != nil && *payload.DateTo != ""

	parseInputDate := func(dateStr string) string {
		if dateStr == "" {
			return ""
		}
		if t, err := time.Parse(time.RFC3339, dateStr); err == nil {
			return t.UTC().Format(initializers.Format)
		}
		if t, err := time.ParseInLocation(initializers.Format, dateStr, initializers.LocParis); err == nil {
			return t.UTC().Format(initializers.Format)
		}
		return strings.ReplaceAll(dateStr, "T", " ")
	}

	dateAdd := parseInputDate(payload.DateAdd)
	var dateTo *string
	if payload.DateTo != nil && *payload.DateTo != "" {
		dt := parseInputDate(*payload.DateTo)
		dateTo = &dt
	}

	_, err = initializers.DB.Exec(
		"UPDATE "+taskSetup["table"]+" SET title=?, date_add=?, is_done=?, date_to=? WHERE id=? AND ref_user=?",
		payload.Title, dateAdd, isDone, dateTo, id, userID,
	)
	if err != nil {
		wrapper.Error(err.Error(), http.StatusInternalServerError)
		return
	}

	wrapper.Render(map[string]any{"message": "Task updated successfully"})
}

func GetTasks(wrapper *initializers.Wrapper) {
	date := wrapper.Request.URL.Query().Get("date")
	query := "SELECT " + taskSetup["payload"] + " FROM " + taskSetup["table"] + " WHERE ref_user=?"
	args := []any{wrapper.ReturnUser()}

	if date != "" {
		query += " AND DATE(date_add) = ?"
		args = append(args, date)
	}

	query += " ORDER BY date_add DESC LIMIT 50"

	rows, err := initializers.DB.Query(query, args...)
	if err != nil {
		wrapper.Error(err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	tasks := []Task{}
	for rows.Next() {
		task, err := scanTask(rows)
		if err != nil {
			wrapper.Error(err.Error(), http.StatusInternalServerError)
			return
		}
		if task.DateTo != nil {
			*task.DateTo, _ = wrapFormat(task.DateTo)
		}
		*task.DateAdd, _ = wrapFormat(task.DateAdd)
		tasks = append(tasks, task)
	}
	wrapper.Render(map[string]any{"data": tasks})
}

func GetTask(wrapper *initializers.Wrapper) {
	id := chi.URLParam(wrapper.Request, "id")
	row := initializers.DB.QueryRow("SELECT "+taskSetup["payload"]+" FROM "+taskSetup["table"]+" WHERE id=? AND ref_user=?", id, wrapper.ReturnUser())

	task, err := scanTask(row)
	if err != nil {
		wrapper.Error("Task not found", http.StatusNotFound)
		return
	}
	wrapper.Render(map[string]any{"data": task})
}

func PatchTask(wrapper *initializers.Wrapper) {
	id := chi.URLParam(wrapper.Request, "id")
	userID := wrapper.ReturnUser()

	row := initializers.DB.QueryRow("SELECT is_done FROM "+taskSetup["table"]+" WHERE id=? AND ref_user=?", id, userID)
	var isDone bool
	if err := row.Scan(&isDone); err != nil {
		wrapper.Error("Task not found", http.StatusNotFound)
		return
	}

	newIsDone := !isDone
	var dateTo *string
	if newIsDone {
		now := time.Now().UTC().Format(initializers.Format)
		dateTo = &now
	}

	_, err := initializers.DB.Exec(
		"UPDATE "+taskSetup["table"]+" SET is_done = ?, date_to=? WHERE id=? AND ref_user=?",
		newIsDone, dateTo, id, userID,
	)
	if err != nil {
		wrapper.Error(err.Error(), http.StatusInternalServerError)
		return
	}

	wrapper.Render(map[string]any{"message": "Status updated"})
}

func wrapFormat(dateStr *string) (string, error) {
	if dateStr == nil || *dateStr == "" {
		return "", nil
	}
	parsed, err := time.ParseInLocation(initializers.Format, *dateStr, time.UTC)
	if err != nil {
		return *dateStr, err
	}
	return parsed.In(initializers.LocParis).Format(initializers.Format), nil
}

func DeleteTask(wrapper *initializers.Wrapper) {
	res, err := initializers.DB.Exec(
		"DELETE FROM "+taskSetup["table"]+" WHERE id=? AND ref_user=?",
		chi.URLParam(wrapper.Request, "id"), wrapper.ReturnUser(),
	)
	if err != nil {
		wrapper.Error(err.Error(), http.StatusInternalServerError)
		return
	}
	affected, _ := res.RowsAffected()
	if affected == 0 {
		wrapper.Error("Task not found", http.StatusNotFound)
		return
	}
	wrapper.Render(map[string]any{"message": "Task deleted"})
}
