package handlers

import (
	"errors"
	"net/http"

	"github.com/devgugga/NullTask/internal/models"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// CreateUser handles the creation of a new user.
// It expects a JSON object in the request body with the user's information.
// The function validates the input, checks if the email is already registered,
// hashes the password, and saves the user to the database.
// If successful, it returns a JSON object with the created user and a HTTP status code 201.
// If there are any errors during the process, it returns an appropriate error message and a HTTP status code.
func (h *Handler) CreateUser(c echo.Context) (err error) {
	u := new(models.User)
	// Bind the request body to the User struct
	if err = c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	// Validate the User struct
	if err := c.Validate(u); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	var existingUser models.User
	// Check if the email is already registered
	if err := h.DB.Where("email =?", u.Email).First(&existingUser).Error; err == nil {
		return echo.NewHTTPError(http.StatusConflict, "E-mail já cadastrado")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusInternalServerError, "Erro ao verificar e-mail")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Erro ao gerar hash da senha")
	}
	u.Password = string(hashedPassword)

	// Save the user to the database
	if err := h.DB.Create(&u).Error; err != nil {
		h.Log.Error("Error creating user", "error", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Erro ao criar usuário")
	}

	// Return the created user and a HTTP status code 201
	return c.JSON(http.StatusCreated, u)
}

// GetUserByEmail retrieves a user from the database based on their email.
// It accepts an echo.Context as a parameter, which provides information about the request and response.
// The function extracts the user's email from the request parameters.
// It then queries the database using GORM to find a user with the given email.
// If the user is found, it returns the user as a JSON response with a HTTP status code 200.
// If the user is not found, it returns a HTTP status code 404 with an appropriate error message.
// If there is any error during the process, it logs the error using the provided logger and returns a HTTP status code 500 with an appropriate error message.
func (h *Handler) GetUserByEmail(c echo.Context) (err error) {
	userEmail := c.Param("email")

	var user models.User
	if err := h.DB.Where("email =?", userEmail).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "Usuário não encontrado")
		}
		h.Log.Error("Erro ao buscar o usuário", "error", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Erro ao buscar usuário")
	}
	return c.JSON(http.StatusOK, user)
}

// GetUserById retrieves a user from the database based on their ID.
//
// Parameters:
// c echo.Context: Provides information about the request and response.
//
// Returns:
// err error: An error if any occurred during the process. If successful, it returns nil.
//
// The function extracts the user's ID from the request parameters.
// It then queries the database using GORM to find a user with the given ID.
// If the user is found, it returns the user as a JSON response with a HTTP status code 200.
// If the user is not found, it returns a HTTP status code 404 with an appropriate error message.
// If there is any error during the process, it logs the error using the provided logger and returns a HTTP status code 500 with an appropriate error message.
func (h *Handler) GetUserById(c echo.Context) (err error) {
	userId := c.Param("id")

	var user models.User
	if err := h.DB.Where("id =?", userId).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "Usuário não encontrado")
		}
		h.Log.Error("Erro ao buscar o usuário", "error", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Erro ao buscar usuário")
	}
	return c.JSON(http.StatusOK, user)
}

// GetAllUsers retrieves all users from the database and returns them as a JSON response.
//
// Parameters:
// c echo.Context: Provides information about the request and response.
//
// Returns:
// err error: An error if any occurred during the process. If successful, it returns nil.
//
// The function initializes an empty slice of User structs.
// It then queries the database using GORM's Find method to retrieve all users and stores them in the slice.
// If there is an error during the database query, it logs the error using the provided logger and returns a HTTP status code 500 with an appropriate error message.
// If the query is successful, it returns the retrieved users as a JSON response with a HTTP status code 200.
func (h *Handler) GetAllUsers(c echo.Context) (err error) {
	var users []models.User
	result := h.DB.Find(&users)

	if result.Error != nil {
		h.Log.Error("Erro ao buscar usuários", "error", result.Error)
		return echo.NewHTTPError(http.StatusInternalServerError, "Erro ao buscar usuários")
	}

	return c.JSON(http.StatusOK, users)
}

