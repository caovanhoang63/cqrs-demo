package commands

type CreateReservation struct {
	HotelId   int    `json:"hotel_id"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	RoomType  int    `json:"room_type"`
	RoomNum   int    `json:"room_num"`
}

type UpdateReservation struct {
	// ... thong tin up date
}

type UpdateState struct {
	Id    int `json:"id"`
	State int `json:"state"`
}

type DeleteReservation struct {
	Id int `json:"id"`
}
