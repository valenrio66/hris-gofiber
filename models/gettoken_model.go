package models

type Credentials struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	IdPengguna string `json:"id_pengguna"`
}

type TokenResponse struct {
	Token string `json:"token"`
}
