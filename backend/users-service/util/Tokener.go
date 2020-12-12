package util

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
	"os"
	"time"
)

var (
	accTokenDuration, _ = time.ParseDuration("15m")
	refTokenDuration, _ = time.ParseDuration("30m")
)

var jwtSecret []byte

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	_ = godotenv.Load(os.Getenv("WORK_DIR") + "credentials.env")
	jwtSecret = []byte(os.Getenv("JWT_SECRET"))
}

func ParseToken(tokenString string) (successful bool, uid bson.ObjectId, role int8, err error) {
	var token *jwt.Token
	token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})
	if token == nil || err != nil {
		return false, uid, role, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		uid = bson.ObjectIdHex(claims["uid"].(string))
		role = int8(claims["role"].(float64))
		return true, uid, role, err
	}
	return false, uid, role, err
}

func SignToken(uid bson.ObjectId, role int8, ref bool) (tokenString string, err error) {
	var exp int64
	if ref {
		exp = time.Now().Add(refTokenDuration).Unix()
	} else {
		exp = time.Now().Add(accTokenDuration).Unix()
	}
	tokenString, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"uid": uid, "role": role, "ref": ref, "exp": exp}).SignedString(jwtSecret)
	return tokenString, err
}
