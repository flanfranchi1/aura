# Aura Roadmap: The Path to a Full Screen Reader

This document outlines the strategic milestones for Aura, following the principle: 
"Finishing is more important than starting."

## Phase 1: Foundation (Current) ✅
- [x] Initial AT-SPI2 bus discovery and connection.
- [x] Mapping of basic Roles (Terminal, Window, Button, etc.).
- [x] Non-blocking integration with Speech Dispatcher via SSIP.
- [x] Focus tracking and basic naming logic.

## Phase 2: Semantics & Localization (Next Steps) 🛠️
- [ ] **Internationalization (i18n):** Create a translation layer for Roles and States (PT-BR).
- [ ] **State Mapping:** Expand the engine to track object states (checked, expanded, busy).
- [ ] **Pronunciation Dictionary:** User-definable word corrections.

## Phase 3: Input & Control ⌨️
- [ ] **Global Hotkeys:** Implementation of a "Aura Key" (e.g., Caps Lock/Insert).
- [ ] **Speech Interruption:** Immediate speech cancellation on key press (Ctrl).
- [ ] **Navigation Modes:** Toggle between "Focus Mode" and "Object Navigation Mode".

## Phase 4: Advanced Interrogation 🔍
- [ ] **Text Interface:** Ability to read characters, words, and lines in edit fields.
- [ ] **Review Cursor:** Inspect screen elements without changing system focus.
- [ ] **Earcons:** Audio cues for system events (window open/close, errors).

## Phase 5: UI & UX ⚙️
- [ ] **Configuration Manager:** Persistence of user settings (speed, pitch, voice).
- [ ] **Minimalist UI:** Terminal-based or lightweight GUI for setup.

## Phase 6: Web & Documents 🌐
- [ ] **Virtual Buffer:** Support for complex document browsing (HTML/PDF).