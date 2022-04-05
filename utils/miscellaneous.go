package utils

type Wallet struct {
	ID      int64  `json:"id"`
	UserId  string `json:"user_id"`
	Balance int64  `json:"balance"`
}
