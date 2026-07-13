#!/bin/bash
set -e

BODY=$(cat << 'EOF'
[
    {
        "correlation_id": "a",
        "original_url": "http://a.com"
    },
    {
        "correlation_id": "b",
        "original_url": "http://b.com"
    },
    {
        "correlation_id": "c",
        "original_url": "http://c.com"
    },
    {
        "correlation_id": "d",
        "original_url": "http://d.com"
    },
    {
        "correlation_id": "e",
        "original_url": "http://e.com"
    },
    {
        "correlation_id": "f",
        "original_url": "http://f.com"
    },
    {
        "correlation_id": "g",
        "original_url": "http://g.com"
    },
    {
        "correlation_id": "h",
        "original_url": "http://h.com"
    },
    {
        "correlation_id": "i",
        "original_url": "http://i.com"
    },
    {
        "correlation_id": "j",
        "original_url": "http://j.com"
    },
    {
        "correlation_id": "k",
        "original_url": "http://k.com"
    },
    {
        "correlation_id": "l",
        "original_url": "http://l.com"
    }
]
EOF
)

curl -X POST -d "$BODY" -H "Content-Type: application/json" http://localhost:8080/api/shorten/batch

wrk -t4 -c50 -d30s -s ./profiles/random.lua http://localhost:8080 &

curl -o ./profiles/result.pprof "http://localhost:8080/debug/pprof/profile?seconds=30"
