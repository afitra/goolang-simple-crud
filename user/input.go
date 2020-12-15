package user

type RegisterUserInput struct {
	Username     string `json:"username" binding:"required"`
	Password     string `json:"password"  binding:"required"`
	Nama_Lengkap string `json:"namaLengkap"  binding:"required"`
}
type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password"  binding:"required"`
}
