package postgres

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
	Volume string `json:"volume" binding:"required"`
	Type   string `json:"type" binding:"required"`
}

type Monitor struct{}
