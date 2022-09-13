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
	"fmt"
	"net/url"
	"testing" //"time"

	"github.com/golang/mock/gomock"

	"github.com/ory/fosite/internal"

	//"github.com/golang/mock/gomock"
	"time"

	"github.com/ory/fosite" //"github.com/ory/fosite/internal"
	"github.com/ory/fosite/storage"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDeviceAuthorizeCode_PopulateTokenEndpointResponse(t *testing.T) {
	for k, strategy := range map[string]CoreStrategy{
		"hmac": &hmacshaStrategy,
	} {
		t.Run("strategy="+k, func(t *testing.T) {
			store := storage.NewMemoryStore()

			var h DeviceAuthorizationHandler
			for _, c := range []struct {
				areq        *fosite.AccessRequest
				description string
				setup       func(t *testing.T, areq *fosite.AccessRequest, config *fosite.Config)
				check       func(t *testing.T, aresp *fosite.AccessResponse)
				expectErr   error
			}{
				{
					areq: &fosite.AccessRequest{
						GrantTypes: fosite.Arguments{"123"},
					},
					description: "should fail because not responsible",
					expectErr:   fosite.ErrUnknownRequest,
				},
				{
					areq: &fosite.AccessRequest{
						GrantTypes: fosite.Arguments{"urn:ietf:params:oauth:grant-type:device_code"},
						Request: fosite.Request{
							Form: url.Values{},
							Client: &fosite.DefaultClient{
								GrantTypes: fosite.Arguments{"urn:ietf:params:oauth:grant-type:device_code"},
							},
							Session:     &fosite.DefaultSession{},
							RequestedAt: time.Now().UTC(),
						},
					},
					description: "should fail because device code not found",
					setup: func(t *testing.T, areq *fosite.AccessRequest, config *fosite.Config) {
						code, _, err := strategy.GenerateDeviceCode(context.TODO())
						require.NoError(t, err)
						areq.Form.Set("device_code", code)
					},
					expectErr: fosite.ErrServerError,
				},
				{
					areq: &fosite.AccessRequest{
						GrantTypes: fosite.Arguments{"urn:ietf:params:oauth:grant-type:device_code"},
						Request: fosite.Request{
							Form: url.Values{},
							Client: &fosite.DefaultClient{
								GrantTypes: fosite.Arguments{"urn:ietf:params:oauth:grant-type:device_code", "refresh_token"},
							},
							GrantedScope: fosite.Arguments{"foo", "offline"},
							Session:      &fosite.DefaultSession{},
							RequestedAt:  time.Now().UTC(),
						},
					},
					setup: func(t *testing.T, areq *fosite.AccessRequest, config *fosite.Config) {
						code, sig, err := strategy.GenerateDeviceCode(context.TODO())
						require.NoError(t, err)
						areq.Form.Add("device_code", code)

						require.NoError(t, store.CreateDeviceCodeSession(context.TODO(), sig, areq))
					},
					description: "should pass with offline scope and refresh token",
					check: func(t *testing.T, aresp *fosite.AccessResponse) {
						assert.NotEmpty(t, aresp.AccessToken)
						assert.Equal(t, "bearer", aresp.TokenType)
						assert.NotEmpty(t, aresp.GetExtra("refresh_token"))
						assert.NotEmpty(t, aresp.GetExtra("expires_in"))
						assert.Equal(t, "foo offline", aresp.GetExtra("scope"))
					},
				},
				{
					areq: &fosite.AccessRequest{
						GrantTypes: fosite.Arguments{"urn:ietf:params:oauth:grant-type:device_code"},
						Request: fosite.Request{
							Form: url.Values{},
							Client: &fosite.DefaultClient{
								GrantTypes: fosite.Arguments{"urn:ietf:params:oauth:grant-type:device_code", "refresh_token"},
							},
							GrantedScope: fosite.Arguments{"foo"},
							Session:      &fosite.DefaultSession{},
							RequestedAt:  time.Now().UTC(),
						},
					},
					setup: func(t *testing.T, areq *fosite.AccessRequest, config *fosite.Config) {
						config.RefreshTokenScopes = []string{}
						code, sig, err := strategy.GenerateDeviceCode(context.TODO())
						require.NoError(t, err)
						areq.Form.Add("device_code", code)

						require.NoError(t, store.CreateDeviceCodeSession(context.TODO(), sig, areq))
					},
					description: "should pass with refresh token always provided",
					check: func(t *testing.T, aresp *fosite.AccessResponse) {
						assert.NotEmpty(t, aresp.AccessToken)
						assert.Equal(t, "bearer", aresp.TokenType)
						assert.NotEmpty(t, aresp.GetExtra("refresh_token"))
						assert.NotEmpty(t, aresp.GetExtra("expires_in"))
						assert.Equal(t, "foo", aresp.GetExtra("scope"))
					},
				},
				{
					areq: &fosite.AccessRequest{
						GrantTypes: fosite.Arguments{"urn:ietf:params:oauth:grant-type:device_code"},
						Request: fosite.Request{
							Form: url.Values{},
							Client: &fosite.DefaultClient{
								GrantTypes: fosite.Arguments{"urn:ietf:params:oauth:grant-type:device_code"},
							},
							GrantedScope: fosite.Arguments{},
							Session:      &fosite.DefaultSession{},
							RequestedAt:  time.Now().UTC(),
						},
					},
					setup: func(t *testing.T, areq *fosite.AccessRequest, config *fosite.Config) {
						config.RefreshTokenScopes = []string{}
						code, sig, err := strategy.GenerateDeviceCode(context.TODO())
						require.NoError(t, err)
						areq.Form.Add("device_code", code)

						require.NoError(t, store.CreateDeviceCodeSession(context.TODO(), sig, areq))
					},
					description: "should pass with no refresh token",
					check: func(t *testing.T, aresp *fosite.AccessResponse) {
						assert.NotEmpty(t, aresp.AccessToken)
						assert.Equal(t, "bearer", aresp.TokenType)
						assert.Empty(t, aresp.GetExtra("refresh_token"))
						assert.NotEmpty(t, aresp.GetExtra("expires_in"))
						assert.Empty(t, aresp.GetExtra("scope"))
					},
				},
				{
					areq: &fosite.AccessRequest{
						GrantTypes: fosite.Arguments{"urn:ietf:params:oauth:grant-type:device_code"},
						Request: fosite.Request{
							Form: url.Values{},
							Client: &fosite.DefaultClient{
								GrantTypes: fosite.Arguments{"urn:ietf:params:oauth:grant-type:device_code"},
							},
							GrantedScope: fosite.Arguments{"foo"},
							Session:      &fosite.DefaultSession{},
							RequestedAt:  time.Now().UTC(),
						},
					},
					setup: func(t *testing.T, areq *fosite.AccessRequest, config *fosite.Config) {
						code, sig, err := strategy.GenerateDeviceCode(context.TODO())
						require.NoError(t, err)
						areq.Form.Add("device_code", code)

						require.NoError(t, store.CreateDeviceCodeSession(context.TODO(), sig, areq))
					},
					description: "should not have refresh token",
					check: func(t *testing.T, aresp *fosite.AccessResponse) {
						assert.NotEmpty(t, aresp.AccessToken)
						assert.Equal(t, "bearer", aresp.TokenType)
						assert.Empty(t, aresp.GetExtra("refresh_token"))
						assert.NotEmpty(t, aresp.GetExtra("expires_in"))
						assert.Equal(t, "foo", aresp.GetExtra("scope"))
					},
				},
			} {
				t.Run("case="+c.description, func(t *testing.T) {
					config := &fosite.Config{
						ScopeStrategy:            fosite.HierarchicScopeStrategy,
						AudienceMatchingStrategy: fosite.DefaultAudienceMatchingStrategy,
						AccessTokenLifespan:      time.Minute,
						RefreshTokenScopes:       []string{"offline"},
					}
					h = DeviceAuthorizationHandler{
						CoreStorage:          store,
						DeviceCodeStrategy:   strategy,
						UserCodeStrategy:     strategy,
						AccessTokenStrategy:  strategy,
						RefreshTokenStrategy: strategy,
						Config:               config,
					}

					if c.setup != nil {
						c.setup(t, c.areq, config)
					}

					aresp := fosite.NewAccessResponse()
					err := h.PopulateTokenEndpointResponse(context.TODO(), c.areq, aresp)

					if c.expectErr != nil {
						require.EqualError(t, err, c.expectErr.Error(), "%+v", err)
					} else {
						require.NoError(t, err, "%+v", err)
					}

					if c.check != nil {
						c.check(t, aresp)
					}
				})
			}
		})
	}
}

