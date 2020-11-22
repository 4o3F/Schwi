package def

type User struct {
	Recaptcha  string `json:"recaptcha,omitempty"`
	Uid        int    `json:"uid,omitempty"`
	Username       string `json:"username"`
	Password   string `json:"password,omitempty"`
	Email      string `json:"email"`
	Experience int    `json:"experience"`
}
