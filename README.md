# go-sqltest

![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/cybergarage/go-sqltest)
[![test](https://github.com/cybergarage/go-sqltest/actions/workflows/make.yml/badge.svg)](https://github.com/cybergarage/go-sqltest/actions/workflows/make.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/cybergarage/go-sqltest.svg)](https://pkg.go.dev/github.com/cybergarage/go-sqltest)

The go-sqltest is a scenario test framework for SQL-compatible databases. 
The go-sqltest imports the scenario test files and runs the scenario test to the target SQL-compatible databases as the following.

![](doc/img/framework.png)

The scenario testing framework will be used in SQL-compatible database projects such as [go-mysql](https://github.com/cybergarage/go-mysql).

## Client

The go-sqltest is implemented using Go and uses standard generic [database/sql](https://pkg.go.dev/database/sql) interfaces, mainly SQL (or SQL-like) databases for clients. Therefore, go-sqltest can run scenario tests against a variety of SQL databases such as MySQL and PostgreSQL by providing drivers for the [database/sql](https://pkg.go.dev/database/sql) interface.

## Adding scenario tests

The go-sqltest imports a scenario test file in the `test` directory and embeds the file in the helper test function. The scenario test file must have the extension `qst` (query scenario test).

### Scenario File Format

A scenario test consists of a combination of a query and an expected response; the EBNF specification is as follows.

```
scenario-test = (query, response)*
query = SQL
response = JSON response
```

