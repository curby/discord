{{/*
    Handles messages for the leveling system.
    See <https://yagpdb-cc.github.io/leveling/message-handler> for more information.

    Author: jo3-l <https://github.com/jo3-l>
*/}}

{{ $settings := 0 }} {{/* Instantiate settings at nil value */}}
{{ $roleRewards := sdict "type" "stack" }} {{/* Default role reward settings */}}
{{ with (dbGet 0 "xpSettings") }} {{ $settings = sdict .Value }} {{ end }} {{/* If in db, then we update value */}}
{{ with (dbGet 0 "roleRewards") }} {{ $roleRewards = sdict .Value }} {{ end }} {{/* See above */}}

{{ $cooldown := false }} {{/* We presume that user is not on cooldown */}}
{{ if (dbGet .User.ID "xpCooldown") }} {{ $cooldown = true }} {{ end }} {{/* Make user on cooldown if there is cooldown DB entry */}}

{{ if and (not $cooldown) $settings }} {{/* Make sure that both the user is not on cooldown and settings exist */}}
    {{/* Only give xp to high value posts, but always set timeout regardless */}}
    {{ $trimmed := reReplace `(?i)([a-z\d]+://)([\w_-]+(?:(?:\.[\w_-]+)+))([\w.,@?^=%&:/~+#-]*[\w@?^=%&/~+#-])` (trimSpace .Message.Content) "l" }} {{/* Ignore posts with just a link. */}}
    {{ $trimmed = reReplace `<a?:[\w~]+:(\d+)>` $trimmed "e" }} {{/* Ignore posts with just an emoji. */}}
    {{ if gt (len $trimmed) 20 }} {{/* Ignore short posts. */}}
        {{ sendMessage 1126983699704066239 (cembed "fields" (cslice (sdict "name" "msg cont" "value" (print "`" .Message.Content "`") "inline" false) (sdict "name" "trimmed" "value" $trimmed "inline" false) (sdict "name" "Link" "value" .Message.Link "inline" false) (sdict "name" "Len" "value" (str (len $trimmed)) "inline" false) ) ) }}
        {{ $amtToGive := randInt $settings.min $settings.max }} {{/* Amount of XP to give */}}
        {{ $currentXp := 0 }} {{/* User current XP */}}
        {{ with (dbGet .User.ID "xp") }}
            {{ $currentXp = .Value }}
        {{ end }} {{/* Update XP amount if present */}}

        {{ $currentLvl := roundFloor (mult 0.1 (sqrt $currentXp)) }} {{/* Calculate level */}}
        {{ $newXp := dbIncr .User.ID "xp" $amtToGive }} {{/* Increment the xp */}}
        {{ $newLvl := roundFloor (mult 0.1 (sqrt $newXp)) }} {{/* Calculate new level */}}
        {{ $channel := or $settings.channel .Channel.ID }}
        {{ if not (.Guild.GetChannel $channel) }} {{ $channel = .Channel.ID }} {{ end }}

        {{ if ne $newLvl $currentLvl }} {{/* If the level changed / user ranked up */}}
            {{ $newLoc := " Keep it up!" }}
            {{ $type := $roleRewards.type }} {{/* Type of role giving (highest / stack) */}}
            {{ $toAdd := or ($roleRewards.Get (json $newLvl)) 0 }} {{/* Try to get the role reward for this level */}}
            {{ range $level, $reward := $roleRewards }} {{/* Loop over role rewards */}}
                {{- if and (ge (toInt $newLvl) (toInt $level)) (not (hasRoleID $reward)) (eq $type "stack") (ne $level "type") }} {{- addRoleID $reward }}
                {{- else if and (hasRoleID $reward) (eq $type "highest") $toAdd }} {{- removeRoleID $reward }} {{- end -}}
            {{ end }}
            {{ if $toAdd }} {{ addRoleID $toAdd }} {{ $newLoc = (printf "\n\nYou're now in %s!" (getRole $toAdd).Name) }} {{ end }}
            {{ $embed := cembed 
                "title" "❯ Level up!"
                "thumbnail" (sdict "url" "https://webstockreview.net/images/emoji-clipart-celebration-4.png")
                "description" (printf "Congratulations **%s**! You've leveled up to level %d.%s" .User.String (toInt $newLvl) $newLoc)
                "color" 14232643
            }}
                {{ if $settings.announcements }}
                {{ sendMessage $channel (complexMessage "content" .User.Mention "embed" $embed) }} {{/* Send levelup notification */}}
                {{ end }}
        {{ end }}
    {{ end }}

    {{ $cooldownSeconds := div $settings.cooldown 1000000000 }} {{/* Convert cooldown to seconds */}}
    {{ dbSetExpire .User.ID "xpCooldown" true $cooldownSeconds }} {{/* Set cooldown entry */}}
{{ end }}

{{/* vim: set ts=4 sw=4 et: */}}
