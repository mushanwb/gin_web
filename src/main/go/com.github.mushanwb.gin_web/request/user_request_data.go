package request

type RegisterRequestData struct {
	Name     string `form:"name" binding:"required,min=2,max=30"`
	Phone    string `form:"phone" binding:"required,len=11"`
	Password string `form:"password" binding:"required,min=6,alphanum"`
}
