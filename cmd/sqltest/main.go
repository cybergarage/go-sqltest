// Copyright (C) 2022 The go-sqltest Authors All rights reserved.
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

// nolint: forbidigo
package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/cybergarage/go-logger/log"
	"github.com/cybergarage/go-sqltest/sqltest"
)

func printUsage() {
	println("Usage: sqltest [options] <scenario path>")
	println("Options:")
	println("  -host <host>")
	println("  -port <port>")
	println("  -protocol <protocol>")
}

func printError(err error) {
	log.Error(err)
}

func main() {
	log.SetStdoutDebugEnbled(true)

	var (
		host     = flag.String("host", "localhost", "Host")
		port     = flag.Int("port", 0, "Port")
		protocol = flag.String("protocol", "", "<mysql|pg>")
	)
	flag.Parse()

	args := os.Args
	if len(args) < 2 {
		printUsage()
		return
	}

	var client sqltest.Client
	switch *protocol {
	case "mysql":
		client = sqltest.NewMySQLClient()
		if *port == 0 {
			*port = 3306
		}
	case "pg":
		client = sqltest.NewPostgresClient()
		if *port == 0 {
			*port = 5432
		}
	default:
		printUsage()
		return
	}

	client.SetHost(*host)
	client.SetPort(*port)

	scenarioPath := args[1]

	scenarioTest, err := sqltest.NewScenarioTestWithFile(scenarioPath)
	if err != nil {
		printError(err)
		return
	}

	err = client.Open()
	if err != nil {
		printError(err)
		return
	}

	defer func() {
		err := client.Close()
		if err != nil {
			printError(err)
		}
	}()

	testDBName := fmt.Sprintf("%s%d", scenarioTest.Name, time.Now().UnixNano())
	client.SetDatabase(testDBName)

	err = client.CreateDatabase(testDBName)
	if err != nil {
		printError(err)
		return
	}

	defer func() {
		err := client.DropDatabase(testDBName)
		if err != nil {
			printError(err)
		}
	}()

	scenarioTest.SetClient(client)
	err = scenarioTest.Run()
	if err != nil {
		printError(fmt.Errorf("%s : %s", scenarioTest.Name(), err.Error()))
	}
}
