package children

type Children struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	BirthDate string `json:"birthDate"`
}

type ChildrenEvent struct {
	Id        string  `json:"id"`
	ChildId   string  `json:"childId"`
	Name      string  `json:"name"`
	CreatedAt string  `json:"createdAt"`
	Value     float32 `json:"value"`
	Comment   string  `json:"comment"`
}

var ChildrenList = []Children{
	{Id: "1", FirstName: "Toto", LastName: "Tata", BirthDate: "2022-12-23"},
	{Id: "2", FirstName: "titi", LastName: "Titi", BirthDate: "2022-12-23"},
	{Id: "3", FirstName: "Tutu", LastName: "Tutu", BirthDate: "2022-12-23"},
}
