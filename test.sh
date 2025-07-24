#!/bin/bash
go run . reset
go run . register William
go run . login William
go run . addfeed techcrunch https://techcrunch.com/feed/
go run . addfeed hackernews https://news.ycombinator.com/rss
go run . addfeed bootsblog https://blog.boot.dev/index.xml
timeout --signal=INT 11s go run . agg 10s