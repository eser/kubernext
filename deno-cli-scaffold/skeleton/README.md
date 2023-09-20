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

- Ensure that `deno` is installed on your machine

```bash
$ deno --version
deno 1.36.4 (release, aarch64-apple-darwin)
v8 11.6.189.12
typescript 5.1.6
```

## Execution

- Run the project in dev mode

```bash
$ deno task dev
Task dev deno run --unstable --watch ./src/main.ts
Watcher Process started.
Add 2 + 3 = 5
```
