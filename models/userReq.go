package models

type UserReq struct {
	Nama     string `json:"nama"`
	Kelas    string `json:"kelas"`
	Semester string `json:"semester"`
	Prodi    string `json:"prodi"`
	Wa       string `json:"wa"`
}
