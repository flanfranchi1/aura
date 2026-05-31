Project Aura: Architecture Overview

Aura utilizes a pragmatic, layered architecture written in idiomatic Go. The primary goal is to isolate the complex, stateful business logic of the screen reader from the chaotic, asynchronous nature of Linux.

Architecture favors explicit data flow, minimal abstractions, and observable behavior to support sustainable AI-assisted development. accessibility APIs.

Layered Design
ADR-007: Layer Ownership and Dependency Rules

Decision:
Define dependency boundaries among Domain, Provider, Adapter, Speech, and UX layers.

Context:
Prevent accidental coupling between business logic and infrastructure concerns.
1. Domain Layer (internal/domain)

The core of Aura. This layer defines the canonical interfaces and models. It contains the logic for what should be spoken, how focus is tracked, and how the user's configuration alters output.

Dependencies: None. This layer imports only standard Go libraries. It does not know about AT-SPI, DBus, or Speech Dispatcher.

2. Provider Layer (internal/provider)

Defines the contracts used by the Domain to interact with accessibility backends and external services.

Event Stream: Channels or simple interfaces for receiving accessibility events.

Tree Navigation: Interfaces for querying nodes (Parent, Children).

3. Adapters Layer (internal/adapters)

The concrete implementations of the Provider interfaces.

AT-SPI Adapter (internal/adapters/atspi): Connects to DBus, listens for AT-SPI signals, and translates AT-SPI structs into Aura's canonical Domain models. This layer handles DBus connection pooling, signal debouncing, and caching.

4. Speech Layer (internal/speech)

Handles text-to-speech output.

Speech Dispatcher Adapter: Connects to the speech-dispatcher daemon, managing speech queues, interruption (canceling current speech on new focus), and rate/pitch configuration.

5. UX/Interaction Layer (internal/ux)

Manages keyboard hooking, the global keymap, and user commands. It translates physical key presses into Domain actions (interaction policy lives here, not inside providers)

Canonical Model

To maintain the backend-agnostic principle, Aura uses a canonical model to represent the accessibility tree.

Node: An struct representing an accessible object. Exposes methods to get ID, Name, Description, Role, States, and relationships (Parent, Children).

Event: A structured representation of something that happened (e.g., FocusEvent, StateChangeEvent).

Role/State: Typed constants representing standard UI concepts (e.g., RoleButton, StateFocused), abstracted away from AT-SPI specific integers.

Event Flow (MVP)

The OS triggers a focus change in an application.

The AT-SPI Adapter receives a DBus signal.

The EventNormalization treats what is coming fromATSPI, discarding what is noise and filtering what is relevant to be processed. 

The Adapter fetches the node details, translates them into a canonical FocusEvent, and pushes it to an internal Go channel.

The Domain Layer's main event loop consumes the FocusEvent.

The Domain Layer applies user configuration and verbosity rules to determine the text to be spoken.

The Domain Layer sends the finalized text string to the Speech Layer.

The Speech Layer manages speech policy and outpu    t delivery then instructs Speech Dispatcher to synthesize the audio, interrupting any currently playing audio.