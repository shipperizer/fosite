/*
 * Copyright © 2015-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * @author		Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @copyright 	2015-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @license 	Apache-2.0
 *
 */

package oauth2

import (
	"context"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/ory/fosite"
	"github.com/ory/fosite/token/jwt"
	"github.com/ory/x/errorsx"
)

// DefaultJWTStrategy is a JWT RS256 strategy.
type DefaultJWTStrategy struct {
	jwt.Signer
	HMACSHAStrategy *HMACSHAStrategy
	Config          interface {
		fosite.AccessTokenIssuerProvider
		fosite.JWTScopeFieldProvider
	}
}

func (h DefaultJWTStrategy) signature(token string) string {
	split := strings.Split(token, ".")
	if len(split) != 3 {
		return ""
	}

	return split[2]
}

func (h DefaultJWTStrategy) AccessTokenSignature(ctx context.Context, token string) string {
	return h.signature(token)
}

func (h *DefaultJWTStrategy) GenerateAccessToken(ctx context.Context, requester fosite.Requester) (token string, signature string, err error) {
	return h.generate(ctx, fosite.AccessToken, requester)
}

func (h *DefaultJWTStrategy) ValidateAccessToken(ctx context.Context, _ fosite.Requester, token string) error {
	_, err := validate(ctx, h.Signer, token)
	return err
}

func (h DefaultJWTStrategy) RefreshTokenSignature(ctx context.Context, token string) string {
	return h.HMACSHAStrategy.RefreshTokenSignature(ctx, token)
}

func (h DefaultJWTStrategy) AuthorizeCodeSignature(ctx context.Context, token string) string {
	return h.HMACSHAStrategy.AuthorizeCodeSignature(ctx, token)
}

func (h *DefaultJWTStrategy) GenerateRefreshToken(ctx context.Context, req fosite.Requester) (token string, signature string, err error) {
	return h.HMACSHAStrategy.GenerateRefreshToken(ctx, req)
}

func (h *DefaultJWTStrategy) ValidateRefreshToken(ctx context.Context, req fosite.Requester, token string) error {
	return h.HMACSHAStrategy.ValidateRefreshToken(ctx, req, token)
}

func (h *DefaultJWTStrategy) GenerateAuthorizeCode(ctx context.Context, req fosite.Requester) (token string, signature string, err error) {
	return h.HMACSHAStrategy.GenerateAuthorizeCode(ctx, req)
}

func (h *DefaultJWTStrategy) ValidateAuthorizeCode(ctx context.Context, req fosite.Requester, token string) error {
	return h.HMACSHAStrategy.ValidateAuthorizeCode(ctx, req, token)
}

func validate(ctx context.Context, jwtStrategy jwt.Signer, token string) (t *jwt.Token, err error) {
	t, err = jwtStrategy.Decode(ctx, token)
	if err == nil {
		err = t.Claims.Valid()
		return
	}

	var e *jwt.ValidationError
	if err != nil && errors.As(err, &e) {
		err = errorsx.WithStack(toRFCErr(e).WithWrap(err).WithDebug(err.Error()))
	}

	return
}

func toRFCErr(v *jwt.ValidationError) *fosite.RFC6749Error {
	switch {
	case v == nil:
		return nil
	case v.Has(jwt.ValidationErrorMalformed):
		return fosite.ErrInvalidTokenFormat
	case v.Has(jwt.ValidationErrorUnverifiable | jwt.ValidationErrorSignatureInvalid):
		return fosite.ErrTokenSignatureMismatch
	case v.Has(jwt.ValidationErrorExpired):
		return fosite.ErrTokenExpired
	case v.Has(jwt.ValidationErrorAudience |
		jwt.ValidationErrorIssuedAt |
		jwt.ValidationErrorIssuer |
		jwt.ValidationErrorNotValidYet |
		jwt.ValidationErrorId |
		jwt.ValidationErrorClaimsInvalid):
		return fosite.ErrTokenClaim
	default:
		return fosite.ErrRequestUnauthorized
	}
}

func (h *DefaultJWTStrategy) generate(ctx context.Context, tokenType fosite.TokenType, requester fosite.Requester) (string, string, error) {
	if jwtSession, ok := requester.GetSession().(JWTSessionContainer); !ok {
		return "", "", errors.Errorf("Session must be of type JWTSessionContainer but got type: %T", requester.GetSession())
	} else if jwtSession.GetJWTClaims() == nil {
		return "", "", errors.New("GetTokenClaims() must not be nil")
	} else {
		claims := jwtSession.GetJWTClaims().
			With(
				jwtSession.GetExpiresAt(tokenType),
				requester.GetGrantedScopes(),
				requester.GetGrantedAudience(),
			).
			WithDefaults(
				time.Now().UTC(),
				h.Config.GetAccessTokenIssuer(ctx),
			).
			WithScopeField(
				h.Config.GetJWTScopeField(ctx),
			)

		return h.Signer.Generate(ctx, claims.ToMapClaims(), jwtSession.GetJWTHeader())
	}
}
