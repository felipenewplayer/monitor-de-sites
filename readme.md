# 🕵️‍♂️ Site Monitor Go

Aplicativo em Go para monitorar a disponibilidade de sites. Ele verifica periodicamente se os sites estão online e registra os status em um arquivo de log.

---

## 🚀 Funcionalidades

- Checagem periódica de múltiplos sites
- Verifica status HTTP (200 = OK)
- Log de cada verificação com data e hora
- Simples, rápido e eficiente
- Docker-ready 🐳

---

## 🛠️ Tecnologias

- [Go](https://golang.org/)
- HTTP Client nativo
- Arquivo de log (`log.txt`)
- Docker (opcional)

---

## 🧪 Como usar localmente

```bash
# Clonar o repositório
git clone https://github.com/felipenewplayer/site-monitor-go.git
cd site-monitor-go

# Rodar localmente
go run main.go
