# Plain

Proof of concept - simple static site generator featuring templates

## Getting started 
Install go
`brew install go`

Build
`go run plain.go`

## Project Structure

```
.
├── _templates/              # Contains the base templates
│   └── baseof.html          # The main template used for rendering pages
├── _gen/                    # Output directory for generated HTML files
├── index.html               # Example of a page using baseof.html
└── plain.go                 # The Go code for building and serving the site

```