# E Terminal

E Terminal is a CLI designed for easliy accessing socio-economic. The design is centered around accessing data from public APIs, (Census, Fred), 
but with the possiblity to configure other APIs during runtime.

## Installation

In the root directory run the following:

```bash

 go build -o eterm.exe

```

## Commands

#### init

```bash

./eterm init

```

creates a config.toml file in the current working directory with some preconfigured APIs

#### add

```bash

./eterm add

```

addes an API to the config.toml

#### edit

```bash

./eterm edit

```

opens the config.toml file in a editable form
