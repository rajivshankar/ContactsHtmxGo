package contacts

type Contact struct {
	Id    int    `form:"id"`
	Name  string `form:"name" binding:"required"`
	Email string `form:"email" binding:"required"`
}

type Contacts []*Contact
