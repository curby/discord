{{/**************************************************************************\

    YAGPDB Custom Command - "wroteIntro" status toggle
    --------------------------------------------------

    Allows a moderator to toggle a member's "wroteIntro" status (indicating
    whether they've written an intro post.

    Trigger type: Command
    Trigger string: wroteIntro

    Setup: Change intro channel ID below
    Setup: Restrict to only run in botspam channels
    Setup: Restrict to only moderators (e.g. put in moderator cc group)

  \**************************************************************************/}}

{{ $args := parseArgs 1 "Usage: wroteIntro <userID> [introMessageID]"
    (carg "userid" "userID")
    (carg "string" "messageID")
}}

{{ $introChannelID := 975085554905514014 }}
{{ $key := "wroteIntro" }}

{{ $user := ($args.Get 0) }}

{{ $wroteIntro := dbGet $user $key }}
{{ $username := (userArg $user).Username }}

{{ if $wroteIntro }}
    {{ dbDel $user $key }}
    Removed `wroteIntro` flag from {{ $username }}.
{{ else }}
    {{ $messageID := ($args.Get 1) }}
    {{ $message := getMessage $introChannelID $messageID }}
    {{ if $message }}
        {{ if eq $message.Author.ID (toInt $user) }}
            {{ dbSet $user $key $messageID }}
            Added `wroteIntro` flag to {{ $username }}.
        {{ else }}
            The given message's author is not {{ $username }}.
        {{ end }}
    {{ else }}
        {{ $username }} doesn't have a recorded intro messageID. To add one, give the messageID as an argument:
        `Usage: wroteIntro <userID> [introMessageID]`.
    {{ end }}
{{ end }}
