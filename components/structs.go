package components

type ResultT struct {
	Result ResultsT `json:"result"`
}

type ResultsT struct {
	Results []TaipeiJson `json:"results"`
}

type TaipeiJson struct {
	Id          int    `json:"_id"`
	Name        string `json:"name"`
	Category    string `json:"CAT"`
	Description string `json:"description"`
	Address     string `json:"address"`
	Transport   string `json:"direction"`
	Mrt         string `json:"MRT"`
	Lat         string `json:"latitude"`
	Lng         string `json:"longitude"`
	Imgs        string `json:"file"`
}

type Cat struct {
	Category_name string `json:"category_name"`
}
