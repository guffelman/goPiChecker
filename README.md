Simple GoLang program to check if the Raspberry Pi is in stock. When it is found, it sends a post request to a link of your choice with the link to purchase.

Can be run on Windows as an .exe, or as a docker container.

# Building the application
## Windows:
```go build -o GoPiChecker.exe .```

## Docker:
```docker build -t <tag> .```
