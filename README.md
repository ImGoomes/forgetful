<div align="center">
  <h1>Forgetful</h1>
  <p><em>Never forget a terminal command again.</em></p>
</div>

<p align="center">
  <img src="img/forg.png" alt="Dog with question mark" />
</p>

## Install

### Homebrew (macOS and Linux)

```sh
brew install imgoomes/tap/forg
```

### Download binary (Linux and macOS)

Download the latest release from the [releases page](https://github.com/imgoomes/forgetful/releases), extract and move to your PATH:

```sh
# macOS (Apple Silicon)
curl -L https://github.com/imgoomes/forgetful/releases/latest/download/forgetful_darwin_arm64.tar.gz | tar -xz
sudo mv forg /usr/local/bin/

# macOS (Intel)
curl -L https://github.com/imgoomes/forgetful/releases/latest/download/forgetful_darwin_amd64.tar.gz | tar -xz
sudo mv forg /usr/local/bin/

# Linux (amd64)
curl -L https://github.com/imgoomes/forgetful/releases/latest/download/forgetful_linux_amd64.tar.gz | tar -xz
sudo mv forg /usr/local/bin/

# Linux (arm64)
curl -L https://github.com/imgoomes/forgetful/releases/latest/download/forgetful_linux_arm64.tar.gz | tar -xz
sudo mv forg /usr/local/bin/
```

### Go install

```sh
go install github.com/imgoomes/forgetful/cmd/forg@latest
```

### Build from source

```sh
git clone https://github.com/imgoomes/forgetful
cd forgetful
go build -o forg ./cmd/forg
```

## Usage

### Save a command

```sh
forg add -c "git config user.name 'Gabriel'" -t git -d "Change local git user"
forg add -c "docker ps -a" -d "List all containers"
forg add -c "ffmpeg -i input.mp4 -vn output.mp3"
```

Flags:

- `-c` — command to save (required)
- `-t` — tag (optional; defaults to the first word of the command)
- `-d` — description (optional)

### List saved commands

```sh
forg list          # all commands
forg list -t git   # filter by tag
forg ls -t docker
```

### Run a command by ID

```sh
forg run 1
```

### Delete a command

```sh
forg delete 3   # prompts for confirmation
forg rm 3
forg remove 3
```

### List tags

```sh
forg tags
```

## Storage

Commands are stored in `~/.forg/commands.json`.

## License

MIT
