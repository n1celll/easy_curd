package serializer

import (
	"easy_curd/model"
	"easy_curd/util"
)

// User 用户序列化器
type User struct {
	ID        uint   `json:"id"`
	UserName  string `json:"username"`
	Nickname  string `json:"nickname"`
	Avatar    string `json:"avatar"`
	CreatedAt int64  `json:"created_at"`
}

type UserWithToken struct {
	UserInfo User `json:"user_info"`
	Token string `json:"token"`
}


// BuildUser 序列化用户
func BuildUser(user model.User) User {
	return User{
		ID:        user.ID,
		UserName:  user.UserName,
		Nickname:  user.Nickname,
		Avatar:    user.Avatar,
		CreatedAt: user.CreatedAt.Unix(),
	}
}

// 返回带token的用户信息
func BuildUserWithTokenResponse(user model.User) Response  {
	return Response{
		Status: 200,
		Data: UserWithToken{
			UserInfo: BuildUser(user),
			Token: util.GenJwt(user.ID),
		},
	}
}

// BuildUserResponse 序列化用户响应
func BuildUserResponse(user model.User) Response {
	return Response{
		Data: BuildUser(user),
	}
}
