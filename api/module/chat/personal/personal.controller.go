package personal

import (
	"github.com/cluster05/chat/pkg/response"
	"github.com/cluster05/chat/pkg/validation"

	"github.com/gin-gonic/gin"
)

type PersonalChatController interface {
	getChatHandler(*gin.Context)
}

type personalChatController struct {
	service PersonalChatService
}

func NewPersonalChatController(service PersonalChatService) PersonalChatController {
	return &personalChatController{
		service: service,
	}
}

func (pcc *personalChatController) getChatHandler(ctx *gin.Context) {
	var getPersonalChatDTO GetPersonalChatDTO
	if valid := validation.Bind(ctx, &getPersonalChatDTO); !valid {
		return
	}

	chats, err := pcc.service.getChat(ctx.Request.Context(), getPersonalChatDTO.FriendshipId)
	if err != nil {
		response.InternalServerError(ctx, err.Error())
		return
	}

	response.OK(ctx, chats)

}
