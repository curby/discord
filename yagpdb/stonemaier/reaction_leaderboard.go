{{/**************************************************************************\

    YAGPDB Custom Command - Reaction Leaderboard
    --------------------------------------------

    Edits a post with the current reactions leaderboard (this could be a pinned
    post for example). Make it run in a bot spam channel, even if it's editing
    a post in another channel.

    Trigger type: Minute Interval
    Interval: 15 minutes

    Setup: Change channel and post IDs below
    Setup: Restrict to only run in botspam channels
    Setup: Restrict to only admins (e.g. put in administrator cc group)

  \**************************************************************************/}}

{{ $channelID := 1142498711931474091 }}
{{ $postID := 1142506464435519572 }}

{{/* Initialize */}}
{{ $position := 0 }}
{{ $leaderID := 0 }}
{{ $emojis := cslice "zero" "one" "two" "three" "four" "five" "six" "seven" "eight" "nine" "keycap_ten" }}
{{ $output := "" }}

{{/* Iterate through top 10 reactees */}}
{{ $lb := dbTopEntries "reaction_counter_%" 10 0 }}
{{ range $lb }}
    {{- $position = add $position 1 }}
    {{- $leaderID = (slice .Key 17) }}
    {{- $output = joinStr "" $output ":" (index $emojis $position) ": <@" $leaderID "> (" (toString (toInt .Value)) ")\n" }}
{{- end }}

{{ $embed := cembed
    "title" "Reaction Leaderboard"
    "color" 13585960
    "description" (print $output "\nType `?reactions` below to see the number of times people have reacted to your posts.")
    "footer" (sdict "text" "Last update (every 15m):")
    "timestamp" currentTime.UTC
}}
{{ editMessage $channelID $postID (complexMessageEdit "embed" $embed "content" "") }}

{{/* copy and paste these lines to change other posts too */}}
{{/*
{{ $channelID = 971816708841046076 }}
{{ $postID = 981959484861714532 }}
{{ editMessage $channelID $postID (complexMessageEdit "embed" $embed "content" "") }}
*/}}

{{/* alternate display style, using the embed's description:
"description" (print "Total Members: **" .Guild.MemberCount "**\n" "Online Members: **" onlineCount "**")
 */}}

{{/* vim: set ts=4 sw=4 et: */}}

