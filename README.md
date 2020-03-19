# Go-Wiki

Wiki web application models built using `Go`

![Go](https://img.shields.io/github/go-mod/go-version/rjNemo/go-wiki) [![License](https://img.shields.io/github/license/rjNemo/go-wiki)](LICENSE.md) ![Release](https://img.shields.io/github/v/release/rjNemo/go-wiki) ![Tag](https://img.shields.io/github/v/tag/rjNemo/go-wiki)

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

You need `Go v1.14+` to install `Go-Wiki`. [Get official installation](https://golang.org/doc/install)

### Installing

A step by step series of examples that tell you how to get a development env running

Use official go package manager to install `Go-Wiki`.

```bash
go get -u github.com/rjNemo/go-wiki
```

## Usage

In the main application folder launch the server using

```bash
go run main.go
```

## Deployment

An example app is deployed on [Heroku](https://www.heroku.com/).

Follow these steps to deploy. Compile the application:

```bash
go build -o bin/go-wiki
```

Make sure that any unused modules have been removed from your application and is resistant to erosion:

```bash
go mod tidy
go mod vendor
```

You will need an `.env` file such as:

```env
PORT=8080
TMPLDIR=views/templates/
DATABASE_URL=postgres://oshiervtwfrqrr:f177bc7ff355c142931aeb3f3
```

## Built With

- [Go](https://golang.org/) - Build simple, reliable, and efficient software
- [Bootstrap](https://getbootstrap.com/) - The most popular HTML, CSS, and JS library in the world
- [PostgreSQL](https://www.postgresql.org/) - The world's most advanced open source database
  <!-- - [Quilljs](https://quilljs.com/) - Your powerful rich text editor -->

## Contributing

Please read [CONTRIBUTING.md](contributing.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](https://semver.org/) for versioning. For the versions available, see the [CHANGELOG.md](CHANGELOG.md) and the [tags on this repository](https://github.com/rjNemo/go-wiki/tags).

## Authors

- **Ruidy Nemausat** - _Initial work_ - [rjNemo](https://github.com/rjNemo)

See also the list of [contributors](https://github.com/rjNemo/go-wiki/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

For inspiration:

- [Go MVC](https://www.calhoun.io/)
- [Alex Edwards](https://www.alexedwards.net/blog/organising-database-access)
