package dao

type (
	Order struct {
		CategoryRoom string `json:"category_room"`
		UserEmail    string `json:"email"`
		From         string `json:"from"`
		To           string `json:"to"`
	}

	Room struct {
		Type  string `json:"type"`
		Count int    `json:"count"`
		From  string `json:"from"`
		To    string `json:"to"`
	}
)
