package sql

type Employees struct {
	Id         int64   `gorm:"primaryKey;autoIncrement;column:id"`
	Name       string  `gorm:"type:varchar(100);not null;column:name"`
	Department string  `gorm:"type:varchar(100);not null;column:department"`
	Salary     float64 `gorm:"not null;column:salary"`
}
