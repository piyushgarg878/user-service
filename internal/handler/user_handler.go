package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"user-service/internal/domain"
	"user-service/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
    service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
    return &UserHandler{service: s}
}


func (h *UserHandler) Register(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.Register(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user created", "user_id": user.ID})	
}
func (h *UserHandler) GetUserByName(c *gin.Context) {
	name := c.Param("name")

	user, err := h.service.GetUserByName(name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
func (h *UserHandler) GetProfile(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    profile, err := h.service.ProfileRepo.GetByUserID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "profile not found"})
        return
    }
    c.JSON(http.StatusOK, profile)
}
func (h *UserHandler) UpdateProfile(c *gin.Context) {
	var profile domain.UserProfile

	if err := c.ShouldBindJSON(&profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// get user ID from URL param
	userIDParam := c.Param("id")
	var userID uint
	_, err := fmt.Sscan(userIDParam, &userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	if err := h.service.UpdateProfile(userID, &profile); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "profile updated"})
}
