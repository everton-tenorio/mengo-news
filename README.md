# Mengo News

### Scraping and sending Flamengo news using Golang and Telegram.

# Instructions
<p>1. <a href="https://go.dev/dl/">Download</a> and <a href="https://go.dev/doc/install">install Go</a>
<p>2. Create .env file 
<pre>
  TELEGRAM_TOKEN= "token"
  ID= "chat_id"
</pre>
</p>  
<p>3. Run mengonews.go <pre>go run mengonews.go</pre></p>

---

# Run with docker
<p>1. Create .env file
<pre>
  TELEGRAM_TOKEN= "token"
  ID= "chat_id"
</pre>
</p>

<p>2. Run mengo-news 
  
```bash
docker run -dt --rm --name mengo-news -v $(pwd):/app -w /app golang:alpine go build -v && ./mengo-news
```
</p>

---
 
<div align="center">
  <img src="./assets/example.gif" width="20%">
  <img src="./assets/go.png" width="10%">
  <p><img align="center" src="./assets/flamengo-news.png" width="65%"></p>
</div>
