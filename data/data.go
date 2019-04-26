package data

type User struct {
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
}

type TestRes struct {
	Ret     int    `json:"ret"`
	Version int    `json:"version,omitempty"`
	Action  string `json:"action,omitempty"`
	Result  int    `json:"result,omitempty"`
}
type TUser struct {
	FirstName string
	LastName  string
	Age       int
}

type AuthenRes struct {
	Ret      int    `json:"ret,omitempty"`
	Name     string `json:"name,omitempty"`
	UserId   string `json:"userid,omitempty"`
	Greeting string `json:"greeting,omitempty"`
}
