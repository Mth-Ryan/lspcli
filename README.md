# lspcli
Manage all your LSPs and programming tools from one cli.

https://github.com/Mth-Ryan/lspcli/assets/46976272/991f917a-ce09-4775-8295-7347165cc26b

## Usage

You can see all commands and examples with the `--help` flag. Here some commands examples:
```bash
lspcli list                               # list all tools
lspcli install typescript-language-server # install one tool
lspcli update omnisharp                   # update one tool
lspcli remove tailwind-language-server    # remove one tool
lspcli describe svelte-language-server    # show all info of one tool
```

## Installing

### From Source:

For the installation you will need some dependencies: `go`, `git` and `GNU Make`:
```bash
# Clone this repo:
git clone https://github.com/Mth-Ryan/lspcli

# Install:
make install

# You can change the the installation folder (not recommended) using:
# make install INSTALL_DIR=other_location.
```

Then, add the respective env script to your shell init file. Example `~/.bashrc` or `~/.zshrc`:

```bash
. ~/.local/share/lspcli/env.bash # For GNU Bash
. ~/.local/share/lspcli/env.zsh  # For Zsh
. ~/.local/share/lspcli/env.ps1  # For Powershell
```

This will only set the lspcli bin folder to your path, you can do
manually for any other shell and the program will work as expected.
  

## How to build for development

You will need the same dependencies of the [installation](#installing). With this dependencies installed, build the
cli and the scripts:

```bash
# Clone this repo:
git clone https://github.com/Mth-Ryan/lspcli

# Build all:
make # or make all

# Run with the current dir runtime:
./bin/lspcli --runtime ./runtime
```

## How to contribute

Just create a fork of the repository and send a pull request.
