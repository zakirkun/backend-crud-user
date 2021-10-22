package models

type Profile struct {
	ID                 int `gorm:"primaryKey"`
	AlamatKtp          string
	Pekerjaan          string
	NamaLengkap        string
	PendidikanTerakhir string
	NoHp               string
	User               User `gorm:"foreignKey:ID"`
}
