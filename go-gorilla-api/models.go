package main

type Expression struct {
	ID     int     `json:"id"`
	Status string  `json:"status"`
	Result float64 `json:"result"`
}

var expressions = []Expression{
	{ID: 1, Status: "Active", Result: 0},
	{ID: 2, Status: "Done", Result: 5},
}

type Data struct {
	ID         int    `json:"id"`
	Expression string `json:"expressions"`
}

type Task struct {
	ID             int     `json:"id"`             // <идентификатор задачи>,
	Arg1           int     `json:"arg1"`           //: <имя первого аргумента>,
	Arg2           int     `json:"arg2"`           //: <имя второго аргумента>,
	Operation      string  `json:"operation"`      //: <операция>,
	Operation_time float64 `json:"operation_time"` //: <время выполнения операции>
}

var data = []Data{
	{ID: 1, Expression: "2+2*2"},
	{ID: 2, Expression: "2*2+3"},
}

var task = []Task{
	{ID: 1, Arg1: 2, Arg2: 2, Operation: "+", Operation_time: 2},
	{ID: 2, Arg1: 2, Arg2: 2, Operation: "*", Operation_time: 3},
}
