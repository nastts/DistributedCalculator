package task

import(
	"github.com/nastts/calc_go2/internal/structs"
	
	//"github.com/nastts/calc_go2/internal/orchestrator"
)

func GetTask(node *structs.Node, exprID string){
	if node == nil{
		return
	}
	if node.Operator != ""{
		GetTask(node.Left,exprID)
		GetTask(node.Right,exprID)
		if node.Left != nil && node.Right != nil && node.Left.Computed && node.Right.Computed && !node.Computed{
			task := &structs.Task{
				ID: structs.GetID(),
				Arg1: node.Left.Value,
				Arg2: node.Right.Value,
				Operation: node.Operator,
				OperationTime: 1000,
			}
			node.TaskID = task.ID
			structs.TasksQueue = append(structs.TasksQueue, task)
		}
	}
}

func SolveTask(node *structs.Node, exprID string){
	if node == nil{
		return
	}
	if node.Operator != "" && !node.Computed{
		if node.Operator != "" && !node.Computed && node.Left != nil && node.Right != nil && node.Left.Computed && node.Right.Computed{
			task := &structs.Task{
				ID: structs.GetID(),
				Arg1: node.Left.Value,
				Arg2: node.Right.Value,
				Operation: node.Operator,
				OperationTime: 1000,
			}
				structs.TasksQueue = append(structs.TasksQueue, task)
			}
		}
		SolveTask(node.Left,exprID )
		SolveTask(node.Right,exprID)
	}

func UpdateTask(expr *structs.ExpressionID, taskID string, result float64) bool {
	found := false
	var traverse func(node *structs.Node)
	traverse = func(node *structs.Node) {
		if node == nil || found {
			return
		}
		if node.TaskID == taskID && !node.Computed {
			node.Value = result
			node.Computed = true
			found = true
			node.TaskID = ""
		}
		traverse(node.Left)
		traverse(node.Right)
	}
	traverse(expr.Root)
	return found
}