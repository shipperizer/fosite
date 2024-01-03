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

package fosite

type DeviceAuthorizeResponse struct {
	deviceCode              string
	userCode                string
	verificationURI         string
	verificationURIComplete string
	interval                int
	expiresIn               int64
}

func NewDeviceAuthorizeResponse() *DeviceAuthorizeResponse {
	return &DeviceAuthorizeResponse{}
}

func (d *DeviceAuthorizeResponse) GetDeviceCode() string {
	return d.deviceCode
}

// GetUserCode returns the response's user code
func (d *DeviceAuthorizeResponse) SetDeviceCode(code string) {
	d.deviceCode = code
}

func (d *DeviceAuthorizeResponse) GetUserCode() string {
	return d.userCode
}

func (d *DeviceAuthorizeResponse) SetUserCode(code string) {
	d.userCode = code
}

// GetVerificationURI returns the response's verification uri
func (d *DeviceAuthorizeResponse) GetVerificationURI() string {
	return d.verificationURI
}

func (d *DeviceAuthorizeResponse) SetVerificationURI(uri string) {
	d.verificationURI = uri
}

// GetVerificationURIComplete returns the response's complete verification uri if set
func (d *DeviceAuthorizeResponse) GetVerificationURIComplete() string {
	return d.verificationURIComplete
}

func (d *DeviceAuthorizeResponse) SetVerificationURIComplete(uri string) {
	d.verificationURIComplete = uri
}

// GetExpiresIn returns the response's device code and user code lifetime in seconds if set
func (d *DeviceAuthorizeResponse) GetExpiresIn() int64 {
	return d.expiresIn
}

func (d *DeviceAuthorizeResponse) SetExpiresIn(seconds int64) {
	d.expiresIn = seconds
}

// GetInterval returns the response's polling interval if set
func (d *DeviceAuthorizeResponse) GetInterval() int {
	return d.interval
}

func (d *DeviceAuthorizeResponse) SetInterval(seconds int) {
	d.interval = seconds
}
