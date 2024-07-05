{{/**************************************************************************\

    YAGPDB Custom Command - Reaction Monthly Upkeep
    -----------------------------------------------

    Edits a post with the current reactions leaderboard (this could be a
    pinned post for example). Then, call a script to reset all reaction
    counts. Make it run in a bot spam channel, even if it's editing a post in
    another channel.

    Trigger type: Hour Interval
    Interval: 1 hour

    Setup: Change channel and post IDs below
    Setup: Restrict to only run in botspam channels
    Setup: Restrict to only admins (e.g. put in administrator cc group)

  \**************************************************************************/}}

{{ $channelID := 1142498711931474091 }}
{{ $postID := 1148361322887663757 }}
{{ $RID := 26 }}         {{/* database id for reaction counts */}}
{{ $resetCCID := 31 }}   {{/* ID of the CC that will reset counts */}}

{{/* Get the date and hour. Only run when they are 1 and 00 respectively (start of month). */}}
{{ if eq 100 (toInt (formatTime currentTime "0215")) }}
    {{/* Initialize */}}
    {{ $position := 0 }}
    {{ $leaderID := 0 }}
    {{ $emojis := cslice "zero" "one" "two" "three" "four" "five" "six" "seven" "eight" "nine" "keycap_ten" }}
    {{ $lastMonth := formatTime (timestampToTime (sub currentTime.Unix 1296000)) "January" }}

    {{ $output := "" }}

    {{/* Iterate through top 10 reactees */}}
    {{ $lb := dbTopEntries "reaction_counter_%" 10 0 }}
    {{ range $lb }}
        {{- $position = add $position 1 }}
        {{- $leaderID = (slice .Key 17) }}
        {{- $output = joinStr "" $output ":" (index $emojis $position) ": <@" $leaderID "> (" (toString (toInt .Value)) ")\n" }}
    {{- end }}

    {{ $embed := cembed
        "title" (print $lastMonth "'s Leaderboard")
        "color" 13585960
        "description" $output
    }}
    {{ editMessage $channelID $postID (complexMessageEdit "embed" $embed "content" "") }}

    {{/* Call child, telling it to remove stuff */}}
    {{ $reactees := dbCount $RID }}
    {{ execCC $resetCCID nil 3 $reactees }}

{{ end }}

{{/* vim: set ts=4 sw=4 et: */}}
