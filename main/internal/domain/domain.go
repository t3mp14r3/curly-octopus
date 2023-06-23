package domain

type User struct {
    ID          string  `json:"-" db:"id"`
    Login       string  `json:"login" db:"login"`
    Password    string  `json:"password,omitempty" db:"password"`
    Name        string  `json:"name" db:"name"`
    Email       string  `json:"email" db:"email"`
}

type Product struct {
    ID          string  `json:"id" db:"id"`
    Name        string  `json:"name" db:"name"`
    Desc        string  `json:"desc" db:"desc"`
    Cost        int     `json:"cost" db:"cost"`
    Barcode     string  `json:"barcode" db:"barcode"`
    UserID      string  `json:"-" db:"user_id"`
}
