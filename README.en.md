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

☁️ **Atomically Lightweight**: Consumes less than **15MB** of memory with an image size under **50MB**, powered by a single-file SQLite architecture  
🚀 **Instant Deployment**: Zero configuration required — from installation to operation in just one command  
✍️ **Distraction-Free Writing**: A clean, online Markdown editor with rich plugin support and real-time preview  
📦 **Data Sovereignty**: All content is stored locally in SQLite, with full RSS feed support  
🔐 **Secure Backup Mechanism**: One-click full-data export and backup via Web, TUI, or CLI  
♻️ **Seamless Recovery**: Supports TUI/CLI snapshot restoration and Web-based zero-downtime recovery, ensuring data safety with ease  
🎉 **Forever Free**: Open-sourced under the AGPL-3.0 license — no tracking, no subscriptions, no external dependencies  
🌍 **Cross-Platform Adaptation**: Fully responsive design optimized for desktop, tablet, and mobile browsers  
👾 **PWA Ready**: Installable as a web application, offering a near-native experience  
🏷️ **Elegant Tag Management & Filtering**: Intelligent tagging system with fast filtering and precise search for effortless organization  
☁️ **S3 Storage Integration** — Native support for S3-compatible object storage enables efficient cloud synchronization  
🌐 **ActivityPub Federation** — Seamlessly federates with Mastodon, Misskey, and other decentralized platforms  
🔑 **OAuth2 Integration** — Native support for OAuth2, simplifying third-party login and API authorization  
🪶 **Highly Available Webhook**: Enables real-time integration and collaboration with external systems, supporting event-driven automated workflows  
📝 **Built-in Todo Management**: Easily capture and manage daily tasks to stay organized and productive  
🧰 **Command-Line Powerhouse**: A built-in high-availability CLI that empowers developers and advanced users with precision control and seamless automation  
📟 **Refined TUI Experience**: A beautifully designed terminal interface offering intuitive management of Ech0  
🔗 **Ech0 Connect**: A multi-instance connectivity feature that enables real-time status sharing and synchronization between Ech0 nodes  
🎵 **Seamless Music Integration**: Lightweight embedded music player providing immersive soundscapes and focus modes  
🎥 **Instant Video Sharing**: Natively supports intelligent parsing of Bilibili and YouTube videos  
🃏 **Rich Smart Cards**: Instantly share websites, GitHub projects, and other media in visually engaging cards  
⚙️ **Advanced Customization**: Easily personalize styles and scripts for expressive, unique content presentation  
💬 **Comment System**: Quick Twikoo integration for lightweight, instant, and non-intrusive interactions  
💻 **Cross-Platform Compatibility**: Runs natively on Windows, Linux, and ARM devices like Raspberry Pi for stable deployment anywhere  
🔗 **Ech0 Hub Integration**: Connect to the official Ech0 Hub to discover, subscribe, and share high-quality content  
📦 **Self-Contained Binary**: Includes all required resources — no extra dependencies, no setup hassle  
🔗 **Rich API Support**: Open APIs for seamless integration with external systems and workflows  
🃏 **Dynamic Content Display**: Supports Twitter-like card layouts with likes and social interactions  
👤 **Multi-Account & Permission Management**: Flexible user and role-based access control ensuring privacy and security  


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

### ☸️ Kubernetes (Helm)

If you want to deploy Ech0 in a Kubernetes cluster, you can use the Helm Chart provided in this project.

Since this project does not provide an online Helm repository, you need to clone the repository to your local machine first, and then install from the local directory.

1.  **Clone the repository:**
    ```shell
    git clone https://github.com/lin-snow/Ech0.git
    cd Ech0
    ```

2.  **Install with Helm:**
    ```shell
    # helm install <release-name> <chart-directory>
    helm install ech0 ./charts/ech0
    ```

    You can also customize the release name and namespace:
    ```shell
    helm install my-ech0 ./charts/ech0 --namespace my-namespace --create-namespace
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
go run main.go

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
