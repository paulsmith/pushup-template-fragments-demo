package build

type Contact struct {
	id       int
	email    string
	archived bool
}

var contact Contact

func init() {
	contact = Contact{id: 42, email: "foo@bar.net"}
	contact.archived = false
}

var pushupGreeting = "Welcome to Pushup!"
