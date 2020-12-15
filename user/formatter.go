package user

type UserFormatter struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Nama_Lengkap string `json:"namaLengkap"`
	Foto         string `json:"foto"`
	Token        string `json:"token"`
}

func FormatUser(user User, token string) UserFormatter {

	formatter := UserFormatter{
		ID:           user.ID,
		Username:     user.Username,
		Password:     user.Password,
		Nama_Lengkap: user.Nama_Lengkap,
		Token:        token,
		Foto:         user.Foto,
	}
	return formatter
}
