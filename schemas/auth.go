package schemas

type LoginFormRequest struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

type ClientRegiterRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ClientRegiterResponse struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type ArrayClientRegisterResponse struct {
	Data []ClientRegiterResponse `json:"data"`
}
