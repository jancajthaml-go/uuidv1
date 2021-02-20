## universally unique identifier

[![Go Report Card](https://goreportcard.com/badge/jancajthaml-go/uuid)](https://goreportcard.com/report/jancajthaml-go/uuidv1)

Algorithm generate 128bit UUID with node and time based randomness in RFC4122 format.

### Usage ###

```
import "github.com/jancajthaml-go/uuidv1"

uuidv1.Generate()
```

### Performance ###

- 328 B/op
- 17 allocs/op


verify your performance by running `make benchmark`

### Resources ###

* [Wikipedia - Universally Unique IDentifier](https://en.wikipedia.org/wiki/Universally_unique_identifier)
