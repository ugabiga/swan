# Swan Golang Web Development Boilerplate (Under Development)

Packages used:

- [echo](https://echo.labstack.com)(web framework)
- [fx](https://uber-go.github.io/fx)(dependency injection)
- [atlas](https://atlasgo.io)(migration)
- [swag](https://github.com/swaggo/swag)(swagger)
- [watermill](https://watermill.io)(pub/sub)
- [cobra](https://github.com/spf13/cobra)(CLI)
- [godotenv](https://github.com/joho/godotenv)(env)

Inspired by [caesar](http://github.com/caesar-rocks)

## Installation

```bash
go install github.com/ugabiga/swan/swctl@latest
```

## Usage

Create a new project:

```bash
swctl new <project-name>
```

All make commands are generate files in the given directory and add dependencies to config/app

Create a new handler:

command:
```bash
swctl make:handler [folder-path] [api-prefix] [endpoint-name]
```

example:
```bash
swctl make:handler todos /api/v1 todos
```

Create a new command:
This command will generate a new cobra command in the given folder path named command.go

command:
```bash
swctl make:command [folder-path]
```

example:
```bash
swctl make:command todos
```

Create a new event:
This command will generate a new event in the given folder path named event.go

command:
```bash
swctl make:event [folder-path]
```

example:
```bash
swctl make:event todos
```

Create a struct:
This command will generate a new struct in the given folder path named by the struct name
and add it to container.go

command:
```bash
swctl make:struct [folder-path] [struct-name]
```

example:
```bash
swctl make:struct todos Todo
```


