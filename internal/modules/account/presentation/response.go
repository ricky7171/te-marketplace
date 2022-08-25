package accountpresent

// ini lagi misahin object response. Pokoknya yang dikembaliin oleh app service itu tetep model, bukan loginResponse
// loginResponse ini dikembaliin oleh handler.
type LoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
	User         UserData
}

type UserData struct {
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}
