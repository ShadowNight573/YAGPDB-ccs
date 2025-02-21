{{/*
		Made by Ranger (765316548516380732)

	Trigger Type: `Regex`
	Trigger: `\A(-|<@!?204255221017214977>\s*)((edit|modify)-?item)(\s+|\z)`

    ©️ Ranger 2020-Present
    GNU, GPLV3 License
    Repository: https://github.com/Ranger-4297/YAGPDB-ccs
*/}}


{{/* Only edit below if you know what you're doing (: rawr */}}

{{/* Initiates variables */}}
{{$successColor := 0x00ff7b}}
{{$errorColor := 0xFF0000}}
{{$prefix := index (reFindAllSubmatches `.*?: \x60(.*)\x60\z` (execAdmin "Prefix")) 0 1 }}

{{/* Create item */}}

{{/* Response */}}
{{$embed := sdict}}
{{$embed.Set "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "1024"))}}
{{$embed.Set "timestamp" currentTime}}
{{$perms := split (index (split (exec "viewperms") "\n") 2) ", "}}
{{if or (in $perms "Administrator") (in $perms "ManageServer")}}
	{{with (dbGet 0 "EconomySettings")}}
		{{$a := sdict .Value}}
		{{$symbol := $a.symbol}}
		{{with (dbGet 0 "store")}}
			{{$store := sdict .Value}}
			{{$items := sdict}}
			{{$value := ""}}
			{{if ($store.Get "Items")}}
				{{$items = $store.Get "Items"}}
				{{with $.CmdArgs}}
					{{$name := (index . 0)}}
					{{if $items.Get $name}}
						{{$options := cslice "description" "role-given" "name" "price" "quantity"}}
						{{if gt (len $.CmdArgs) 1}}
							{{$option := (index . 1) | lower}}
							{{if in $options $option}}
								{{$cont := 0}}
								{{if eq $option "name"}}
									{{if gt (len $.CmdArgs) 2}}
										{{$value = (index . 2)}}
										{{$items.Set $value (($store.Get "Items").Get $name)}}
										{{$items.Del $name}}
										{{$store.Set "Items" $items}}
										{{dbSet 0 "store" $store}}
										{{$cont = 1}}
									{{else}}
										{{$embed.Set "description" (print "No name argument provided :(\nSyntax is `" $.Cmd " <Name> <Option:String> <Value>`")}}
										{{$embed.Set "color" $errorColor}}
									{{end}}
								{{else if eq $option "quantity"}}
									{{if gt (len $.CmdArgs) 2}}
										{{if toInt (index . 2)}}
											{{if ge (toInt (index . 2)) 1}}
												{{$value = (index . 2)}}
												{{$cont = 1}}
											{{else}}
												{{$value = "Infinite"}}
												{{$cont = 1}}
											{{end}}
											{{if $cont}}
												{{$item := $items.Get $name}}
												{{$item.Set "quantity" $value}}
												{{$items.Set $name $item}}
												{{$store.Set "Items" $items}}
												{{dbSet 0 "store" $store}}
											{{end}}
										{{else}}
											{{if eq (lower (index . 2)) "infinite" "infinity" "inf"}}
												{{$value = "Infinite"}}
												{{$item := $items.Get $name}}
												{{$item.Set "quantity" $value}}
												{{$items.Set $name $item}}
												{{$store.Set "Items" $items}}
												{{dbSet 0 "store" $store}}
												{{$cont = 1}}
											{{else}}
												{{$embed.Set "description" (print "Invalid quantity argument provided :(\nSyntax is `" $.Cmd " <Name> <Option:String> <Value>`")}}
												{{$embed.Set "color" $errorColor}}
											{{end}}
										{{end}}
									{{else}}
										{{$embed.Set "description" (print "No quantity argument provided :(\nSyntax is `" $.Cmd " <Name> <Option:String> <Value>`")}}
										{{$embed.Set "color" $errorColor}}
									{{end}}
								{{else if eq $option "price"}}
									{{if gt (len $.CmdArgs) 2}}
										{{if toInt (index . 2)}}
											{{$value = (toInt (index . 2))}}
											{{$item := $items.Get $name}}
											{{$item.Set "price" $value}}
											{{$items.Set $name $item}}
											{{$store.Set "Items" $items}}
											{{dbSet 0 "store" $store}}
											{{$cont = 1}}
										{{else}}
											{{$embed.Set "description" (print "Invalid price argument provided :(\nSyntax is `" $.Cmd " <Name> <Option:String> <Value>`")}}
											{{$embed.Set "color" $errorColor}}
										{{end}}
									{{else}}
										{{$embed.Set "description" (print "No price argument provided :(\nSyntax is `" $.Cmd " <Name> <Option:String> <Value>`")}}
										{{$embed.Set "color" $errorColor}}
									{{end}}
								{{else if eq $option "description"}}
									{{if gt (len $.CmdArgs) 2}}
										{{$value = (joinStr " " (slice $.CmdArgs 2))}}
										{{$item := $items.Get $name}}
										{{$item.Set "description" $value}}
										{{$items.Set $name $item}}
										{{$store.Set "Items" $items}}
										{{dbSet 0 "store" $store}}
										{{$cont = 1}}
									{{else}}
										{{$embed.Set "description" (print "No description argument provided :(\nSyntax is `" $.Cmd " <Name> <Option:String> <Value>`")}}
										{{$embed.Set "color" $errorColor}}
									{{end}}
								{{else if eq $option "role"}}
									{{if gt (len $.CmdArgs) 2}}
										{{$role := (index . 2)}}
										{{if $.Guild.GetRole (toInt64 $role)}}
											{{$value = $role}}
											{{$item := $items.Get $name}}
											{{$item.Set "role-given" $value}}
											{{$items.Set $name $item}}
											{{$store.Set "Items" $items}}
											{{dbSet 0 "store" $store}}
											{{$cont = 1}}
										{{else}}
											{{$embed.Set "description" (print "Invalid role argument provided :(\nSyntax is `" $.Cmd " <Name> <Option:String> <Value>`")}}
											{{$embed.Set "color" $errorColor}}
										{{end}}
									{{else}}
										{{$embed.Set "description" (print "No role argument provided :(\nSyntax is `" $.Cmd " <Name> <Option:String> <Value>`")}}
										{{$embed.Set "color" $errorColor}}
									{{end}}
								{{end}}
								{{if $cont}}
									{{$embed.Set "description" (print $name "'s `" $option "` has been changed to " $value)}}
									{{$embed.Set "color" $successColor}}
								{{end}}
							{{else}}
								{{$embed.Set "description" (print "Invalid option argument provided :(\nSyntax is `" $.Cmd " <Name> <Option:String> <Value>\nAvailable options are: `name`, `description`, `price`, `quantity` and `role``")}}
								{{$embed.Set "color" $errorColor}}
							{{end}}
						{{else}}
							{{$embed.Set "description" (print "No option argument provided :(\nSyntax is `" $.Cmd " <Name> <Option:String> <Value>`\nAvailable options are: `name`, `description`, `price`, `quantity` and `role`")}}
							{{$embed.Set "color" $errorColor}}
						{{end}}
					{{else}}
						{{$embed.Set "description" (print "Invalid item argument provided :(\nSyntax is `" $.Cmd " <Name> <Option:String> <Value>`\nUse `" $prefix "shop` to view the items!")}}
						{{$embed.Set "color" $errorColor}}
					{{end}}
				{{else}}
					{{$embed.Set "description" (print "No item argument provided.\nSyntax is `" $.Cmd " <Name> <Option:String> <Value>`\nUse `" $prefix "shop` to view the items!")}}
					{{$embed.Set "color" $errorColor}}
				{{end}}
			{{else}}
				{{$embed.Set "description" (print "There are no items :(\nAdd some items with `" $prefix "create-item <Name> <Price:Int> <Quantity:Int> <Description:String>`")}}
				{{$embed.Set "color" $errorColor}}
			{{end}}
		{{end}}
	{{else}}
			{{$embed.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" $prefix "set default`")}}
			{{$embed.Set "color" $errorColor}}
	{{end}}
{{else}}
	{{$embed.Set "description" (print "Insufficient permissions.\nTo use this command you need to have either `Administrator` or `ManageServer` permissions")}}
	{{$embed.Set "color" $errorColor}}
{{end}}
{{sendMessage nil (cembed $embed)}}