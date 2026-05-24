package note

import (
	"net/http"
	"strconv"
	"tasker/initializers"
	"time"

	"github.com/go-chi/chi/v5"
)

type Note struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	DateAdd    string `json:"dateAdd"`
	DateUpdate string `json:"dateUpdate"`
	RefUser    int    `json:"refUser"`
}

var noteSetup = map[string]string{
	"payload": "id,date_add,COALESCE(date_update, ''),title,content,ref_user",
	"table":   "note",
}

func scanNote(row interface{ Scan(...any) error }) (Note, error) {
	var t Note
	err := row.Scan(&t.ID, &t.DateAdd, &t.DateUpdate, &t.Title, &t.Content, &t.RefUser)
	return t, err
}

func GetNotes(wrapper *initializers.Wrapper) {
	query := "SELECT " + noteSetup["payload"] + " FROM " + noteSetup["table"] + " WHERE ref_user=? ORDER BY date_add DESC"
	rows, err := initializers.DB.Query(query, wrapper.ReturnUser())
	if err != nil {
		wrapper.Error(err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	notes := []Note{}
	for rows.Next() {
		var n Note
		err := rows.Scan(&n.ID, &n.DateAdd, &n.DateUpdate, &n.Title, &n.Content, &n.RefUser)
		if err != nil {
			wrapper.Error(err.Error(), http.StatusInternalServerError)
			return
		}
		notes = append(notes, n)
	}
	wrapper.Render(map[string]any{"data": notes})
}

func GetNote(wrapper *initializers.Wrapper) {
	id := chi.URLParam(wrapper.Request, "id")
	row := initializers.DB.QueryRow("SELECT "+noteSetup["payload"]+" FROM "+noteSetup["table"]+" WHERE id=? AND ref_user=?", id, wrapper.ReturnUser())

	note, err := scanNote(row)
	if err != nil {
		wrapper.Error("Note not found", http.StatusNotFound)
		return
	}
	wrapper.Render(map[string]any{"data": note})
}

func CreateNote(wrapper *initializers.Wrapper) {
	var body struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	if err := wrapper.ParseJSON(&body); err != nil {
		wrapper.Error("Invalid request body : "+err.Error(), http.StatusBadRequest)
		return
	}

	if body.Title == "" || body.Content == "" {
		wrapper.Error("Title and content are required", http.StatusBadRequest)
		return
	}

	dateAdd := time.Now().UTC().Format(initializers.Format)

	query := "INSERT INTO " + noteSetup["table"] + " (ref_user,date_add,title,content) VALUES (?, ?, ?, ?)"

	r, err := initializers.DB.Exec(query, wrapper.ReturnUser(), dateAdd, body.Title, body.Content)
	if err != nil {
		wrapper.Error("Invalid request body : "+err.Error(), http.StatusInternalServerError)
		return
	}
	insertedID, err := r.LastInsertId()
	if err != nil {
		wrapper.Error("Invalid request body : "+err.Error(), http.StatusInternalServerError)
		return
	}

	wrapper.Render(map[string]any{"data": Note{
		ID:      int(insertedID),
		DateAdd: dateAdd,
		Title:   body.Title,
		Content: body.Content,
		RefUser: wrapper.ReturnUser(),
	}}, http.StatusCreated)
}

func PutNote(wrapper *initializers.Wrapper) {
	idStr := chi.URLParam(wrapper.Request, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		wrapper.Error("Invalid ID format", http.StatusBadRequest)
		return
	}

	row := initializers.DB.QueryRow("SELECT "+noteSetup["payload"]+" FROM "+noteSetup["table"]+" WHERE id=? AND ref_user=?", id, wrapper.ReturnUser())

	note, err := scanNote(row)
	if err != nil {
		wrapper.Error("Note not found + "+err.Error(), http.StatusNotFound)
		return
	}

	var payload struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	if err := wrapper.ParseJSON(&payload); err != nil {
		wrapper.Error(err.Error(), http.StatusBadRequest)
		return
	}

	if payload.Title == "" {
		wrapper.Error("Title is required", http.StatusBadRequest)
		return
	}
	if payload.Content == "" {
		wrapper.Error("Content is required", http.StatusBadRequest)
		return
	}

	dateUpdate := time.Now().UTC().Format(initializers.Format)

	query := "UPDATE " + noteSetup["table"] + " SET title=?, date_update=?, content=? WHERE id=? AND ref_user=?"

	_, err = initializers.DB.Exec(query, payload.Title, dateUpdate, payload.Content, id, wrapper.ReturnUser())
	if err != nil {
		wrapper.Error("Invalid request body : "+err.Error(), http.StatusInternalServerError)
		return
	}

	wrapper.Render(map[string]any{"data": Note{
		ID:         id,
		DateAdd:    note.DateAdd,
		DateUpdate: dateUpdate,
		Title:      payload.Title,
		Content:    payload.Content,
		RefUser:    wrapper.ReturnUser(),
	}})
}

func DeleteNote(wrapper *initializers.Wrapper) {
	idStr := chi.URLParam(wrapper.Request, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		wrapper.Error("Invalid ID format", http.StatusBadRequest)
		return
	}

	query := "DELETE FROM " + noteSetup["table"] + " WHERE id=? AND ref_user=?"

	_, err = initializers.DB.Exec(query, id, wrapper.ReturnUser())
	if err != nil {
		wrapper.Error("Invalid request body : "+err.Error(), http.StatusInternalServerError)
		return
	}

	wrapper.Render(map[string]any{"data": Note{
		ID: id,
	}})
}
