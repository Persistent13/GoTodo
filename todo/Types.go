package todo

type (
	Todo struct {
		ID           uint
		Content      string
		CreatedAtUtc uint
		UpdatedAtUtc uint
		Done         bool
		IsDeleted    bool
	}

	PatchTodoPogo struct {
		ID        uint    `json:"id"`
		Content   *string `json:"content"`
		Done      *bool   `json:"done"`
		IsDeleted *bool   `json:"isDeleted"`
	}

	CreateTodoPogo struct {
		Content string `json:"content"`
	}

	Error struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
	}
)
