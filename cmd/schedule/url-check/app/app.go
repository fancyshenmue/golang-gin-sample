package app

import (
	"fmt"
	customHttp "sre-app/pkg/http"
	customSlack "sre-app/pkg/slack"
)

func CheckURLStatus() {
	var (
		r []domain
		s string
	)

	for _, v := range URLList {
		req := &customHttp.HttpRequestHelper{
			Method:        "GET",
			URL:           v,
			RequestHeader: nil,
			RequestBody:   nil,
		}

		resp := req.HttpStatusCheck()
		if !resp {
			domain := domain{
				Domain: v,
				Status: false,
			}
			r = append(r, domain)
		}
	}

	if len(r) > 0 {
		lengthR := len(r) - 1
		for k, v := range r {
			if k != lengthR {
				s = fmt.Sprintf("%sDomain: %s\nStatus: %t\n", s, v.Domain, v.Status)
			} else if k == lengthR {
				s = fmt.Sprintf("%sDomain: %s\nStatus: %t", s, v.Domain, v.Status)
			}
		}
	}

	if len(s) > 0 {
		customSlack.SlackSendMessage(
			SlackTokenID,
			SlackChannelID,
			"*_App Endpoint Status_*",
			"Failed Endpoint",
			s,
		)
	}
}
