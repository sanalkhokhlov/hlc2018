package store

type Account struct {
	ID    uint32 `json:"id"`
	Fname string `json:"fname,omitempty"`
	Sname string `json:"sname,omitempty"`
	Email string `json:"email" valid:"email"`
	// Interests []string `json:"interests,omitempty"`
	Status string `json:"status,omitempty"`
	// Premium   *struct {
	// 	Start  uint64 `json:"start,omitempty"`
	// 	Finish uint64 `json:"finish,omitempty"`
	// } `json:"premium,omitempty"`
	Sex string `json:"sex,omitempty"`
	// Phone   string `json:"phone,omitempty"`
	// Likes   []Like `json:"likes,omitempty"`
	// Birth   uint64 `json:"birth,omitempty"`
	City    string `json:"city,omitempty"`
	Country string `json:"country,omitempty"`
	// Joined  uint64 `json:"joined,omitempty"`
}

type Like struct {
	Ts uint64 `json:"ts" valid:"required"`
	ID uint32 `json:"id" valid:"required"`
}

type NewLike struct {
	Likee uint32 `json:"likee" valid:"required"`
	Liker uint32 `json:"liker" valid:"required"`
	Ts    uint64 `json:"ts" valid:"required"`
}

type NotFoundError struct{}

func (e *NotFoundError) Error() string {
	return ""
}

type BadRequestError struct{}

func (e *BadRequestError) Error() string {
	return ""
}

type FilterArgs struct {
	Limit int
	Parts [][][]byte
}
