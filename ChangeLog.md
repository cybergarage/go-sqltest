# Changelog

## v1.4.x (2023-xx-xx)
- Added new scenario tests, such as pgbench
- Added new scenario tests
  - Added ALTER queries

## v1.4.0 (2024-95-12)
- Updated 
  - Client interface
    - Add TLS settings

## v1.3.2 (2023-12-02)
- Updated
  - Client interface
    - Add Ping() for testing

## v1.3.1 (2023-11-28)
- Updated
  - Client interface
    - Add SetUser() and SetPassword() for authentication testing

## v1.3.0 (2023-10-20)
- Updated
  - Extented RunEmbedSuites to support regular expressions
- Added new scenario tests
  - BEGIN, COMMIT and ROLLBACK

## v1.2.4 (2023-09-28)
- Updated
  - ScenarioTest
    - Supported timestamp data comparison
  - Scenario tests
    - Added timestamp data type for PICT tests

## v1.2.3 (2023-09-24)
- Updated client interface
  - Added user and password configurations
- Updated PostgreSQL clients
  - Added Pgx client

## v1.2.2 (2023-09-19)
- Added new scenario tests
  - UPDATE
    - Added tests for arithmetic operations

## v1.2.1 (2023-09-18)
- Added new scenario tests
  - SELECT
    - Added tests for ORDER BY
    - Added tests for LIMIT

## v1.2.0 (2023-09-16)
- Added command line utility
  - Added `sqltest` command
- Improved
  - Aggregation function tests
    - COUNT, SUM, AVG, MIN and MAX
      - Added tests for empty tables
      - Removed insertred rows after each test

## v1.1.1 (2023-09-12)
- Improved 
  - Suite::Run() 
    - Allowed floating point comparisons to account for precision errors
- Added new scenario tests
  - Added math function tests
    - ABS, FLOOR and CEIL

## v1.1.0 (2023-09-12)
- Added new scenario tests
  - Added aggregation function tests
    - COUNT, SUM, AVG, MIN and MAX
- Updated scenario tests
  - Updated YCSB-like workload test

## v1.0.1 (2023-09-05)
- Updated ScenarioTest to compare rows more accurately
- Added new scenario tests
  - Added CRUD tests using PICT

## v1.0.0 (2023-08-20)
- Fixed embed test interfaces

## v0.9.6 (2023-07-28)
- Updated error messages
- Added new client interface
  - Independent from MySQL protocol
  - Added PostgreSQL client

## v0.9.5 (2023-06-11)
- Updated RunEmbedSuites() to run only specified tests

## v0.9.4 (2023-04-23)
- Updated embed test interfaces
- Updated to use an empty database for each scenario test
- Fixed compiler warnings

## v0.9.3 (2023-04-09)
- Updated scenario test interfaces

## v0.9.2 (2023-03-18)
- Added Client interface

## v0.9.1 (2023-03-17)
- Added
  - Embed scenario test files
- Improved
  - Removed go-mysql dependency

## v0.9.0 (2023-02-28)
- Initial public release
- Initial scenario tests
  - Imported scenario testing modules from go-mysql
