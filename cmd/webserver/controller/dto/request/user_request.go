package request

type UserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,containsany=@#$%*()_+!;:?"`
	Name     string `json:"name" binding:"required,min=3,max=80"`
	Age      int8   `json:"age" binding:"gte=0,lte=140"`
}

type UserUpdateRequest struct {
	Name string `json:"name" binding:"omitempty,min=3,max=80"`
	Age  int8   `json:"age" binding:"omitempty,gte=0,lte=140"`
}
