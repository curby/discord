{{/* Deletes a database entry specified by id and key.
     See dbDelUser to delete all of a user's database entries. */}}

{{ $args := parseArgs 2 "Usage: dbDel <ID> <Key>"
  (carg "userid" "ID")
  (carg "string" "key") }}
{{ dbDel ($args.Get 0) ($args.Get 1) }}
Finished
