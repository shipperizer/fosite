package oauth2

import (
	"context"
	"time"

	"github.com/ory/fosite"
	"github.com/ory/x/errorsx"
)

// DeviceAuthorizationHandler is a response handler for the Device Authorisation Grant as
// defined in https://tools.ietf.org/html/rfc8628#section-3.1
type DeviceAuthorizationHandler struct {
	DeviceCodeStorage  DeviceCodeStorage
	UserCodeStorage    UserCodeStorage
	DeviceCodeStrategy DeviceCodeStrategy
	UserCodeStrategy   UserCodeStrategy
	Config             fosite.Configurator
}

func (d *DeviceAuthorizationHandler) HandleDeviceAuthorizeEndpointRequest(ctx context.Context, dar fosite.Requester, resp fosite.DeviceAuthorizeResponder) error {
	deviceCode, deviceCodeSignature, err := d.DeviceCodeStrategy.GenerateDeviceCode(ctx)
	if err != nil {
		return errorsx.WithStack(fosite.ErrServerError.WithWrap(err).WithDebug(err.Error()))
	}

	userCode, userCodeSignature, err := d.UserCodeStrategy.GenerateUserCode(ctx)
	if err != nil {
		return errorsx.WithStack(fosite.ErrServerError.WithWrap(err).WithDebug(err.Error()))
	}

	// Set User Code expiry time
	dar.GetSession().SetExpiresAt(fosite.UserCode, time.Now().UTC().Add(d.Config.GetDeviceAndUserCodeLifespan(ctx)).Round(time.Second))
	dar.SetID(deviceCodeSignature)

	// Store the User Code session (this has no real data other that the uer and device code), can be converted into a 'full' session after user auth
	if err := d.UserCodeStorage.CreateUserCodeSession(ctx, userCodeSignature, dar); err != nil {
		return errorsx.WithStack(fosite.ErrServerError.WithWrap(err).WithDebug(err.Error()))
	}

	// Populate the response fields
	resp.SetDeviceCode(deviceCode)
	resp.SetUserCode(userCode)
	resp.SetVerificationURI(d.Config.GetDeviceVerificationURL(ctx))
	resp.SetVerificationURIComplete(d.Config.GetDeviceVerificationURL(ctx) + "?user_code=" + userCode)
	resp.SetExpiresIn(int64(time.Until(dar.GetSession().GetExpiresAt(fosite.UserCode)).Seconds()))
	resp.SetInterval(int(d.Config.GetDeviceAuthTokenPollingInterval(ctx).Seconds()))
	return nil
}
