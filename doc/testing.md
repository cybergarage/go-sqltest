# Testing

Creating tests is one of the most important parts of software development. It is the only way to ensure that the software works as expected and that it continues to work as expected as it evolves.

## Test Harnesses

**go-sqltest** is a library that provides a set of helpers to write tests for SQL databases. It is designed to be used with the `testing` package from the Go standard library. **go-sqltest** provides the follwing category tests:

### Scenario Tests

Scenario tests are used to test the behavior of the database in a specific scenario. They are written in a declarative language that describes the scenario and the expected results.

Scenario tests are defined in a `.qst` file. Each test consists of a sequence of SQL statements and a JSON object that describes the expected results. Scenario tests are located in the ../sqltest/scenarios\[sqltest/scenarios\] directory.

## Reference

- [SQLite Testing](https://www.sqlite.org/testing.html)

- [The MySQL Test Framework](https://dev.mysql.com/doc/dev/mysql-server/9.2.0/PAGE_MYSQL_TEST_RUN.html)

- [Test Frameworks - PostgreSQL wiki](https://wiki.postgresql.org/wiki/Test_Frameworks)

- [Cassandra Testing](https://cassandra.apache.org/_/development/testing.html)

- [Scylla Testing](https://www.scylladb.com/product/technology/scylla-testing/)
