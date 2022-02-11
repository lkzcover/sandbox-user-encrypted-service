package data

type Driver struct {
	OrderID *string // Уникальный идентификатор заказа
	Addr    string  // Адрес для связи
	Name    string
	Rating  uint8
	Cars    string

	Encrypted string
}

type DriverPersonalData struct {
	Phone string
}
