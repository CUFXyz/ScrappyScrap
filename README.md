# ScrappyScrap
Simple program what scrapping web pages.

Created just for learning purposes




# Usage/Examples

```Bash
make build
```

For Single link
```Bash
./cmd/ScrappyScrap <link>
```
For Multiple links
```Bash
./cmd/ScrappyScrap -m <link> <link> <...>
```

For example:
```Bash
./cmd/ScrappyScrap -m https://en.wikipedia.org/wiki/Compiler https://en.wikipedia.org/wiki/Source_code
```
That's creating txt file in main directory with separator **[Link %number of link%]**

