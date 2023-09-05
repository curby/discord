{{/**************************************************************************\

    YAGPDB Custom Command - Reaction Leaderboard (Deprecated on Stonemaier)
    --------------------------------------------

    Gives the top 10 recipients of reactions and the reaction counts of those
    members as well as that of the caller or an optionally specified member ID.

    Note: @mentions are intentionally unsupported to reduce notification spam.

    Trigger type: Command
    Trigger name: ?reactions

    Setup: None

  \**************************************************************************/}}


{{/* Initialize */}}
{{ $missing := true }}
{{ $position := 0 }}
{{ $leaderID := 0 }}
{{ $emojis := cslice "zero" "one" "two" "three" "four" "five" "six" "seven" "eight" "nine" "keycap_ten" }}
{{ $output := "" }}

{{/* Look for optional user (if none, look for triggering user) */}}
{{ $args := parseArgs 0 "Give an optional user ID" (carg "int" "targetID") }}
{{ $targetID := .User.ID }}
{{ if ($args.IsSet 0) }}
    {{ $targetID = ($args.Get 0) }}
{{ end -}}

{{/* Iterate through top 10 reactees */}}
{{ $lb := dbTopEntries "reaction_counter_%" 10 0 }}
{{ range $lb }}
    {{- $position = add $position 1 }}
    {{- $leaderID = (slice .Key 17) }}
    {{- if eq (toInt $leaderID) $targetID }}
        {{- $missing = false }}
    {{- end }}
    {{- $output = joinStr "" $output ":" (index $emojis $position) ": <@" $leaderID "> (" (toString (toInt .Value)) ")\n" }}
{{- end }}
{{ if $missing }}
    {{ $userReactions := (dbGet 26 (print "reaction_counter_" $targetID)).Value }}
    {{ $output = joinStr "" $output "⋮\n:asterisk: <@" $targetID "> (" (toString (toInt $userReactions)) ")\n" }}
{{ end }}

{{ sendMessage nil (cembed "title" "Reaction Leaderboard" "color" 26367 "description" $output) }}

{{/* vim: set ts=4 sw=4 et: */}}
