# prometheus-nginx-http-status-code-exporter
nginx all status code exporter

This exporter get nginx access log and then after parse from it the returned status code, put it into prometheus exporter client.

This project is just born and this code is partial and under testing, so pay attenction if you want to use it

# INSTALLATION
go build nginx-httpcode-exporter.go
