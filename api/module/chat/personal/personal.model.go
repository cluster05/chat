package personal

type PersonalChat struct {
	PersonalChatId string `bson:"personalChatId" json:"personalChatId"`
	FriendshipId   string `bson:"friendshipId" json:"friendshipId"`
	From           string `bson:"from" json:"from"`
	To             string `bson:"to" json:"to"`
	Message        string `bson:"message" json:"message"`
	CreatedAt      int64  `bson:"createdAt" json:"createdAt"`
	UpdatedAt      int64  `bson:"updatedAt" json:"updatedAt"`
	IsDeleted      bool   `bson:"isDeleted" json:"isDeleted"`
}

type PersonalChatDTO struct {
	FriendshipId string `json:"friendshipId" binding:"required"`
	From         string `json:"from" binding:"required"`
	To           string `json:"to" binding:"required"`
	Message      string `json:"message" binding:"required"`
}

type GetPersonalChatDTO struct {
	FriendshipId string `json:"friendshipId" binding:"required"`
}
