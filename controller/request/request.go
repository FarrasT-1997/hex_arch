package request

type PostAuthorRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Sex      string `json:"sex" form:"sex"`
	Password string `json:"password" form:"password"`
}

type EditAuthorRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Sex      string `json:"sex" form:"sex"`
	Password string `json:"password" form:"password"`
}

type LoginAuthorRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
