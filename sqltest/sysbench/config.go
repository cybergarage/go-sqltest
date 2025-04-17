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

import "strconv"

const (
	// https://github.com/akopytov/sysbench
	ConfigThreads   = "threads"
	ConfigEvents    = "events"
	ConfigTime      = "time"
	ConfigTableSize = "table-size"
	ConfigDBDriver  = "db-driver"
	ConfigDBDebug   = "db-debug"
)

const (
	// https://github.com/akopytov/sysbench/tree/master/src/lua
	OltpDelete      = "oltp_delete"
	OltpInsert      = "oltp_insert"
	OltpReadOnly    = "oltp_read_only"
	OltpReadWrite   = "oltp_read_write"
	OltpUpdateIndex = "oltp_update_index"
	OltpWriteOnly   = "oltp_write_only"
	OltpCommon      = "oltp_common"
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
	cfg.SetThreads(DefaultThreads)
	cfg.SetEvents(DefaultEvents)
	cfg.SetTime(DefaultTime)
	cfg.SetTableSize(DefaultTableSize)
	return cfg
}

// Set sets a config value.
func (config Config) Set(name string, value string) {
	config[name] = value
}

// SetBool sets a config value as a boolean.
func (config Config) SetBool(key string, value bool) {
	if value {
		config.Set(key, "on")
		return
	}
	config.Set(key, "off")
}

// SetThreads sets the number of threads.
func (config Config) SetThreads(v int) {
	config.Set(ConfigThreads, strconv.Itoa(v))
}

// SetEvents sets the number of events.
func (config Config) SetEvents(v int) {
	config.Set(ConfigEvents, strconv.Itoa(v))
}

// SetTime sets the time.
func (config Config) SetTime(v int) {
	config.Set(ConfigTime, strconv.Itoa(v))
}

// SetTableSize sets the table size.
func (config Config) SetTableSize(v int) {
	config.Set(ConfigTableSize, strconv.Itoa(v))
}

// SetDBDriver sets the database driver.
func (config Config) SetDBDriver(value string) {
	config.Set(ConfigDBDriver, value)
}

// SetDBDebug sets the database debug.
func (config Config) SetDBDebug(value bool) {
	config.SetBool(ConfigDBDebug, value)
}
