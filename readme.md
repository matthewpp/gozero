# Prerequisite

1. Go compiler
1. VS Code with [Go Extension](https://marketplace.visualstudio.com/items?itemName=golang.Go) and [SQLite Viewer Extension](https://marketplace.visualstudio.com/items?itemName=qwtel.sqlite-viewer)
1. Make
1. (Optional) Docker or Docker compatible [colima , podman]

## Windows Prepare

### 1. Go Compiler

🧩 Step-by-Step Installation For Window

1. Make sure winget is available, Open PowerShell (as Administrator) and run:

```shell
winget --version
```

If it shows a version (like v1.9.0), you’re good.
If not, update Windows or install App Installer from the Microsoft Store.

2. Search for Go

You can see what’s available:

```shell
winget search Go
```

You should see something like:

```
| Name | Id        | Version | Source |
| ---- | --------- | ------- | ------ |
| Go   | GoLang.Go | x.y.z   | winget |
```

3. Install Go

Run:

```shell
winget install -e --id GoLang.Go
```

Explanation:

- `--id GoLang.Go` → installs the official Go package

- `-e` → ensures exact match (avoid similar names)

This will automatically:

- Download the latest Go release

- Add `C:\Program Files\Go\bin` to your PATH

4. Verify the installation

After it finishes, restart PowerShell (or open a new one) and run:

```shell
go version
```

Expected output example:

```
go version go1.23.2 windows/amd64
```

5. Set Path

```shell
setx PATH "$env:PATH;C:\Program Files\Go\bin"
```

run go version again to confirm.

```shell
go version
```

### 2. VSCode extension

Install from this link

1. [Go Extension](https://marketplace.visualstudio.com/items?itemName=golang.Go)
2. [SQLite Viewer Extension](https://marketplace.visualstudio.com/items?itemName=qwtel.sqlite-viewer)

### 3. Make

```shell
winget install -e --id GnuWin32.Make
```

### 4. Docker compatible

- Install podman follow this instruction https://github.com/containers/podman/blob/main/docs/tutorials/podman-for-windows.md

**You're all set on windows !**

## Mac Prepare

### 1. Go Compiler

```shell
brew install go
```

Check Go version

```shell
go version
```

Expected output example:

```
go version go1.23.2 darwin/arm64
```

### 2. VSCode extension

Install from this link

1. [Go Extension](https://marketplace.visualstudio.com/items?itemName=golang.Go)
2. [SQLite Viewer Extension](https://marketplace.visualstudio.com/items?itemName=qwtel.sqlite-viewer)

### 3. Make

```shell
brew install make
```

### 4. Docker compatible

- Install colima follow this instruction https://github.com/abiosoft/colima

OR

- Install podman follow this instruction https://github.com/containers/podman/blob/main/docs/tutorials/podman-for-windows.md

**You're all set on mac !**
