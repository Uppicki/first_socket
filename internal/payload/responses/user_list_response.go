package responses

import "first_socket/internal/domain"

type UserListResponse struct {
	Users map[string]*domain.User `json:"users"`
}
