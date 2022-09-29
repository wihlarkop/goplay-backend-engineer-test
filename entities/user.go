package entities

type User struct {
	Id        int    `json:"id" db:"id"`
	Username  string `json:"username" db:"username"`
	Password  string `json:"password" db:"password"`
	CreatedAt string `json:"created_at" db:"created_at"`
}
