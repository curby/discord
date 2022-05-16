{{/**********************************************************************
     Voyages channel manager
   **********************************************************************/}}

{{/* Initialize */}}
{{ $gameNumber := 0 }}
{{ $gameStamp := 0 }}
{{ $gameEnd := 0 }}
{{ $roundNumber := -1 }}
{{ $roundStamp := 0 }}
{{ $gameDelay := 1200 }}
{{ $roundDelay := 6 }}

{{/* Grab game metadata */}}
{{ $game := (dbGet .CCID "game").Value }}
{{ if $game }}
	{{ $gameNumber = toInt (index $game 0) }}
	{{ $gameStamp = toInt (index $game 1) }}
	{{ $gameEnd = toInt (index $game 2) }}
{{ end }}

{{ if eq "new game" (lower .Cmd) }}
	{{/* See if we can make a new game */}}
	{{ if lt currentTime.Unix (add $gameStamp $gameDelay) }}
		Wait a little bit before starting another game.
		{{ deleteResponse 10 }}
	{{ else }}
		{{/* Update voter list */}}
		{{ $voters := (dbGet .CCID "voters").Value }}
		{{ if $voters }}
			{{ if $voters.HasKey (str .User.ID) }}
				You've already voted. Another <@&930737911626883112> must vote.
				{{ deleteResponse 10 }}
			{{ else }}
				{{ $voters.Set (str .User.ID) 1 }}
			{{ end }}
		{{ else }}
			{{ $voters = (sdict (str .User.ID) 1) }}
			<@{{.User.ID}}> voted to start a new game. One more vote needed to start.
				{{ deleteResponse 120 }}
		{{ end }}
		{{/* Start game if we have enough votes */}}
		{{ if eq 2 (len $voters) }}
			{{ dbDel .CCID "voters" }}
			{{ $gameNumber = add 1 $gameNumber }}
			{{ $gameEnd = 0 }}
			{{ dbSet .CCID "game" (cslice $gameNumber currentTime.Unix $gameEnd) }}
			{{ dbSet .CCID "round" (cslice 0 $roundStamp) }}
			{{ sendMessage nil (print "**Game " (str $gameNumber) " begins!**") }}
		{{ else }}
			{{ dbSet .CCID "voters" $voters }}
		{{ end }}
	{{ end }}
{{ end }}

{{/* Grab round metadata */}}
{{ $round := (dbGet .CCID "round").Value }}
{{ if $round }}
	{{ $roundNumber = toInt (index $round 0) }}
	{{ $roundStamp = toInt (index $round 1) }}
{{ end }}

{{ if eq "three stars" (lower .Cmd) }}
	{{ if or (eq $gameNumber 0) (eq $gameEnd 2) }}
		No games running. Use `new game` to start one.
		{{ deleteResponse 10 }}
	{{ else }}
		{{ dbSet .CCID "game" (cslice $gameNumber $gameStamp 1) }}
			{{ $status := print "<@&930737911626883112>: " .User.Mention " has three stars, so the next round will be the last round this game!" }}
			{{sendMessageNoEscape nil ($status) }}
	{{ end }}
{{ end }}


{{ if or (eq "next round" (lower .Cmd)) (eq $roundNumber 0) }}
	{{ if lt currentTime.Unix (add $roundStamp $roundDelay) }}
		Wait a little bit before starting the next round.
		{{ deleteResponse 10 }}
	{{ else if or (eq $gameNumber 0) (eq $gameEnd 2) }}
		No games running. Use `new game` to start one.
		{{ deleteResponse 10 }}
	{{ else }}
		{{ $roundNumber = add 1 $roundNumber }}
		{{ dbSet .CCID "round" (cslice $roundNumber currentTime.Unix) }}
		{{ $status := print "<@&930737911626883112>: Game " (str $gameNumber) ", Round " (str $roundNumber) ": **" }}
		{{ $pips1 := "" }}
		{{ $pips2 := "" }}
		{{ $pips3 := "" }}
		{{ range (seq 1 4) }}
			{{ $d6 := randInt 1 7 }}
			{{ $status = print $status (str $d6) }}
			{{ if eq $d6 1 }}
				{{ $pips1 = print $pips1 "|       |" }}
				{{ $pips2 = print $pips2 "|   O   |" }}
				{{ $pips3 = print $pips3 "|       |" }}
			{{ else if eq $d6 2 }}
				{{ $pips1 = print $pips1 "|     O |" }}
				{{ $pips2 = print $pips2 "|       |" }}
				{{ $pips3 = print $pips3 "| O     |" }}
			{{ else if eq $d6 3 }}
				{{ $pips1 = print $pips1 "|     O |" }}
				{{ $pips2 = print $pips2 "|   O   |" }}
				{{ $pips3 = print $pips3 "| O     |" }}
			{{ else if eq $d6 4 }}
				{{ $pips1 = print $pips1 "| O   O |" }}
				{{ $pips2 = print $pips2 "|       |" }}
				{{ $pips3 = print $pips3 "| O   O |" }}
			{{ else if eq $d6 5 }}
				{{ $pips1 = print $pips1 "| O   O |" }}
				{{ $pips2 = print $pips2 "|   O   |" }}
				{{ $pips3 = print $pips3 "| O   O |" }}
			{{ else if eq $d6 6 }}
				{{ $pips1 = print $pips1 "| O   O |" }}
				{{ $pips2 = print $pips2 "| O   O |" }}
				{{ $pips3 = print $pips3 "| O   O |" }}
			{{ end }}
			{{ if ne . 3 }}
				{{ $status = print $status ", " }}
				{{ $pips1 = print $pips1 "  " }}
				{{ $pips2 = print $pips2 "  " }}
				{{ $pips3 = print $pips3 "  " }}
			{{ else }}
				{{ $status = print $status "**" }}
			{{ end }}
		{{ end }}
		{{ sendMessageNoEscape nil (joinStr "\n" $status "```.-------.  .-------.  .-------." $pips1 $pips2 $pips3 "'-------'  '-------'  '-------'```") }}
		{{ if eq $gameEnd 1 }}
			{{ sendMessage nil (print "**Game " (str $gameNumber) " is over!**") }}
			{{ dbSet .CCID "game" (cslice $gameNumber $gameStamp 2) }}
		{{ end }}
	{{ end }}
{{ end }}
{{ deleteTrigger 1 }}
