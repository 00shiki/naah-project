package deliveries

type Province struct {
	ID   string
	Name string
}

type City struct {
	ID         string
	Name       string
	Type       string
	PostalCoda string
	Province
}
