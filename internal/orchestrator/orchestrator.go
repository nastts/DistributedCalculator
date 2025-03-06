package orchestrator

import (
	"encoding/json"
	"net/http"
	//"sync"
	

	"github.com/nastts/calc_go2/pkg/calculation"
	"github.com/nastts/calc_go2/internal/structs"
	"github.com/nastts/calc_go2/internal/task"
)


var(
	
	
	//m sync.Mutex
	expressions = make(map[string]*structs.ExpressionID)
	//tasks = make([]*structs.Task, 0)
	
)





func  IDhandler(w http.ResponseWriter, r *http.Request){
	var req structs.Request
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil{
		http.Error(w, "expression is not valid", http.StatusUnprocessableEntity)
		return
	}
	id := structs.GetID()
	res, err := calculation.Calc(req.Expression)
	//root, err := calculation.ParseExpression(req.Expression)
	if err != nil{
		http.Error(w, "expression is not valid", http.StatusUnprocessableEntity)
		return
	}
	expr := &structs.ExpressionID{
		ID: id,
		Status: "in process",
		Result: res,
	}
	expressions[id] = expr
	task.GetTask(expr.Root, expr.ID)
	if len(structs.TasksQueue) >0 {
		expr.Status = "done"
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(structs.IDDD{ID: id})
}



func  ExpressionListHandler(w http.ResponseWriter, r *http.Request){
	defer r.Body.Close()
	if len(expressions) == 0{
		http.Error(w, "expressions is not found", http.StatusNotFound)
		return
	}
	list := make([]*structs.ExpressionID, 0, len(expressions))
	for _, ex := range expressions{
		list = append(list, ex)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string][]*structs.ExpressionID{"expressions":list})
}






func ExpressionListIDHandler(w http.ResponseWriter, r *http.Request){
	defer r.Body.Close()
	id := r.PathValue("id")
	main, found := expressions[id]
	if !found{
		http.Error(w, "id not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]*structs.ExpressionID{"expression":main})

}


func GetTaskHandler(w http.ResponseWriter, r *http.Request){
	defer r.Body.Close()
	if len(structs.TasksQueue) == 0{
		http.Error(w, "tasks is not found", http.StatusNotFound)
		return
	}
	task := structs.TasksQueue[0]
	structs.TasksQueue = structs.TasksQueue[1:]
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]*structs.Task{"task":task})
}


func ResultTaskHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var result structs.Result
	err := json.NewDecoder(r.Body).Decode(&result)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusUnprocessableEntity)
		return
	}

	expr, found := expressions[result.ID]
	if !found {
		http.Error(w, "expression ID not found", http.StatusNotFound)
		return
	}

	updated := task.UpdateTask(expr, result.ID, result.Result)
	if !updated {
		http.Error(w, "task ID not found or already computed", http.StatusUnprocessableEntity)
		return
	}

	if expr.Root.Computed {
		expr.Status = "done"
	}

	task.SolveTask(expr.Root, expr.ID)
	if len(structs.TasksQueue) > 0 {
		expr.Status = "in progress"
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": expr.Status})
}
