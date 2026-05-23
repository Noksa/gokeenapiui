# Gokeenapi UI

A graphical user interface for managing AWG (AmneziaWG) VPN connections on Keenetic routers.

This project is a GUI version of the [gokeenapi](https://github.com/Noksa/gokeenapi) utility.

## Feature Scope

This application covers **AWG (AmneziaWG) connection management only**. It does not support standard WireGuard or other VPN types. Supported operations:

- Create a new AWG VPN connection from a `.conf` file
- Add static routes to a WireGuard/AWG interface (from `.bat` files or URLs)
- Delete all static routes from an interface
- View all WireGuard interfaces on the router

Requires a Keenetic router running in **Router mode** with the **WireGuard VPN** component installed. Extender/repeater mode is not supported.

## Description

The application allows you to easily configure AWG VPN connections on Keenetic routers through a convenient graphical interface. Supports connection to both local routers and via the internet using KeenDNS.

## Installation

### Pre-built Binaries

Download the latest release from the [Releases page](https://github.com/Noksa/gokeenapiui/releases).

| Platform | File | Notes |
|----------|------|-------|
| Windows (x64) | `gokeenapiui.exe` | No installer needed — run directly |
| macOS (Apple Silicon) | Built from source | See [Build from Source](#build-from-source) |

Latest release: [v0.2.1](https://github.com/Noksa/gokeenapiui/releases/tag/v0.2.1)

### Build from Source

**Prerequisites:**

- [Go 1.21+](https://go.dev/dl/)
- [Node.js 18+](https://nodejs.org/)
- [Wails v2](https://wails.io/docs/gettingstarted/installation): `go install github.com/wailsapp/wails/v2/cmd/wails@latest`

**Build:**

```bash
# Clone the repository
git clone https://github.com/Noksa/gokeenapiui.git
cd gokeenapiui

# Build for your current platform
wails build

# Or build specific platforms (requires cross-compilation toolchain)
wails build -platform windows/amd64
wails build -platform darwin/arm64
```

The binary is written to `build/bin/`.

Alternatively, use the Makefile to build all supported platforms at once:

```bash
make binaries
```

## Features

- 🌐 Connect to router via HTTP/HTTPS
- 📄 Load AWG configuration files
- ⚡ Automatic VPN connection creation
- 🔒 Secure authentication
- 🛣️ Manage static routes (add/delete) from BAT files or URLs
- 📋 View all WireGuard interfaces on the router
- 💡 Intuitive interface with tooltips
- ✅ Input field validation
- 🎨 Modern design with animations
