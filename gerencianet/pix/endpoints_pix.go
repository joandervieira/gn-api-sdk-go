package pix

type Endpoints struct {
	Requester interface {
		Request(endpoint string, httpVerb string, requestParams map[string]string, body map[string]interface{}) (string, error)
		RequestWithHeaders(endpoint string, httpVerb string, requestParams map[string]string, body map[string]interface{}, headers map[string]string) (string, error)
	}
}

func (e Endpoints) CreateImmediateCharge(body map[string]interface{}) (string, error) {
	return e.Requester.Request("/v2/cob", "POST", nil, body)
}

func (e Endpoints) CreateCharge(txid string, body map[string]interface{}) (string, error) {
	params := map[string]string{"txid": (txid)}
	return e.Requester.Request("/v2/cob/:txid", "PUT", params, body)
}

func (e Endpoints) UpdateCharge(txid string) (string, error) {
	params := map[string]string{"txid": (txid)}
	return e.Requester.Request("/v2/cob/:txid", "PATCH", params, nil)
}

func (e Endpoints) DetailCharge(txid string) (string, error) {
	params := map[string]string{"txid": (txid)}
	return e.Requester.Request("/v2/cob/:txid", "GET", params, nil)
}

func (e Endpoints) ListCharges(inicio string, fim string) (string, error) {
	params := map[string]string{
		"inicio": (inicio),
		"fim":    (fim),
	}
	return e.Requester.Request("/v2/cob?inicio=:inicio&fim=:fim", "GET", params, nil)
}

func (e Endpoints) PixDevolution(e2eid string, id string, body map[string]interface{}) (string, error) {
	params := map[string]string{
		"e2eid": (e2eid),
		"id":    (id)}
	return e.Requester.Request("/v2/pix/:e2eid/devolucao/:id", "PUT", params, body)
}

func (e Endpoints) PixDetailDevolution(e2eid string, id string) (string, error) {
	params := map[string]string{
		"e2eid": (e2eid),
		"id":    (id)}
	return e.Requester.Request("/v2/pix/:e2eid/devolucao/:id", "GET", params, nil)
}

func (e Endpoints) PixSend(idEnvio string, body map[string]interface{}) (string, error) {
	params := map[string]string{"idEnvio": (idEnvio)}
	return e.Requester.Request("/v2/gn/pix/:idEnvio", "PUT", params, body)
}

func (e Endpoints) PixSendList(e2eid string) (string, error) {
	params := map[string]string{"e2eid": (e2eid)}
	return e.Requester.Request("/v2/pix/:e2eid", "GET", params, nil)
}

func (e Endpoints) PixSendDetail(e2eid string) (string, error) {
	params := map[string]string{"e2eid": (e2eid)}
	return e.Requester.Request("/v2/gn/pix/enviados/:e2eid", "GET", params, nil)
}

func (e Endpoints) PixCreateLocation(body map[string]interface{}) (string, error) {
	return e.Requester.Request("/v2/loc", "POST", nil, body)
}

func (e Endpoints) PixUnlinkTxidLocation(id string, body map[string]interface{}) (string, error) {
	params := map[string]string{"id": (id)}
	return e.Requester.Request("/v2/loc/:id/txid", "DELETE", params, body)
}

func (e Endpoints) PixDetailLocation(id string) (string, error) {
	params := map[string]string{"id": (id)}
	return e.Requester.Request("/v2/loc/:id", "GET", params, nil)
}

func (e Endpoints) PixLocationList(inicio string, fim string) (string, error) {
	params := map[string]string{
		"inicio": (inicio),
		"fim":    (fim),
	}
	return e.Requester.Request("/v2/loc?inicio=:inicio&fim=:fim", "GET", params, nil)
}

func (e Endpoints) PixGenerateQRCode(id string) (string, error) {
	params := map[string]string{"id": (id)}
	return e.Requester.Request("/v2/loc/:id/qrcode", "GET", params, nil)
}

func (e Endpoints) GetAccountBalance(body map[string]interface{}) (string, error) {
	return e.Requester.Request("/v2/gn/saldo", "GET", nil, body)
}

func (e Endpoints) ListAccountConfig(body map[string]interface{}) (string, error) {
	return e.Requester.Request("/v2/gn/config", "GET", nil, body)
}

func (e Endpoints) UpdateAccountConfig(body map[string]interface{}) (string, error) {
	return e.Requester.Request("/v2/gn/config", "PUT", nil, body)
}

func (e Endpoints) PixCreateEvp(body map[string]interface{}) (string, error) {
	return e.Requester.Request("/v2/gn/evp", "POST", nil, body)
}

func (e Endpoints) PixListEvp(body map[string]interface{}) (string, error) {
	return e.Requester.Request("/v2/gn/evp", "GET", nil, body)
}

func (e Endpoints) PixDeleteEvp(key string, body map[string]interface{}) (string, error) {
	params := map[string]string{"key": (key)}
	return e.Requester.Request("/v2/gn/evp/:key", "DELETE", params, body)
}

