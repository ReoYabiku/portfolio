package entity

type Biography struct {
	ID          int64  `json:"id"`
	Src         string `json:"src"`
	Alt         string `json:"alt"`
	Summary     string `json:"summary"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Achievement struct {
	ID          int64  `json:"id"`
	Src         string `json:"src"`
	Alt         string `json:"alt"`
	Summary     string `json:"summary"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
