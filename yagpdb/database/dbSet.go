{{/* Sets a database entry to a value, specified by id and key.
     See dbGet/dbDel to see/delete an entry respectively. */}}

{{ $args := parseArgs 3 "Usage: dbSet <ID> <Key> <Value>"
  (carg "userid" "ID")
  (carg "string" "key")
  (carg "string" "value") }}
{{ dbSet ($args.Get 0) ($args.Get 1) ($args.Get 2) }}
Finished
