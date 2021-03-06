# [go-cvss] - Common Vulnerability Scoring System (CVSS)

[![check vulns](https://github.com/spiegel-im-spiegel/go-cvss/workflows/vulns/badge.svg)](https://github.com/spiegel-im-spiegel/go-cvss/actions)
[![lint status](https://github.com/spiegel-im-spiegel/go-cvss/workflows/lint/badge.svg)](https://github.com/spiegel-im-spiegel/go-cvss/actions)
[![GitHub license](https://img.shields.io/badge/license-Apache%202-blue.svg)](https://raw.githubusercontent.com/spiegel-im-spiegel/go-cvss/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/spiegel-im-spiegel/go-cvss.svg)](https://github.com/spiegel-im-spiegel/go-cvss/releases/latest)

Importing CVSS vector and scoring.

- Supoort CVSS version 3.0 and 3.1
- Exporting CVSS information with template string

## Sample Code

### Base Metrics

```go
package main

import (
    "fmt"
    "os"

    "github.com/spiegel-im-spiegel/go-cvss/v3/metric"
)

func main() {
    bm, err := metric.NewBase().Decode("CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:C/C:H/I:H/A:H") //CVE-2020-1472: ZeroLogon
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    fmt.Printf("Severity: %v (%v)\n", bm.Severity(), bm.Score())
    // Output:
    // Severity: Critical (10)
}
```

### Temporal Metrics

```go
package main

import (
    "fmt"
    "os"

    "github.com/spiegel-im-spiegel/go-cvss/v3/metric"
)

func main() {
    tm, err := metric.NewTemporal().Decode("CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:C/C:H/I:H/A:H/E:F/RL:W/RC:R") //CVE-2020-1472: ZeroLogon
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    fmt.Printf("Base Severity: %v (%v)\n", tm.BaseMetrics().Severity(), tm.BaseMetrics().Score())
    fmt.Printf("Temporal Severity: %v (%v)\n", tm.Severity(), tm.Score())
    // Output:
    // Base Severity: Critical (10)
    // Temporal Severity: Critical (9.1)
}
```

### Reporting with template

ref: [sample.go](https://github.com/spiegel-im-spiegel/go-cvss/blob/master/sample/sample.go)

## Reference

- [CVSS v3.0 Specification Document](https://www.first.org/cvss/v3.0/specification-document)
- [CVSS v3.1 Specification Document](https://www.first.org/cvss/v3.1/specification-document)

[go-cvss]: https://github.com/spiegel-im-spiegel/go-cvss
