{{/**************************************************************************\

    YAGPDB Custom Command - Twitter Auto-embeds
    -------------------------------------------

    Reposts Twitter links so they trigger Discord previews.

    Trigger type: Regex
    Trigger string: ([^<]|^)https?:\/\/(twitter|x)\.com\/\S+\/status\/\S+

    Original from Discord user standardquip
    https://discord.com/channels/166207328570441728/384011387132706816/1169758137310523413

    Modified by cur.by:
        Ignore links in <angle brackets>
        Support HTTP links (but translate them into HTTPS)
        More comments

  \**************************************************************************/}}

{{/* Only fix if message had NO embeds (this could false positive if something else triggers an embed) */}}
{{if .Message.Embeds}}
    {{return}}
{{end}}

{{$msg := .Message.Content}}

{{/* Rehome to vxtwitter */}}
{{$vx := reReplace `https?:\/\/(twitter|x)\.com` $msg "https://vxtwitter.com" }}

{{/* Remove query string (potential tracking) */}}
{{$vx = reReplace `(https?\S*?)\?.*?(\s|$)` $vx "$1$2" }}

{{/* Extract links from message */}}
{{$links := reFindAll `https:\/\/vxtwitter.com\/\S+\/status\/\S+` $vx}}
{{$list := ""}}
{{range $links -}}
    {{- $list = print $list . " \n" -}}
{{- end}}

{{sendMessage nil (complexMessage 
    "content" (print "Fixed embeds below:\n" $list) 
    "reply" .Message.ID
)}}
