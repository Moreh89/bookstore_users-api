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
	result, saveError := services.UsersService.CreateUser(user)
	if saveError != nil {
		c.JSON(saveError.Status, saveError)
		return
	}
	// curl -X POST localhost:8081/users -d '{"id":1,"first_name":"damian","email":"d@asd.com","created_date":"2"}'
	c.JSON(http.StatusCreated, result.Marshall((c.GetHeader("X-Public") == "true")))
}

func GetUser(c *gin.Context) {
	userId, idErr := GetUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}
	user, getErr := services.UsersService.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user.Marshall((c.GetHeader("X-Public") == "true")))
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}

func UpdateUser(c *gin.Context) {
	userId, idErr := GetUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		// curl -X POST localhost:8081/users -d '{"id":1,"first_name":"damian","email":"d@asd.com","created_date":"04/06/2022}'
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	user.Id = userId

	isPartial := c.Request.Method == http.MethodPatch
	result, err := services.UsersService.UpdateUser(isPartial, user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, result.Marshall((c.GetHeader("X-Public") == "true")))
}

func DeleteUser(c *gin.Context) {
	userId, idErr := GetUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	if err := services.UsersService.DeleteUser(userId); err != nil {
		c.JSON(err.Status, idErr)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func GetUserId(userIdParam string) (int64, *errors.RestError) {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequestError("user id should be a number")
	}
	return userId, nil
}

func Search(c *gin.Context){
	status := c.Query("status")
	users, err := services.UsersService.Search(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, users.Marshall((c.GetHeader("X-Public") == "true")))
}