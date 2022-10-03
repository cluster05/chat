package friend

type Friendship struct {
	FriendshipId string `bson:"friendshipId" json:"friendshipId"`
	MeId         string `bson:"meId" json:"meId"`
	FriendId     string `bson:"friendId" json:"friendId"`
	CreatedAt    int64  `bson:"createdAt" json:"createdAt"`
	UpdatedAt    int64  `bson:"updatedAt" json:"updatedAt"`
	IsDeleted    bool   `bson:"isDeleted" json:"isDeleted"`
}

type FriendshipResult struct {
	FriendshipId string `bson:"friendshipId" json:"friendshipId"`
	MeId         string `bson:"meId" json:"meId"`
	FriendId     string `bson:"friendId" json:"friendId"`
	IsDeleted    bool   `bson:"isDeleted" json:"isDeleted"`
	FriendName   string `bson:"friendName" json:"friendName"`
}

type CreateFriendshipDTO struct {
	MeId     string `json:"meId" binding:"required"`
	FriendId string `json:"friendId" binding:"required"`
}

type GetFriendListDTO struct {
	MeId string `json:"meId" binding:"required"`
}

type DeleteFriendshipDTO struct {
	FriendshipId string `json:"friendshipId" binding:"required"`
}

type SearchFriendshipDTO struct {
	Filter string `json:"filter" binding:"required"`
}

type Search struct {
	AuthId   string `bson:"authId" json:"authId"`
	Username string `bson:"username" json:"username"`
}
