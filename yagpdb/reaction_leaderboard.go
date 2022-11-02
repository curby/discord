{{ $missing := true }}
{{ $lb := dbTopEntries "reaction_counter_%" 10 0 }}
{{ $position := 0 }}
{{ $leaderID := 0 }}
{{ $emojis := cslice "zero" "one" "two" "three" "four" "five" "six" "seven" "eight" "nine" "keycap_ten" }}
{{ $output := "" }}
{{ range $lb }}
    {{- $position = add $position 1 }}
    {{- $leaderID = (slice .Key 17) }}
    {{- if eq (toInt $leaderID) .User.ID }}
        {{- $missing = false }}
    {{- end }}
    {{- $output = joinStr "" $output ":" (index $emojis $position) ": <@" $leaderID "> (" (toString (toInt .Value)) ")\n" }}
{{- end }}
{{ if $missing }}
    {{ $userReactions := (dbGet 26 (print "reaction_counter_" .User.ID)).Value }}
    {{ $output = joinStr "" $output "⋮\n:asterisk: <@" .User.ID "> (" (toString (toInt $userReactions)) ")\n" }}
{{ end }}

{{ sendMessage nil (cembed "title" "Reaction Leaderboard" "color" 26367 "description" $output) }}


{{/* vim: set tabstop=4:shiftwidth=4   */}}
