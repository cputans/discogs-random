# Discogs Random Album Selector

A Go-based CLI tool that fetches a random album from your Discogs collection.

## Prerequisites

- Go 1.21 or higher
- A Discogs account
- A Discogs personal access token

## Getting a Personal Access Token

1. Go to https://www.discogs.com/settings/developers
2. Click "Generate new token"
3. Give it a name (e.g., "Random Album Selector")
4. Copy the generated token

## Installation

### Option 1: Download Pre-built Binary

Download the latest release for your platform from the [Releases page](../../releases):
- **macOS (Intel)**: `discogs-random-darwin-amd64`
- **macOS (Apple Silicon)**: `discogs-random-darwin-arm64`
- **Linux**: `discogs-random-linux-amd64`
- **Windows**: `discogs-random-windows-amd64.exe`

Then make it executable (macOS/Linux):
```bash
chmod +x discogs-random
```

### Option 2: Build from Source

Prerequisites: Go 1.21 or higher

```bash
git clone https://github.com/yourusername/discogs-random.git
cd discogs-random
go build -o discogs-random
```

## Usage

### Option 1: Command-line flags
```bash
./discogs-random -username <your-username> -token <your-token>
```

### Option 2: Environment variables
```bash
export DISCOGS_USERNAME=<your-username>
export DISCOGS_TOKEN=<your-token>
./discogs-random
```

### Example
```bash
export DISCOGS_USERNAME=myusername
export DISCOGS_TOKEN=AbCdEfGhIjKlMnOpQrStUvWxYz
./discogs-random
```

## Output

The tool will display:
- Album title
- Artist(s)
- Release year
- Your rating (if set)
- Date added to collection
- Discogs ID and direct link

## Features

- Handles pagination automatically (fetches all collection items)
- Uses Discogs personal access token for authentication
- Displays formatted album information
- Per-page limit of 250 items for efficiency

## Notes

- The tool requires internet connection to access the Discogs API
- Your personal access token should be kept private and not committed to version control

## CI/CD

This project uses GitHub Actions to automatically build binaries for all platforms on every push and pull request:
- Linux (x86-64)
- macOS (Intel x86-64 and Apple Silicon ARM64)
- Windows (x86-64)

Build artifacts are automatically uploaded to GitHub releases when a git tag is pushed. To create a release and attach the built binaries:

1. Create a release on GitHub (via the web UI or with `gh release create`)
2. Push a git tag to trigger the build:

```bash
git tag v1.0.0
git push origin v1.0.0
```

The build artifacts will be automatically uploaded to the release corresponding to the pushed tag.
