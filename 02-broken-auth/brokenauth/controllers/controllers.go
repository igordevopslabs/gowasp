package controllers

import (
	"brokenauth/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var db *gorm.DB

var jwtKey = []byte("secret_key")

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

func RegisterRoutes(r *gin.Engine, database *gorm.DB) {
	db = database

	r.POST("/register", register)
	r.POST("/login", login)
	r.GET("/protected", protected)
}

// registra user
func register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//não podemos salvar a senha em texto plano, da maneira que recebemos
	//precisamos salva-la minimanete criptada no banco
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//salva a senha com hash no banco
	user.Password = string(hashedPassword)

	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	/*	//esse modelo é inseguro pois faz a query considerando a senha em texto plano
		var foundUser models.User
		if err := db.Where("username = ? AND password = ?", user.Username, user.Password).First(&foundUser).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}*/

	//separa a query em user
	var foundUser models.User
	if err := db.Where("username = ?", user.Username).First(&foundUser).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	//separa query para comparar o hash das senhas
	if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	//seta um periodo de expiração do token
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		UserID: foundUser.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

/*
	//abandona a geração amadora de token
	func genToken() string {

	bytes := make([]byte, 16)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
	}
*/

func protected(c *gin.Context) {

	tokenString := c.GetHeader("Authorization")

	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "No token provided!"})
	}

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	/*
		//não precisa mais validar o token no banco
			var userToken models.Token
		    if err := db.Where("token = ?", token).First(&userToken).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}
	*/

	c.JSON(http.StatusOK, gin.H{"message": "Congratulations! You have access to this resource"})
}
