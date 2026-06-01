# Aura Core

Aura is a lightweight, high-performance screen reader for Linux, written in Go.

The project targets a future where accessibility feedback is fast, reliable, observable, and deeply integrated with modern Linux desktop environments through AT-SPI2 and D-Bus.

Aura is developed through an AI-assisted workflow, combining human direction, runtime experimentation, and iterative implementation.

---

## Vision

Aura aims to become a low-latency, reliability-focused Linux screen reader designed around:

* **Real-time accessibility feedback**
* **Concurrency-friendly event processing**
* **Robust speech integration**
* **Observable runtime behavior**
* **Developer-friendly debugging and experimentation**

Long-term architectural goals include:

* Low-latency AT-SPI event processing
* Non-blocking speech orchestration
* Clean separation of runtime concerns
* High reliability under desktop event load
* Practical extensibility without unnecessary complexity

---

## Current Status

Aura is currently in an **experimental CLI diagnostic phase**.

Already implemented:

* AT-SPI accessibility bus discovery (`doctor`)
* Accessibility tree inspection (`list-apps`)
* Raw AT-SPI event observation (`watch-events`)
* Real-time focus event monitoring (`watch-focus`)
* Focus metadata enrichment:

  * application name
  * accessible role
  * accessible name / description fallback
* Audible runtime diagnostics using Speech Dispatcher
* Linux runtime validation against Fedora / GNOME environments

Current implementation favors:

* small bounded slices
* runtime evidence over assumptions
* direct implementations before abstraction
* fast developer feedback loops

---

## Key Characteristics

* **Built with Go:** native performance and strong tooling.
* **AT-SPI / D-Bus Native:** direct interaction with Linux accessibility infrastructure.
* **Low-latency Diagnostics:** real-time runtime inspection and feedback.
* **Speech-enabled Runtime Exploration:** accessibility events can already be observed audibly during development.
* **AI-assisted Development Workflow:** architecture, planning, and implementation are developed collaboratively between human guidance and AI agents.

---

## Tech Stack

* **Language:** Go (Golang)
* **Accessibility:** AT-SPI2
* **Communication:** D-Bus
* **Speech:** Speech Dispatcher
* **Target Platform:** Linux (currently optimized for Fedora + GNOME)

---

## Getting Started

Prerequisites:

* Linux desktop with AT-SPI enabled
* `speech-dispatcher`
* Go toolchain installed

Clone:

```bash
git clone https://github.com/flanfranchi1/aura.git
cd aura
```

Install dependencies:

```bash
go mod tidy
```

Build:

```bash
make build
```

---

## CLI Examples

Environment validation:

```bash
go run ./cmd/aura doctor
```

List top-level accessible applications:

```bash
go run ./cmd/aura list-apps
```

Observe raw accessibility events:

```bash
go run ./cmd/aura watch-events
```

Observe real-time focus changes with audible diagnostics:

```bash
go run ./cmd/aura watch-focus
```

---

## Development

Run tests:

```bash
make test
```

Run build + validation checks:

```bash
make check
```

Format code:

```bash
make fmt
```

---

## License

GNU GPL v3.0 — see [LICENSE](LICENSE).
