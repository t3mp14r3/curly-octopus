package domain

type User struct {
    ID          int64   `json:"-" db:"id"`
    Login       string  `json:"login" db:"login"`
    Password    string  `json:"password,omitempty" db:"password"`
    Name        string  `json:"name" db:"name"`
    Email       string  `json:"email" db:"email"`
}
