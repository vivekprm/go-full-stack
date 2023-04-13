package viewmodel

type Login struct {
	Title string
	Active string
}

func NewLogin() Login {
	return Login{
		Title: "Todo - Login",
		Active: "login",
	}
}