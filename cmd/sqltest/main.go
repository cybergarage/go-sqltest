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

	"github.com/cybergarage/go-logger/log"
	"github.com/cybergarage/go-sqltest/sqltest"
)

func printError(err error) {
	log.Error(err)
}

func main() {
	log.SetSharedLogger(log.NewStdoutLogger(log.LevelInfo))

	var (
		host     = flag.String("host", "localhost", "Database host")
		protocol = flag.String("protocol", "pg", "Database type (mysql|pg)")
		port     = flag.Int("port", 0, "Database port")
		user     = flag.String("user", "", "Database user")
		password = flag.String("password", "", "Database password")
	)
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: sqltest [options] <scenario file>\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		printError(fmt.Errorf("missing scenario file"))
		flag.Usage()
		return
	}

	var client sqltest.Client
	var queryDialect sqltest.QueryDialect

	switch *protocol {
	case "mysql":
		client = sqltest.NewMySQLClient()
		queryDialect = sqltest.QueryDialectMySQL
		if *port == 0 {
			*port = 3306
		}
	case "pg":
		client = sqltest.NewPostgresClient()
		queryDialect = sqltest.QueryDialectPostgreSQL
		if *port == 0 {
			*port = 5432
		}
	default:
		flag.Usage()
		return
	}

	client.SetHost(*host)
	client.SetPort(*port)
	if 0 < len(*user) {
		client.SetUser(*user)
	}
	if 0 < len(*password) {
		client.SetPassword(*password)
	}

	scenarioPath := args[0]

	scenarioTest, err := sqltest.NewScenarioTesterWith(
		sqltest.WithScenarioTesterFile(scenarioPath),
		sqltest.WithScenarioTesterQueryDialect(queryDialect),
	)
	if err != nil {
		printError(err)
		return
	}
	log.Infof("scenario loaded : %s", scenarioTest.Name())

	testDBName := sqltest.GenerateTempDBName(sqltest.TestDBNamePrefix)

	log.Infof("test database name : %s", testDBName)

	// Create a test database

	err = client.Open()
	if err != nil {
		printError(err)
		return
	}

	createDBErr := client.CreateDatabase(testDBName)

	err = client.Close()
	if err != nil {
		printError(err)
	}

	if createDBErr != nil {
		printError(fmt.Errorf("failed to create database %s : %s", testDBName, createDBErr.Error()))
		return
	}

	// Run the scenario test

	client.SetDatabase(testDBName)
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

		client.SetDatabase("")
		err = client.Open()
		if err != nil {
			printError(err)
			return
		}

		err = client.DropDatabase(testDBName)
		if err != nil {
			printError(err)
		}

		err = client.Close()
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
