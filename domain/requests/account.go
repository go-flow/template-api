package requests

// AccountRegister request model
type AccountRegister struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

// AccountLogin request model
type AccountLogin struct {
	Email    string
	Password string
}
