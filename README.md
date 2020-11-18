# fresbi

[![build-img]][build-url]
[![pkg-img]][pkg-url]
[![reportcard-img]][reportcard-url]
[![coverage-img]][coverage-url]

Fresbi (R) ElasticSearch Bulk Indexer.

## Rationale

Two most popular Elasticsearch (ES) clients ([official](https://github.com/elastic/go-elasticsearch) and [community](https://github.com/olivere/elastic)) have API for ES Bulk indexing ([docs](https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-bulk.html)).
Sadly it's not very friendly: asynchronous where it cannot be, too verbose for simple things and is memory greedy. This package addresses all of the problems.

## Features

* Simple API.
* Clean and tested code.
* Optimized for speed.
* Dependency-free.

Also R in `fresbi` can be:

- `ReliableClient` 
- `RawClient`

or as a feature:

- `Retryable`
- `Reliable`

## Install

Go version 1.15+

```
go get github.com/cristalhq/fresbi
```

## Example

```go
// TODO
```

Also see examples: [this above](https://github.com/cristalhq/fresbi/blob/master/example_test.go).

## Documentation

See [these docs][doc-url].

## License

[MIT License](LICENSE).

[build-img]: https://github.com/cristalhq/fresbi/workflows/build/badge.svg
[build-url]: https://github.com/cristalhq/fresbi/actions
[pkg-img]: https://pkg.go.dev/badge/cristalhq/fresbi
[pkg-url]: https://pkg.go.dev/github.com/cristalhq/fresbi
[reportcard-img]: https://goreportcard.com/badge/cristalhq/fresbi
[reportcard-url]: https://goreportcard.com/report/cristalhq/fresbi
[coverage-img]: https://codecov.io/gh/cristalhq/fresbi/branch/master/graph/badge.svg
[coverage-url]: https://codecov.io/gh/cristalhq/fresbi
