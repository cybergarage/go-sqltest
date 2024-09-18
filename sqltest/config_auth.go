// Copyright (C) 2020 The go-sqltest Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sqltest

// AuthMethod represents an authentication method.
type AuthMethod int

const (
	AuthdNone AuthMethod = iota
	AuthPlain
	AuthMD5
	AuthSCRAMSHA256
)

// AuthConfig stores authentication configuration parameters.
type AuthConfig struct {
	User     string
	Password string
	Auth     AuthMethod
}

// NewAuthConfig returns a default authentication configuration instance.
func NewAuthConfig() *AuthConfig {
	config := &AuthConfig{
		User:     "",
		Password: "",
		Auth:     AuthdNone,
	}

	return config
}

// SetUser sets a user name.
func (config *AuthConfig) SetUser(user string) {
	config.User = user
}

// SetPassword sets a password.
func (config *AuthConfig) SetPassword(password string) {
	config.Password = password
}

// SetAuth sets an authentication method.
func (config *AuthConfig) SetAuth(auth AuthMethod) {
	config.Auth = auth
}
