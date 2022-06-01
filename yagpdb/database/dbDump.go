{{/* Displays the first 100 database entries after some $offset */}}

{{ $args := parseArgs 0 "Give an optional offset" (carg "int" "offset") }}
{{ $offset := 0 }}
{{ $count := 25 }}
{{ $processed := 0 }}
{{ $output := "" }}
{{ $found := false }}

{{ if ($args.IsSet 0) }}
	{{ $offset = ($args.Get 0) }}
{{ end -}}

{{ $lb := dbTopEntries "%" 100 $offset }}
{{ range $lb }}
	{{ $output = joinStr "" $output "\n" .UserID "`:`" .Key "`:`" (json .Value) }}
	{{ $processed = add $processed 1 }}
	{{ if eq $processed $count }}
		{{ sendMessage nil (joinStr "" "**Entries "  $offset " - " (add $offset $processed -1) ":**" $output) }}
		{{ $processed = 0 }}
		{{ $offset = add $offset $count }}
		{{ $output = "" }}
		{{ $found := true }}
	{{ end }}
{{ end }}

{{ if eq $processed 0 }}
	{{ if not $found }}
		Offset too high!
	{{ end }}
{{ else }}
	{{ sendMessage nil (joinStr "" "**Entries "  $offset " - " (add $offset $processed -1) ":**" $output) }}
{{ end }}
