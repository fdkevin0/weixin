package core

type User struct {
	CardNumber   string `json:"card_number"` //学工号
	Name         string `json:"name"`        //姓名
	EnName       string `json:"en_name"`
	IdentityType string `json:"identity_type"`
	HeadImage    string
	Gender       string
	Birthday     string
	StartAt      string
	ExpireAt     string
	Organization string
	Grade        string
	Collage      string
	Profession   string
	Class        string
	Campus       string
	DormNumber   string
	Telephone    string
	OfficePhone  string
	Email        string
	QQ           string
}
