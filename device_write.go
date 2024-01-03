// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package fosite

import (
	"context"
	"net/http"
)

// TODO: Do documentation

func (f *Fosite) WriteDeviceResponse(ctx context.Context, rw http.ResponseWriter, requester DeviceRequester, responder DeviceResponder) {
	// Set custom headers, e.g. "X-MySuperCoolCustomHeader" or "X-DONT-CACHE-ME"...
	wh := rw.Header()
	rh := responder.GetHeader()
	for k := range rh {
		wh.Set(k, rh.Get(k))
	}

	rw.Header().Set("Content-Type", "application/json;charset=UTF-8")
	rw.Header().Set("Cache-Control", "no-store")
	rw.Header().Set("Pragma", "no-cache")

	deviceResponse := &DeviceResponse{
		deviceResponse{
			DeviceCode:              responder.GetDeviceCode(),
			UserCode:                responder.GetUserCode(),
			VerificationURI:         responder.GetVerificationURI(),
			VerificationURIComplete: responder.GetVerificationURIComplete(),
			ExpiresIn:               responder.GetExpiresIn(),
			Interval:                responder.GetInterval(),
		},
	}

	_ = deviceResponse.ToJson(rw)
}