func TestDeviceAuthorizeCode_HandleTokenEndpointRequest(t *testing.T) {
	for k, strategy := range map[string]CoreStrategy{
		"hmac": &hmacshaStrategy,
	} {
		t.Run("strategy="+k, func(t *testing.T) {
			store := storage.NewMemoryStore()

			h := DeviceAuthorizationHandler{
				CoreStorage:        store,
				DeviceCodeStrategy: &hmacshaStrategy,
				UserCodeStrategy:   &hmacshaStrategy,
				Config: &fosite.Config{
					ScopeStrategy:            fosite.HierarchicScopeStrategy,
					AudienceMatchingStrategy: fosite.DefaultAudienceMatchingStrategy,
					AuthorizeCodeLifespan:    time.Minute,
				},
			}
			for i, c := range []struct {
				areq        *fosite.AccessRequest
				authreq     *fosite.DeviceAuthorizeRequest
				description string
				setup       func(t *testing.T, areq *fosite.AccessRequest, authreq *fosite.DeviceAuthorizeRequest)
				check       func(t *testing.T, areq *fosite.AccessRequest, authreq *fosite.DeviceAuthorizeRequest)
				expectErr   error
			}{
				{
					areq: &fosite.AccessRequest{
						GrantTypes: fosite.Arguments{"12345678"},
					},
					description: "should fail because not responsible",
					expectErr:   fosite.ErrUnknownRequest,
				},
				{
					areq: &fosite.AccessRequest{
						GrantTypes: fosite.Arguments{"urn:ietf:params:oauth:grant-type:device_code"},
						Request: fosite.Request{
							Client:      &fosite.DefaultClient{ID: "foo", GrantTypes: []string{""}},
							Session:     &fosite.DefaultSession{},
							RequestedAt: time.Now().UTC(),
						},
					},
					description: "should fail because client is not granted this grant type",
					expectErr:   fosite.ErrUnauthorizedClient,
				},
				{
					areq: &fosite.AccessRequest{
						GrantTypes: fosite.Arguments{"urn:ietf:params:oauth:grant-type:device_code"},
						Request: fosite.Request{
							Client:      &fosite.DefaultClient{GrantTypes: []string{"urn:ietf:params:oauth:grant-type:device_code"}},
							Session:     &fosite.DefaultSession{},
							RequestedAt: time.Now().UTC(),
						},
					},
					description: "should fail because device code could not be retrieved",
					setup: func(t *testing.T, areq *fosite.AccessRequest, authreq *fosite.DeviceAuthorizeRequest) {
						deviceCode, _, err := strategy.GenerateDeviceCode(context.TODO())
						require.NoError(t, err)
						areq.Form = url.Values{"device_code": {deviceCode}}
					},
					expectErr: fosite.ErrInvalidGrant,
				},
				{
					areq: &fosite.AccessRequest{
						GrantTypes: fosite.Arguments{"urn:ietf:params:oauth:grant-type:device_code"},
						Request: fosite.Request{
							Form:        url.Values{"device_code": {"AAAA"}},
							Client:      &fosite.DefaultClient{GrantTypes: []string{"urn:ietf:params:oauth:grant-type:device_code"}},
							Session:     &fosite.DefaultSession{},
							RequestedAt: time.Now().UTC(),
						},
					},
					description: "should fail because device code validation failed",
					expectErr:   fosite.ErrInvalidGrant,
				},
				{
					areq: &fosite.AccessRequest{
						GrantTypes: fosite.Arguments{"urn:ietf:params:oauth:grant-type:device_code"},
						Request: fosite.Request{
							Client:      &fosite.DefaultClient{ID: "foo", GrantTypes: []string{"urn:ietf:params:oauth:grant-type:device_code"}},
							Session:     &fosite.DefaultSession{},
							RequestedAt: time.Now().UTC(),
						},
					},
					authreq: &fosite.DeviceAuthorizeRequest{
						Request: fosite.Request{
							Client:         &fosite.DefaultClient{ID: "bar"},
							RequestedScope: fosite.Arguments{"a", "b"},
						},
					},
					description: "should fail because client mismatch",
					setup: func(t *testing.T, areq *fosite.AccessRequest, authreq *fosite.DeviceAuthorizeRequest) {
						token, signature, err := strategy.GenerateDeviceCode(context.TODO())
						require.NoError(t, err)
						areq.Form = url.Values{"device_code": {token}}

						require.NoError(t, store.CreateDeviceCodeSession(context.TODO(), signature, authreq))
					},
					expectErr: fosite.ErrInvalidGrant,
				},
				{
					areq: &fosite.AccessRequest{
						GrantTypes: fosite.Arguments{"urn:ietf:params:oauth:grant-type:device_code"},
						Request: fosite.Request{
							Client:      &fosite.DefaultClient{ID: "foo", GrantTypes: []string{"urn:ietf:params:oauth:grant-type:device_code"}},
							Session:     &fosite.DefaultSession{},
							RequestedAt: time.Now().UTC(),
						},
					},
					authreq: &fosite.DeviceAuthorizeRequest{
						Request: fosite.Request{
							Client:         &fosite.DefaultClient{ID: "foo", GrantTypes: []string{"urn:ietf:params:oauth:grant-type:device_code"}},
							Session:        &fosite.DefaultSession{},
							RequestedScope: fosite.Arguments{"a", "b"},
							RequestedAt:    time.Now().UTC(),
						},
					},
					description: "should pass",
					setup: func(t *testing.T, areq *fosite.AccessRequest, authreq *fosite.DeviceAuthorizeRequest) {
						token, signature, err := strategy.GenerateDeviceCode(context.TODO())
						require.NoError(t, err)

						areq.Form = url.Values{"device_code": {token}}
						require.NoError(t, store.CreateDeviceCodeSession(context.TODO(), signature, authreq))
					},
				},
				{
					areq: &fosite.AccessRequest{
						GrantTypes: fosite.Arguments{"urn:ietf:params:oauth:grant-type:device_code"},
						Request: fosite.Request{
							Form: url.Values{},
							Client: &fosite.DefaultClient{
								GrantTypes: fosite.Arguments{"urn:ietf:params:oauth:grant-type:device_code"},
							},
							GrantedScope: fosite.Arguments{"foo", "offline"},
							Session:      &fosite.DefaultSession{},
							RequestedAt:  time.Now().UTC(),
						},
					},
					check: func(t *testing.T, areq *fosite.AccessRequest, authreq *fosite.DeviceAuthorizeRequest) {
						assert.Equal(t, time.Now().Add(time.Minute).UTC().Round(time.Second), areq.GetSession().GetExpiresAt(fosite.AccessToken))
						assert.Equal(t, time.Now().Add(time.Minute).UTC().Round(time.Second), areq.GetSession().GetExpiresAt(fosite.RefreshToken))
					},
					setup: func(t *testing.T, areq *fosite.AccessRequest, authreq *fosite.DeviceAuthorizeRequest) {
						code, sig, err := strategy.GenerateAuthorizeCode(nil, nil)
						require.NoError(t, err)
						areq.Form.Add("code", code)

						require.NoError(t, store.CreateAuthorizeCodeSession(nil, sig, areq))
						require.NoError(t, store.InvalidateAuthorizeCodeSession(nil, sig))
					},
					description: "should fail because code has been used already",
					expectErr:   fosite.ErrInvalidGrant,
				},
				{
					areq: &fosite.AccessRequest{
						GrantTypes: fosite.Arguments{"urn:ietf:params:oauth:grant-type:device_code"},
						Request: fosite.Request{
							Form: url.Values{},
							Client: &fosite.DefaultClient{
								GrantTypes: fosite.Arguments{"urn:ietf:params:oauth:grant-type:device_code"},
							},
							GrantedScope: fosite.Arguments{"foo", "offline"},
							Session:      &fosite.DefaultSession{},
							RequestedAt:  time.Now().UTC(),
						},
					},
					check: func(t *testing.T, areq *fosite.AccessRequest, authreq *fosite.DeviceAuthorizeRequest) {
						assert.Equal(t, time.Now().Add(time.Minute).UTC().Round(time.Second), areq.GetSession().GetExpiresAt(fosite.AccessToken))
						assert.Equal(t, time.Now().Add(time.Minute).UTC().Round(time.Second), areq.GetSession().GetExpiresAt(fosite.RefreshToken))
					},
					setup: func(t *testing.T, areq *fosite.AccessRequest, authreq *fosite.DeviceAuthorizeRequest) {
						code, sig, err := strategy.GenerateDeviceCode(context.TODO())
						require.NoError(t, err)
						areq.Form.Add("device_code", code)
						areq.GetSession().SetExpiresAt(fosite.DeviceCode, time.Now().Add(-time.Hour).UTC().Round(time.Second))
						require.NoError(t, store.CreateDeviceCodeSession(context.TODO(), sig, areq))
					},
					description: "should fail because device code has expired",
					expectErr:   fosite.ErrDeviceExpiredToken,
				},
			} {
				t.Run(fmt.Sprintf("case=%d/description=%s", i, c.description), func(t *testing.T) {
					if c.setup != nil {
						c.setup(t, c.areq, c.authreq)
					}

					t.Logf("Processing %+v", c.areq.Client)

					err := h.HandleTokenEndpointRequest(context.Background(), c.areq)
					if c.expectErr != nil {
						require.EqualError(t, err, c.expectErr.Error(), "%+v", err)
					} else {
						require.NoError(t, err, "%+v", err)
						if c.check != nil {
							c.check(t, c.areq, c.authreq)
						}
					}
				})
			}
		})
	}
}

