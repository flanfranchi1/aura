# Aura Core

Aura is a lightweight, high-performance screen reader for Linux (Fedora/GNOME), developed in Go. 

Designed for speed and reliability, Aura leverages the AT-SPI2 bus via D-Bus to provide real-time 
accessibility feedback. Unlike traditional screen readers, Aura is built with a concurrency-first 
approach, ensuring that system events and speech synthesis never block each other.

## 🚀 Key Features

- **Built with Go:** Native performance and efficient memory management.
- **Asynchronous Speech:** Integrated with Speech Dispatcher via Unix Sockets using a non-blocking architecture.
- **Low Latency:** High-speed event processing through the AT-SPI2 Registry.
- **Developer Centric:** Clean code structure, focused on extensibility and business value.

## 🛠️ Tech Stack

- **Language:** Go (Golang)
- **Communication:** D-Bus (AT-SPI2 Protocol)
- **Speech Synthesis:** Speech Dispatcher (SSIP Protocol)
- **OS Target:** Linux (optimized for Fedora with GNOME)

## 📦 Getting Started

1. Ensure `speech-dispatcher` is running on your system.
2. Clone the repository:
   ```bash
   git clone [https://github.com/flanfranchi1/aura.git](https://github.com/flanfranchi1/aura.git)

## License
This project is licensed under the GNU General Public License v3.0 - see the [LICENSE](LICENSE) file for details.