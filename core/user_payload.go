/*
 * FileBox server 0.0.1
 * Copyright (c) 2020. Mikhail Lazarev
 */

package core

type (
	TokenPair struct {
		Access  string `json:"access"`
		Refresh string `json:"refresh"`
	}

	RefreshTokenReq struct {
		Token string `json:"refresh" binding:"required"`
	}

	UserRes struct {
		Name  string `json:"name"`
		Score string `json:"score"`
	}

	AppleCodeReq struct {
		Name string `json:"name"`
		Code string `json:"code"`
	}

	AppleAuthResponse struct {
		Email           string
		EmailIsVerified bool
		IsPrivateEmail  bool
	}
)
