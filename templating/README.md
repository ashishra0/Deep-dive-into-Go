### Type Template
<h3>template.Template</h3>

```go
template.Template
```
---

### Parsing Templates

<h3>template.ParseFiles</h3>

```go
func ParseFiles(filenames ...string) (*Template, error)
```

<h3>template.ParseGlob</h3>

```go
func parseGlob(pattern string) (*Template, error)
```
---
<h3>template.ParseFiles</h3>

```go
func (t *Template) ParseFiles(filename ...string) (*Template, error)
```
<h3>template.ParseGlob</h3>

```go
func (t *Template) ParseGlob(filename ...string) (*Template, error)
```
---

### Executing Templates

<h3>template.Execute</h3>

```go
func (t *Template) Execute(wr io.writer, data interface{}) error
```

<h3>template.ExecuteTemplate</h3>

```go
func (t *Template) ExecuteTemplate(wr io.writer, name string, data interface{}) error
```
---

### Helper methods

<h3>template.Must</h3>

```go
func Must(t *Template, err error) *Template
```

<h3>template.New</h3>

```go
func New(name string) *Template
```
---