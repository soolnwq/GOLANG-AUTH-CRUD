package entities

type User struct {
	ID        int    `db:"id"`
	Email     string `db:"email"`
	Username  string `db:"username"`
	Password  string `db:"password"`
	CreatedAt string `db:"created_at"`
}
