package db

type Location struct {
	ID       string
	Address  string
	Postcode int
	State    string
	Country  string
}

type Store struct {
	ID   string
	Name string
}

type MenuItem struct {
	ID       string
	Name     string
	Calories int
}
