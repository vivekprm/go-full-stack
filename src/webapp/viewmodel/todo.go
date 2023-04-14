package viewmodel

type TodoPage struct {
	Title string
	Active string
}

func NewTodo() TodoPage {
	return TodoPage {
		Title: "Todo - New",
		Active: "todo",
	}
}