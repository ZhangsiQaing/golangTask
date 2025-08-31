package sql

type Books struct {
	Id     int64   `gorm:"primaryKey;autoIncrement;column:id"`
	Title  string  `gorm:"type:varchar(100);not null;column:title"`
	Author string  `gorm:"type:varchar(100);not null;column:author"`
	Price  float64 `gorm:"not null;column:price"`
}
