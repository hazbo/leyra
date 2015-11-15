# Leyra

Lerya is a web application framework of sorts. It's being activly developed and
because of that at this stage there are a few dependencies (Go packages mainly)
that come as part of it. The aim is to continue making this more modular to the
point where certain dependencies are only pulled in when they are needed.

As of now, Leyra is being built on the back of an example blog web application
which may give you an idea on how it can be used.

Along-side the development of Leyra, a cli tool called Gresh is also being
developed to manage the framework itself and assit the user when developing a
web application usig Leyra. Although Gresh is still being activly developed,
the usage section here will use Gresh to show you how to get started.

### Requirements

  - Go (tested on v1.5)
  - Make
  - Gresh (currently optional)

### Install

As of right now Gresh can create a new blank project for you or pull down an
open source one from GitHub. I'll cover both in this example:

##### Installing Gresh

Set your `GOPATH`

```bash
$ git clone git@github.com:leyra/gresh.git
$ cd gresh
$ make
$ make install
```

##### New Project

```bash
$ gresh new my_app
$ export GOPATH=$(pwd)/my_app
$ cd my_app/src/leyra
$ make run
```

Head over to http://localhost:3000 in your browser now.

### Usage

The `GOPATH` for a new Leyra project is a local path, relative to the project.
This happens to be the case for now, but may change at some point. It keeps
everything a bit cleaner for now, especially as it is being developed.

At this point you'll have a pretty empty version of Leyra running, with just
the text "Hello, World!" being displayed in your browser. Although things may
change, for now this is how you can get srated with your app.

###### Routes

Routes are located in `app/http/routes.go`. Adding new routes is simple, you'll
see a Home Controller route in there as an example, here's how we'd add another
one:

```go
func Route() *echo.Echo {
	e := echo.New()
	e.Get("/", routeHome)
	e.Get("/about", routeAbout)

	return e
}

func routeHome(c *echo.Context) error {
	return new(controller.Home).Home(c)
}

func routeAbout(c *echo.Context) error {
	return new(controller.Home).About(c)
}
```

This will route `/about` to the Home Controller and will call `About(c)`. So the
Home Controller file located at `app/http/controllers/home.go` will not have
an About func yet. We can add one though:

```go
type Home struct {
}

func (h Home) Home(c *echo.Context) error {
	return c.HTML(http.StatusOK, "Hello, World!")
}

func (h Home) About(c *echo.Context) error {
	return c.HTML(http.StatusOK, "My about page!")
}
```

At this point, running `make run` inside your `leyra` directory will compile the
code and start the server. Heading to `http://localhost:3000/about` will now
display the text you return in this controller.

###### Views

Views make use of Go's `html/template` package. They can be located at
`app/views`. Each view is just a regular HTML file, although they can make use
of variables or anything else supported within the `html/template` package.

Views are loaded in and cached initially on runtime. Eventually I'll get round
to allowing caching to be turned on / off - probably somewhere in `etc/rc.conf`.

Views can be accessed from the `Store` object (`S`), like so:

```go
app.S.View.ExecuteTemplate(buffer, "index.html", nil)
```

This will load in `app/views/index.html`, using no data to be passed through. A
buffer can easily be created which is where the executed contents of your
template will be stored. For example, in a controller you could do something
like this:

```go
package controller

import (
	"bytes"
	"net/http"

	"gopkg.in/leyra/echo.v1"

	"leyra/app"
)

type Home struct {
}

type IndexView struct {
	FirstName string
	LastName  string
}

func (h Home) Home(c *echo.Context) error {
	buff := new(bytes.Buffer)

	app.S.View.ExecuteTemplate(buff, "hello.html", IndexView{
		FirstName: "John",
		LastName: "Smith",
	})

	return c.HTML(http.StatusOK, buff.String())
}
```

With a corresponding 'app/views/hello.html' looking something like this:

```html
<html>
	<head>
		<title>Hello!</title>
	</head>
	<body>
		<h1>Hello {{ .FirstName }} {{ .LastName }}!</h1>
	</body>
</html>
```

This ties our controllers and views together. Next we'll look at using some real
data here.

###### Models

Coming soon...

### License

MIT
