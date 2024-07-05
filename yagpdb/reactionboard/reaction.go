{{/**************************************************************************\

    YAGPDB Custom Command - Reaction Counter
    ----------------------------------------

    Counts the number of reactions each member gets.
    Use a separate command to see the leaderboard, and a third command to
    clean up a user's database entries when they leave the server.

    Trigger type: Reaction
    Trigger event: Added reactions only

    Setup: Just configuration below

  \**************************************************************************/}}

{{/* Configuration */}}
{{ $cooldown := 30 }}   {{/* time between reactions from reactor to reactee */}}
{{ $CID := .CCID }}      {{/* database id for cooldown timers */}}
{{ $RID := 26 }}         {{/* database id for reaction counts */}}
{{ $debugChannel := 971603266003664986 }}
{{ $botRoleID := 971613059812589638 }}
{{ $debug := true }}

{{ if $debug }}
    {{ sendMessage $debugChannel ( print "[CC: " .CCID "] " (userArg .Reaction.UserID).Username " :point_right: " .ReactionMessage.Author.Username " (" .ReactionMessage.Link ")" ) }}
{{ end }}

{{/* Skip people reacting to themselves */}}
{{ if ne .Reaction.UserID .ReactionMessage.Author.ID }}

    {{/* Skip people reacting to bots */}}
    {{ if not (targetHasRoleID .ReactionMessage.Author.ID $botRoleID) }}

        {{/* Get cooldown from database */}}
        {{ $Ckey := print .Reaction.UserID "_" .ReactionMessage.Author.ID }}
        {{ $result := (dbGet $CID $Ckey).Value }}
        {{ $lastTime := 0 }}
        {{ $thisTime := currentTime.Unix }}
        {{ if $result }}
            {{ $lastTime = toInt $result }}
            {{ if $debug }}
                {{ sendMessage $debugChannel ( print "[CC: " .CCID "] Found cooldown entry " $lastTime " (now " $thisTime ")" ) }}
            {{ end }}
            {{/* Clean up dirty database entry */}}
            {{ if gt $thisTime (add $lastTime $cooldown) }}
                {{ sendMessage $debugChannel ( print "[CC: " .CCID "] Cleaning up database entry: " $CID " " $Ckey ) }}
                {{ dbDel $CID $Ckey }}
            {{ end }}
        {{ end }}

        {{/* Check cooldown (skip someone's spammed reactions to same target) */}}
        {{ if gt $thisTime (add $lastTime $cooldown) }}
            {{ $result = dbIncr $RID (print "reaction_counter_" .ReactionMessage.Author.ID) 1 }}
            {{ dbSetExpire $CID $Ckey $thisTime 30 }}
            {{ if $debug }}
                {{ sendMessage $debugChannel ( print "[CC: " .CCID "] " .ReactionMessage.Author.Username " now has " $result ) }}
            {{ end }}

        {{ else if $debug }}
            {{ sendMessage $debugChannel ( print "[CC: " .CCID "] Cooldown active, noop" ) }}
        {{ end }}

    {{ else if $debug }}
        {{ sendMessage $debugChannel ( print "[CC: " .CCID "] Reaction to bot, noop" ) }}
    {{ end }}

{{ else if $debug }}
    {{ sendMessage $debugChannel ( print "[CC: " .CCID "] Reaction to self, noop" ) }}
{{ end }}

{{/* vim: set ts=4 sw=4 et: */}}
