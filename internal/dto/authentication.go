package dto

type ConsumerDto struct {
	Username string `json:"username"`
}

type JwtResponseDto struct {
	Key      string `json:"key"`
	Secret   string `json:"secret"`
	CreateAt int64  `json:"created_at"`
}
