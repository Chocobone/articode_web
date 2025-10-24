package user

import (
	"strconv"

	userRepo "github.com/chocobone/articode_web/user/repository"
	"github.com/chocobone/articode_web/util"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *UserService
}

func NewUserHandler(service *UserService) *UserHandler {
	return &UserHandler{service: service}
}

// GET, /api/users/
func (h *UserHandler) GetUserInfo(c *gin.Context) {
	ctx := c.Request.Context()

	userIDStr, err := util.GetUserIDFromContext(c)
	if err != nil {
		util.RespondUnauthorized(c, "Unauthorized access")
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		util.RespondBadRequest(c, "Bad Request")
		return
	}

	userInfoPtr, err := h.service.GetUserInfo(ctx, userID)
	if err != nil {
		util.RespondInternalError(c, err.Error())
		return
	}

	userInfoResponse := *userInfoPtr
	util.RespondSuccess(c, userInfoResponse)
}

// POST, /api/users
func (h *UserHandler) PostUserInfo(c *gin.Context) {
	ctx := c.Request.Context()

	userIDStr, err := util.GetUserIDFromContext(c)
	if err != nil {
		util.RespondUnauthorized(c, "Unauthorized access")
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		util.RespondBadRequest(c, "Bad Request")
		return
	}

	var newUser userRepo.UserInfoResponse
	if err := c.ShouldBindJSON(&newUser); err != nil {
		util.RespondBadRequest(c, "Invalid request body")
		return
	}

	newUser.UserID = userID

	createdUser, err := h.service.PostUserInfo(ctx, &newUser)
	if err != nil {
		util.RespondInternalError(c, err.Error())
		return
	}

	util.RespondSuccess(c, createdUser)
}

func (h *UserHandler) DeleteUserInfo(c *gin.Context) {
	ctx := c.Request.Context()

	userIDStr, err := util.GetUserIDFromContext(c)
	if err != nil {
		util.RespondUnauthorized(c, "Unauthorized access")
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		util.RespondBadRequest(c, "Bad Request")
		return
	}

	if err := h.service.DeleteUserInfo(ctx, userID); err != nil {
		util.RespondInternalError(c, err.Error())
		return
	}

	util.RespondSuccess(c, gin.H{"deleted": true})
}
