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