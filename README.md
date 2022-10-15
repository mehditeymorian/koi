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
```Go
// create a pond 
pond := koi.NewPond()

// create a worker
worker := koi.Worker{
    ConcurrentCount: CONCURRENT_RUNNING_GOROUTINE_FOR_THIS_WORKER,
    QueueSize:       REQUEST_QUEUE_SIZE,
    Work: func(a any) any {
        // do some work
        return RESULT
    },
}

// register worker to a unique id
err = pond.RegisterWorker("workerID", worker)

// add job to worker
// this is non-blocking unless the queue is full.
resultChan, err := pond.AddWork("workerID", requestData)

// read results from worker
for res := range resultChan {
    // do something with result
}  
```
**Note**: `pond.AddJob` is non-blocking unless worker queue is full.


# Terminology
- **Koi**: Koi is an informal name for the colored variants of C. rubrofuscus kept for ornamental purposes.
- **Pond**: an area of water smaller than a lake, often artificially made.

