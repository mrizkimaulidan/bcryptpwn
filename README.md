# bcryptpwn

Brute force your bcrypt hashes based on your wordlist.

Clone :

```bash
$ git clone https://github.com/mrizkimaulidan/bcryptpwn.git
```

Download required dependencies:

```bash
$ go mod download
```

Build:

```bash
$ go build main.go
```

Help:

```bash
$ ./main -help
```

```bash
Usage of ./main:
  -hash string
        your bcrypt hashes
  -path string
        your wordlist path
```

Usage:

```bash
$ ./main bcrypt-hashes /path/to/wordlist.txt
```

Example :

```bash
$ ./main -hash='$2a$12$3Y3NvQQ7H57miY5LFidbmeHaGgIKVX1ZMiftRpm50Cfm9/0oyJCaq' -path=wordlist.txt
```

Do not forget to wrap hashes with single quote.