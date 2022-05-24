{{/**************************************************************************\

    YAGPDB Custom Command - Song template for karaoke machine
    ---------------------------------------------------------

    Change the CCID to that of the karaoke machine script on your server
    (this only has to be done once per server). Then, for each song,
    copy this template, enter the lyrics, and set an appropriate trigger.

    See samplesong.go for a detailed example with sample trigger regex.

    Trigger type: Regex
    Trigger string: [put regex here for documentation]

    Setup: Change CCID to the karaoke machine script's ID.

  \**************************************************************************/}}

{{ $lyrics := cslice
"[put lyrics here]"
}}

{{/* Set to the number of the CC that will display lyrics */}}
{{ $lyricsCCID := 31 }}
{{/* This logic *must* exist outside of the karaoke script or else
     the mutex warning will never trigger (triggers will just queue up). */}}
{{ if $mutex := dbGet $lyricsCCID "mutex" }}
    Sorry, I can't sing two songs at once!
    {{ deleteResponse 10 }}
{{ else }}
    {{ dbSet $lyricsCCID "mutex" "mutex" }}
    {{ execCC $lyricsCCID nil 0 (sdict "triggerID" .Message.ID "lyrics" $lyrics) }}
{{ end }}
