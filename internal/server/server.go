package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"testauth/internal/database"
	"testauth/internal/entity/product"
	"testauth/internal/entity/user"
	"testauth/pkg/crypto"
	"testauth/pkg/encoder"
	"testauth/pkg/validator"

	"github.com/gin-gonic/gin"
)

func loginHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var u user.User

		if err := c.BindJSON(&u); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		targetUser, err := user.GetUserByEmail(db, u.Email)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
			return
		}

		fmt.Println("targetUser email: ", targetUser.Email)
		fmt.Println("targetUser pass: ", targetUser.Password)
		fmt.Println("u email: ", u.Email)
		fmt.Println("u pass: ", u.Password)

		matchPassword := encoder.CheckPasswordHash(u.Password, targetUser.Password)

		if targetUser.Email != u.Email || !matchPassword {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong email or password"})
			return
		}

		token, err := crypto.CreateToken(u.Username, u.Id, u.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Token error"})
			return
		}

		c.IndentedJSON(http.StatusOK, gin.H{
			"token": token,
		})

	}
}

func signupHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var u user.User
		//перевірка баді
		if err := c.BindJSON(&u); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request body"})
			return
		}
		//валідація пошти
		if !validator.ValidateEmail(u.Email) {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid email"})
			return
		}
		//валідація ніку
		if !validator.ValidateUsername(u.Username) {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Username must contain at least 3 characters"})
			return
		}
		//валідація паролю
		if !validator.ValidatePassword(u.Password) {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Password must contain at least 8 characters"})
			return
		}
		//перевірка у базі
		if err := user.CheckUser(db, u.Email); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "User already exists"})
			return
		}

		//хешування паролю
		hashPassword, err := encoder.HashPassword(u.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Hash password error"})
			return
		}

		u.Password = hashPassword
		fmt.Print("HashPass: ", hashPassword)
		fmt.Print("u.Password: ", u.Password)

		//інша помилка
		if err := user.NewUser(db, u); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "User with this username already exists"})
			return
		}

		c.IndentedJSON(http.StatusOK, gin.H{

			"Message": fmt.Sprintf("Welcome, %s", u.Username),
		})
	}
}

func protectedHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		list := product.GetProducts(db)
		c.JSON(200, list)
	}
}

func AuthMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		// Перевіряємо токен
		claims, err := crypto.VerifyTokenBearer(authHeader)
		if err != nil {
			c.JSON(401, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Перевіряємо, чи є email у claims
		email, ok := claims["email"].(string)
		if !ok || email == "" {
			c.JSON(401, gin.H{"error": "Invalid token payload"})
			c.Abort()
			return
		}

		// Шукаємо користувача в базі
		_, err = user.GetUserByEmail(db, email)
		if err != nil {
			c.JSON(401, gin.H{"error": "User not authorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func Init(ip, port string) {

	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Fatal error: ", err)
	}

	r := gin.Default()

	r.POST("/login", loginHandler(db))
	r.POST("/signup", signupHandler(db))
	r.GET("/protected", AuthMiddleware(db), protectedHandler(db))

	r.Run(ip + ":" + port)

}
