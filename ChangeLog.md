# Changelog

## v1.0.3 (2023-xx-xx)
- Added new scenario tests, such as pgbench
- Added new scenario tests
  - Added ALTER queries

## v1.0.2 (2023-xx-xx)
- Added new scenario tests
  - Added aggregation function tests

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
