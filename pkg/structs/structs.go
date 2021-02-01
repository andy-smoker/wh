package structs

type User struct {
	ID       int    `json:"-" db:"id"`
	Login    string `json:"login" binding:"required"` //  binding:"required" для валидации полей в теле запроса
	Username string `json:"username" binding:"required"`
	Pass     string `json:"pass" binding:"required"`
	//Role     string `json:"role" binding:"required"`
}

type WHitem struct {
	ID       int        `json:"-"`
	ItemName string     `json:"item_type" binding:"required"`
	Item     WHitemType `json:"item" binding:"required"`
}

type WHlist struct {
	ID     int
	ItemID int
}

type WHitemType struct {
	Strorage `json:"storage"`
	Monitor  `json:"monitor"`
}

type Strorage struct {
	Title  string `json:"title" binding:"required"`
	Size   string `json:"size" binding:"required"`
	Volume int    `json:"volume" binding:"required"`
	Type   string `json:"type" binding:"required"`
}

type Monitor struct{}
