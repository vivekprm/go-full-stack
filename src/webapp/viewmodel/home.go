package viewmodel

import "time"

type Todo struct {
	Name string
	Category string
	Description string
	Done bool
	Deleted bool
	CreatedOn *time.Time
	UpdatedOn *time.Time
}
type Home struct {
	Title string
	Active string
	Todos []Todo
	Total int
}

func NewHome() Home {
	return Home {
		Title: "Todo - Home",
		Active: "home",
		Total: 8,
		Todos: []Todo{
			{
				Name: "Todo1",
				Category: "Work",
				Description: "Some quick example text to build on the card title and make up the bulk of the card's content.",
				Done: false,
				Deleted: false,
			},
			{
				Name: "Todo2",
				Category: "Work",
				Description: "Some quick example text to build on the card title and make up the bulk of the card's content.",
				Done: false,
				Deleted: false,
			},
			{
				Name: "Todo3",
				Category: "Work",
				Description: "Some quick example text to build on the card title and make up the bulk of the card's content.",
				Done: false,
				Deleted: false,
			},
			{
				Name: "Todo4",
				Category: "Work",
				Description: "Some quick example text to build on the card title and make up the bulk of the card's content.",
				Done: false,
				Deleted: false,
			},
			{
				Name: "Todo5",
				Category: "Work",
				Description: "Some quick example text to build on the card title and make up the bulk of the card's content.",
				Done: false,
				Deleted: false,
			},
			{
				Name: "Todo6",
				Category: "Work",
				Description: "Some quick example text to build on the card title and make up the bulk of the card's content.",
				Done: false,
				Deleted: false,
			},
			{
				Name: "Todo7",
				Category: "Personal",
				Description: "Some quick example text to build on the card title and make up the bulk of the card's content.",
				Done: false,
				Deleted: false,
			},
			{
				Name: "Todo8",
				Category: "Personal",
				Description: "Some quick example text to build on the card title and make up the bulk of the card's content.",
				Done: true,
				Deleted: false,
			},
		},
	}
}