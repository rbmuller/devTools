# devTools

This is an open package to help developers solve daily trivial coding problems.

There are functions to print custom error messages indicating file and line number, custom functions like Stringfy or Chunk an Array.

Feel free to join our community and help to improve **DevTools**

## How to use

Import it into your code and start using.

```bash
go get github.com/rbmuller/devtools
```

Features Catalog:

1. Errors
"github.com/rbmuller/devtools/errors"

2. Filters
"github.com/rbmuller/devtools/filters"

3. Logging
"github.com/rbmuller/devtools/logging"

4. Mail
"github.com/rbmuller/devtools/mail"

5. Process
"github.com/rbmuller/devtools/process"

**Errors** pkg can help finding the file and line of code with problem implementing a custom Error

```bash
devErrors.NewError(err)
```

For more example visit [example.go](./example.go)
