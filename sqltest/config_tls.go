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

// TLSConfig represents a TLS configuration.
type TLSConfig struct {
	ClientCertFile string
	ClientKeyFile  string
	RootCertFile   string
}

// NewTLSConfig returns a new TLS configuration.
// nolint: gosec, exhaustruct
func NewTLSConfig() *TLSConfig {
	return &TLSConfig{
		ClientCertFile: "",
		ClientKeyFile:  "",
		RootCertFile:   "",
	}
}

// SetClientKeyFile sets a SSL client key file.
func (config *TLSConfig) SetClientKeyFile(file string) {
	config.ClientKeyFile = file
}

// SetClientCertFile sets a SSL client certificate file.
func (config *TLSConfig) SetClientCertFile(file string) {
	config.ClientCertFile = file
}

// SetRootCertFile sets a SSL root certificate file.
func (config *TLSConfig) SetRootCertFile(file string) {
	config.RootCertFile = file
}
