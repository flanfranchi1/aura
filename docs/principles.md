Project Aura: Core Principles

These principles guide all technical, architectural, and product decisions for Project Aura. When evaluating a new feature or technical approach, it must be weighed against these axioms.
1. User Control First

The system works for the user, not the other way around. Defaults should be sensible, but behavior must be configurable. Avoid hidden magic that the user cannot disable or tweak.

    Reliability Before Sophistication

Reliability over feature count. A screen reader that correctly and predictably announces focus 100% of the time is infinitely more valuable than one with an advanced browse mode that crashes or lags. We will master the basics before expanding scope.
3. Backend-Agnostic Core

Business and accessibility domain logic must remain independent from the underlying OS accessibility API. The core engine should not know if it is talking to AT-SPI, Windows UIAutomation, or a Wayland-native API.
4. AT-SPI First Implementation

While the core is agnostic, the initial provider implementation will target AT-SPI over DBus. This is the current standard for Linux desktop accessibility and provides the fastest path to MVP.
5. Desktop-First MVP

We prioritize the standard desktop computing experience (web browsers, file managers, system settings) over specialized environments like advanced terminal emulators or embedded devices for the initial releases.
Decisions should balance implementation effort, operational complexity, maintainability, and expected benefits. Simpler solutions should be preferred unless additional complexity provides clear, justified value.
7. Sustainable AI-First Development

We recognize that "finishing is more important than starting." The architecture must remain clean, idiomatic, and scannable so that human developers and AI assistants can collaborate efficiently. We reject over-engineering, premature abstractions, and complex design patterns (like CQRS or heavy DI frameworks) that hinder rapid, maintainable progress.
8. Observable Systems

Internal behavior should be inspectable and diagnosable. Event streams, focus transitions, and provider interactions must be debuggable without requiring invasive instrumentation.