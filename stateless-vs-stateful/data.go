package stateless_vs_stateful

type Geo struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

type Address struct {
	Street  string `json:"street"`
	Suite   string `json:"suite"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
	Geo     Geo    `json:"geo"`
}

type Company struct {
	Name        string `json:"name"`
	CatchPhrase string `json:"catchPhrase"`
	Bs          string `json:"bs"`
}

type User struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Address  Address `json:"address"`
	Phone    string  `json:"phone"`
	Website  string  `json:"website"`
	Company  Company `json:"company"`
}

var UsersData = []User{
	{
		ID:       1,
		Name:     "Leanne Graham",
		Username: "Bret",
		Email:    "Sincere@april.biz",
		Address: Address{
			Street:  "Kulas Light",
			Suite:   "Apt. 556",
			City:    "Gwenborough",
			Zipcode: "92998-3874",
			Geo: Geo{
				Lat: "-37.3159",
				Lng: "81.1496",
			},
		},
		Phone:   "1-770-736-8031 x56442",
		Website: "hildegard.org",
		Company: Company{
			Name:        "Romaguera-Crona",
			CatchPhrase: "Multi-layered client-server neural-net",
			Bs:          "harness real-time e-markets",
		},
	},
	{
		ID:       2,
		Name:     "Ervin Howell",
		Username: "Antonette",
		Email:    "Shanna@melissa.tv",
		Address: Address{
			Street:  "Victor Plains",
			Suite:   "Suite 879",
			City:    "Wisokyburgh",
			Zipcode: "90566-7771",
			Geo: Geo{
				Lat: "-43.9509",
				Lng: "-34.4618",
			},
		},
		Phone:   "010-692-6593 x09125",
		Website: "anastasia.net",
		Company: Company{
			Name:        "Deckow-Crist",
			CatchPhrase: "Proactive didactic contingency",
			Bs:          "synergize scalable supply-chains",
		},
	},
	{
		ID:       3,
		Name:     "Clementine Bauch",
		Username: "Samantha",
		Email:    "Nathan@yesenia.net",
		Address: Address{
			Street:  "Douglas Extension",
			Suite:   "Suite 847",
			City:    "McKenziehaven",
			Zipcode: "59590-4157",
			Geo: Geo{
				Lat: "-68.6102",
				Lng: "-47.0653",
			},
		},
		Phone:   "1-463-123-4447",
		Website: "ramiro.info",
		Company: Company{
			Name:        "Romaguera-Jacobson",
			CatchPhrase: "Face to face bifurcated interface",
			Bs:          "e-enable strategic applications",
		},
	},
	{
		ID:       4,
		Name:     "Patricia Lebsack",
		Username: "Karianne",
		Email:    "Julianne.OConner@kory.org",
		Address: Address{
			Street:  "Hoeger Mall",
			Suite:   "Apt. 692",
			City:    "South Elvis",
			Zipcode: "53919-4257",
			Geo: Geo{
				Lat: "29.4572",
				Lng: "-164.2990",
			},
		},
		Phone:   "493-170-9623 x156",
		Website: "kale.biz",
		Company: Company{
			Name:        "Robel-Corkery",
			CatchPhrase: "Multi-tiered zero tolerance productivity",
			Bs:          "transition cutting-edge web services",
		},
	},
	{
		ID:       5,
		Name:     "Chelsey Dietrich",
		Username: "Kamren",
		Email:    "Lucio_Hettinger@annie.ca",
		Address: Address{
			Street:  "Skiles Walks",
			Suite:   "Suite 351",
			City:    "Roscoeview",
			Zipcode: "33263",
			Geo: Geo{
				Lat: "-31.8129",
				Lng: "62.5342",
			},
		},
		Phone:   "(254)954-1289",
		Website: "demarco.info",
		Company: Company{
			Name:        "Keebler LLC",
			CatchPhrase: "User-centric fault-tolerant solution",
			Bs:          "revolutionize end-to-end systems",
		},
	},
	{
		ID:       6,
		Name:     "Mrs. Dennis Schulist",
		Username: "Leopoldo_Corkery",
		Email:    "Karley_Dach@jasper.info",
		Address: Address{
			Street:  "Norberto Crossing",
			Suite:   "Apt. 950",
			City:    "South Christy",
			Zipcode: "23505-1337",
			Geo: Geo{
				Lat: "-71.4197",
				Lng: "71.7478",
			},
		},
		Phone:   "1-477-935-8478 x6430",
		Website: "ola.org",
		Company: Company{
			Name:        "Considine-Lockman",
			CatchPhrase: "Synchronised bottom-line interface",
			Bs:          "e-enable innovative applications",
		},
	},
	{
		ID:       7,
		Name:     "Kurtis Weissnat",
		Username: "Elwyn.Skiles",
		Email:    "Telly.Hoeger@billy.biz",
		Address: Address{
			Street:  "Rex Trail",
			Suite:   "Suite 280",
			City:    "Howemouth",
			Zipcode: "58804-1099",
			Geo: Geo{
				Lat: "24.8918",
				Lng: "21.8984",
			},
		},
		Phone:   "210.067.6132",
		Website: "elvis.io",
		Company: Company{
			Name:        "Johns Group",
			CatchPhrase: "Configurable multimedia task-force",
			Bs:          "generate enterprise e-tailers",
		},
	},
	{
		ID:       8,
		Name:     "Nicholas Runolfsdottir V",
		Username: "Maxime_Nienow",
		Email:    "Sherwood@rosamond.me",
		Address: Address{
			Street:  "Ellsworth Summit",
			Suite:   "Suite 729",
			City:    "Aliyaview",
			Zipcode: "45169",
			Geo: Geo{
				Lat: "-14.3990",
				Lng: "-120.7677",
			},
		},
		Phone:   "586.493.6943 x140",
		Website: "jacynthe.com",
		Company: Company{
			Name:        "Abernathy Group",
			CatchPhrase: "Implemented secondary concept",
			Bs:          "e-enable extensible e-tailers",
		},
	},
	{
		ID:       9,
		Name:     "Glenna Reichert",
		Username: "Delphine",
		Email:    "Chaim_McDermott@dana.io",
		Address: Address{
			Street:  "Dayna Park",
			Suite:   "Suite 449",
			City:    "Bartholomebury",
			Zipcode: "76495-3109",
			Geo: Geo{
				Lat: "24.6463",
				Lng: "-168.8889",
			},
		},
		Phone:   "(775)976-6794 x41206",
		Website: "conrad.com",
		Company: Company{
			Name:        "Yost and Sons",
			CatchPhrase: "Switchable contextually-based project",
			Bs:          "aggregate real-time technologies",
		},
	},
	{
		ID:       10,
		Name:     "Clementina DuBuque",
		Username: "Moriah.Stanton",
		Email:    "Rey.Padberg@karina.biz",
		Address: Address{
			Street:  "Kattie Turnpike",
			Suite:   "Suite 198",
			City:    "Lebsackbury",
			Zipcode: "31428-2261",
			Geo: Geo{
				Lat: "-38.2386",
				Lng: "57.2232",
			},
		},
		Phone:   "024-648-3804",
		Website: "ambrose.net",
		Company: Company{
			Name:        "Hoeger LLC",
			CatchPhrase: "Centralized empowering task-force",
			Bs:          "target end-to-end models",
		},
	},
}
