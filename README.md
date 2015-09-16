# golang-common-ddd
Common interfaces (and implementations) for usage in Domain Driven Design (DDD) with golang web API's.

## TODO
- ~~Add basic `EncryptionService`~~
- ~~Add basic `ErrorsService`~~
- ~~Add basic `HttpRenderHelperService`~~
- ~~Add basic `HttpRequestHelperService`~~
- Add `AuthorizationHelperService` (this must be simple and easy to use)
- Add `Logger` with multiple implementations like `ConsoleLogger`, `FileLogger`, `MultiLogger` (just a wrapper containing a list of Loggers), etc
- Add `Storage` with a default implementation of `DatabaseSqlx`