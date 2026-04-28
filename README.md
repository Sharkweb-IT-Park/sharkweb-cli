# 🦈 Sharkweb CLI

**Sharkweb CLI** is a modular full-stack development toolkit that lets you build scalable applications using plug-and-play modules.

---

# 🚀 Installation

## ⚡ One-line Install (Recommended)

```bash
curl -fsSL https://raw.githubusercontent.com/Sharkweb-IT-Park/sharkweb-cli/main/install.sh | bash
```

👉 This will:

* Download latest Sharkweb CLI
* Build/install binary
* Add it to your system

---

## 🧪 Verify Installation

```bash
sharkweb --help
```

---

## 🛠 Alternative (Manual Install)

### 1. Clone repo

```bash
git clone https://github.com/Sharkweb-IT-Park/sharkweb-cli.git
cd sharkweb-cli
```

### 2. Build

```bash
go build -o sharkweb
```

### 3. Add to PATH

#### Linux / Mac

```bash
mv sharkweb /usr/local/bin/
```

#### Windows

```bash
move sharkweb.exe C:\Windows\System32\
```

---

# 🚀 Usage

# 🦈 Sharkweb CLI Commands

Sharkweb CLI provides a set of powerful commands to manage modular applications.

---

# 🚀 Core Commands

## 📦 Add Module

Install a module from the registry:

```bash
sharkweb add module <name>
```

### Example:

```bash
sharkweb add module crm
```

---

## 🔄 Upgrade Module

Upgrade an installed module to the latest version:

```bash
sharkweb upgrade module <name>
```

### Example:

```bash
sharkweb upgrade module crm
```

---

## ❌ Remove Module

Remove a module from the project:

```bash
sharkweb remove module <name>
```

### Example:

```bash
sharkweb remove module crm
```

---

# 🛠 Development Commands

## ⚙️ Generate Module

Generate a new module from template:

```bash
sharkweb generate module <name>
```

### Example:

```bash
sharkweb generate module crm
```

---

## 🚀 Publish Module

Publish your module to GitHub:

```bash
sharkweb publish module <name> --repo <repo-url> --version <version>
```

### Example:

```bash
sharkweb publish module crm --repo https://github.com/user/crm --version 1.0.0
```

---

## 🧪 Dev Mode (Local Development)

Run development utilities:

```bash
sharkweb dev
```

---

# 📁 Project Commands

## 🆕 Create Project

Create a new Sharkweb project:

```bash
sharkweb create <project-name>
```

---

# ℹ️ Info Commands

## 📌 Version

Check CLI version:

```bash
sharkweb version
```

---

## 🆘 Help

```bash
sharkweb --help
```

---

# 🧠 Command Flow

```text
create → add → wiring → run → upgrade/remove → publish
```

---

# 💡 Notes

* Modules are installed from a registry (GitHub-based)
* Backend uses **Gin**
* Frontend uses **Next.js App Router**
* Wiring is automatically generated after install

---

# 🔥 Pro Tip

After adding a module:

```bash
sharkweb add module crm
```

You immediately get:

* API: `/api/crm`
* Route: `/crm`

---


# 🤝 Contributing

PRs welcome 🚀

---

# 📄 License

MIT License © Sharkweb IT Park
