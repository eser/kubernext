# ${{ values.name }}

{%- if values.description %}
${{ values.description }}
{%- endif %}

## How to get contributed to this project?

- Clone this repository into your projects folder
- Execute required installation commands (see [Installation](#installation) below)
- Create a new branch for your feature
- Make your changes
- Test your changes (see [Execution](#execution) below)
- Commit your changes
- Push your changes to your branch
- Create a pull request to the `main` branch
- Wait for review and merge
- Delete your branch

## Installation

- Ensure that `golang` tools are installed on your machine

```bash
$ go version
go version go1.21.1 darwin/arm64
```

- Ensure that you can access private dependencies

You need to get a Personal Access Token from your GitHub account in order to download private dependencies.

To get these, visit https://github.com/settings/tokens/new and create a new token with the `read:packages` scope.

Then, you need to create or edit the `.netrc` file in your home directory with the following content:

```
machine github.com login <your-github-username> password <your-github-access-token>
```

- Download required modules for the project

```bash
$ go mod download
```

## Execution

- Run the project in test mode

```bash
$ go build && ./go-service
[GIN-debug] Listening and serving HTTP on :8080
```

- You can access http://localhost:8080/ to check if the project is running
