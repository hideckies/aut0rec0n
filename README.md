# aut0rec0n

A automatic reconnaissance tool.

- **DNS**
- **Port Scanning**
- **Subdomain**
    
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

