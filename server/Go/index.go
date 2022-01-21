package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    // Allowing all
    router.Use(cors.Default())

	router.POST("/decode", decodeEndpoint)
    router.POST("/encode", encodeEndpoint)

	router.Run()
}

func encodeEndpoint(c *gin.Context) {
    json, err := c.GetRawData()
    fmt.Printf("Received a hit on encode endpoint %s", json)
    if err != nil{
        fmt.Printf("Error: %s", err)
     }
    fmt.Printf("%s", json)
    c.JSON(200, gin.H{
        "json":"json",
    })
}

func decodeEndpoint(c *gin.Context) {
    //Get raw data because client is not sending structured json
    //Just assume that rawData is a string, anything else is not supported
    type Body struct {
        Jwt string `json:"jwt"`
    }
    
    var body Body

    // rawData, err := c.GetRawData()

    err := c.ShouldBindJSON(&body)
    
    if err != nil {
        fmt.Printf("Error: %s", err)
    }

    fmt.Printf("token: %s", body.Jwt)
    
    // tokenString, err := strconv.Unquote(string(body.jwt))

    // fmt.Printf("This is the token %s", tokenString)

    // if err != nil {
    //     fmt.Printf("Error: %s", err)
    // }

    //  token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	// 	return []byte("secret"), nil
	// })

    // if err != nil {
    //     fmt.Printf("Error: %s", err)
    // }

    // // validate the essential claims
    // if !token.Valid {
    //     c.JSON(200, "Invalid token")

    // } else {
    //     fmt.Printf("%s", token)
    //     c.JSON(200, token)
    // }

}
