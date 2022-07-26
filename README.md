# Gocrypt

![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/minghsu0107/gocrypt?label=Version&sort=semver)

Gocrypt is a TUI based application that monitors cryptocurrency prices in real time, written in Go.

![image](https://user-images.githubusercontent.com/50090692/160075949-2bd0ef01-5a4a-433f-9fef-a2357f646cb5.png)

## Installation
### Using Binary
For Mac (amd64):
```bash
VERSION=v1.0.2
wget "https://github.com/minghsu0107/gocrypt/releases/download/${VERSION}/gocrypt-${VERSION}-darwin-amd64.tar.gz"
tar -xzvf "gocrypt-${VERSION}-darwin-amd64.tar.gz"
```
For Mac (arm64):
```bash
VERSION=v1.0.2
wget "https://github.com/minghsu0107/gocrypt/releases/download/${VERSION}/gocrypt-${VERSION}-darwin-arm64.tar.gz"
tar -xzvf "gocrypt-${VERSION}-darwin-arm64.tar.gz"
```
For Linux:
```bash
VERSION=v1.0.2
wget "https://github.com/minghsu0107/gocrypt/releases/download/${VERSION}/gocrypt-${VERSION}-linux-amd64.tar.gz"
tar -xzvf "gocrypt-${VERSION}-linux-amd64.tar.gz"
```
### Using Go

```bash
go get -u github.com/minghsu0107/gocrypt
```

### Using Docker

Set `VERSION` to a specific version for stable builds. Omitting `VERSION` uses the latest stable version or setting `main` as version provides the latest development version.

```bash
# Pull Image
make docker-pull

# Pull specific version of image
VERSION=v1.0.2 make docker-pull

# Run image
make docker-run

# Run specific version of image
VERSION=v1.0.2 make docker-run

# Run image with portfolio command
ARG=portfolio make docker-run

# Run image with help command
ARG=help make docker-run

```

### Building Image locally

```bash
# Clone the repository
git clone https://github.com/minghsu0107/gocrypt

# Navigate into repository
cd gocrypt

# Build image
make docker-build

# Run image
make docker-run
```

### Building From Source

Building requires [Go](https://golang.org) to be installed.

```bash
# Clone the repository
git clone https://github.com/minghsu0107/gocrypt

# Navigate into repository
cd gocrypt

# Build executable
make build
```

## Usage
```
Usage:
  gocrypt [flags]
  gocrypt [command]

Available Commands:
  help        Help about any command
  portfolio   Track your portfolio
  version     Print the current version

Flags:
  -c, --config string               config file (default is $HOME/.gocrypt.yaml)
  -i, --currency-init-unit string   initial currency unit
  -h, --help                        help for gocrypt
  -u, --user string                 root user
```

Gocrypt makes use of the API provided by [CoinCap.io](https://coincap.io/) and [CoinGecko](https://www.coingecko.com/en) to provide the required details.
### Main Page
- Top 3 currencies (as ranked by Market Cap) are displayed with their graphs on top.
- A coin table is provided with relevant information about other currencies.
- `gocrypt` allows you to keep track of your favorite currencies by adding them to the favourites table.
- A selected coin from either the coin table or favorites can be further inspected in detail.

Key-bindings can be found by pressing `?`. This displays the help prompt.

-	**Quit**: `q` or `<Ctrl-c>`
-	**Table Navigation**
	-	`k` and `<Up>`: up
	-	`j` and `<Down>`: down
	-	`<Ctrl-u>`: half page up
	-	`<Ctrl-d>`: half page down
	-	`<Ctrl-b>`: full page up
	-	`<Ctrl-f>`: full page down
	-	`gg` and `<Home>`: jump to top
	-	`G` and `<End>`: jump to bottom
	-	`f`: focus favourites table
	-	`F`: focus coin table
-	**Searching/Filtering**
	-	`/`: Open search box
	-	`Esc`: Clear filter text
-	**Sorting**
	-	Use column number to sort ascending.
	-	Use `F-column number` to sort descending.
		-  Mac user should press `fn + F-column number`
	-	Eg: 1 to sort ascending on 1st Col and F1 for descending
-	**Actions (Coin Table)**
	-	`c`: Select Currency (from popular list)
	-	`C`: Select Currency (from full list)
	-	`e`: Add/Edit coin to Portfolio
	-	`P`: View portfolio
	-	`s`: Star, save to favourites
	-	`S`: UnStar,remove from favourites
	-	`<Enter>`: View Coin Information
	-	`%`: Select Duration for Percentage Change

### Coin Page
- The coin page gives more details of a particular coin.
- It can be navigated to from either the favourites or coin table.
- The price history is displayed on top and can be viewed through different intervals, as provided by the Graph Interval table on the bottom.
- A live price is streamed in the price box and additional details are described in the details table.

Key-bindings can be found by pressing `?`. This displays the help prompt.

-	**Quit**: `q` or `<Ctrl-c>`
-	**Table Navigation**
	-	`k` and `<Up>`: up
	-	`j` and `<Down>`: down
	-	`<Ctrl-u>`: half page up
	-	`<Ctrl-d>`: half page down
	-	`<Ctrl-b>`: full page up
	-	`<Ctrl-f>`: full page down
	-	`gg` and `<Home>`: jump to top
	-	`G` and `<End>`: jump to bottom
	-	`f`: focus favourites table
	-	`F`: focus explorers table
-	**Sorting (Favourites Table)**
	-	Use column number to sort ascending.
	-	Use `<F-column number>` to sort descending.
	-	Eg: `1` to sort ascending on 1st Col and `F1` for descending
-	**Actions**
	-	`c`: Select Currency (from popular list)
	-	`C`: Select Currency (from full list)
	-	`d`: Select time interval

### Portfolio Page
- Gocrypt allows you to track your crypto portfolio through a separately defined page.
- This page can be accessed with the command `gocrypt portfolio`.

```
Usage:
  gocrypt portfolio [flags]

Flags:
  -h, --help           help for portfolio
  -p, --puser string   portfolio user

Global Flags:
  -c, --config string               config file (default is $HOME/.gocrypt.yaml)
  -i, --currency-init-unit string   initial currency unit
```

Key-bindings:

-	**Quit: `q` or `<Ctrl-c>`**
-	**Table Navigation**
	-	`k` and `<Up>`: up
	-	`j` and `<Down>`: down
	-	`<Ctrl-u>`: half page up
	-	`<Ctrl-d>`: half page down
	-	`<Ctrl-b>`: full page up
	-	`<Ctrl-f>`: full page down
	-	`gg` and `<Home>`: jump to top
	-	`G` and `<End>`: jump to bottom
-	**Sorting**
	-	Use column number to sort ascending.
	-	Use `<F-column number>` to sort descending.
	-	Eg: `1` to sort ascending on 1st Col and `F1` for descending
-	**Actions**
	-	`c`: Select Currency (from popular list)
	-	`C`: Select Currency (from full list)
	-	`e`: Add/Edit coin to Portfolio
	-	`<Enter>`: View Coin Information

### Mini Portfolio
-	Gocrypt also allows you to view your holdings through a mini portfolio from other pages.
-	Coins can be added/modified/removed by pressing `e` on a coin in the main page. (Set Holding Amount as 0 to remove)
-	Holdings can be modified either through the main page or through the portfolio itself. The below image shows the edit box when modifying holdings.
