module github.com/storage-lock/go-event-listener-stdout

go 1.23

require (
	github.com/storage-lock/go-events v0.0.3
	github.com/storage-lock/go-memory-locks v0.0.0-20230830164812-86517fab5b4a
	github.com/storage-lock/go-storage-lock v0.0.3
	github.com/storage-lock/go-utils v0.0.2
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/golang-infrastructure/go-goroutine-id v0.0.0-20230917075314-f1af231c0b06 // indirect
	github.com/golang-infrastructure/go-iterator v0.0.0-20230524171120-56988a9b127c // indirect
	github.com/golang-infrastructure/go-pointer v0.0.5 // indirect
	github.com/golang-infrastructure/go-reflect-utils v0.0.0-20221130143747-965ef2eb09c3 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/storage-lock/go-memory-storage v0.0.0-20230815161904-4e2a1ca32314 // indirect
	github.com/storage-lock/go-storage v0.0.2 // indirect
	github.com/storage-lock/go-storage-events v0.0.5 // indirect
	github.com/storage-lock/go-storage-lock-factory v0.0.2-0.20230830164658-ef9d99aed23d // indirect
	github.com/stretchr/testify v1.11.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/storage-lock/go-storage => ../go-storage

replace github.com/storage-lock/go-storage-lock => ../go-storage-lock

replace github.com/storage-lock/go-memory-locks => ../go-memory-locks

replace github.com/storage-lock/go-memory-storage => ../go-memory-storage
