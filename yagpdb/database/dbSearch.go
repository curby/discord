{{/* Searches for all database entries with the query string in the userid or key fields */}}

{{ $args := parseArgs 1 "Usage: dbSearch <searchString>" (carg "string" "query") }}
{{ $query := $args.Get 0 }}
{{ $offset := 0 }}
{{ $processed := 0 }}
{{ $output := "" }}
{{ $found := false }}

{{ if gt (toInt $query) 0 }}
	{{ $lb := dbGetPattern (toInt64 $query) "%" 100 0 }}
	{{ range $lb }}
		{{ $output = joinStr "" $output "\n" .UserID "`:`" .Key "`:`" (json .Value) }}
		{{ $processed = add $processed 1 }}
	{{ end }}

	{{ if eq $processed 0 }}
		{{ if not $found }}
			{{ sendMessage nil "Query not found in UserID field." }}
		{{ end }}
	{{ else }}
		{{ sendMessage nil (joinStr "" "**Results in UserID Field ("  $offset " - " (add $offset $processed -1) ")**" $output) }}
	{{ end }}
{{ end }}

{{ $processed = 0 }}
{{ $output = "" }}
{{ $found = false }}

{{ $lb := dbTopEntries (print "%" $query "%") 100 0 }}
{{ range $lb }}
	{{ $output = joinStr "" $output "\n" .UserID "`:`" .Key "`:`" (json .Value)}}
	{{ $processed = add $processed 1 }}
{{ end }}

{{ if eq $processed 0 }}
	{{ if not $found }}
		{{ sendMessage nil "Query not found in key field." }}
	{{ end }}
{{ else }}
	{{ sendMessage nil (joinStr "" "**Results in Key Field ("  $offset " - " (add $offset $processed -1) ")**" $output) }}
{{ end }}
