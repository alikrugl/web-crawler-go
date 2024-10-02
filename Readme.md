[<img src="https://img.icons8.com/?size=512&id=55494&format=png" align="right" width="25%" padding-right="350">]()

# `WEB-CRAWLER-GO`

<p  style="display: flex; align-items: left;">
  <em>Built with:</em>
  <img src="https://img.shields.io/badge/Go-00ADD8.svg?style=flat&logo=Go&logoColor=white" alt="Go" style="margin-left: 10px;">
</p>

## 📍 Overview

The `web-crawler-go` project is a web crawler written in Go. It is designed to traverse web pages, extract URLs, and perform various web scraping tasks. The project aims to provide a simple and efficient way to gather data from the web, making use of Go's concurrency features to handle multiple requests simultaneously.

(Just a small project to get familiar with go and scraping)

---


## 🚀 Getting Started

### 📦 Installation

Build the project from source:

1. Clone the web-crawler-go repository:
```sh
❯ git clone https://github.com/alikrugl/web-crawler-go
```

2. Navigate to the project directory:
```sh
❯ cd web-crawler-go
```

3. Build the binary:
```sh
❯ go build -o crawler
```

### 🤖 Usage

To run the project, execute the following command:

```sh
❯ ./crawler rootUrl maxConcurrency maxPages
```
- `rootUrl`: The starting URL for the web crawler.
- `maxConcurrency`: The maximum number of concurrent requests the crawler can make.
- `maxPages`: The maximum number of pages the crawler should visit.

### 🧪 Tests

Execute the test suite using the following command:

```sh
❯ go test
```