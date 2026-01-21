package main

import (
	"fmt"
	"strings"
	// 'google/gopacket' kütüphanesi bu iş için standarttır
)

func main() {
	fmt.Println("[*] NetSniffer IDS Başlatıldı... Ağ trafiği analiz ediliyor.")

	// Basit mantık: Gelen paket içeriğinde şüpheli bir string ara
	fakePacketData := "GET /etc/passwd HTTP/1.1" // Örnek bir paket verisi

	analyzePacket(fakePacketData)
}

func analyzePacket(data string) {
	// Suricata/Snort kural mantığı
	rules := []string{"/etc/passwd", "UNION SELECT", "nmap"}

	for _, rule := range rules {
		if strings.Contains(data, rule) {
			fmt.Printf("[ALERT] Saldırı Girişimi Tespit Edildi! Kural: %s\n", rule)
		}
	}
}
