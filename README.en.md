<div align="right">
  <a title="en" href="./README.md"><img src="https://img.shields.io/badge/-简体中文-545759?style=for-the-badge" alt="简体中文"></a>
  <img src="https://img.shields.io/badge/-English-F54A00?style=for-the-badge" alt="english">
</div>

<div align="center">
  <img alt="Ech0" src="./docs/imgs/logo.svg" width="150">

  [Preview](https://memo.vaaat.com/) | [Official Site & Doc](https://echo.soopy.cn/) | [Ech0 Hub](https://echohub.soopy.cn/)

  # Ech0
</div>

<div align="center">

[![GitHub release](https://img.shields.io/github/v/release/lin-snow/Ech0)](https://github.com/lin-snow/Ech0/releases) ![License](https://img.shields.io/github/license/lin-snow/Ech0) [![Ask DeepWiki](https://deepwiki.com/badge.svg)](https://deepwiki.com/lin-snow/Ech0)

</div>

> A next-generation open-source, self-hosted, lightweight federated publishing platform focused on personal idea sharing.

Ech0 is a new-generation open-source self-hosted platform designed for individual users. It is ultra-lightweight and low-cost, supporting the ActivityPub protocol to let you easily publish and share ideas, writings, and links. With a clean, intuitive interface and powerful command-line tools, content management becomes simple and flexible. Your data is fully owned and controlled by you, always connected to the world, building your own network of thoughts.

![Interface Preview](./docs/imgs/screenshot.png)

---

<details>
   <summary><strong>Table of Contents</strong></summary>

- [Ech0](#ech0)
  - [Highlights](#highlights)
  - [Quick Deployment](#quick-deployment)
    - [🐳 Docker (Recommended)](#-docker-recommended)
    - [🐋 Docker Compose](#-docker-compose)
  - [Upgrading](#upgrading)
    - [🔄 Docker](#-docker)
    - [💎 Docker Compose](#-docker-compose-1)
  - [Access Modes](#access-modes)
    - [🖥️ TUI Mode](#️-tui-mode)
    - [🔐 SSH Mode](#-ssh-mode)
  - [FAQ](#faq)
  - [Feedback \& Community](#feedback--community)
  - [Architecture](#architecture)
  - [Development Guide](#development-guide)
    - [Backend Requirements](#backend-requirements)
    - [Frontend Requirements](#frontend-requirements)
    - [Start Backend \& Frontend](#start-backend--frontend)
  - [Acknowledgements](#acknowledgements)
  - [Star History](#star-history)
  - [Support](#support)
</details>

---

## Highlights

☁️ **Atomically Lightweight**: Uses less than 15 MB of memory, image size under 50 MB, powered by a single SQLite file architecture.  
🚀 **Blazing-Fast Deployment**: No configuration required—ready to use with a single command.  
🧰 **CLI Power Tool**: Built-in high-availability command-line interface with one-click backup, restore, and export.  
📟 **Refined TUI Experience**: A beautifully crafted terminal user interface for effortless Ech0 management.  
✍️ **Distraction-Free Writing**: A clean online Markdown editor with rich plugin and live preview support.  
📦 **Data Sovereignty**: All content stored locally in a SQLite database, with built-in RSS subscription support.  
🔐 **Secure Backup System**: One-click export and full backup available via Web, TUI, or CLI.  
♻️ **Seamless Recovery**: Instantly restore any backup from TUI, CLI, or Web, ensuring complete data safety.  
🎉 **Forever Free**: Open source under the AGPL-3.0 license — no tracking, no subscription, no external dependencies.  
🌍 **Cross-Platform Design**: Fully responsive on desktop and mobile browsers — works flawlessly on phone, iPad, and PC.  
👾 **PWA Ready**: Installable as a web app for a near-native experience.  
☁️ **S3 Storage Integration**: Natively supports S3-compatible object storage for efficient cloud archiving.  
🌐 **ActivityPub Federation**: Interoperable with Mastodon, Misskey, and other federated platforms for a decentralized ecosystem.  
🔑 **OAuth2 Integration**: Native support for OAuth2, enabling easy third-party login and API authorization.  
📝 **Built-in To-Do Manager**: Effortlessly record and manage daily tasks for efficient planning and progress tracking.  
🔗 **Ech0 Connect**: A new multi-instance communication system that enables cross-instance status subscriptions and tracking.  
🎵 **Seamless Music Integration**: Ultra-lightweight built-in music player for immersive background audio and focus mode.  
🎥 **Instant Video Sharing**: Native intelligent parsing for Bilibili and YouTube links.  
🃏 **Rich Shortcut Cards**: One-click sharing for websites, GitHub projects, and other rich media—making content vivid and interactive.  
⚙️ **Advanced Customization**: Power users can easily add custom styles and scripts for expressive, personalized sharing.  
💬 **Comment System**: Lightweight integration with Twikoo comment service for instant, non-intrusive interaction and feedback.  
💻 **Cross-Platform Compatibility**: Natively supports Windows, Linux, and ARM devices like Raspberry Pi for stable, versatile deployment.  
🔗 **Ech0 Hub Integration**: Submit your instance to the official Ech0 Hub ecosystem to discover, subscribe to, and share high-quality content.  
📦 **Self-Contained Binary**: All resources bundled in a single executable—no external dependencies or configuration required.  
🔗 **Extensive API Support**: Open APIs for seamless integration with external systems and flexible application development.  
🃏 **Dynamic Content Display**: Supports X-style (Twitter-like) cards with social interactions such as likes.  
👤 **Multi-Account Access Control**: Flexible user and permission management system for secure and private content access.  

---

## Quick Deployment

### 🐳 Docker (Recommended)

```shell
docker run -d \
  --name ech0 \
  -p 6277:6277 \
  -p 6278:6278 \
  -v /opt/ech0/data:/app/data \
  -v /opt/ech0/backup:/app/backup \
  -e JWT_SECRET="Hello Echos" \
  sn0wl1n/ech0:latest
```

> 💡 After deployment, access `ip:6277` to use  
> 🚷 It is recommended to change `JWT_SECRET="Hello Echos"` to a secure secret  
> 📍 The first registered user will be set as administrator  
> 🎈 Data stored under `/opt/ech0/data`

### 🐋 Docker Compose

1. Create a new directory and place `docker-compose.yml` inside.  
2. Run:

```shell
docker-compose up -d
```

---

## Upgrading

### 🔄 Docker

```shell
docker stop ech0
docker rm ech0
docker pull sn0wl1n/ech0:latest
docker run -d \
  --name ech0 \
  -p 6277:6277 \
  -p 6278:6278 \
  -v /opt/ech0/data:/app/data \
  -v /opt/ech0/backup:/app/backup \
  -e JWT_SECRET="Hello Echos" \
  sn0wl1n/ech0:latest
```

### 💎 Docker Compose

```shell
cd /path/to/compose
docker-compose pull && \
docker-compose up -d --force-recreate
docker image prune -f
```

---

## Access Modes

### 🖥️ TUI Mode

![TUI Mode](./docs/imgs/tui.png)

Run the binary directly (for example, on Windows double-click `Ech0.exe`).

### 🔐 SSH Mode

Connect to the instance via port 6278:

```shell
ssh -p 6278 ssh.vaaat.com
```

---

## FAQ

1. **What is Ech0?**  
   A lightweight, open-source self-hosted platform for quickly sharing thoughts, writings, and links. All content is locally stored.  

2. **What Ech0 is NOT?**  
   Not a professional note-taking app like Obsidian or Notion; its core function is similar to social feed/microblog.  

3. **Is Ech0 free?**  
   Yes, fully free and open-source under AGPL-3.0, no ads, tracking, subscription, or service dependency.  

4. **How do I back up and restore data?**  
  Since all content is stored in a local SQLite file, you only need to back up the files in the `/opt/ech0/data` directory (or the mapped path you chose during deployment). To restore, simply replace the data files with your backup. You can also use the online data management features in the settings under "Data Management" to quickly create, export, or restore snapshots. If the latest content does not appear after restoring, try manually restarting the Docker container.

5. **Does Ech0 support RSS?**  
   Yes, content updates can be subscribed via RSS.  

6. **Why can't I publish content?**  
   Only administrators can publish. First registered user is admin.  

7. **Why no detailed permission system?**  
   Ech0 emphasizes simplicity: admin vs non-admin only, for smooth experience.  

8. **Why Connect avatars may not show?**  
   Set your instance URL in `System Settings - Service URL` (with `http://` or `https://`).  

9. **What is MetingAPI?**  
   Used to parse music streaming URLs for music cards. If empty, default API provided by Ech0 is used.  

10. **Why not all Connect items show?**  
    Instances that are offline or unreachable are ignored; only valid instances are displayed.  

11. **What content is not recommended?**  
    Avoid publishing dense content mixing text + images + extension cards. Long posts or extension cards alone are okay.  

12. **How to enable comments?**  
    Set up Twikoo backend URL in settings. Only Twikoo is supported.  

13. **How to configure S3?**  
    Fill in endpoint (without http/https) and bucket with public access.

14. **How to join the Fediverse?**  
  You need to bind Ech0 to a domain name and fill in the domain in the server address field in the settings page. Once set, Ech0 will automatically join the Fediverse. Example: `https://memo.vaaat.com`

---

## Feedback & Community

- Report bugs via [Issues](https://github.com/lin-snow/Ech0/issues).
- Propose features or share ideas in [Discussions](https://github.com/lin-snow/Ech0/discussions).

---

## Architecture

![Architecture Diagram](./docs/imgs/Ech0技术架构图.svg)  
> by ExcaliDraw

---

## Development Guide

### Backend Requirements
- Go 1.25.1+  
- C Compiler for CGO (`go-sqlite3`):
  - Windows: [MinGW-w64](https://winlibs.com/)  
  - macOS: `brew install gcc`  
  - Linux: `sudo apt install build-essential`  
- Google Wire: `go install github.com/google/wire/cmd/wire@latest`  
- Golangci-Lint: `golangci-lint run` / `golangci-lint fmt`  
- Swagger: `swag init -g internal/server/server.go -o internal/swagger`  

### Frontend Requirements
- NodeJS v24.5.0+, PNPM v10.17.1+  
- Use [fnm](https://github.com/Schniz/fnm) if multiple Node versions needed

### Start Backend & Frontend
```shell
# Backend
go run cmd/ech0/main.go

# Frontend
cd web
pnpm install
pnpm dev
```

Preview: Backend `http://localhost:6277`, Frontend `http://localhost:5173`

> When importing layered packages, prefer consistent aliases such as `xxxModel`, `xxxService`, `xxxRepository`, and so on.

---

## Acknowledgements

- [Gin](https://github.com/gin-gonic/gin)  
- [Md-Editor-V3](https://github.com/imzbf/md-editor-v3)  
- [Figma](https://www.figma.com/)  
- [VSCode](https://code.visualstudio.com/) & [GoLand](https://www.jetbrains.com/go/)  
- Open-source community contributors

---

## Star History

<a href="https://www.star-history.com/#lin-snow/Ech0&Timeline">
 <picture>
   <source media="(prefers-color-scheme: dark)" srcset="https://api.star-history.com/svg?repos=lin-snow/Ech0&type=Timeline&theme=dark" />
   <source media="(prefers-color-scheme: light)" srcset="https://api.star-history.com/svg?repos=lin-snow/Ech0&type=Timeline" />
   <img alt="Star History Chart" src="https://api.star-history.com/svg?repos=lin-snow/Ech0&type=Timeline" />
 </picture>
</a>

---

## Support

🌟 If you like **Ech0**, please give it a Star! 🚀  
Ech0 is completely free and open-source. Support helps the project continue improving.  

| Platform | QR Code |
| :------: | :------ |
| [**Afdian**](https://afdian.com/a/l1nsn0w) | <img src="./docs/imgs/pay.jpeg" alt="Pay" width="200"> |

---

```cpp

███████╗     ██████╗    ██╗  ██╗     ██████╗ 
██╔════╝    ██╔════╝    ██║  ██║    ██╔═████╗
█████╗      ██║         ███████║    ██║██╔██║
██╔══╝      ██║         ██╔══██║    ████╔╝██║
███████╗    ╚██████╗    ██║  ██║    ╚██████╔╝
╚══════╝     ╚═════╝    ╚═╝  ╚═╝     ╚═════╝ 

``` 