func TestDeviceAuthorizeCodeTransactional_HandleTokenEndpointRequest(t *testing.T) {
	var mockTransactional *internal.MockTransactional
	var mockCoreStore *internal.MockCoreStorage
	strategy := hmacshaStrategy
	request := &fosite.AccessRequest{
		GrantTypes: fosite.Arguments{"urn:ietf:params:oauth:grant-type:device_code"},
		Request: fosite.Request{
			Client: &fosite.DefaultClient{
				GrantTypes: fosite.Arguments{"urn:ietf:params:oauth:grant-type:device_code", "refresh_token"},
			},
			GrantedScope: fosite.Arguments{"offline"},
			Session:      &fosite.DefaultSession{},
			RequestedAt:  time.Now().UTC(),
		},
	}
	token, _, err := strategy.GenerateDeviceCode(context.TODO())
	require.NoError(t, err)
	request.Form = url.Values{"device_code": {token}}
	response := fosite.NewAccessResponse()
	propagatedContext := context.Background()

	// some storage implementation that has support for transactions, notice the embedded type `storage.Transactional`
	type transactionalStore struct {
		storage.Transactional
		CoreStorage
	}

	for _, testCase := range []struct {
		description string
		setup       func()
		expectError error
	}{
		{
			description: "transaction should be committed successfully if no errors occur",
			setup: func() {
				mockCoreStore.
					EXPECT().
					GetDeviceCodeSession(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(request, nil).
					Times(1)
				mockTransactional.
					EXPECT().
					BeginTX(propagatedContext).
					Return(propagatedContext, nil)
				mockCoreStore.
					EXPECT().
					InvalidateDeviceCodeSession(gomock.Any(), gomock.Any()).
					Return(nil).
					Times(1)
				mockCoreStore.
					EXPECT().
					CreateAccessTokenSession(propagatedContext, gomock.Any(), gomock.Any()).
					Return(nil).
					Times(1)
				mockCoreStore.
					EXPECT().
					CreateRefreshTokenSession(propagatedContext, gomock.Any(), gomock.Any()).
					Return(nil).
					Times(1)
				mockTransactional.
					EXPECT().
					Commit(propagatedContext).
					Return(nil).
					Times(1)
			},
		},
		{
			description: "transaction should be rolled back if `InvalidateDeviceCodeSession` returns an error",
			setup: func() {
				mockCoreStore.
					EXPECT().
					GetDeviceCodeSession(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(request, nil).
					Times(1)
				mockTransactional.
					EXPECT().
					BeginTX(propagatedContext).
					Return(propagatedContext, nil)
				mockCoreStore.
					EXPECT().
					InvalidateDeviceCodeSession(gomock.Any(), gomock.Any()).
					Return(errors.New("Whoops, a nasty database error occurred!")).
					Times(1)
				mockTransactional.
					EXPECT().
					Rollback(propagatedContext).
					Return(nil).
					Times(1)
			},
			expectError: fosite.ErrServerError,
		},
		{
			description: "transaction should be rolled back if `CreateAccessTokenSession` returns an error",
			setup: func() {
				mockCoreStore.
					EXPECT().
					GetDeviceCodeSession(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(request, nil).
					Times(1)
				mockTransactional.
					EXPECT().
					BeginTX(propagatedContext).
					Return(propagatedContext, nil)
				mockCoreStore.
					EXPECT().
					InvalidateDeviceCodeSession(gomock.Any(), gomock.Any()).
					Return(nil).
					Times(1)
				mockCoreStore.
					EXPECT().
					CreateAccessTokenSession(propagatedContext, gomock.Any(), gomock.Any()).
					Return(errors.New("Whoops, a nasty database error occurred!")).
					Times(1)
				mockTransactional.
					EXPECT().
					Rollback(propagatedContext).
					Return(nil).
					Times(1)
			},
			expectError: fosite.ErrServerError,
		},
		{
			description: "should result in a server error if transaction cannot be created",
			setup: func() {
				mockCoreStore.
					EXPECT().
					GetDeviceCodeSession(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(request, nil).
					Times(1)
				mockTransactional.
					EXPECT().
					BeginTX(propagatedContext).
					Return(nil, errors.New("Whoops, unable to create transaction!"))
			},
			expectError: fosite.ErrServerError,
		},
		{
			description: "should result in a server error if transaction cannot be rolled back",
			setup: func() {
				mockCoreStore.
					EXPECT().
					GetDeviceCodeSession(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(request, nil).
					Times(1)
				mockTransactional.
					EXPECT().
					BeginTX(propagatedContext).
					Return(propagatedContext, nil)
				mockCoreStore.
					EXPECT().
					InvalidateDeviceCodeSession(gomock.Any(), gomock.Any()).
					Return(errors.New("Whoops, a nasty database error occurred!")).
					Times(1)
				mockTransactional.
					EXPECT().
					Rollback(propagatedContext).
					Return(errors.New("Whoops, unable to rollback transaction!")).
					Times(1)
			},
			expectError: fosite.ErrServerError,
		},
		{
			description: "should result in a server error if transaction cannot be committed",
			setup: func() {
				mockCoreStore.
					EXPECT().
					GetDeviceCodeSession(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(request, nil).
					Times(1)
				mockTransactional.
					EXPECT().
					BeginTX(propagatedContext).
					Return(propagatedContext, nil)
				mockCoreStore.
					EXPECT().
					InvalidateDeviceCodeSession(gomock.Any(), gomock.Any()).
					Return(nil).
					Times(1)
				mockCoreStore.
					EXPECT().
					CreateAccessTokenSession(propagatedContext, gomock.Any(), gomock.Any()).
					Return(nil).
					Times(1)
				mockCoreStore.
					EXPECT().
					CreateRefreshTokenSession(propagatedContext, gomock.Any(), gomock.Any()).
					Return(nil).
					Times(1)
				mockTransactional.
					EXPECT().
					Commit(propagatedContext).
					Return(errors.New("Whoops, unable to commit transaction!")).
					Times(1)
				mockTransactional.
					EXPECT().
					Rollback(propagatedContext).
					Return(nil).
					Times(1)
			},
			expectError: fosite.ErrServerError,
		},
	} {
		t.Run(fmt.Sprintf("scenario=%s", testCase.description), func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockTransactional = internal.NewMockTransactional(ctrl)
			mockCoreStore = internal.NewMockCoreStorage(ctrl)
			testCase.setup()

			handler := DeviceAuthorizationHandler{
				CoreStorage: transactionalStore{
					mockTransactional,
					mockCoreStore,
				},
				AccessTokenStrategy:  &strategy,
				RefreshTokenStrategy: &strategy,
				DeviceCodeStrategy:   &strategy,
				UserCodeStrategy:     &strategy,
				Config: &fosite.Config{
					ScopeStrategy:             fosite.HierarchicScopeStrategy,
					AudienceMatchingStrategy:  fosite.DefaultAudienceMatchingStrategy,
					DeviceAndUserCodeLifespan: time.Minute,
				},
			}

			if err := handler.PopulateTokenEndpointResponse(propagatedContext, request, response); testCase.expectError != nil {
				assert.EqualError(t, err, testCase.expectError.Error())
			}
		})
	}
}
