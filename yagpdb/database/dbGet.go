{{/* Fetches a database entry specified by id and key.
     See dbDump to list all database entries (paginated). */}}

{{ $args := parseArgs 2 "Usage: dbGet <ID> <Key>"
  (carg "userid" "ID")
  (carg "string" "key") }}
{{ dbGet ($args.Get 0) ($args.Get 1) }}
