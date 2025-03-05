package main

import (
	//"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/nastts/calc_go2/internal/agent"
	"github.com/nastts/calc_go2/internal/orchestrator"
	//"github.com/nastts/calc_go2/pkg/calculation"
)

	
func main(){
	pon := os.Getenv("computing_power")
	am, err := strconv.Atoi(pon)
	if err != nil || am < 1{
		am = 2
	}
	for i := 0; i < am; i++{
		go agent.Worker()
	}
	http.HandleFunc("/api/v1/calculate", orchestrator.IDhandler)
	http.HandleFunc("/api/v1/expressions", orchestrator.ExpressionListHandler)
	http.HandleFunc("/api/v1/expressions/{id}", orchestrator.ExpressionListIDHandler)
	http.HandleFunc("localhost/internal/task", orchestrator.GetTaskHandler)
	http.HandleFunc("/internal/task", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			orchestrator.GetTaskHandler(w, r)
		case http.MethodPost:
			orchestrator.ResultTaskHandler(w, r)
		default:
			http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		}
	})
	log.Printf("сервер запущен")
	if err := http.ListenAndServe(":8080", nil); err != nil{
		log.Fatal("ошибка при запуске сервера", err)
	}
}




