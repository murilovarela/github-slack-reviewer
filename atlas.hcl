data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "./store/scripts/atlas-gorm-loader.go",
  ]
}

env "gorm" {
  src = data.external_schema.gorm.url
  dev = "postgres://testuser:PGU2yYtqZ+pyJraDtdj2Tkb3GgW4KcqT@localhost:5432/github_slack_bot?sslmode=disable"

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
