# Slugify

Go library which transforms a text in a form appropriate for a URL.

## Usage

Get the package with:

```bash
go get github.com/mgechev/slugify
```

Use in your source code with:

```go
import "github.com/mgechev/slugify

slugify.Transform("Random string") // random-string
slugify.Transform("Случаен низ") // sluchaen-niz
slugify.Transform("Случаен низ с <3") // sluchaen-niz-s-love
slugify.Transform("I don't like 🕺") // i-dont-like
```

## License

MIT
