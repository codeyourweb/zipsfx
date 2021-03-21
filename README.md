# zipsfx (Self Extracting Archive)
`zipsfx` is [Go](http://golang.org/) package that allows creating __Self Extracting Archive (sfx)__ for Windows.

## Usage
This package provides a bundler which append an archive into an sfx bootstrapper. 

```go
package main

import (
	"github.com/codeyourweb/zipsfx"
)

func main() {

    // BuildSFX("inputFolder", "windows-cmd-to-execute-after-unzipping-folder", "output-sfx-executable-name")
	err := zipsfx.BuildSFX("myFolder/", "my-binary.exe", "my-sfx-package.exe")
	if err != nil {
		panic(err)
	}
}
```

## SFX bootstrapper
This package relies on Winrar zip sfx in order to provide a basic self extracting archive library

## License
[MIT](https://github.com/codeyourweb/zipsfx/LICENSE)
