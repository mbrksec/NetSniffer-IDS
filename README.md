🛡️ NetSniffer IDS (Intrusion Detection System)

NetSniffer, ağ trafiğini ham paket (raw packet) seviyesinde analiz eden, imza tabanlı (signature-based) bir Saldırı Tespit Sistemi (IDS) prototipidir. Modern siber güvenlik araçlarından olan Suricata ve Snort mimarilerini temel alarak, ağ arayüzünden geçen verileri gerçek zamanlı olarak tarar ve önceden tanımlanmış saldırı paternlerini tespit eder.

🚀 Öne Çıkan Özellikler

L7 Packet Inspection: Uygulama katmanındaki (HTTP, DNS vb.) paket içeriklerini derinlemesine inceler.

Signature-Based Logic: SQL Injection, Directory Traversal ve Malicious User-Agent gibi saldırı imzalarını içeren özel bir kural motoruna sahiptir.

Concurrency with Go: Go dilinin goroutine yapısı sayesinde ağ trafiğini eşzamanlı ve yüksek performansla işler.

Real-time Alerting: Şüpheli bir aktivite algılandığında anlık olarak uyarı (alert) üretir.

🛠️ Teknik Altyapı

Dil: Go (Golang).

Mimari: Signature-based Detection Engine.

Kapsam: Ağ Güvenliği, Paket Analizi, Saldırı Önleme Sistemleri.

📖 Çalışma Mantığı

NetSniffer, ağ arayüzünü dinlemeye başladıktan sonra her bir paketi yakalar ve içerisindeki veriyi aşağıdaki kurallar dizisiyle karşılaştırır:

Reconnaissance: Nmap veya diğer tarayıcı izlerini yakalar.

Web Attacks: /etc/passwd erişimi veya UNION SELECT gibi enjeksiyon denemelerini tespit eder.

Anomali Tespiti: Beklenmeyen veri paternlerini raporlar.

🚀 Kurulum

# Projeyi klonlayın

git clone https://github.com/mbrksec/NetSniffer-IDS.git

# Proje dizinine gidin

cd NetSniffer-IDS

# Bağımlılıkları yükleyin ve çalıştırın

go run sniffer.go

⚖️ Yasal Uyarı

Bu araç, eğitim ve siber güvenlik farkındalığı oluşturma amacıyla geliştirilmiştir. Sadece yetkiniz olan ağlar üzerinde test edilmelidir. İzinsiz ağ izleme faaliyetleri yasal sorumluluk doğurabilir.