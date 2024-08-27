package wsresponses

type userInfo struct {
	Username string `json:"username"`
	IsActive bool   `json:"is_active"`
}

type UsersInfoResponse struct {
	Users map[string]userInfo `json:"users"`
	Len   int                 `json:"len"`
}

func NewUsersInfoResponse(
	users []string,
	connectedUsers map[string]bool,
) UsersInfoResponse {
	res := make(map[string]userInfo, len(users))

	for _, user := range users {
		userInfoEntry := userInfo{
			Username: user,
			IsActive: false,
		}
		if _, ok := connectedUsers[user]; ok {
			userInfoEntry.IsActive = true
		}
		res[user] = userInfoEntry
	}

	return UsersInfoResponse{
		Users: res,
		Len:   len(users),
	}
}
