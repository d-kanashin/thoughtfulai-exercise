## Description
This repository is a 30-minute solution of the exercise described at [this page](https://thoughtfulautomation.notion.site/Core-Engineering-Technical-Screen-b61b6f6980714c198dc49b91dd23d695)

## How to run

As the main purpose of the programm is the `sort` function, the `main` function doesn't do anything useful. Please run tests to evaludate the `sort` function.

Note: Make sure `go` is installed.
```bash
go test *.go
```

## Note
New package type `INVALID` was introduced to handle negative inputs (like `-5Kg`) and to keep the `sort` function's signature as requested.
Though, this isn't idiomatic for Go.

The Go's solution would be:
```go
type Package struct {
	Width  int
	Height int
	Length int
	Mass   int
}

func (p Package) isHeavy() bool {
	return p.Mass > 20
}

func (p Package) isBulky() bool {
	return (p.Width*p.Height*p.Length) >= 1_000_000 ||
		p.Width >= 150 ||
		p.Height >= 150 ||
		p.Length >= 150
}

func (p Package) Sort() (packageType, error) {
	if p.Width < 0 || p.Height < 0 || p.Length < 0 || p.Mass < 0 {
		return "", fmt.Errorf("negative values are not allowed")
	}

	isBulky := p.isBulky()
	isHeavy := p.isHeavy()

	if isBulky && isHeavy {
		return rejected, nil
	} else if isBulky || isHeavy {
		return special, nil
	}
	return stardard, nil
}

func main() {
	fmt.Println("This is a package sorting program")
	p := Package{
		Width:  10,
		Height: 160,
		Length: 10,
		Mass:   15,
	}
	pType, err := p.Sort()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(pType)
}
```
