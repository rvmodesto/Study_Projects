package token

import (
	"context"
	"e-commerce/database"
	"log"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SignedDetails struct{
	Email string 
	First_Name string
	Last_Name string
	Uid string
	jwt.StandardClaims
}

var UserData *mongo.Collection = (*mongo.Collection)(database.Client, "Users")

// secret-key = "introduce some noise into the token, making it safer"
var SERCRET_KEY = os.Getenv("SECRET_KEY")

//func name (o que é necessario)(o que retorna){}
func TokenGenerator(email string, firstname string, lastname string, uid string)(signedtoken string, signedfreshtoken string, err error){

	claims := &SignedDetails{
		Email: email,
		First_Name: firstname,
		Last_Name: lastname,
		Uid: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshclaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SERCRET_KEY))

	if err != nil {
		return "", "", err //token, refresh token, err -> (signedtoken string, signedfreshtoken string, err error) do inicio da func
	}

	refreshtoken, err := jwt.NewWithClaims(jwt.SigningMethodHS384, refreshclaims).SignedString([]byte(SERCRET_KEY))
	if err != nil {
		log.Panic(err)
		return
	}
	return token, refreshtoken, err

}

func ValidateToken(signedtoken string)(claims *SignedDetails, msg string){
	token, err := jwt.ParseWithClaims(signedtoken, &SignedDetails{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SERCRET_KEY), nil
	})
	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = "The token is invalid"
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix(){
		msg = "token is already expired"
	}
	return claims, msg

}

func UpdateAllTokens(signedtoken string, signedrefreshtoken string, userid string){

	var ctx, cancel = context.WithTimeout(context.Background(), 100 *time.Second)

	var updateobj primitive.D

	updateobj = append(updateobj, bson.E{Key:"token", Value: signedtoken})
	updateobj = append(updateobj, bson.E{Key:"refresh_token", Value: signedrefreshtoken})
	updated_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	updateobj = append(updateobj, bson.E{Key:"updateat", Value: updated_at})

	upsert := true

	filter:= bson.M{"user_id":userid}
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}
	_, err := UserData.UpdateOne(ctx, filter, bson.D{
		{Key:"$set", Value: updateobj},
	}, &opt)
	defer cancel()

	if err != nil {
		log.Panic(err)
		return
	}


}