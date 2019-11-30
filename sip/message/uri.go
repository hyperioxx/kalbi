package message

import "regexp"

type URI struct {
	Scheme string
	User string
	Password string
	Host string
	Port string
	Tag string
	Params string
	Headers string
	Branch string

}


func parseURI(uri string) *URI{
	newuri := new(URI)
	re := regexp.MustCompile(`(?P<scheme>\w+):(?:(?P<user>[+\w\.\-]+):?(?P<password>[\w\.]+)?@)?\[?(?P<host>(?:\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})|(?:(?:[0-9a-fA-F]{1,4}):){7}[0-9a-fA-F]{1,4}|(?:(?:[0-9A-Za-z]+\.)+[0-9A-Za-z]+))\]?:?(?P<port>\d{1,6})?(.*;)?(tag=(?P<tag>.*))?\;?(?:\?(?P<headers>.*))?`)
	match := re.FindStringSubmatch(uri)
	for i, name := range re.SubexpNames() {
        if i != 0 && i <= len(match) {
			switch name {
			case "scheme":
				newuri.Scheme = match[i]
			case "user":
				newuri.User = match[i]
			case "password":
				newuri.Password = match[i]
			case "host":
				newuri.Host = match[i]
			case "port":
				newuri.Port = match[i]
			case "tag":
				newuri.Tag = match[i]
			default:

			}  
        }
    }
    
	return newuri
}