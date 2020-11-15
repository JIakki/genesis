package controllers

import (
	"github.com/JIakki/genesis/modules/payment/services"
)

type GetButtonsResponse struct {
	Buttons []services.GetButtonResponse `json:"buttons"`
	AppLink string                       `json:"appLink"`
}

type ButtonsFormatter struct {
}

func (ctrl *ButtonsFormatter) Format(buttons []services.GetButtonResponse, appLink string) GetButtonsResponse {
	return GetButtonsResponse{
		Buttons: buttons,
		AppLink: appLink,
	}
}

func NewButtonsFormatter() *ButtonsFormatter {
	return &ButtonsFormatter{}
}
