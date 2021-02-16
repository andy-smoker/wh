package structs

type User struct {
	ID       int    `json:"-" db:"id"`
	Login    string `json:"login" binding:"required"` //  binding:"required" для валидации полей в теле запроса
	Username string `json:"username" binding:"required"`
	Pass     string `json:"pass" binding:"required"`
	//Role     string `json:"role" binding:"required"`
}

type WHitem struct {
	ID int `json:"id"`

	ItemProps WHitemProps `json:"item"`
	ItemsType string      `json:"items_type"`
	InStock   bool        `json:"in_stock"`
}

type WHitemProps struct {
	ID     int    `json:"-"`
	Title  string `json:"title"`
	Vendor string `json:"vendor" `
	Strorage
	Monitor
}

type Strorage struct {
	Size   string `json:"size" `
	Volume int    `json:"volume" `
	Type   string `json:"type" `
}

type Monitor struct {
	Diagonal int    `json:"diagonal" `
	Matrix   string `json:"matrix" `
}

type Printer struct {
	Type string `json:"type" `
}

type CPU struct {
	CPUType   string  `json:"type" `
	CPUFreq   float64 `json:"freq" `
	CPUSocket string  `json:"socket" `
}

type GPU struct {
	GPUMemory int `json:"memory" `
}

type RUM struct {
	RUNMemory int `json:"memory" `
}
