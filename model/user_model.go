package model

type User struct {
	BaseModel
	OpenId     string `json:"open_id"`
	NickName   string `json:"nick_name"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"union_id"`
}
