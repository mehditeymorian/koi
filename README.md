<h1 align="center">
<img alt="Koi logo" src="asset/logo.webp" width="500px"/><br/>
KOI
</h1>
<p align="center">Goroutine and Worker Manager</p>

<p align="center">
<a href="https://pkg.go.dev/github.com/mehditeymorian/koi/v3?tab=doc"target="_blank">
    <img src="https://img.shields.io/badge/Go-1.18+-00ADD8?style=for-the-badge&logo=go" alt="go version" />
</a>&nbsp;
<img src="https://img.shields.io/badge/license-apache_2.0-red?style=for-the-badge&logo=none" alt="license" />

<img src="https://img.shields.io/badge/Version-1.0.1-informational?style=for-the-badge&logo=none" alt="version" />
</p>

# Installation

```bash
go get github.com/mehditeymorian/koi
```

# Usage

```go
package main

import (
	"log"
	"time"

	"github.com/mehditeymorian/koi"
)

func main() {
	pond := koi.NewPond[int, int]()

	printWorker := koi.Worker[int, int]{
		ConcurrentCount: 2,
		QueueSize:       10,
		Work: func(a int) *int {
			time.Sleep(1 * time.Second)
			log.Println(a)

			return nil
		},
	}

	_ = pond.RegisterWorker("printer", printWorker)

	for i := 0; i < 10; i++ {
		_, err := pond.AddWork("printer", i)
		if err != nil {
			log.Printf("error while adding job: %v\n", err)
		}
	}

	log.Println("all job added")

	for {

	}
}
```

**Note**: `pond.AddJob` is non-blocking unless worker queue is full.

# Terminology

- **Koi**: Koi is an informal name for the colored variants of C. rubrofuscus kept for ornamental purposes.
- **Pond**: an area of water smaller than a lake, often artificially made.
