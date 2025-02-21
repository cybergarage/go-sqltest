# Changelog

## v1.6.x (2025-xx-xx)
### Added
- New scenario tests:
  - `pgbench`
  - ALTER query tests

## v1.5.0 (2025-02-21)
### Improved
- Enhanced scenario format to support bind parameters.
### Updated
- Updated `RunEmbedSuites` to include bind parameter tests. 

## v1.4.3 (2024-12-11)
### Added
- Secondary index tests.
### Updated
- Enhanced `Suite::Test()` to allow setting steps and error handlers.
- Expanded support for additional data types.

## v1.4.2 (2024-11-08)
### Added
- Authentication method configurations.
### Updated
- `PqClient` and `PgxClient` to support authentication method parameters.
- `MySQLClient::Open()` to set `parseTime=true`.

## v1.4.1 (2024-05-13)
### Fixed
- MySQL client:
  - Disabled inappropriate TLS parameters.

## v1.4.0 (2024-05-12)
### Updated
- Client interface:
  - Added TLS settings.

## v1.3.2 (2023-12-02)
### Updated
- Client interface:
  - Added `Ping()` for testing.

## v1.3.1 (2023-11-28)
### Updated
- Client interface:
  - Added `SetUser()` and `SetPassword()` for authentication testing.

## v1.3.0 (2023-10-20)
### Added
- New scenario tests:
  - `BEGIN`, `COMMIT`, and `ROLLBACK`.

### Updated
- Extended `RunEmbedSuites` to support regular expressions.

## v1.2.4 (2023-09-28)
### Updated
- `ScenarioTest`:
  - Supported timestamp data comparison.
- Scenario tests:
  - Added timestamp data type for PICT tests.

## v1.2.3 (2023-09-24)
### Updated
- Client interface:
  - Added user and password configurations.
- PostgreSQL clients:
  - Added Pgx client.

## v1.2.2 (2023-09-19)
### Added
- New scenario tests:
  - `UPDATE`:
    - Added tests for arithmetic operations.

## v1.2.1 (2023-09-18)
### Added
- New scenario tests:
  - `SELECT`:
    - Added tests for `ORDER BY`.
    - Added tests for `LIMIT`.

## v1.2.0 (2023-09-16)
### Added
- Command-line utility:
  - `sqltest` command.

### Improved
- Aggregation function tests:
  - `COUNT`, `SUM`, `AVG`, `MIN`, and `MAX`:
    - Added tests for empty tables.
    - Removed inserted rows after each test.

## v1.1.1 (2023-09-12)
### Improved
- `Suite::Run()`:
  - Allowed floating-point comparisons to account for precision errors.

### Added
- New scenario tests:
  - Math function tests:
    - `ABS`, `FLOOR`, and `CEIL`.

## v1.1.0 (2023-09-12)
### Added
- New scenario tests:
  - Aggregation function tests:
    - `COUNT`, `SUM`, `AVG`, `MIN`, and `MAX`.
### Updated
- Scenario tests:
  - YCSB-like workload test.

## v1.0.1 (2023-09-05)
### Updated
- `ScenarioTest` to compare rows more accurately.
### Added
- New scenario tests:
  - CRUD tests using PICT.

## v1.0.0 (2023-08-20)
### Fixed
- Embed test interfaces.

## v0.9.6 (2023-07-28)
### Updated
- Error messages.
- Client interface:
  - Made independent from the MySQL protocol.
  - Added PostgreSQL client.

## v0.9.5 (2023-06-11)
### Updated
- `RunEmbedSuites()` to run only specified tests.

## v0.9.4 (2023-04-23)
### Updated
- Embed test interfaces.
- Scenario tests to use an empty database for each test.
### Fixed
- Compiler warnings.

## v0.9.3 (2023-04-09)
### Updated
- Scenario test interfaces.

## v0.9.2 (2023-03-18)
### Added
- Client interface.

## v0.9.1 (2023-03-17)
### Added
- Embedded scenario test files.
### Improved
- Removed `go-mysql` dependency.

## v0.9.0 (2023-02-28)
### Initial Release
- Initial public release.
### Added
- Initial scenario tests:
  - Imported scenario testing modules from `go-mysql`.
