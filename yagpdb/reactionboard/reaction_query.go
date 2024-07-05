{{/**************************************************************************\

    YAGPDB Custom Command - My Reactions Query
    ------------------------------------------

    Returns the reaction count for the caller.

    Trigger type: Regex
    Trigger Regex: .*

    Setup: Restrict to a particular channel, since this removes all posts.

  \**************************************************************************/}}

{{ $RID := .CCID }}         {{/* database id for reaction counts */}}

{{ if eq "?reactions" (trimSpace (lower .Cmd)) }}
    {{ $userReactions := toInt (dbGet $RID (print "reaction_counter_" .User.ID)).Value }}
    {{ $rank := "" }}
    {{ $plural := "s" }}
    {{ if gt $userReactions 0 }}
        {{ $userRank := dbRank (sdict "userID" $RID) $RID (print "reaction_counter_" .User.ID) }}
        {{ $reactees := dbCount $RID }}
        {{ $rank = printf " (ranked #%d out of %d people with reactions)" $userRank $reactees }}
    {{ end }}
    {{ if eq $userReactions 1 }}
        {{ $plural = "" }}
    {{ end }}
    {{ $response := sendMessageRetID nil (print .User.Mention " has " (toString $userReactions) " reaction" $plural $rank ".") }}
    {{ deleteMessage nil $response 35 }}
{{ else }}
    {{ sendDM "The reaction-counts channel only supports the `?reactions` command; use it to view the number of times that others have reacted to your posts." }}
{{ end }}
{{ deleteTrigger 1 }}

{{/* vim: set ts=4 sw=4 et: */}}
