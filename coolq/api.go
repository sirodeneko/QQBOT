package coolq

type Api string

type ApiBase struct {
	Action Api         `json:"action"`
	Params interface{} `json:"params"`
	Echo   string      `json:"echo"`
}
