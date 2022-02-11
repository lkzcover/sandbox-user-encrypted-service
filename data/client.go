package data

type Client struct {
	ID           uint64
	PersonalData struct {
		Name  string
		Phone string
	}
	Rating uint8
}
