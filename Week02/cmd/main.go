package main

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"strconv"

	"xll.com/go-000/Week02/internal/service"

	"github.com/gin-gonic/gin"
	xerrors "github.com/pkg/errors"
)

// GetUserHander creates the UserHandler
func GetUserHander(userService service.UserService) func(c *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			log.Printf("Encounter error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"Message": "Invalid request",
			})
			return
		}
		log.Printf("[GetUser] Request id: %v", id)
		user, err := userService.GetUser(id)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				log.Printf("Cannot find the user with id: %v", id)
				c.JSON(http.StatusNotFound, gin.H{
					"Message": "Cannot find the user",
				})

			} else {
				log.Printf("Encountered error: %T %v\n", xerrors.Cause(err), xerrors.Cause(err))
				log.Printf("Stack trace:\n%+v", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"Message": "Server Internal Error",
				})
			}
			return
		}

		c.JSON(http.StatusOK, user)
		log.Printf("[GetUser] Response: %#v", user)
	}
}

func main() {
	r := gin.Default()
	// initialize userService instance
	userService, err := service.NewUserService(service.UserServiceInMemoryType)
	if err != nil {
		log.Fatalf("Cannot initialize userService due to: %v", err)
	}
	r.GET("/user/:id", GetUserHander(userService))
	r.Run()
}