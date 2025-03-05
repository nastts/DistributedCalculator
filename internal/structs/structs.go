package structs

import (
	"strconv"
	"sync"
)


type IDDD struct{
	ID string `json:"id"`
}
type ExpressionID struct{
	ID string `json:"id"`
	Status string `json:"status"`
	Result float64 `json:"result"`
	Root *Node `json:"-"`
}

type Request struct{
	Expression string `json:"expression"`
}
type Error struct{
	Error string `json:"error"`
}
type ExpressionList struct{
	Expressions []ExpressionID `json:"expressions"`
}

type Task struct{
	ID string `json:"id"`
	//ExpressionID string `json:"ExprID"`
	Arg1 float64 `json:"arg1"`
	Arg2 float64 `json:"arg2"`
	Operation string `json:"operation"`
	OperationTime int `json:"operationTime"`
	Root *Node `json:"-"`
}

type Result struct{
	ID string  `json:"id"`
	Result float64 `json:"result"`
}

type Node struct {
	Operator string  
	Value    float64 
	Left     *Node
	Right    *Node
	Computed bool  
	TaskID   string
	
}

type Times struct {
	TimeAdditionMs int `json:"time_addition_ms"`
	TimeSubtractionMs int `json:"time_subtraction_ms"`
	TimeMultiplicationsMs int `json:"time_multiplications_ms"`
	TimeDivisionsMs int `json:"time_divisions_ms"`
}

var(
	TasksQueue  = make([]*Task, 0)
	m sync.Mutex
	idCount = 0
)

func GetID()string{
	m.Lock()
	defer m.Unlock()
	idCount +=1
	return strconv.Itoa(idCount)
}