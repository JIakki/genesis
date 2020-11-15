package services

type GetButtonResponse struct {
	Url  string `json:"url"`
	Name string `json:"name"`
}

type ButtonFormatter struct {
}

func (ctrl *ButtonFormatter) Format(button *PaymentButton) GetButtonResponse {
	return GetButtonResponse{
		Url:  button.URL,
		Name: button.Name,
	}
}

func NewButtonFormatter() *ButtonFormatter {
	return &ButtonFormatter{}
}
