Project Aura: Vision and Scope

Mission

Create a modern Linux desktop screen reader inspired by the NVDA philosophy, prioritizing user control, reliability, configurability, and sustainable architecture.

Product Philosophy

Aura puts the user firmly in control. We reject the paradigm of mandatory workflows or forced simplifications. The system must be highly configurable, allowing users to mold the screen reader's behavior to their specific mental models and workflows.

While we may selectively draw inspiration from other platforms (like VoiceOver) for specific UX improvements, these features will only be implemented if they enhance the experience without reducing explicit user control.

Aura values predictable behavior over opaque automation or excessive heuristics.

Aura treats stability, reliability, and predictability as core accessibility requirements, not optional quality attributes.

The system must behave consistently under real-world desktop usage. Feature expansion, experimentation, or UX sophistication must never compromise operational trustworthiness.

MVP Scope: Desktop Focus Reader

The initial Minimal Viable Product (MVP) is explicitly scoped to act as a robust desktop focus reader. It is not an immediate, feature-complete replacement for existing tools, but rather a rock-solid foundation.

The MVP is not intended to be a full NVDA replacement, nor to immediately compete on feature breadth with mature screen readers.

Target Environments

Modern Linux desktop environments (GNOME, KDE).

Supported Use Cases (MVP)

Focus navigation via keyboard (desktop usage, Alt+Tab).

Basic interaction with GTK and Qt (KDE/GNOME) menus, dialogs, buttons, lists, and basic text fields.

Reading web content within Firefox and Chromium (relying on standard accessibility tree exposure).

Core Features (MVP)

Reliable speech output via Speech Dispatcher.

Robust focus tracking and event normalization.

Announcement of UI Role, Name, and State.

Basic verbosity configuration.

Minimal flat review capabilities.

A basic, customizable keymap.

Deferred Features (Post-MVP)

Browse mode and virtual buffers.

Scripting and plugin ecosystems.

Optical Character Recognition (OCR).

Braille display support.

Deep terminal integration and advanced console workflows.

AI-driven vision features.

Advanced overlays.