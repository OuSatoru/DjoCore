package setIE

import (
	"golang.org/x/sys/windows/registry"
	"log"
)

func GetProxy() string {
	k, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Internet Settings`, registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()

	s, _, err := k.GetStringValue("ProxyServer")
	if err != nil {
		log.Fatal(err)
	}
	return s
}

func GetWhether() bool {
	//Enable proxy or not
	k, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Internet Settings`, registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()

	s, _, err := k.GetIntegerValue("ProxyEnable")
	if err != nil {
		log.Fatal(err)
	}
	if s == 1 {
		return true
	} else {
		return false
	}
}

func SetProxy(link string) {
	SetWhether(true)
	k, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Internet Settings`, registry.SET_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()

	k.SetStringValue("ProxyServer", link)

}

func SetWhether(q bool) {
	//Enable proxy or not
	k, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Internet Settings`, registry.SET_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()
	if q {
		k.SetDWordValue("ProxyEnable", 0x00000001)
	} else {
		k.SetDWordValue("ProxyEnable", 0x00000000)
	}

}

