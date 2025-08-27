package user

import (
	"strconv"

	"github.com/Chocobone/articode_web/v2/util"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *UserService
}

func NewUserHandler(service *UserService) *UserHandler {
	return &UserHandler{service: service}
}

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
	if err != nil{
		util.RespondInternalError(c, err.Error())
		return
	}

	userInfoResponse := *userInfoPtr
	util.RespondSuccess(c, userInfoResponse)
}