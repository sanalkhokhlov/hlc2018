package store

type Account struct {
	ID        uint32   `json:"id"`
	Fname     string   `json:"fname"`
	Sname     string   `json:"sname"`
	Email     string   `json:"email" valid:"email"`
	Interests []string `json:"interests"`
	Status    string   `json:"status" valid:"required"`
	Premium   struct {
		Start  uint64 `json:"start"`
		Finish uint64 `json:"finish"`
	} `json:"premium"`
	Sex     string `json:"sex" valid:"sex,required"`
	Phone   string `json:"phone"`
	Likes   []Like `json:"likes"`
	Birth   uint64 `json:"birth" valid:"required"`
	City    string `json:"city"`
	Country string `json:"country"`
	Joined  uint64 `json:"joined" valid:"required"`
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
	Pairs [][][]byte
}