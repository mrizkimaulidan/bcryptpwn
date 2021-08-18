# bcryptpwn

Brute force your bcrypt hashes based on your wordlist.

Clone :

```
$ git clone https://github.com/mrizkimaulidan/bcryptpwn.git
```

Download require depedencies :
```
$ go mod download
```

Build :

```
$ go build main.go
```

Usage :
```
$ main.exe bcrypt-hashes /path/to/wordlist.txt
```

Example :
```
$ main.exe $2a$10$s.qv0TNSsSa6//RYuTXAJOn4UZVgiJA4MQ0Gi851SrBvjaM56xTV2 D:\wordlist.txt
```

Only tested on Windows, I don't know if it can run smoothly on Linux. Just open an issue if something wrong on Linux machine.