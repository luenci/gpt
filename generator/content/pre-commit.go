package content

var PreCommit = `# See https://pre-commit.com for more information
# See https://pre-commit.com/hooks.html for more hooks
repos:
  - repo: https://github.com/pre-commit/pre-commit-hooks
    rev: v4.1.0
    hooks:
      - id: trailing-whitespace
      - id: end-of-file-fixer
      - id: check-yaml
      - id: check-added-large-files

  - repo: https://github.com/dnephin/pre-commit-golang
    rev: v0.4.0
    hooks:
      - id: go-fmt
      - id: go-vet
      - id: go-lint
      - id: go-imports
        # - id: go-cyclo
        # args: [-over=15]
      - id: validate-toml
      # - id: no-go-testing
      - id: golangci-lint
      # - id: go-critic
      # - id: go-unit-tests
      # - id: go-build
      - id: go-mod-tidy
`
