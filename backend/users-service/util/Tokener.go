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
	accTokenDuration, _ = time.ParseDuration("5m")
	refTokenDuration, _ = time.ParseDuration("20m")
)

var jwtSecret []byte

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	_ = godotenv.Load(os.Getenv("WORK_DIR") + "credentials.env")
	jwtSecret = []byte(os.Getenv("JWT_SECRET"))
}

func CheckToken(tokenString string) (successful bool, uid bson.ObjectId, role int8, exp time.Time, err error) {
	var token *jwt.Token
	token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})
	if token == nil || err != nil {
		return false, uid, role, exp, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		uid, uidOk := claims["uid"].(bson.ObjectId)
		role, roleOk := claims["role"].(int8)
		exp, expOk := claims["exp"].(time.Time)
		ref, refOk := claims["ref"].(bool)
		if uidOk && roleOk && expOk && refOk && !ref && exp.After(time.Now()) {
			return true, uid, role, exp, err
		}
	}
	return false, uid, role, exp, err
}

func SignToken(uid bson.ObjectId, role int8, isRefreshToken bool) (result string, err error) {
	currentTime := time.Now()
	var expireTime time.Time
	if isRefreshToken {
		expireTime = currentTime.Add(refTokenDuration)
	} else {
		expireTime = currentTime.Add(accTokenDuration)
	}
	result, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"uid": uid, "role": role, "exp": expireTime, "ref": isRefreshToken}).SignedString(jwtSecret)
	return result, err
}
