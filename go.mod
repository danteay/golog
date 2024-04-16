module github.com/danteay/golog

go 1.22

replace (
	github.com/danteay/golog/adapters/zerolog => ./adapters/zerolog
	github.com/danteay/golog/fields => ./fields
	github.com/danteay/golog/levels => ./levels
)

require (
	github.com/danteay/golog/adapters/zerolog v0.0.0
	github.com/danteay/golog/fields v0.0.0
	github.com/danteay/golog/levels v0.0.0
	github.com/magefile/mage v1.15.0
	github.com/stretchr/testify v1.9.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rs/zerolog v1.32.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
