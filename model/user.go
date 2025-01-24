package model

import (
    "golang.org/x/crypto/bcrypt"
)

type UserInfo struct {
    NickName string `json:"nick_name"`
    Desc     string `json:"desc"`
    ShareId  string `json:"share_id"`
    Avatar   string `json:"avatar"`
    Version  int    `json:"version,omitempty"`
    Username string `json:"username"`
    Password string `json:"password"`
    Role     string `json:"role"`
}

// HashPassword hashes the user's password using bcrypt.
func (u *UserInfo) HashPassword() error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    u.Password = string(hashedPassword)
    return nil
}

// CheckPassword compares the hashed password with a plain text password.
func (u *UserInfo) CheckPassword(password string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
    return err == nil
}