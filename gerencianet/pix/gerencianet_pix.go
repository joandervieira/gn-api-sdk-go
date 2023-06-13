package pix

type PixClient struct {
	Endpoints
}

func NewGerencianet(configs map[string]interface{}) *PixClient {
	clientID := configs["client_id"].(string)
	clientSecret := configs["client_secret"].(string)
	CA := configs["CA"].(string)
	Key := configs["Key"].(string)
	sandbox := configs["sandbox"].(bool)
	timeout := configs["timeout"].(int)

	requester := NewRequester(clientID, clientSecret, CA, Key, sandbox, timeout)
	gn := PixClient{}
	gn.Requester = *requester
	return &gn
}
