module example.com/hello3

go 1.20

replace example.com/greetings => ./greetings

require (
    example.com/greetings v0.0.0-00010101000000-000000000000
)
