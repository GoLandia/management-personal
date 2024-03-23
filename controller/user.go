package controller

type User struct {
	Name     string
	Email    string
	Password string
}

func (u *User) Create() (string, string, string) {
	if u.Name == "" || u.Email == "" || u.Password == "" {
		return "", "", ""
	}

	return u.Name, u.Email, u.Password
}
