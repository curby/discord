{{/********************************************************************** 
    Update the embeds used in role commands
   **********************************************************************/}}

{{/* 
    To find the UID of a custom emoji, type \emoji in the chat.
    For example, if you type: \:g_smitten: 
    Discord shows: <:g_smitten:1018221505705955410>
 */}}


{{ if eq (len .Args) 3 }}
    {{ $postID := toInt (index .CmdArgs 1) }}
    {{ if eq (toInt (index .CmdArgs 0)) 1 }}
        {{ $embed := cembed
            "title" "Stonemaier Champion" 
            "color" 13585960
            "description" "If you're a Stonemaier Champion, click the <:champion:972218587454517298> reaction below to grant yourself the <@&971618075600367616> role and show your support.\n\nThis is just for fun and is not linked to your real Stonemaier Champion account. You can find out more about becoming a Stonemaier Champion here:\n<https://stonemaier-games.myshopify.com/pages/stonemaier-champion>"
            "footer" (sdict "text" "Note: Click the reaction again to remove the role.")
        }}
        {{ editMessage nil $postID (complexMessageEdit "embed" $embed "content" "") }}
        Refreshed {{ (index .CmdArgs 0) }}
    {{ else if eq (toInt (index .CmdArgs 0)) 2 }}
        {{ $embed := cembed
            "title" "Stonemaier Games (1/3)"
            "color" 13585960
            "description" "To let others know what you're playing, or just show your love of specific games, click one or more reactions below to add those roles to your profile. Games are sorted chronologically; other posts below list more games.\n\n<:g_viticulture:973685945859731487> Viticulture\n<:g_euphoria:973685945926840350> Euphoria: Build a Better Dystopia\n<:g_b2cities:973685945972977694> Between Two Cities\n<:g_scythe:973685945855508510> Scythe\n\n<:g_charterstone:973685945817784350> Charterstone\n<:g_mylittlescythe:973685945553530994> My Little Scythe\n<:g_b2castles:973685945826160680> Between Two Castles of Mad King Ludwig\n<:g_wingspan:973694985629229116> Wingspan"
            "footer" (sdict "text" "Note: Click a reaction again to remove its role. These cosmetic roles are just for fun; you can add or remove them at any time.")
        }}
        {{ editMessage nil $postID (complexMessageEdit "embed" $embed "content" "") }}
        Refreshed {{ (index .CmdArgs 0) }}
    {{ else if eq (toInt (index .CmdArgs 0)) 3 }}
        {{ $embed := cembed
            "title" "Stonemaier Games (2/3)" 
            "color" 13585960
            "description" "<:g_tapestry:973685945553547275> Tapestry\n<:g_pendulum:973685945918435368> Pendulum\n<:g_redrising:973685945859706890> Red Rising\n<:g_rollingrealms:973685945947803668> Rolling Realms\n\n<:g_libertalia:973685945830350909> Libertalia: Winds of Galecrest\n<:g_smitten:1018221505705955410> Smitten\n<:g_expeditions:1070427602822631464> Expeditions\n<:g_apiary:1148982076901691422> Apiary"
        }}
        {{ editMessage nil $postID (complexMessageEdit "embed" $embed "content" "") }}
        Refreshed {{ (index .CmdArgs 0) }}
    {{ else if eq (toInt (index .CmdArgs 0)) 4 }}
        {{ $embed := cembed
            "title" "Other Custom Roles" 
            "color" 10066329
            "description" "**Automa Fan:** Click the <:automa:973334516410249297> reaction below to show your support of the Automa Factory and their great solo modes!\n\n**Play Async:** If you like to play board games on BGA or other online platforms, click the <:bga:983483565327151125> reaction below to add this role, and be notified when someone mentions the role (usually to announce a new game)."
            "footer" (sdict "text" "Note: Click a reaction again to remove its role. These roles are just for fun; you can add or remove them at any time.")
        }}
        {{ editMessage nil $postID (complexMessageEdit "embed" $embed "content" "") }}
        Refreshed {{ (index .CmdArgs 0) }}
    {{ else if eq (toInt (index .CmdArgs 0)) 5 }}
        {{ $embed := cembed
            "title" "Stonemaier Games (3/3)" 
            "color" 13585960
            "description" "<:g_wyrmspan:1192124259405930517> Wyrmspan\n<:g_vantage:1233528548308947038> Vantage\n<:g_stampswap:1273261115886927943> Stamp Swap"
        }}
        {{ editMessage nil $postID (complexMessageEdit "embed" $embed "content" "") }}
        Refreshed {{ (index .CmdArgs 0) }}
    {{ else }}
        Usage: update-role-text <1|2|3|4|5> <postID>
    {{ end }}
{{ else }}
    Usage: update-role-text <1|2|3|4|5> <postID>
{{ end }}
{{ deleteTrigger 10 }}
{{ deleteResponse 10 }}

{{/* vim: set ts=4 sw=4 et: */}}
