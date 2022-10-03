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
	searchFriendshipHandler(*gin.Context)
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
	if err != nil && err.Error() == "already friend" {
		response.BadRequest(ctx, err.Error())
		return
	}

	if err != nil {
		response.InternalServerError(ctx, err.Error())
		return
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
		return
	}

	response.OK(ctx, friendslist)

}
func (fc *friendshipController) deleteFriendshipHandler(ctx *gin.Context) {
	var deleteFriendshipDTO DeleteFriendshipDTO

	if valid := validation.Bind(ctx, &deleteFriendshipDTO); !valid {
		return
	}

	err := fc.service.deleteFriendship(ctx.Request.Context(), deleteFriendshipDTO)
	if err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	response.OK(ctx, "unfriend done")
}

func (fc *friendshipController) searchFriendshipHandler(ctx *gin.Context) {
	var searchFriendshipDTO SearchFriendshipDTO
	if valid := validation.Bind(ctx, &searchFriendshipDTO); !valid {
		return
	}

	searchlist, err := fc.service.searchFriendship(ctx.Request.Context(), searchFriendshipDTO)
	if err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	response.OK(ctx, searchlist)

}
