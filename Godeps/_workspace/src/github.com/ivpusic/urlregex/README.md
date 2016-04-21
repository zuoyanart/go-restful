urlregex
======
[![Build Status](https://travis-ci.org/ivpusic/urlregex.svg?branch=master)](https://travis-ci.org/ivpusic/urlregex)

express-like named url parameters extracting from url

Library will generate regex based on provided url pattern. Later you will be able to match against some url, and read named params values if they are present.

## Example
```Go
package main

import (
	"fmt"
	"github.com/ivpusic/urlregex"
)

func main() {
package main

import (
	"fmt"
	"github.com/ivpusic/urlregex"
)

func main() {
	reg := urlregex.Pattern("some/:name/path/:other/")
	fmt.Println("regex: " + reg.Regex.String())

	res, err := reg.Match("some/123/path/456/")
	if err != nil {
		fmt.Println("no matches")
		return
	}

	fmt.Println("found matches")
	for k, v := range res {
		fmt.Println(k + ": " + v)
	}
}
```

This will output:
```
regex: ^some\/(?P<name>.[^\/]*)\/path\/(?P<other>.[^\/]*)\/$
found matches
name: 123
other: 456
```

You can also pass url pattern without named parameters, and later check if given url matches.
```Go
package main

import (
	"fmt"
	"github.com/ivpusic/urlregex"
)

func main() {
	reg := urlregex.Pattern("some/123/path/456")
	
	// in this case let's pass something invalid
	_, err := reg.Match("some/123/path/invalid")

	if err != nil {
		fmt.Println("not matched")
		return
	}

	fmt.Println("matched")
}
```
This will output
```
not matched
```
However if we passed ``some/123/path/456`` to ``Match`` method, then it would output ``matched``.

#### Access actual go regex instance
```Go
package main

import (
	"fmt"
	"github.com/ivpusic/urlregex"
)

func main() {
	reg := urlregex.Pattern("some/:name/path/:other/")
	// native generated *Regex instance
	fmt.Println(reg.Regex)
}
```

#### Wildcards
```Go
package main

import (
	"fmt"
	"github.com/ivpusic/urlregex"
)

func main() {
	reg := urlregex.Pattern("some/*")
	
	// this will be matched
	reg.Match("some/123")
	
	// this also
	reg.Match("some/blabla/blabla")
	
	// but this won't
	reg.Match("some")
	
	// you can combine wildcards and url params
	reg = urlregex.Pattern("/some/:name/*/path")
	
	// this will be matched
	reg.Match("/some/user/blabla/path")
	
	// but this won't
	reg.Match("/some/user/blabla/missing")
	
	// named wildcards
	reg := urlregex.Pattern("some/*key")
	res, _ := reg.Match("/some/this/is/a/key")
	fmt.Println(res["key"])
	//--> this/is/a/key
}
```

# License
MIT
