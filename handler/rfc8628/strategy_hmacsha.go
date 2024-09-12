// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package rfc8628

import (
	"context"

	"github.com/ory/x/randx"
	"github.com/patrickmn/go-cache"
	"golang.org/x/time/rate"

	"github.com/ory/fosite"
	enigma "github.com/ory/fosite/token/hmac"
)

type DefaultDeviceStrategy struct {
	Enigma           *enigma.HMACStrategy
	RateLimiterCache *cache.Cache
	Config           interface {
		fosite.DeviceProvider
		fosite.DeviceAndUserCodeLifespanProvider
	}
}

var _ RFC8628CodeStrategy = (*DefaultDeviceStrategy)(nil)

func (h *DefaultDeviceStrategy) GenerateUserCode(ctx context.Context) (token string, signature string, err error) {
	seq, err := randx.RuneSequence(8, []rune(randx.AlphaUpper))
	if err != nil {
		return "", "", err
	}
	userCode := string(seq)
	signUserCode, signErr := h.UserCodeSignature(ctx, userCode)
	if signErr != nil {
		return "", "", err
	}
	return userCode, signUserCode, nil
}

func (h *DefaultDeviceStrategy) UserCodeSignature(ctx context.Context, token string) (signature string, err error) {
	return h.Enigma.Signature(token), nil
}

func (h *DefaultDeviceStrategy) ValidateUserCode(ctx context.Context, r fosite.Requester, code string) (err error) {
	// TODO
	return nil
}

func (h *DefaultDeviceStrategy) GenerateDeviceCode(ctx context.Context) (token string, signature string, err error) {
	token, sig, err := h.Enigma.Generate(ctx)
	if err != nil {
		return "", "", err
	}

	return "ory_dc_" + token, sig, nil
}

func (h *DefaultDeviceStrategy) DeviceCodeSignature(ctx context.Context, token string) (signature string, err error) {
	return h.Enigma.Signature(token), nil
}

func (h *DefaultDeviceStrategy) ValidateDeviceCode(ctx context.Context, r fosite.Requester, code string) (err error) {
	// TODO
	return nil
}

func (t *DefaultDeviceStrategy) ShouldRateLimit(context context.Context, code string) bool {
	key := code + "_limiter"

	if x, found := t.RateLimiterCache.Get(key); found {
		return !x.(*rate.Limiter).Allow()
	}

	rateLimiter := rate.NewLimiter(
		rate.Every(
			t.Config.GetDeviceAuthTokenPollingInterval(context),
		),
		1,
	)

	t.RateLimiterCache.Set(key, rateLimiter, cache.DefaultExpiration)
	return false
}
