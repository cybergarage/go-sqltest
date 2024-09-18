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

type AuthConfig struct {
	User     string
	Password string
	Auth     string
}

func NewAuthConfig() *AuthConfig {
	config := &AuthConfig{
		User:     "",
		Password: "",
		Auth:     "",
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
