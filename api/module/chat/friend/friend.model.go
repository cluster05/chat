package friend

type Friendship struct {
	FriendshipId string `bson:"friendshipId" json:"friendshipId"`
	MeId         string `bson:"meId" json:"meId"`
	FriendId     string `bson:"friendId" json:"friendId"`
	CreatedAt    int64  `bson:"createdAt" json:"createdAt"`
	UpdatedAt    int64  `bson:"updatedAt" json:"updatedAt"`
	IsDeleted    bool   `bson:"isDeleted" json:"isDeleted"`
}

type CreateFriendshipDTO struct {
	MeId     string `json:"meId"`
	FriendId string `json:"friendId"`
}

type GetFriendListDTO struct {
	MeId string `json:"meId"`
}

type DeleteFriendshipDTO struct {
	FriendshipId string `json:"friendshipId"`
}
