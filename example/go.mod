module example

go 1.20

replace github.com/danteay/golog => ../

require github.com/danteay/golog v0.0.0-00010101000000-000000000000

require (
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/rs/zerolog v1.31.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
)
