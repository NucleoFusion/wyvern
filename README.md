# 🐉🔥 Wyvern - A Github-Integrated Chat App!

**Ever felt that communication in Github was a pain? Well fret not, Wyvern solves this!**

Wyvern brings real-time chat to your repositories — so you can discuss issues, ideas, and progress right where the code lives.
With Wyvern, you can:
- 💬 Create dedicated chat spaces for repositories
- 🧑‍🤝‍🧑 Collaborate with contributors, followers, maintainers, and owners
- 🛠️ Manage your repo directly: create issues, PRs, view files, and more
- 🔗 Stay connected — all inside a unified interface

> One app to talk, track, and build better GitHub projects.

## 🗨️ Introduction

Wyvern is a real-time chat and collaboration app built around GitHub repositories. It allows teams, contributors, and maintainers to communicate, manage, and work on projects — all in one place.

GitHub Issues and Discussions are powerful, but they're not built for fast, fluid communication. Project members often end up using third-party tools like Slack or Discord — separating code from conversation.

With Wyvern you can:-
- 💬 Create chats linked to repositories
- 🧑‍🤝‍🧑 Talk to contributors, followers and owners
- 🛠️ Open Issues or PR's from the chat
- 📂 Browser the repo structure and files
- 📊 Look at Repo Insights and Data

## 🛠️ Installation

⚠️ Installation instructions are coming soon.
Wyvern is currently in active development — stay tuned for setup guides and deployment steps.

## ✨ Features

> 🚧 Wyvern is currently under active development.  
> Core features like real-time chat, GitHub integration, and repo management are being built.  
> Stay tuned for the first release!

## 🚧 Roadmap

- 💬 **Repository-specific chats** – Real-time messaging for any GitHub repo
- 🧑‍🤝‍🧑 **Chat with contributors, followers, or owners** – Powered by GitHub integration
- 🛠️ **Create issues and pull requests from chat** – Turn discussion into action
- 📂 **Browse repo files and structure** – View folders, files, and code directly
- 📊 **View repo insights and activity** – Commits, contributors, stars, forks & more
- 🔔 **GitHub notifications integration** – Stay updated on repo events
- 🔐 **GitHub OAuth login** – Secure sign-in with your GitHub account

## ⚙️ Configuration

> 🔧 Configuration details will be documented soon.  
> Wyvern is currently in active development and may require environment variables or GitHub API credentials.

## 🤝 Contributing

Contributions are welcome and appreciated!

📄 A full [CONTRIBUTING.md](CONTRIBUTING.md) guide will be added soon.

## 🗂️ Code Structure

This project is structured as a **monorepo** with a `client` and `server`.

```
wyvern/
├── client/                  # Frontend apps (SvelteKit + Tauri)
│   ├── cross-platform/      # Tauri + SvelteKit app
│   └── tui/                 # (Planned) Terminal UI
│
├── server/                  # Backend API server (Go)
│   └── internal/            # Internal packages
│       └── routes/          # API route handlers
│
├── README.md
└── ...others
```

> 🧪 More modules and packages will be added as development progresses.

## 🪪 License

This project is licensed under the terms of the [MIT License](LICENSE).
