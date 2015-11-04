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

Coming soon...

###### Models

Coming soon...

### License

MIT
