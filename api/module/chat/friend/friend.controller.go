package friend

import (
	"time"
	"web-chat/pkg/response"
	"web-chat/pkg/validation"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/ksuid"
)

type FriendshipController interface {
	createFriendshipHandler(*gin.Context)
	getFriendshipHandler(*gin.Context)
	deleteFriendshipHandler(*gin.Context)
}

type friendshipController struct {
	service FriendshipService
}

func NewFrinshipController(service FriendshipService) FriendshipController {
	return &friendshipController{
		service: service,
	}
}

func (fc *friendshipController) createFriendshipHandler(ctx *gin.Context) {
	var createFriendshipDTO CreateFriendshipDTO
	if valid := validation.Bind(ctx, &createFriendshipDTO); !valid {
		return
	}

	friendship := Friendship{
		FriendshipId: ksuid.New().String(),
		MeId:         createFriendshipDTO.MeId,
		FriendId:     createFriendshipDTO.FriendId,
		CreatedAt:    time.Now().Unix(),
		UpdatedAt:    time.Now().Unix(),
		IsDeleted:    false,
	}

	err := fc.service.createFriendship(ctx.Request.Context(), friendship)
	if err != nil {
		response.InternalServerError(ctx, err.Error())
	}

	response.Created(ctx, friendship)
}
func (fc *friendshipController) getFriendshipHandler(ctx *gin.Context) {

	var getFriendListDTO GetFriendListDTO

	if valid := validation.Bind(ctx, &getFriendListDTO); !valid {
		return
	}

	friendslist, err := fc.service.getFriendship(ctx.Request.Context(), getFriendListDTO.MeId)
	if err != nil {
		response.BadRequest(ctx, err.Error())
	}

	response.OK(ctx, friendslist)

}
func (fc *friendshipController) deleteFriendshipHandler(ctx *gin.Context) {
	var deleteFriendshipDTO DeleteFriendshipDTO

	if valid := validation.Bind(ctx, &deleteFriendshipDTO); !valid {
		return
	}

	err := fc.service.deleteFriendship(ctx.Request.Context(), deleteFriendshipDTO.FriendshipId)
	if err != nil {
		response.BadRequest(ctx, err.Error())
	}

	response.OK(ctx, "unfriend done")
}
