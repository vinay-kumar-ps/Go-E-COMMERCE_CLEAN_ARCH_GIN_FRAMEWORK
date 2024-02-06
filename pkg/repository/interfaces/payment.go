package interfaces


type OtpRepository interface {
	FindUserByMobileNumber(phone string) bool
	UserDetailsUsingPhone(phone string) (models.UserDetailsResponse, error)
}