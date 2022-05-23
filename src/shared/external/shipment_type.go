package external

// ShipmentCost struct
type ShipmentCost struct {
	CityID          string `json:"cityID"`
	ProvinceID      string `json:"provinceID"`
	CityOrigin      string `json:"cityOrigin"`
	CityDestination string `json:"cityDestination"`
	Weight          int    `json:"weight"`
	Courier         string `json:"courier"`
	SellerID        int    `json:"sellerID"`
}
