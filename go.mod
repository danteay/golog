module github.com/danteay/golog

go 1.20

replace (
	github.com/danteay/golog/adapters/slog => ./adapters/slog
	github.com/danteay/golog/fields => ./fields
	github.com/danteay/golog/levels => ./levels
)

require (
	github.com/danteay/golog/adapters/slog v0.0.0
	github.com/danteay/golog/fields v0.0.0
	github.com/danteay/golog/levels v0.0.0
	github.com/stretchr/testify v1.9.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
