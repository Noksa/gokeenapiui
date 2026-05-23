<div align="center">

# 🎨 Gokeenapi UI

**Графический интерфейс для управления AWG VPN соединениями на роутерах Keenetic**

[![GitHub release](https://img.shields.io/github/release/Noksa/gokeenapiui.svg)](https://github.com/Noksa/gokeenapiui/releases)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

*Не любите командную строку? Управляйте AWG VPN соединениями на роутере Keenetic через удобный графический интерфейс.*

[🚀 Установка](#-установка) • [✨ Возможности](#-возможности) • [🔧 Сборка из исходников](#-сборка-из-исходников) • [💻 CLI версия](https://github.com/Noksa/gokeenapi)

</div>

---

Данный проект является GUI версией утилиты [gokeenapi](https://github.com/Noksa/gokeenapi).

## 📋 Поддерживаемые функции

Приложение охватывает **только управление AWG (AmneziaWG) соединениями**. Стандартный WireGuard и другие типы VPN не поддерживаются. Доступные операции:

- Создание нового AWG VPN соединения из `.conf` файла
- Добавление статических маршрутов к интерфейсу WireGuard/AWG (из `.bat` файлов или URL)
- Удаление всех статических маршрутов с интерфейса
- Просмотр всех WireGuard интерфейсов на роутере

Требуется роутер Keenetic в режиме **Router** с установленным компонентом **WireGuard VPN**. Режим Extender/Repeater не поддерживается.

---

## 📖 Описание

Приложение позволяет легко настраивать AWG VPN соединения на роутерах Keenetic через удобный графический интерфейс. Поддерживается подключение как к локальному роутеру, так и через интернет с использованием KeenDNS.

---

## 🚀 Установка

### Готовые сборки

Скачайте последний релиз со [страницы Releases](https://github.com/Noksa/gokeenapiui/releases).

| Платформа | Файл | Примечания |
|-----------|------|------------|
| Windows (x64) | `gokeenapiui.exe` | Установка не нужна — запускайте напрямую |
| macOS (Apple Silicon) | Сборка из исходников | См. [Сборка из исходников](#-сборка-из-исходников) |

Последний релиз: [v0.2.1](https://github.com/Noksa/gokeenapiui/releases/tag/v0.2.1)

---

## 🔧 Сборка из исходников

**Предварительные требования:**

- [Go 1.21+](https://go.dev/dl/)
- [Node.js 18+](https://nodejs.org/)
- [Wails v2](https://wails.io/docs/gettingstarted/installation): `go install github.com/wailsapp/wails/v2/cmd/wails@latest`

**Сборка:**

```bash
# Клонировать репозиторий
git clone https://github.com/Noksa/gokeenapiui.git
cd gokeenapiui

# Собрать для текущей платформы
wails build

# Или собрать для конкретной платформы (требует кросс-компиляционный тулчейн)
wails build -platform windows/amd64
wails build -platform darwin/arm64
```

Бинарный файл будет помещён в `build/bin/`.

Либо используйте Makefile для сборки всех поддерживаемых платформ сразу:

```bash
make binaries
```

---

## ✨ Возможности

- 🌐 Подключение к роутеру по HTTP/HTTPS
- 📄 Загрузка AWG конфигурационных файлов
- ⚡ Автоматическое создание VPN соединения
- 🔒 Безопасная аутентификация
- 🛣️ Управление статическими маршрутами (добавление/удаление) из BAT файлов или URL
- 📋 Просмотр всех WireGuard интерфейсов на роутере
- 💡 Интуитивный интерфейс с подсказками
- ✅ Валидация полей ввода
- 🎨 Современный дизайн с анимациями

---

## 💻 Предпочитаете командную строку?

Если вам нужна автоматизация, расширенные команды или запуск через Docker — используйте CLI версию:

<div align="center">

### [💻 **CLI версия gokeenapi** 🚀](https://github.com/Noksa/gokeenapi)

</div>
