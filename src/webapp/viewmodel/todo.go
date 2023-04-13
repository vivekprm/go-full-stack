package viewmodel

type Todo struct {
	Title string
	Active string
}

func NewTodo() Todo {
	return Todo {
		Title: "Todo - New",
		Active: "todo",
	}
}