package entities

type Account struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Email            string `json:"email"`
	Password         string
	TotalPostLike    int `json:"total_post_like"`
	TotalChatMessage int `json:"total_chat_message"`
}
