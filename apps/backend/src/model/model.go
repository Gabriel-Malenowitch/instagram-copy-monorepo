package model

type User struct {
	Id         string
	User       string
	Created_at string
	Updated_at string
}

type Login struct {
	User_id string
}

type UserInsertPayload struct {
	User string `validate:"required"`
}

type UserUpdatePayload struct {
	User string
}

type UserResponse struct {
	Id              string
	User            string
	Created_at      string
	Updated_at      string
	FollowersLength string
	FollowingLength string
	Img             string
}

type UserConfiguration struct {
	User_id   string
	Font_size int
	Theme     string
}

type UserConfigurationPayload struct {
	Font_size int
	Theme     string
}

type Followers struct {
	User_id    string
	User_refer string
}

type Post struct {
	Id         string
	User_id    string
	Post_refer string
	Created_at string
	Updated_at string
	Likes      int
}

type Comment struct {
	Post_id string
	Comment string
	Likes   int
}
