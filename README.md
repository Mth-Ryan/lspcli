# lspcli
Manage all your LSPs and programming tools from one cli.

https://github.com/Mth-Ryan/lspcli/assets/46976272/991f917a-ce09-4775-8295-7347165cc26b

## Usage

You can see all commands and examples with the `--help` flag. Here some comands examples:
```bash
lspcli list                               # list all tools
lspcli install typescript-language-server # install one tool
lspcli update omnisharp                   # update one tool
lspcli remove tailwind-language-server    # remove one tool
lspcli describe svelte-language-server    # show all info of one tool
```

## How to install

**Comming soon...**

## How to build for development

First you need to clone the repository:
```bash
git clone https://github.com/Mth-Ryan/lspcli
```

Then Create the runtime folders (This step will not be necessary in the future):
```bash
mkdir -p runtime/cache
mkdir -p runtime/installs
mkdir -p runtime/bin
```

Finaly build and run the app:
```bash
go build                     # build
./lspcli --runtime ./runtime # running with the working directory runtime
```

## How to contribute

Just create a fork of the repository and send a pull request.
