{{/**************************************************************************\

    YAGPDB Custom Command - Reaction Monthly Reset
    ----------------------------------------------

    Called manually or via another CC to reset all reaction counts.

    Trigger type: Command

    Setup: Restrict to only admins (e.g. put in administrator cc group)

  \**************************************************************************/}}

{{ $RID := 26 }}         {{/* database id for reaction counts */}}
{{ $limit := 100 }}
{{ $commandChannel := 971603266003664986 }}
{{ $adminChannel := 971522137510785094 }}

{{ if not .ExecData }}
    {{ $deleted1 := dbDelMultiple (sdict "userID" $RID) $limit 0 }}
    {{ $remaining := dbCount $RID }}
    {{ if gt $remaining 0 }}
        {{ sendMessage nil (printf "Deleted %d entries, %d still remain. Please run this command again." $deleted1 $remaining) }}
    {{ else }}
        {{ sendMessage nil (printf "Deleted all %d entries. Thanks for helping to reset the leaderboard!" $deleted1) }}
    {{ end }}
{{ else }}
    {{ $total := toInt .ExecData }}
    {{ $deleted1 := dbDelMultiple (sdict "userID" $RID) $limit 0 }}
    {{ sendMessage nil (printf "Deleted %d entries." $deleted1) }}
    {{ if gt $total $limit }}
        {{ $deleted2 := dbDelMultiple (sdict "userID" $RID) $limit 0 }}
        {{ sendMessage nil (printf "Deleted %d entries in the second set." $deleted2) }}
        {{ if gt $total (mult 2 $limit) }}
            {{ sendMessage $adminChannel (printf "I'm updating the monthly reaction leaderboard. Please run `?reactionLBReset` in <#%d> until all %d remaining entries are gone. Check the pins for more info." $commandChannel (sub $total $deleted1 $deleted2)) }}
        {{ end }}
    {{ end }}
{{ end }}

{{/* vim: set ts=4 sw=4 et: */}}
