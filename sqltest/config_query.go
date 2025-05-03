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

// QueryConfig represents a query configuration.
type QueryConfig struct {
	preparedStatementEnabled bool
}

// NewQueryConfig returns a new QueryConfig instance with default values.
func NewQueryConfig() *QueryConfig {
	return &QueryConfig{
		preparedStatementEnabled: true,
	}
}

// SetPreparedStatementEnabled sets the prepared statement enabled flag.
func (config *QueryConfig) SetPreparedStatementEnabled(enabled bool) {
	config.preparedStatementEnabled = enabled
}

// IsPreparedStatementEnabled returns true if prepared statements are enabled.
func (config *QueryConfig) IsPreparedStatementEnabled() bool {
	return config.preparedStatementEnabled
}
