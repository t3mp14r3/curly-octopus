package domain

import (
    "errors"
    "net/mail"
)

type RegisterRequest struct {
    Login       string  `json:"login" db:"login"`
    Password    string  `json:"password" db:"password"`
    Name        string  `json:"name" db:"name"`
    Email       string  `json:"email" db:"email"`
}

func (r RegisterRequest) Validate() error {
    if len(r.Login) == 0 {
        return errors.New("login is required")
    }
    
    if len(r.Password) == 0 {
        return errors.New("password is required")
    }
    
    if len(r.Name) == 0 {
        return errors.New("name is required")
    }
    
    if len(r.Email) == 0 {
        return errors.New("email is required")
    }

    if _, err := mail.ParseAddress(r.Email); err != nil {
        return errors.New("invalid email")
    }

    return nil
}

type LoginRequest struct {
    Login       string  `json:"login" db:"login"`
    Password    string  `json:"password" db:"password"`
}

func (r LoginRequest) Validate() error {
    if len(r.Login) == 0 {
        return errors.New("login is required")
    }
    
    if len(r.Password) == 0 {
        return errors.New("password is required")
    }

    return nil
}

type CreateProductRequest struct {
    Name        string  `json:"name" db:"name"`
    Desc        string  `json:"desc" db:"desc"`
    Cost        int64   `json:"cost" db:"cost"`
    Barcode     string  `json:"barcode" db:"barcode"`
}

func (r CreateProductRequest) Validate() error {
    if len(r.Name) == 0 {
        return errors.New("product name is required")
    }
    
    if r.Cost == 0 {
        return errors.New("cost must be greater than zero")
    }
    
    if len(r.Barcode) == 0 {
        return errors.New("product barcode is required")
    }
    
    return nil
}
