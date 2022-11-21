package models

type Student struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	UserName    string `json:"user_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	CreatedAt   string `json:"created_at"`
}

type GetStudentsQueryParam struct {
	FirstName string
	LastName  string
	UserName  string
	Page   int64
	Limit  int64
}

type GetStudentResult struct {
	Students []*Student `json:"student"`
	Count    int64      `json:"count"`
}

type ResponseError struct {
	Message string `json:"message"`
}

type ResponseOK struct {
	Message string `json:"message"`
}
type CreateStudentRequest struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	UserName    string `json:"user_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

type UpdateStudentRequest struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	UserName    string `json:"user_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}
