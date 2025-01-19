package internal

import "net/http"

func NewServer(moodHandler *MoodHandler) http.Handler {
	mux := http.NewServeMux()

	addRoutes(mux, moodHandler)
	var handler http.Handler = mux
	return handler
}

func addRoutes(mux *http.ServeMux, moodHandler *MoodHandler) {
	mux.HandleFunc("GET /healthz", handleHealth)
	mux.HandleFunc("POST /api/mood/register", moodHandler.Register)
	mux.HandleFunc("GET /api/mood/report?", moodHandler.Report)
	//mux.HandleFunc("GET /api/challenge/{challengeId}/data/history", moodHandler.HandleGetHistory)
	//mux.HandleFunc("GET /api/challenge/{challengeId}/data/daily", moodHandler.HandleGetDailyData)
	//mux.HandleFunc("POST /api/challenge/add", moodHandler.HandleAddChallenge)
	//mux.HandleFunc("POST /api/challenge/{challengeId}/data/add", moodHandler.HandleAddDailyData)
}

func handleHealth(writer http.ResponseWriter, _ *http.Request) {
	writer.WriteHeader(http.StatusAccepted)
	if _, err := writer.Write([]byte("healthy")); err != nil {
		return
	}
}
