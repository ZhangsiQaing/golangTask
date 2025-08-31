package sql

type Students struct {
	Id    int64  `gorm:"primaryKey;autoIncrement;column:id"`
	Name  string `gorm:"type:varchar(100);not null;column:name"`
	Age   int64  `gorm:"not null;column:age"`
	Grade string `gorm:"type:varchar(50);column:grade"`
}
