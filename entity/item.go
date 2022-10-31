package entity

type Item struct {
	Id         int    `json:"-" gorm:"primary key"`
	OrderRefer string `json:"-"`

	ChrtId      uint   `json:"chrt_id" validate:"required;numeric"`
	TrackNumber string `json:"track_number" validate:"required"`
	Price       uint   `json:"price" validate:"required;numeric"`
	Rid         string `json:"rid" validate:"required;len=12"`
	Name        string `json:"name" validate:"required"`
	Sale        uint   `json:"sale" validate:"required;numeric"`
	Size        string `json:"size" validate:"required"`
	TotalPrice  uint   `json:"total_price" validate:"required;numeric"`
	NmId        uint   `json:"nm_id" validate:"required;numeric"`
	Brand       string `json:"brand" validate:"required"`
	Status      uint   `json:"status" validate:"required;numeric"`
}
