package jwt

import (
    "time"
    "github.com/dgrijalva/jwt-go"
    "os"
)

var secretKey = os.Getenv("JWT_SECRET_KEY")

func GenerateToken(userID string) (string, error) {
    claims := jwt.MapClaims{
        "userID": userID,
        "exp":    time.Now().Add(time.Hour * 24).Unix(), 
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    signedToken, err := token.SignedString([]byte(secretKey))
    if err != nil {
        return "", err
    }

    return signedToken, nil
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
    return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, jwt.ErrInvalidKey
        }
        return []byte(secretKey), nil
    })
}
