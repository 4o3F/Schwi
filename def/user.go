package def

type User struct {
	Recaptcha  string `json:"recaptcha,omitempty"`
	Uid        int    `json:"uid,omitempty"`
	Name       string `json:"name"`
	Password   string `json:"password,omitempty"`
	Email      string `json:"email"`
	Experience int    `json:"experience"`
}
