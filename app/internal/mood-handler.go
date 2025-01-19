package internal

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
)

type MoodHandler struct {
	pg *pgxpool.Pool
}

func (h MoodHandler) Register(writer http.ResponseWriter, request *http.Request) {

}

func (h MoodHandler) Report(writer http.ResponseWriter, request *http.Request) {
	
}

func NewMoodHandler(pg *pgxpool.Pool) MoodHandler {
	return MoodHandler{pg: pg}
}
