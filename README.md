# Pushup template fragments demo

This project demos Pushup's support for [template fragments][essay], which
Pushup calls "inline partials", but it's the same concept.

`app/pages/demo.up`:

```pushup
<html>
    <body>
        <div hx-target="this">
          ^partial archiveui {
            ^if contact.archived {
              <button hx-patch="/contacts/^contact.id/unarchive">Unarchive</button>
            } else {
              <button hx-delete="/contacts/^contact.id/">Archive</button>
            }
          }
        </div>
        <h3>Contact</h3>
        <p>^contact.email</p>
    </body>
</html>
```

What is Pushup? It's a new page-oriented web framework for Go. Pushup is both
a compiler (you run the `pushup` command to compile and build your app), and
a new template-like language (you write Pushup template language in `.up`
files). Pushup is trying to make developers more productive for certain
kinds of server-side web apps.

What does "page-oriented" mean? It means that the main unit of abstraction
in Pushup is the page (a `.up` file), which is a combination of Go code and
HTML markup in the same file (yes, PHP-style). The page is also the means
of composition, because Pushup uses file-based routing, mapping filenames
and directories to URLs paths (again, PHP-style).

So you write Pushup pages, and then you run Pushup which compiles them
and generates pure Go code. It builds a single artifact for your app,
a statically-linked executable which also bundles in any static files like CSS.
This makes deployment trivial, just copy the executable somewhere and run it.

## Inline partials

An inline partial is a block within a larger Pushup page that has been named
and wrapped in curly braces. In the normal course of page rendering, inline
partials are completely transparent, and just contribute their content to the
page as if they were not so wrapped. However, their content can be rendered
independently from the rest of the containing page, or any containing partial
(inline partials can be nested arbitrarily deeply). They participate in the
file-based routing, so an inline partial is available on its own via the
page's route with the partial's name (and the name of any ancestor partial)
appended as a URL path segment.

For example, if you had a page named `foo.up`, and an inline partial
named `bar`, the page would be routed at `/foo`, and the partial at
`/foo/bar`. Clients could access either as needed.

## Demo app

This Pushup app has the following project tree:

```
app/
├── layouts
│   └── default.up
├── pages
│   ├── contacts
│   │   └── $id
│   │       ├── index.up
│   │       └── unarchive.up
│   ├── demo.up
│   ├── index.up
│   └── view-source.up
├── pkg
│   └── app.go
└── static
    ├── custom.css
    ├── htmx.min.js
    └── pico.min.css
```

Run the app, either by building with `pushup build` and then executing
the resulting Go binary, or directly with `pushup run`. Visit
[http://localhost:8080/demo](http://localhost:8080/demo)

[essay]: https://htmx.org/essays/template-fragments/
