package value

type AccountCreationRequest struct {
	AccountName string `json:"account_name" valid:"alphanum,required~account_name required and must be alphanumeric"`
	FirstName   string `json:"first_name" valid:"alpha,required~first_name required and must be alphabetic"`
	LastName    string `json:"last_name" valid:"alpha,required~last_name required and must be alphabetic"`
	Email       string `json:"email" valid:"email,required~email required and must be a valid email address"`
}
