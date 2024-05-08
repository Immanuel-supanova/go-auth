package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/immanuel-supanova/go-auth/database"
	"github.com/immanuel-supanova/go-auth/models"
	"golang.org/x/crypto/bcrypt"
)

func UserCreate(c *gin.Context) {
	// Get Email and Password
	var data struct {
		Email     string
		Password  string
		Password2 string
	}

	if c.Bind(&data) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read data",
		})

		return

	}

	if data.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email not Provided",
		})
		return
	} else if data.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Password not Provided",
		})
		return
	} else if data.Password2 == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Confirm Password not Provided",
		})
		return
	}

	// Compare passwords
	if data.Password != data.Password2 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Passwords do not match",
		})
		return

	}

	// Hash Password
	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})

		return
	}

	// Create User
	user := models.User{Email: data.Email, Password: string(hash), IsActive: true}
	result := database.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})

		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"Message": "User Created",
	})
}

func UserCreateStaff(c *gin.Context) {
	// Get Email and Password
	var data struct {
		Email     string
		Password  string
		Password2 string
	}

	if c.Bind(&data) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read data",
		})

		return

	}

	if data.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email not Provided",
		})
		return
	} else if data.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Password not Provided",
		})
		return
	} else if data.Password2 == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Confirm Password not Provided",
		})
		return
	}

	// Compare passwords
	if data.Password != data.Password2 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Passwords do not match",
		})
		return

	}

	// Hash Password
	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})

		return
	}

	// Create User
	user := models.User{Email: data.Email, Password: string(hash), IsActive: true, IsStaff: true}
	result := database.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})

		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"Message": "User Created",
	})
}

func UserCreateAdmin(c *gin.Context) {
	// Get user inputs
	var data struct {
		Email     string
		Password  string
		Password2 string
	}

	if c.Bind(&data) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read data",
		})

		return

	}

	if data.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email not Provided",
		})
		return
	} else if data.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Password not Provided",
		})
		return
	} else if data.Password2 == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Confirm Password not Provided",
		})
		return
	}

	// Compare passwords
	if data.Password != data.Password2 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Passwords do not match",
		})
		return

	}

	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	// check if there is an admin
	result := database.DB.Where(&models.User{IsAdmin: true}).First(&models.User{})

	if result.Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "An admin already exists",
		})
		return
	}

	// create admin
	dev2 := models.User{Email: data.Email, Password: string(hash), IsActive: true, IsAdmin: true}
	result = database.DB.Create(&dev2)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create admin:",
		})
		return

	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"Message": "Admin Created",
	})
}

func UserRead(c *gin.Context) {
	// Get Email
	var data struct {
		Email string
	}

	if c.Bind(&data) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read data",
		})

		return

	}

	// Get user data from database
	var user = models.User{Email: data.Email}
	database.DB.Omit("Password").Find(&user)

	// Check if User exists
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user",
		})

		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})

}

func UserDelete(c *gin.Context) {
	// Get Email
	var data struct {
		Email string
	}

	if c.Bind(&data) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read data",
		})

		return

	}

	// Check if User exists
	var user = models.User{Email: data.Email}
	database.DB.First(&user)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user",
		})

		return
	}

	// Delete User
	result := database.DB.Delete(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to delete user",
		})

		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"message": "Deletion successfull",
	})
}

func UserList(c *gin.Context) {
	// Get all users from database
	var users []models.User
	database.DB.Omit("Password").Find(&users)

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func UserIsActiveChange(c *gin.Context) {
	// Get Email
	var data struct {
		Email    string
		isActive bool
	}

	if c.Bind(&data) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read data",
		})

		return

	}

	// Check if user exists
	var user = models.User{Email: data.Email}
	database.DB.First(&user)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user",
		})

	}

	// Update User
	result := database.DB.Model(&user).Updates(models.User{
		IsActive: data.isActive,
	})

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to update user",
		})

		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"message": "Update successfull",
	})
}

func UserIsStaffChange(c *gin.Context) {
	// Get Email
	var data struct {
		Email   string
		isStaff bool
	}

	if c.Bind(&data) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read data",
		})

		return

	}

	// Check if user exists
	var user = models.User{Email: data.Email}
	database.DB.First(&user)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user",
		})

	}

	// Update User
	result := database.DB.Model(&user).Updates(models.User{
		IsStaff: data.isStaff,
	})

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to update user",
		})

		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"message": "Update successfull",
	})
}

func UserIsAdminChange(c *gin.Context) {
	// Get Email
	var data struct {
		Email   string
		isAdmin bool
	}

	if c.Bind(&data) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read data",
		})

		return

	}

	// Check if user exists
	var user = models.User{Email: data.Email}
	database.DB.First(&user)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user",
		})

	}

	// Update User
	result := database.DB.Model(&user).Updates(models.User{
		IsAdmin: data.isAdmin,
	})

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to update user",
		})

		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"message": "Update successfull",
	})
}

func UserResetPassword(c *gin.Context) {
	// Get Email, OldPassword, NewPassword, ConfirmPassword
	var data struct {
		Email           string
		OldPassword     string `json:"oldpassword"`
		NewPassword     string `json:"newpassword"`
		ConfirmPassword string `json:"confirmpassword"`
	}

	if c.Bind(&data) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read data",
		})
	}

	if data.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email not Provided",
		})
		return
	} else if data.OldPassword == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Old Password not Provided",
		})
		return
	} else if data.NewPassword == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "New Password not Provided",
		})
		return
	} else if data.ConfirmPassword == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Confirm Password not Provided",
		})
		return
	}

	// Check if Developer exists
	var dev = models.User{Email: data.Email}
	database.DB.First(&dev)

	if dev.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid dev",
		})

		return
	}

	// Check the OldPassword if it matches in the database
	err := bcrypt.CompareHashAndPassword([]byte(dev.Password), []byte(data.OldPassword))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid old password",
		})

		return
	}
	// Check if NewPassword and ConfirmPassword are the same
	if data.NewPassword != data.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Passwords do not match",
		})
	}

	// Hash the NewPassword
	hash, err := bcrypt.GenerateFromPassword([]byte(data.NewPassword), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})

		return
	}

	// Update the Developer
	result := database.DB.Model(&dev).Updates(models.User{
		Password: string(hash),
	})

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to update password",
		})

		return
	}
	// Respond
	c.Status(http.StatusOK)
}
