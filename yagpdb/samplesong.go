{{ $lyrics := cslice
"It's not unusual to be loved by anyone"
"It's not unusual to have fun with anyone"
"But when I see you hanging about with anyone"
"It's not unusual to see me cry"
"Oh I wanna die"
"It's not unusual to go out at any time"
"But when I see you out and about it's such a crime"
"If you should ever want to be loved by anyone"
"It's not unusual it happens every day"
"No matter what you say"
"You find it happens all the time"
"Love will never do, what you want it to"
"Why can't this crazy love be mine?"
"[Saxophone]"
"It's not unusual, to be mad with anyone"
"It's not unusual, to be sad with anyone"
"But if I ever find that you've changed at anytime"
"It's not unusual to find out I'm in love with you"
"Whoa-oh-oh-oh-oh"
}}

{{/* Set to the number of the CC that will display lyrics */}}
{{ $lyricsCCID := 31 }}
{{ if $mutex := dbGet $lyricsCCID "mutex" }}
	Sorry, I can't sing two songs at once!
	{{ deleteResponse 10 }}
{{ else }}
	{{ dbSet $lyricsCCID "mutex" "mutex" }}
	{{ execCC $lyricsCCID nil 0 (sdict "triggerID" .Message.ID "lyrics" $lyrics) }}
{{ end }}
