package reservation

type Reservation struct {
	Id         int    `json:"id"`
	HotelId    int    `json:"hotel_id"`
	CustomerId int    `json:"customer_id"`
	StartDate  string `json:"start_date"`
	EndDate    string `json:"end_date"`
	RoomType   int    `json:"room_type"`
	RoomNum    int    `json:"room_num"`
	Amount     int    `json:"amount"`
	State      int    `json:"state"`
}

type Hotel struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}
