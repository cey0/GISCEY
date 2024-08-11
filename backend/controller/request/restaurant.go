package request

type Create struct {
	Name      string  `json:"name" binding:"required"`
	TypeID    uint    `json:"type_id" binding:"required"`
	Deskripsi string  `json:"deskripsi" binding:"required"`
	Latitude  float64 `json:"latitude" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required"`
}
