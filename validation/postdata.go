package validation

type InsertData struct {
	Username           string `json:"username" binding:"required"`
	Email              string `json:"email" binding:"required"`
	PendidikanTerakhir string `json:"pendidikan_terakhir" binding:"required"`
	Pekerjaan          string `json:"pekerjaan" binding:"required"`
	NamaLengkap        string `json:"nama_lengkap" binding:"false"`
	AlamatKtp          string `json:"alamat" binding:"false"`
	NoHp               string `json:"no_hp" binding:"false"`
}
