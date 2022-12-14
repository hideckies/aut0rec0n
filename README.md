# aut0rec0n

A automatic reconnaissance tool.

- **DNS** - Golang `net` package
- **Subdomain** - Google, [rocketlaunchr/google-search](https://github.com/rocketlaunchr/google-search)
- **Web Archives** - Wayback Machine API
- **WHOIS** - [domainr/whois](https://github.com/domainr/whois)
    

<br />

## Usage

```sh
aut0rec0n example.com

# Specify scripts
aut0rec0n example.com --script dns,subdomain
# Output results to given folder
aut0rec0n example.com -o results
```

To print all scripts:

```sh
aut0rec0n --script-list
```

<br />

## Installation

- **Option 1. Go install**

    ```sh
    go install github.com/hideckies/aut0rec0n@latest
    ```

- **Option 2. Clone the Repo**

    ```sh
    git clone https://github.com/hideckies/aut0rec0n.git
    cd aut0rec0n
    go get ; go build
    ```