// UpdateUserByEmail updates an existing user in the database based on their email.
//
// Parameters:
// c echo.Context: Provides information about the request and response.
//
// Returns:
// err error: An error if any occurred during the process. If successful, it returns nil.
//
// The function extracts the user's email from the request parameters.
// It then queries the database using GORM to find a user with the given email.
// If the user is not found, it returns a HTTP status code 404 with an appropriate error message.
// If the user is found, it binds the request body to the user struct, validates the struct,
// and saves the updated user to the database.
// If there is any error during the binding, validation, or saving process, it logs the error using the provided logger
// and returns a HTTP status code 500 with an appropriate error message.
// If the update is successful, it returns the updated user as a JSON response with a HTTP status code 200.
func (h *Handler) UpdateUserByEmail(c echo.Context) (err error) {
	userEmail := c.Param("email")

	var user models.User
	if err := h.DB.Where("email =?", userEmail).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "Usuário não encontrado")
		}
		h.Log.Error("Erro ao buscar o usuário", "error", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Erro ao buscar usuário")
	}

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Dados inválidos"})
	}

	if err := c.Validate(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	if err := h.DB.Save(&user).Error; err != nil {
		h.Log.Error("Erro ao atualizar o usuário", "error", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Erro ao atualizar usuário")
	}

	return c.JSON(http.StatusOK, user)
}

// UpdateUserById updates an existing user in the database based on their ID.
//
// Parameters:
// c echo.Context: Provides information about the request and response.
//
// Returns:
// err error: An error if any occurred during the process. If successful, it returns nil.
//
// The function extracts the user's ID from the request parameters.
// It then queries the database using GORM to find a user with the given ID.
// If the user is not found, it returns a HTTP status code 404 with an appropriate error message.
// If the user is found, it binds the request body to the user struct, validates the struct,
// and saves the updated user to the database.
// If there is any error during the binding, validation, or saving process, it logs the error using the provided logger
// and returns a HTTP status code 500 with an appropriate error message.
// If the update is successful, it returns the updated user as a JSON response with a HTTP status code 200.
func (h *Handler) UpdateUserById(c echo.Context) (err error) {
	userId := c.Param("id")

	var user models.User
	if err := h.DB.Where("id =?", userId).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "Usuário não encontrado")
		}
		h.Log.Error("Erro ao buscar usuário", "error", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Erro ao buscar usuário")
	}

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Dados inválidos"})
	}

	if err := c.Validate(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	if err := h.DB.Save(&user).Error; err != nil {
		h.Log.Error("Erro ao atualizar o usuário", "error", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Erro ao atualizar usuário")
	}

	return c.JSON(http.StatusOK, user)
}

// DeleteUserByEmail deletes a user from the database based on their email.
//
// Parameters:
// c echo.Context: Provides information about the request and response.
//
// Returns:
// err error: An error if any occurred during the process. If successful, it returns nil.
//
// The function extracts the user's email from the request parameters.
// It then queries the database using GORM to find a user with the given email.
// If the user is not found, it returns a HTTP status code 404 with an appropriate error message.
// If the user is found, it deletes the user from the database.
// If there is any error during the deletion process, it logs the error using the provided logger
// and returns a HTTP status code 500 with an appropriate error message.
// If the deletion is successful, it returns a JSON response with a HTTP status code 200 and a message "Usuário Deletado".
func (h *Handler) DeleteUserByEmail(c echo.Context) (err error) {
	userEmail := c.Param("email")

	var user models.User
	if err := h.DB.Where("email =?", userEmail).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "Usuário não encontrado")
		}
		h.Log.Error("Erro ao buscar usuário", "error", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Erro ao buscar usuário")
	}

	if err := h.DB.Delete(&user).Error; err != nil {
		h.Log.Error("Erro ao deletar o usuário", "error", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Erro ao deletar usuário")
	}

	return c.JSON(http.StatusOK, "Usuário Deletado")
}

// DeleteUserById deletes a user from the database based on their ID.
//
// Parameters:
// c echo.Context: Provides information about the request and response.
//
// Returns:
// err error: An error if any occurred during the process. If successful, it returns nil.
//
// The function extracts the user's ID from the request parameters.
// It then queries the database using GORM to find a user with the given ID.
// If the user is not found, it returns a HTTP status code 404 with an appropriate error message.
// If the user is found, it deletes the user from the database.
// If there is any error during the deletion process, it logs the error using the provided logger
// and returns a HTTP status code 500 with an appropriate error message.
// If the deletion is successful, it returns a JSON response with a HTTP status code 200 and a message "Usuário Deletado".
func (h *Handler) DeleteUserById(c echo.Context) (err error) {
	userId := c.Param("id")

	var user models.User
	if err := h.DB.Where("id =?", userId).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "Usuário não encontrado")
		}
		h.Log.Error("Erro ao buscar usuário", "error", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Erro ao buscar usuário")
	}

	if err := h.DB.Delete(&user).Error; err != nil {
		h.Log.Error("Erro ao deletar o usuário", "error", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Erro ao deletar usuário")
	}

	return c.JSON(http.StatusOK, "Usuário Deletado")
}
