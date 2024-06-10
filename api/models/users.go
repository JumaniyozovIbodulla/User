package models

type location struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"long"`
}

type User struct {
	Id        int64    `json:"id"`
	FullName  string   `json:"full_name"`
	NickName  string   `json:"nick_name,omitempty"`
	Photo     string   `json:"photo"`
	Birthday  string   `json:"birthday"`
	Location  location `json:"location"`
	CreatedAt string   `json:"created_at"`
	DeletedAt string   `json:"deleted_at"`
	UpdatedAt string   `json:"updated_at"`
	CreatedBy string   `json:"created_by"`
	UpdatedBy string   `json:"updated_by"`
	DeletedBy string   `json:"deleted_by"`
}

type UpdateUser struct {
	Id        int64    `json:"id"`
	FullName  string   `json:"full_name"`
	NickName  string   `json:"nick_name,omitempty"`
	Photo     string   `json:"photo"`
	Birthday  string   `json:"birthday"`
	Location  location `json:"location"`
	UpdatedBy string   `json:"updated_by"`
}

type GetUser struct {
	FullName  string   `json:"full_name"`
	NickName  string   `json:"nick_name,omitempty"`
	Photo     string   `json:"photo,omitempty"`
	Birthday  string   `json:"birthday"`
	Location  location `json:"location"`
	CreatedAt string   `json:"created_at"`
	DeletedAt string   `json:"deleted_at"`
	UpdatedAt string   `json:"updated_at"`
	CreatedBy string   `json:"created_by"`
	UpdatedBy string   `json:"updated_by"`
	DeletedBy string   `json:"deleted_by"`
}

type UserList struct {
	FullName  string   `json:"full_name"`
	NickName  string   `json:"nick_name,omitempty"`
	Photo     string   `json:"photo"`
	Birthday  string   `json:"birthday"`
	Location  location `json:"location"`
	CreatedAt string   `json:"created_at,omitempty"`
	UpdatedAt string   `json:"updated_at,omitempty"`
}

type DeleteUser struct {
	Id        int64  `json:"id"`
	DeletedBy string `json:"deleted_by"`
}

type AddUser struct {
	FullName  string   `json:"full_name"`
	NickName  string   `json:"nick_name,omitempty"`
	Photo     string   `json:"photo"`
	Birthday  string   `json:"birthday"`
	Location  location `json:"location"`
	CreatedBy string   `json:"created_by"`
}

type AddUsers struct {
	Users []AddUser
}

type UpdateUsers struct {
	Users []UpdateUser
}

type DeleteUsers struct {
	Users []DeleteUser
}

type GetListRequest struct {
	Search string `json:"search"`
	Page   uint64 `json:"page"`
	Limit  uint64 `json:"limit"`
}

type GetListResponse struct {
	Users []UserList `json:"users"`
	Count int64      `json:"count"`
}

/*
Get User List . you can add user fields that omitted fields to users list for response  .

ex::: created_at,updated_at .  GET  api/?fields=created_at,updated_at

Sort User List . GET /api/sort?full_name:asc,id:desc
*/
