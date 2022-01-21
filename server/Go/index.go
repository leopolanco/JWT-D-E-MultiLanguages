package main

import (
	"encoding/json"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func main() {
    router := gin.Default()

    var trustedProxies = []string{"127.0.0.1"}
    router.SetTrustedProxies(trustedProxies)
    // Allowing all
    router.Use(cors.Default())
    
	router.POST("/decode", decodeEndpoint)
    router.POST("/encode", encodeEndpoint)

	router.Run()
}



func encodeEndpoint(c *gin.Context) {
    var payload map[string]interface{}
    err := c.ShouldBindJSON(&payload)

    if err != nil {
        fmt.Printf("\nError: %s \n", err)
    }

    // Have to pass the payload as jwt.MapClaims to be converted
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(payload))

    //Secret key have to be a []byte
    signedToken, err := token.SignedString([]byte("secret"))

    if err != nil {
        fmt.Printf("\nError: %s \n", err)
        c.JSON(200, string("Invalid Object"))
        return 
    }

    c.JSON(200, string(signedToken))
}

func decodeEndpoint(c *gin.Context) {
    var receivedJwt string
    err := c.ShouldBindJSON(&receivedJwt)
    
    if err != nil {
        fmt.Printf("\nError: %s \n", err)
    }

     //KeyFunc as nil because we are only reading the payload
     //And omitting validating the signature
     //This is an unsecure way of parsing the jwt but we are just doing it for practice
     token, err := jwt.Parse(receivedJwt, nil)

    if err != nil {
        fmt.Printf("\nError parsing the jwt: %s \n", err)
    }

    if token == nil {
        c.JSON(200, string("Invalid Token"))
        return
    }

    claims := token.Claims.(jwt.MapClaims)
    
    // Repeating code, can this be done in a better way?
    if len(claims) == 0 {
        c.JSON(200, string("Invalid Token"))
        return
    }

    jsonToken, err := json.Marshal(claims)

    if err != nil {
        fmt.Printf("\nError jsoning the jwt: %s\n", err)
    }

    fmt.Printf("%s", jsonToken)
    c.JSON(200, string(jsonToken))
}
