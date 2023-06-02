package baby

type Baby struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	BirthDate string `json:"birthDate"`
}

type BabyEvent struct {
	Id        string  `json:"id"`
	BabyId    string  `json:"babyId"`
	Name      string  `json:"name"`
	CreatedAt string  `json:"createdAt"`
	Value     float32 `json:"value"`
	Comment   string  `json:"comment"`
}
