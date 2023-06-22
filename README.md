# aut0rec0n

An automatic reconnaissance tool.

- DNS
- Port Scanning
- Subdomain


<br />

## APIs/Resources

aut0rec0n fetches information from the following APIs and resources.

- Google
- Shodan
- VirusTotal

Some APIs need API keys. We can set them into **`~/.config/aut0rec0n/config.yaml`**. This file will be automatically generated after the first running **aut0rec0n**.
    
<br />

## Usage

```sh
aut0rec0n -H example.com

# Specify a method
aut0rec0n dns -H example.com
aut0rec0n port -H example.com
aut0rec0n subdomain -H example.com
```

<br />

## Installation

### Option 1. Go install

```sh
go install github.com/hideckies/aut0rec0n@latest
```

### Option 2. Clone the Repo

```sh
git clone https://github.com/hideckies/aut0rec0n.git
cd aut0rec0n
go get ; go build
```

<br />

## Set API Keys