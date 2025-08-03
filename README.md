# 🕷️ Go Web Crawler

A fast, concurrent web crawler written in Go.  
It crawls a given website, collects all **internal links**, and generates a **report** showing how many times each page is linked internally.

---

## ✨ Features

- **⚡ Concurrent crawling** – multiple pages fetched in parallel for speed.
- **📌 Configurable max concurrency** – limit simultaneous HTTP requests.
- **📄 Configurable max pages** – stop crawling after a set number of pages.
- **🔒 Domain restriction** – only crawls links within the same domain.
- **📊 Sorted report output** – most linked pages appear first.
- **🛡️ Thread‑safe** – safe access to shared state using `sync.Mutex`.

---

## 📦 Installation

Clone the repository:
```bash
git clone https://github.com/somu-codes/crawler.git
cd crawler
```

Build the binary:
```bash
go build -o crawler
```

Run the binary:

```bash
go run . <url> [maxConcurrency] [maxPages]
```

