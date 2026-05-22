package task

import (
	"net/http"
	"tasker/initializers"

	"github.com/go-chi/chi/v5"
)

func GetCommonTasks(wrapper *initializers.Wrapper) {
	rows, err := initializers.DB.Query(
		"SELECT id, title, date_add, ref_user FROM common_task WHERE ref_user=? ORDER BY date_add DESC",
		wrapper.ReturnUser(),
	)
	if err != nil {
		wrapper.Error(err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	commonTasks := []CommonTask{}
	for rows.Next() {
		var t CommonTask
		if err := rows.Scan(&t.ID, &t.Title, &t.DateAdd, &t.RefUser); err != nil {
			wrapper.Error(err.Error(), http.StatusInternalServerError)
			return
		}
		commonTasks = append(commonTasks, t)
	}
	wrapper.Render(map[string]any{"data": commonTasks})
}

func CreateCommonTask(wrapper *initializers.Wrapper) {
	var payload struct {
		Title string `json:"title"`
	}
	if err := wrapper.ParseJSON(&payload); err != nil {
		wrapper.Error(err.Error(), http.StatusBadRequest)
		return
	}

	dateAdd := getCombinedDate(wrapper.Request.URL.Query().Get("date"))
	userID := wrapper.ReturnUser()

	_, err := initializers.DB.Exec(
		"INSERT INTO common_task (title, date_add, ref_user) VALUES (?,?,?)",
		payload.Title, dateAdd, userID,
	)
	if err != nil {
		wrapper.Error(err.Error(), http.StatusInternalServerError)
		return
	}
	GetCommonTasks(wrapper)
}

func ValidateCommonTask(wrapper *initializers.Wrapper) {
	id := chi.URLParam(wrapper.Request, "id")
	userID := wrapper.ReturnUser()

	var title string
	err := initializers.DB.QueryRow("SELECT title FROM common_task WHERE id=? AND ref_user=?", id, userID).Scan(&title)
	if err != nil {
		wrapper.Error("Common task not found", http.StatusNotFound)
		return
	}

	// Insert into regular tasks table as a pending task for today
	dateAdd := getCombinedDate(wrapper.Request.URL.Query().Get("date"))
	_, err = initializers.DB.Exec(
		"INSERT INTO task (title, date_add, date_to, is_done, is_common, ref_user) VALUES (?,?,?,?,?,?)",
		title, dateAdd, nil, false, true, userID,
	)
	if err != nil {
		wrapper.Error(err.Error(), http.StatusInternalServerError)
		return
	}

	wrapper.Render(map[string]any{"message": "Common task validated and added to today's tasks"})
}

func DeleteCommonTask(wrapper *initializers.Wrapper) {
	id := chi.URLParam(wrapper.Request, "id")
	userID := wrapper.ReturnUser()

	_, err := initializers.DB.Exec("DELETE FROM common_task WHERE id=? AND ref_user=?", id, userID)
	if err != nil {
		wrapper.Error(err.Error(), http.StatusInternalServerError)
		return
	}
	GetCommonTasks(wrapper)
}
