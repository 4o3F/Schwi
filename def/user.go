package def

type User struct {
	Recaptcha  string `json:"recaptcha,omitempty"`
	Uid        int    `json:"uid,omitempty"`
	Username   string `json:"username,omitempty"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	Experience int    `json:"experience,omitempty"`
}
