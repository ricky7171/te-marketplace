package accountpresent

// ini harusnya ada object yang merepresentasikan request, lalu ada validasi bawaaan gin juga
type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
