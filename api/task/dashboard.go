package task

import (
	"net/http"
	"tasker/initializers"
	"time"
)

func GetDashboardStats(wrapper *initializers.Wrapper) {
	userID := wrapper.ReturnUser()
	lastMonth := time.Now().UTC().AddDate(0, -1, 0).Format(initializers.Format)

	// Tasks added in the last month
	var addedCount int
	err := initializers.DB.QueryRow(
		"SELECT COUNT(*) FROM task WHERE ref_user=? AND date_add >= ?",
		userID, lastMonth,
	).Scan(&addedCount)
	if err != nil {
		wrapper.Error(err.Error(), http.StatusInternalServerError)
		return
	}

	// Tasks done in the last month
	var doneCount int
	err = initializers.DB.QueryRow(
		"SELECT COUNT(*) FROM task WHERE ref_user=? AND is_done=1 AND date_to >= ?",
		userID, lastMonth,
	).Scan(&doneCount)
	if err != nil {
		wrapper.Error(err.Error(), http.StatusInternalServerError)
		return
	}

	// Average duration in minutes for tasks done in the last month
	// MySQL: TIMESTAMPDIFF(MINUTE, date_add, date_to)
	var avgDuration float64
	err = initializers.DB.QueryRow(
		"SELECT COALESCE(AVG(TIMESTAMPDIFF(MINUTE, date_add, date_to)), 0) FROM task WHERE ref_user=? AND is_done=1 AND date_to >= ?",
		userID, lastMonth,
	).Scan(&avgDuration)
	if err != nil {
		wrapper.Error(err.Error(), http.StatusInternalServerError)
		return
	}

	wrapper.Render(map[string]any{
		"data": map[string]any{
			"tasksAdded":   addedCount,
			"tasksDone":    doneCount,
			"avgDuration": avgDuration,
		},
	})
}
