# 🐝 Bee

Bee is a local-first CLI agent designed to assist with programming, code generation, and command execution directly from your terminal.

Unlike cloud-first AI tools, Bee is built to run locally and integrate seamlessly with local LLM runtimes such as LM Studio, giving you full control over your data, workflow, and execution environment.

---

## ✨ Key Features

- 🧠 AI-powered programming assistant in your terminal
- 💻 Code generation and editing capabilities
- ⚙️ Ability to execute system commands (with user control)
- 🏠 Local-first architecture (no mandatory cloud dependency)
- 🔌 Designed to work with local LLMs (e.g. LM Studio)
- 🐝 Lightweight CLI with fast startup

---

## 🚀 Philosophy

Bee is built around a simple idea:

> Your AI should live where you work — not in the cloud.

It runs locally, talks to local models, and can evolve into a fully autonomous programming agent while keeping you in control.

---

## 📦 Installation

### Download binary (recommended)

Go to the latest release:

https://github.com/Angel-del-dev/bee/releases

Download the version for your system:

- `bee-linux-amd64`
- `bee-darwin-amd64`
- `bee-windows-amd64.exe`

Then move it to your PATH:

```bash
mv bee-linux-amd64 /usr/local/bin/bee
chmod +x /usr/local/bin/bee
```

## Build from source

```bash
git clone https://github.com/Angel-del-dev/bee.git
cd bee
go build -o bin/bee ./cmd/agent
```

## 🐝 Usage

```bash
$ bee
🐝> hello
📨 received: hello
🐝 Bzz Bzz
🐝> Create a Go HTTP server
```

## 🧠 Local LLM Integration

Bee is designed to work with local inference engines such as:

LM Studio
Ollama (future support planned)
Custom OpenAI-compatible APIs (planned)

Configure your model endpoint inside Bee (coming soon).

## ⚠️ Safety

Bee will eventually be able to execute system commands.

For now:

Execution is controlled
No automatic destructive actions
User confirmation will always be required (planned design)

## 🐝 Vision

Bee aims to become a fully local programming agent:

No mandatory cloud APIs
Fully developer-owned execution
Extensible tool system
Fast, terminal-native workflow