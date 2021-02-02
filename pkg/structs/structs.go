package structs

type User struct {
	ID       int    `json:"-" db:"id"`
	Login    string `json:"login" binding:"required"` //  binding:"required" для валидации полей в теле запроса
	Username string `json:"username" binding:"required"`
	Pass     string `json:"pass" binding:"required"`
	//Role     string `json:"role" binding:"required"`
}

type WHitem struct {
	ID      int        `json:"id"`
	itemID  int        `json:"item_id"`
	Item    WHitemType `json:"item" `
	InStock bool       `json:"in_stock"`
}

type WHitemType struct {
	Strorage `json:"storage"`
	Monitor  `json:"monitor"`
}

type Strorage struct {
	Title  string `json:"title" `
	Size   string `json:"size" `
	Volume int    `json:"volume" `
	Type   string `json:"type" "`
	Vendor string `json:"vendor" `
}

type Monitor struct {
	Title    string `json:"title" `
	Diagonal int    `json:"diagonal" `
	Matrix   string `json:"matrix" `
	Vendor   string `json:"vendor"`
}

type Printer struct {
	Title  string `json:"title" `
	Type   string `json:"type" `
	Vendor string `json:"vendor" `
}

type CPU struct {
	Type   string  `json:"type" `
	Freq   float64 `json:"freq" `
	Socket string  `json:"socket" `
	Vendor string  `json:"vendor" `
}

type GPU struct {
	Type   string `json:"type" `
	Memory int    `json:"memory" `
	Vendor string `json:"vendor" `
}

type RUM struct {
	Type   string `json:"type"`
	Memory int    `json:"memory" `
	Vendor string `json:"vendor"`
}
