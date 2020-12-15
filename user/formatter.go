package user

type UserFormatter struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Nama_Lengkap string `json:"namaLengkap"`
	Foto         string `json:"foto"`
	Token        string `json:"token"`
}

func userFormatter(user User, token string) UserFormatter {

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

func FormatUsers(users []User) []UserFormatter {

	allFormatter := []UserFormatter{}
	for _, user := range users {
		token := ""
		data := userFormatter(user, token)

		allFormatter = append(allFormatter, data)
	}

	return allFormatter
}

func FormatOneUser(user User, token string) UserFormatter {
	return userFormatter(user, token)
}
