{{/**************************************************************************\

    YAGPDB Custom Command - Karaoke Machine
    ---------------------------------------

    Displays a song's lyrics line-by-line. See samplesong.go for a detailed
    example with sample trigger regex, or use songtemplate.go to define
    new songs.

    Trigger type: None (To be called by external script)

  \**************************************************************************/}}

{{/* Unpack data from caller */}}
{{ $triggerID := .ExecData.triggerID }}
{{ $lyrics := .ExecData.lyrics }}

{{ $message := joinStr "" "\n> **" (index $lyrics 0) "**\n> " (index $lyrics 1)  "\n> " (index $lyrics 2) }}
{{ $messageID := sendMessageRetID nil $message }}

{{ $lines := len $lyrics }}
{{ range seq 1 $lines -}}
	{{ $current := . -}}
	{{ $last := sub $current 1 -}}
	{{ $next := add $current 1 -}}
	{{ $pause := div (len (index $lyrics $last)) 15 -}}
	{{ if lt $pause 2 -}}
		{{ $pause = 2 -}}
	{{ end -}}
	{{ sleep $pause -}}
	{{ $message = joinStr "" "> " (index $lyrics $last) "\n> **" (index $lyrics $current) "**" -}}
	{{ if lt $current (sub $lines 1) -}}
		{{ $message = joinStr "" $message "\n> " (index $lyrics $next) -}}
	{{ else -}}
		{{ $message = joinStr "" "> " (index $lyrics (sub $current 2)) "\n" $message -}}
	{{ end -}}
	{{ editMessage nil $messageID $message -}}
{{ end }}

{{ deleteMessage nil $messageID 5 }}
{{ dbDel .CCID "mutex" }}
{{ addMessageReactions nil $triggerID ":musical_note:" }}
