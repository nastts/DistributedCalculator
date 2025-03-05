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
	if err != nil{
		http.Error(w, "expression is not valid", http.StatusUnprocessableEntity)
		return
	}
	expr := &structs.ExpressionID{
		ID: id,
		Status: "in progress",
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

func ResultTaskHandler(w http.ResponseWriter, r *http.Request){
	var res structs.Result
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&res)
	if err != nil{
		http.Error(w, "expression is not valid", http.StatusUnprocessableEntity)
		return
	}
	found := false
	for _, expr := range expressions{
		if task.UpdateTask(expr, res.ID, res.Result){
			found = true
			if expr.Root.Computed{
				expr.Status = "done"
				expr.Result = expr.Root.Value
			}
		}
	}
	if !found{
		http.Error(w, "result is not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}
