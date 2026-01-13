variable "envfile" {
 type    = string
 default = ".env"
}

locals {
 envfile = {
   for line in split("\n", file(var.envfile)): split("=", line)[0] => regex("=(.*)", line)[0]
   if !startswith(line, "#") && length(split("=", line)) > 1
 }
}

data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "./cmd/loader/main.go",
  ]
}

env "gorm" {
  src = data.external_schema.gorm.url
  url = local.envfile["DATABASE_URL"]
  dev = local.envfile["DEV_DATABASE_URL"]
  migration {
    dir = "file://migrations"
    format = golang-migrate
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}
