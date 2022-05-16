{{/********************************************************************** 
     Intro channel manager
   **********************************************************************/}}

{{ $key := "wroteIntro" }}
{{ $introChannelID := 975085554905514014 }}
{{ $modGroupID := 971522397775736833 }}
{{ $generalChannelID := 971484195098599506 }}
{{ $generalChannelID = 972209499123380335 }}
{{ $imageLinkRegex := `https?:\/\/(?:\w+\.)?[\w-]+\.[\w]{2,3}(?:\/[\w-_.]+)+\.(?:png|jpg|jpeg|gif|webp)` }}
{{ $previewLifetime := 3600 }}
{{ $avatar := (joinStr "" "https://cdn.discordapp.com/avatars/" (toString .User.ID) "/" .User.Avatar ".png") }}

{{ $introMessageID := dbGet .User.ID $key }}
{{/* XXX if message no longer exists, update dbentry but don't warn/delete */}}
{{ if $introMessageID }}
    {{ with getMessage $introChannelID $introMessageID.Value }}
        {{ $introLink := (print "https://discord.com/channels/971484194633048235/" $introChannelID "/" $introMessageID.Value) }}
        {{ $embed := cembed
            "color" 12328245
            "description" (print "Please use <#" $generalChannelID "> or another discussion channel for general conversations, so this channel can be dedicated to introductions.\n\nIf you want to update your introduction, edit your [existing introduction](" $introLink ").\n\nIf you've never posted in this channel, there may be a bug (notify <@&" $modGroupID ">).")
    "footer" (sdict "text" "Note: Your message, along with this notification, will be deleted in a minute to keep this channel clean. If you'd like to repost your message elsewhere, copy its text soon so you can paste it into another channel.")
        }}
        {{ $messageID := sendMessageRetID nil $embed }}
        {{ deleteTrigger 70 }}
        {{ deleteMessage nil $messageID 70 }}
    {{ else }}
        {{ dbSet .User.ID $key (str .Message.ID) }}
    {{ end }}
{{ else }}
    {{ dbSet .User.ID $key (str .Message.ID) }}
    {{ $msg := getMessage $introChannelID .Message.ID }}
    {{ $messageLink := (print "https://discord.com/channels/971484194633048235/" $introChannelID "/" (str .Message.ID)) }}
    {{ $embed := sdict
        "color" 13585960
        "description" (print "**[" $msg.Author.String " wrote an intro!](" $messageLink ")**\n\n")
        "fields" (cslice )
        "footer" (sdict "text" "(This will auto-delete in an hour, but #introductions still has the original.)")
    }}

    {{/* If we add an image to the embed, we won't also add a thumbnail. */}}
    {{ $hasImage := false }}
    {{ with $msg.Content }}
        {{/* See if post links to an image. */}}
        {{ with reFind $imageLinkRegex . }}
            {{ $embed.Set "image" (sdict "url" .) }}
            {{ $hasImage = true }}
        {{ end }}
        {{ $content := . }}
        {{ if gt (len .) 1000 }} {{ $content = slice . 0 1000 | printf "%s ..." }} {{ end }}
        {{ $embed.Set "description" (joinStr "" $embed.description $content) }}
    {{ end }}
    {{ with $msg.Attachments }}
        {{ $attachment := (index . 0).URL }}
        {{ $filename := (index . 0).Filename }}
        {{ if reFind `\.(png|jpg|jpeg|gif|webp)$` $attachment }}
            {{ $embed.Set "image" (sdict "url" $attachment) }}
            {{ $hasImage = true }}
        {{ else }}
            {{ $embed.Set "fields" (cslice
                (sdict "name" "File Name" "value" (print "`" $filename "`") "inline" true)
                (sdict "name" "URL" "value" (print "[File Link](" $attachment ")") "inline" true))
            }}
        {{ end }}
    {{ end }}
    {{ with $msg.Embeds }}
        {{ $em := (index . 0) }}
        {{ if $em.Title }}
            {{ $embed.Set "fields" ($embed.fields.Append (sdict "name" "Linked Page" "value" $em.Title)) }}
        {{ end }}

        {{ if not $hasImage }}
            {{ if $em.Image }}
                {{ $embed.Set "thumbnail" ($em.Image) }}
                {{ $hasImage = true }}
            {{ else if $em.Thumbnail }}
                {{ $embed.Set "thumbnail" ($em.Thumbnail) }}
                {{ $hasImage = true }}
            {{ end }}
        {{ end }}
    {{ end }}

    {{/* If no image so far, set to user avatar (only works if user has custom avatar) */}}
    {{ if not $hasImage }}
        {{ $embed.Set "thumbnail" (sdict "url" $avatar) }}
    {{ end }}

    {{/* Create timed-life preview */}}
    {{ $previewID := sendMessageRetID $generalChannelID (cembed $embed) }}
    {{ deleteMessage $generalChannelID $previewID $previewLifetime }}
{{ end }}
{{/* vim: set tabstop=4:softtabstop=4:shiftwidth=4   */}}
