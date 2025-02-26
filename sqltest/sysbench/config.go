// Copyright (C) 2025 The go-sqltest Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sysbench

const (
	// https://github.com/akopytov/sysbench
	CONFIG_THREADS = "threads"
	CONFIG_EVENTS  = "events"
	CONFIG_TIME    = "time"
)

const (
	// https://github.com/akopytov/sysbench/tree/master/src/lua
	OLTP_DELETE       = "oltp_delete"
	OLTP_INSERT       = "oltp_insert"
	OLTP_READ_ONLY    = "oltp_read_only"
	OLTP_READ_WRITE   = "oltp_read_write"
	OLTP_UPDATE_INDEX = "oltp_update_index"
	OLTP_WRITE_ONLY   = "oltp_write_only"
	OLTP_COMMON       = "oltp_common"
)

// Config represents a sysbench configuration.
type Config map[string]string

// NewConfig returns a new config.
func NewConfig() Config {
	return Config{}
}

// NewDefaultConfig returns a new default config.
func NewDefaultConfig() Config {
	cfg := NewConfig()
	cfg.Set(CONFIG_THREADS, "1")
	cfg.Set(CONFIG_EVENTS, "0")
	cfg.Set(CONFIG_TIME, "1")
	return cfg
}

// Set sets a config value.
func (config Config) Set(name string, value string) {
	config[name] = value
}
