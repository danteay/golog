module github.com/danteay/golog

go 1.20

replace github.com/danteay/golog/levels => ./levels

replace github.com/danteay/golog/fields => ./fields

replace github.com/danteay/golog/adapters/zerolog => ./adapters/zerolog

require (
	github.com/danteay/golog/adapters/zerolog v0.0.0-00010101000000-000000000000
	github.com/danteay/golog/fields v0.0.0-00010101000000-000000000000
	github.com/danteay/golog/levels v0.0.0-00010101000000-000000000000
	github.com/magefile/mage v1.15.0
	github.com/stretchr/testify v1.8.4
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rs/zerolog v1.31.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
