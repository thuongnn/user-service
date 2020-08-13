package models

// AuthModel holds information used to authenticate.
type AuthModel struct {
	Username string
	Password string
}

type AccessDetails struct {
	AccessUUID string
	UserID     uint64
}

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUUID   string
	RefreshUUID  string
	AtExpires    int64
	RtExpires    int64
}
