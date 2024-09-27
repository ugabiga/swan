# Swan Golang Web Development Boilerplate (Under Development)

Packages used:

- [echo](https://echo.labstack.com)(web framework)
- [fx](https://uber-go.github.io/fx)(dependency injection)
- [ent](https://entgo.io)(ORM)
- [atlas](https://atlasgo.io)(migration)
- [swag](https://github.com/swaggo/swag)(swagger)
- [watermill](https://watermill.io)(pub/sub)
- [cobra](https://github.com/spf13/cobra)(CLI)
- [godotenv](https://github.com/joho/godotenv)(env)

Inspired by [caesar](http://github.com/caesar-rocks)

## Installation

```bash
go install github.com/ugabiga/swan/cli@latest
```

## Usage

Create a new project:

```bash
swan new <project-name>
```

All make commands are generate files in the given directory and add dependencies to config/app

Create a new handler:

```bash
swan make:handler
```

Create a new command:

```bash
swan make:command
```

Create a new event:

```bash
swan make:event
```

Create a struct:

```bash
swan make:struct
```
