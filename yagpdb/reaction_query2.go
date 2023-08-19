{{/**************************************************************************\

    YAGPDB Custom Command - My Reactions Query
    ------------------------------------------

    Returns the reaction count for the caller.

    Trigger type: Regex
    Trigger Regex: .*

    Setup: Restrict to a particular channel, since this removes all posts.

  \**************************************************************************/}}

{{ if eq "?reactions" (trimSpace (lower .Cmd)) }}
    {{ $userReactions := (dbGet .CCID (print "reaction_counter_" .User.ID)).Value }}
    {{ $response := sendMessageRetID nil (print .User.Mention " has " (toString (toInt $userReactions)) " reactions.") }}
    {{ deleteMessage nil $response 35 }}
{{ else }}
    {{ sendDM "The reactions channel only supports the `?reactions` command; use it to view the number of times that others have reacted to your posts." }}
{{ end }}
{{ deleteTrigger 1 }}

{{/* vim: set ts=4 sw=4 et: */}}
