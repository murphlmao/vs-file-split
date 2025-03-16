# vs-file-split

A simple, efficient file monitor that watches a directory and its subdirectories for files named using `{}` or `,` patterns (e.g., `{file1,file2}.txt` or `file1.txt, file2.txt`) and automatically splits them into separate files.

## Features
- 🔄 **Monitors all subdirectories** dynamically.
- 🚀 **Efficient & optimized** event handling.
- ✅ **Automatically splits files** based on naming patterns.
- 🗑️ **Deletes the original file** after splitting.
- 🔧 **Cross-platform support** (Linux, macOS, Windows).

---

## **🚀 Installation & Usage**

### **1️⃣ Running in Development Mode**
```sh
git clone https://github.com/murphlmao/vs-file-splitter.git
cd vs-file-splitter
go mod tidy
go run src/cmd/main.go
```
_(Watches your current directory and subdirectories.)_

### **2️⃣ Running as a Compiled Binary**
```sh
go build -o vs-file-splitter src/cmd/main.go
./vs-file-splitter
```

---

## **📌 Build**

A `Makefile` is included for easier building.

### **Build for Your Current OS**
```sh
make build
```
✅ **Output:**
- `dist/vs-file-splitter-linux` (Linux)
- `dist/vs-file-splitter-mac` (macOS)
- `dist/vs-file-splitter.exe` (Windows)

### **Cross-Compile for All Platforms**
```sh
make cross-compile
```
✅ **Output (in `dist/` folder):**
```
vs-file-splitter-linux-amd64
vs-file-splitter-windows-amd64.exe
vs-file-splitter-mac-amd64
vs-file-splitter-mac-arm64
```

### **Run Tests**
```sh
make test
```

### **Clean Build Artifacts**
```sh
make clean
```

---

## **📌 Example Usage**
#### **Creates two files from `{file1,file2}.txt`**
```sh
touch "{file1,file2}.txt"
```
✅ **Results:**
```
file1.txt
file2.txt
```

#### **Creates multiple files from `file1.txt, file2.txt`**
```sh
touch "file1.txt, file2.txt"
```
✅ **Results:**
```
file1.txt
file2.txt
```

#### **Works inside subdirectories too!**
```sh
mkdir nested

# touch already handles `touch {config,settings}.yaml` natively,
# but this is meant for VSCode for when you click 'New File'.
touch "nested/{config,settings}.yaml"
```
✅ **Results:**
```
nested/config.yaml
nested/settings.yaml
```

---

## **📌 Cross-Compiling Manually**
If you don't want to use `make cross-compile`, you can manually compile `vs-file-splitter` for different OSes:

```sh
# Linux (x86_64)
GOOS=linux GOARCH=amd64 go build -o dist/vs-file-splitter-linux src/cmd/main.go

# Windows (x86_64)
GOOS=windows GOARCH=amd64 go build -o dist/vs-file-splitter.exe src/cmd/main.go

# macOS (Intel)
GOOS=darwin GOARCH=amd64 go build -o dist/vs-file-splitter-mac src/cmd/main.go

# macOS (M1/M2 ARM)
GOOS=darwin GOARCH=arm64 go build -o dist/vs-file-splitter-mac-arm src/cmd/main.go
```