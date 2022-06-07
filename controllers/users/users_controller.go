package users

import (
	"net/http"
	"strconv"

	"github.com/Moreh89/bookstore_users-api/domain/users"
	"github.com/Moreh89/bookstore_users-api/services"
	"github.com/Moreh89/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User
	// fmt.Println(user)
	// bytes, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	// TODO handle err
	// 	return
	// }
	// if err := json.Unmarshal(bytes, &user);	err != nil {
	// 	// TODO handle err
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// optimization of previos code
	if err := c.ShouldBindJSON(&user); err != nil {
		// curl -X POST localhost:8081/users -d '{"id":1,"first_name":"damian","email":"d@asd.com","created_date":"04/06/2022}'
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveError := services.CreateUser(user)
	if saveError != nil {
		c.JSON(saveError.Status, saveError)
		return
	}
	// curl -X POST localhost:8081/users -d '{"id":1,"first_name":"damian","email":"d@asd.com","created_date":"2"}'
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"),10,64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status, err)
		return
	}
	user, getErr := services.GetUser(userId)
	if getErr != nil{
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}
