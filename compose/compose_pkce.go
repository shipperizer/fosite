// Copyright © 2022 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package compose

import (
	"github.com/ory/fosite"
	"github.com/ory/fosite/handler/oauth2"
	"github.com/ory/fosite/handler/pkce"
	"github.com/ory/fosite/handler/rfc8628"
)

// OAuth2PKCEFactory creates a PKCE handler.
func OAuth2PKCEFactory(config fosite.Configurator, storage interface{}, strategy interface{}) interface{} {
	return &pkce.Handler{
		AuthorizeCodeStrategy: strategy.(oauth2.AuthorizeCodeStrategy),
		DeviceCodeStrategy:    strategy.(rfc8628.DeviceCodeStrategy),
		Storage:               storage.(pkce.PKCERequestStorage),
		Config:                config,
	}
}
