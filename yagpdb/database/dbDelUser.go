{{/* Remove all db entries relating to a user. Made to be called by another script or by hand with the userID. */}}

{{/* todo: search for user and delete all? currently this is more of a framework for manually
specifying each entry to delete */}}

{{ $target := 0 }}

{{ if .ExecData }}
	{{ $target = .ExecData.userID }}
{{ else }}
	{{ $args := parseArgs 1 "Usage: dbDelUser <userID>" (carg "userid" "ID") }}
	{{ $target = $args.Get 0 }}
{{ end }}

{{ dbDel $target "wroteIntro" }}
{{ sendMessage 971644979380383794 (print "Cleaned up database entries for ID " $target) }}
