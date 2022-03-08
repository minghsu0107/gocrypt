# Gocrypt

![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/minghsu0107/gocrypt?label=Version&sort=semver)

Gocrypt is a TUI based application that monitors cryptocurrency prices in real time, written in Go.

<img width="1401" alt="image" src="https://user-images.githubusercontent.com/50090692/157175880-584d8b11-fb64-4826-b9a6-96822c7e9a8a.png">

## Installation

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
VERSION=v1.0.0 make docker-pull

# Run image
make docker-run

# Run specific version of image
VERSION=v1.0.0 make docker-run

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

Gocrypt helps you look at cryptocurrency values, details and track your crypto portfolio straight from your terminal.

It makes use of the API provided by [CoinCap.io](https://coincap.io/) and [CoinGecko](https://www.coingecko.com/en) to provide the required details.

### Main Page
- Top 3 currencies (as ranked by Market Cap) are displayed with their graphs on top.
- A coin table is provided with relevant information about other currencies.
- `gocrypt` allows you to keep track of your favorite currencies by adding them to the favourites table.
- A selected coin from either the coin table or favorites can be further inspected in detail.

Key-bindings can be found by pressing `?`. This displays the help prompt.

-	**Quit**: `q` or `<C-c>`
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
	-	`<c>`: Select Currency (from popular list)
	-	`<C>`: Select Currency (from full list)
	-	`e`: Add/Edit coin to Portfolio
	-	`P`: View portfolio
	-	`<s>`: Star, save to favourites
	-	`<S>`: UnStar,remove from favourites
	-	`<Enter>`: View Coin Information
	-	`%`: Select Duration for Percentage Change

### Coin Page
- The coin page gives more details of a particular coin.
- It can be navigated to from either the favourites or coin table.
- The price history is displayed on top and can be viewed through different intervals, as provided by the Graph Interval table on the bottom.
- A live price is streamed in the price box and additional details are described in the details table.

Key-bindings can be found by pressing `?`. This displays the help prompt.

-	**Quit**: `q` or `<C-c>`
-	**Table Navigation**
	-	`k` and `<Up>`: up
	-	`j` and `<Down>`: down
	-	`<C-u>`: half page up
	-	`<C-d>`: half page down
	-	`<C-b>`: full page up
	-	`<C-f>`: full page down
	-	`gg` and `<Home>`: jump to top
	-	`G` and `<End>`: jump to bottom
	-	`f`: focus favourites table
	-	`F`: focus explorers table
-	**Sorting (Favourites Table)**
	-	Use column number to sort ascending.
	-	Use `<F-column number>` to sort descending.
	-	Eg: `1` to sort ascending on 1st Col and `F1` for descending
-	**Actions**
	-	`<c>`: Select Currency (from popular list)
	-	`<C>`: Select Currency (from full list)

### Portfolio Page
- GOcrypt allows you to track your crypto portfolio through a separately defined page.
- This page can be accessed with the command `gocrypt portfolio`.

Key-bindings:

-	**Quit: `q` or `<C-c>`**
-	**Table Navigation**
	-	`k` and `<Up>`: up
	-	`j` and `<Down>`: down
	-	`<C-u>`: half page up
	-	`<C-d>`: half page down
	-	`<C-b>`: full page up
	-	`<C-f>`: full page down
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
