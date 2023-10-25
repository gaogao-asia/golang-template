# Reference:
- https://github.com/golang-standards/project-layout
- https://gin-gonic.com/docs/quickstart/
- https://gorm.io/docs/connecting_to_the_database.html

# VS code ext:
1. Go
2. go snippet
3. Go struct tag
4. Go test explore
5. Json to go
6. GitLens
7. Github pull request and issue
8. Database client
9. ChatGPT GenieAI
10. Github Copilot

# VS code setting for go test coverage
```
"go.coverageDecorator": {
        "type": "gutter",
        "coveredHighlightColor": "rgba(64,128,128,0.5)",
        "coveredGutterStyle": "blockgreen",
        "uncoveredGutterStyle": "blockred"
    }
```
# Step to run
### RUN via local
1. Init
```
$ make init
```

2. Run
```
$ make run
```

### RUN via vscode
1. Init
```
$ make init
```
2. Run
`Mac: cmd + f5`

### RUN via container
1. Init
```
$ make init
```

2. Build
```
$ make quickbuild
```

3. Run
```
$ make run/container
```