func (e Endpoints) PixConfigWebhook(chave string, body map[string]interface{}, headers map[string]string) (string, error) {
	params := map[string]string{"chave": (chave)}
	return e.Requester.RequestWithHeaders("/v2/webhook/:chave", "PUT", params, body, headers)
}

func (e Endpoints) PixDeleteWebhook(chave string, body map[string]interface{}) (string, error) {
	params := map[string]string{"chave": (chave)}
	return e.Requester.Request("/v2/webhook/:chave", "DELETE", params, body)
}

func (e Endpoints) PixDetailWebhook(chave string) (string, error) {
	params := map[string]string{"chave": (chave)}
	return e.Requester.Request("/v2/webhook/:chave", "GET", params, nil)
}

func (e Endpoints) PixListWebhooks(inicio string, fim string) (string, error) {
	params := map[string]string{
		"inicio": (inicio),
		"fim":    (fim),
	}
	return e.Requester.Request("/v2/webhook?inicio=:inicio&fim=:fim", "GET", params, nil)
}

func (e Endpoints) CreateDueCharge(txid string, body map[string]interface{}) (string, error) {
	params := map[string]string{"txid": (txid)}
	return e.Requester.Request("/v2/cobv/:txid", "PUT", params, body)
}

func (e Endpoints) PixUpdateDueCharge(txid string, body map[string]interface{}) (string, error) {
	params := map[string]string{"txid": (txid)}
	return e.Requester.Request("/v2/cobv/:txid", "PATCH", params, body)
}

func (e Endpoints) DetailDueCharge(txid string) (string, error) {
	params := map[string]string{"txid": (txid)}
	return e.Requester.Request("/v2/cobv/:txid", "GET", params, nil)
}

func (e Endpoints) PixListDueCharges(params map[string]string) (string, error) {
	return e.Requester.Request("/v2/cobv", "GET", params, nil)
}

func (e Endpoints) CreateReport(body map[string]interface{}) (string, error) {
	return e.Requester.Request("/v2/gn/relatorios/extrato-conciliacao", "POST", nil, body)
}

func (e Endpoints) DetailReport(id string) (string, error) {
	params := map[string]string{"id": (id)}
	return e.Requester.Request("/v2/gn/relatorios/:id", "GET", params, nil)
}

func (e Endpoints) PixReceivedList(params map[string]string) (string, error) {
	return e.Requester.Request("/v2/pix", "GET", params, nil)
}

func (e Endpoints) PixDetailReceived(e2eid string) (string, error) {
	params := map[string]string{"e2eid": (e2eid)}
	return e.Requester.Request("/v2/pix/:e2eid", "GET", params, nil)
}

func (e Endpoints) PixSplitDetailCharge(txid string) (string, error) {
	params := map[string]string{"txid": (txid)}
	return e.Requester.Request("/v2/gn/split/cob/:txid", "GET", params, nil)
}

func (e Endpoints) PixSplitLinkCharge(txid string, splitConfigId string) (string, error) {
	params := map[string]string{
		"txid":          (txid),
		"splitConfigId": (splitConfigId),
	}
	return e.Requester.Request("/v2/gn/split/cob/:txid/vinculo/:splitConfigId", "PUT", params, nil)
}

func (e Endpoints) PixSplitUnlinkCharge(txid string, splitConfigId string) (string, error) {
	params := map[string]string{
		"txid":          (txid),
		"splitConfigId": (splitConfigId),
	}
	return e.Requester.Request("/v2/gn/split/cob/:txid/vinculo/:splitConfigId", "DELETE", params, nil)
}

func (e Endpoints) PixSplitDetailDueCharge(txid string) (string, error) {
	params := map[string]string{"txid": (txid)}
	return e.Requester.Request("/v2/gn/split/cobv/:txid", "GET", params, nil)
}

func (e Endpoints) PixSplitLinkDueCharge(txid string, splitConfigId string) (string, error) {
	params := map[string]string{
		"txid":          (txid),
		"splitConfigId": (splitConfigId),
	}
	return e.Requester.Request("/v2/gn/split/cobv/:txid/vinculo/:splitConfigId", "PUT", params, nil)
}

func (e Endpoints) PixSplitUnlinkDueCharge(txid string, splitConfigId string) (string, error) {
	params := map[string]string{
		"txid":          (txid),
		"splitConfigId": (splitConfigId),
	}
	return e.Requester.Request("/v2/gn/split/cobv/:txid/vinculo/:splitConfigId", "DELETE", params, nil)
}

func (e Endpoints) PixSplitConfig(body map[string]interface{}) (string, error) {
	return e.Requester.Request("/v2/gn/split/config", "POST", nil, body)
}

func (e Endpoints) PixSplitConfigId(id string, body map[string]interface{}) (string, error) {
	params := map[string]string{"id": (id)}
	return e.Requester.Request("/v2/gn/split/config/:id", "PUT", params, body)
}

func (e Endpoints) PixSplitDetailConfig(id string) (string, error) {
	params := map[string]string{"id": (id)}
	return e.Requester.Request("/v2/gn/split/config/:id", "GET", params, nil)
}
