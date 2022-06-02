{{/**************************************************************************\

    YAGPDB Custom Command - Server Population Display
    -------------------------------------------------

    Edits a post with the current population stats (this could be a pinned post
    for example). Make it run in a bot spam channel, even if it's editing a
    post in another channel.

    Trigger type: Minute Interval
    Interval: 15 minutes

    Setup: Change channel and post IDs below
    Setup: Restrict to only run in botspam channels
    Setup: Restrict to only admins (e.g. put in administrator cc group)

  \**************************************************************************/}}

{{ $channelID := 971522137510785094 }}
{{ $postID := 981434444105777162 }}

{{ $embed := cembed
    "title" "Server Population"
    "color" 13585960
    "fields" (cslice
        (sdict "name" "Total" "value" (toString .Guild.MemberCount) "inline" true)
        (sdict "name" "Online" "value" (toString onlineCount) "inline" true)
    )
    "footer" (sdict "text" "Last update (every 15m):")
    "timestamp" currentTime.UTC
}}
{{ editMessage $channelID $postID (complexMessageEdit "embed" $embed "content" "") }}

{{/* copy and paste these lines to change other posts too */}}
{{ $channelID = 971816708841046076 }}
{{ $postID = 981959484861714532 }}
{{ editMessage $channelID $postID (complexMessageEdit "embed" $embed "content" "") }}




{{/* alternate display style, using the embed's description:
"description" (print "Total Members: **" .Guild.MemberCount "**\n" "Online Members: **" onlineCount "**")
 */}}
