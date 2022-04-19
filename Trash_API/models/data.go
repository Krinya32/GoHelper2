package models

// Наша "База данных"
var DB []Pizza

type Pizza struct {
	ID       int     `json:"id"`
	Diameter int     `json:"diameter"`
	Price    float64 `json:"price"`
	Title    string  `json:"title"`
}

func init() {
	pizza1 := Pizza{
		ID:       1,
		Diameter: 22,
		Price:    500.50,
		Title:    "Pepperoni",
	}

	pizza2 := Pizza{
		ID:       2,
		Diameter: 25,
		Price:    650.23,
		Title:    "BBQ",
	}

	pizza3 := Pizza{
		ID:       3,
		Diameter: 22,
		Price:    450,
		Title:    "Margarita",
	}
	DB = append(DB, pizza1, pizza2, pizza3)
}

// Helper
type ErrorMessage struct {
	Message string `json:"message"`
}

//Вспомогательная функция для модели (модельный метод)
func FindPizzaById(id int) (Pizza, bool) {
	var pizza Pizza
	var found bool
	for _, p := range DB {
		if p.ID == id {
			pizza = p
			found = true
			break
		}
	}
	return pizza, found
}
