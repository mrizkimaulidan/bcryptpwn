# bcryptpwn

Brute force your bcrypt hashes based on your wordlist.

Clone :

```bash
$ git clone https://github.com/mrizkimaulidan/bcryptpwn.git
```

Download required depedencies :

```bash
$ go mod download
```

Build :

```bash
$ go build main.go
```

Usage :

```bash
$ main.exe bcrypt-hashes /path/to/wordlist.txt
```

Example :

```bash
$ main.exe $2a$10$s.qv0TNSsSa6//RYuTXAJOn4UZVgiJA4MQ0Gi851SrBvjaM56xTV2 D:\wordlist.txt
```

Only tested on Windows, I don't know if it can run smoothly on Linux. Just open an issue if something wrong on Linux machine.
