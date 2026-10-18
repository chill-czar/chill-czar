# Sarthak this side

> Building AI agent runtimes, microVM sandboxes, and high-throughput distributed systems.

Platform engineer focused on systems engineering, virtualization, and real-time backend infrastructure. Active open-source contributor to Google ADK-Go, E2B, and SuperServe.

---

## What I've shipped

**Devin — Autonomous AI Software Engineering Sandbox**  
An isolated code execution runtime and orchestrator. Built around custom Firecracker microVM snapshotting, CNI networking, and a Go-based host-guest daemon (`boxd`) communicating over `AF_VSOCK` without exposed guest ports.  
[GitHub](https://github.com/chill-czar/devin)

**ApexPerp — Perpetual Futures Exchange Engine**  
A high-throughput crypto perpetual exchange featuring an in-memory matching engine with price-time priority, Redis Streams/Pub-Sub pipeline, TimescaleDB continuous candlestick aggregation, and WebSocket order book fanout.  
[Live](https://apex-perp.vercel.app/) · [GitHub](https://github.com/chill-czar/apex-perp)

**Collabify — Distributed Real-Time Collaboration Platform**  
A collaborative workspace with real-time notes and whiteboard canvas. Features Yjs CRDT delta synchronization over WebSockets, a transactional outbox + Redis Streams chat pipeline, and a TOCTOU-safe AWS S3 vault with transactional quota reservation.  
[Server](https://github.com/official-collabify/collabify-server) · [Client](https://github.com/official-collabify/client)

**Laura Voice — Real-Time Voice AI Intake Orchestrator**  
An end-to-end voice AI consultation intake pipeline using LiveKit (WebRTC), Deepgram (Nova STT & Aura TTS), and an OpenRouter LLM state machine. Conducts structured caller qualification with live Server-Sent Events (SSE) telemetry streaming to the frontend.  
[GitHub](https://github.com/chill-czar/laura-voice)

---

## Open Source Contributions

- **[Google ADK-Go](https://github.com/google/adk-go)** (5 Merged PRs): Fixed multi-part artifact loading ([#1320](https://github.com/google/adk-go/pull/1320)), prevented socket and goroutine leaks on context cancellation ([#1351](https://github.com/google/adk-go/pull/1351)), deferred telemetry shutdown on server errors ([#1470](https://github.com/google/adk-go/pull/1470)), and added memory store nil guards ([#1349](https://github.com/google/adk-go/pull/1349), [#1353](https://github.com/google/adk-go/pull/1353)).
- **[E2B Runtime](https://github.com/e2b-dev/runtime)**: Implemented non-positive timeout validation across sandbox create, resume, and fork APIs ([#3419](https://github.com/e2b-dev/runtime/pull/3419)).
- **[SuperServe](https://github.com/superserve-ai/sandbox)**: Resolved quota alert serialization failure to guarantee notification delivery ([#324](https://github.com/superserve-ai/sandbox/pull/324)).

---

## Where I've worked

- **Cartigram** — Full-Stack Developer (Dec 2025 – Present)  
  Engineered multi-source authorization across 8 ERP modules, built a dynamic approval workflow engine, and implemented CRM lead ingestion pipelines.
- **Sheryians Coding School** — Full-Stack Developer Intern (Jul 2025 – Oct 2025)  
  Built candidate-management APIs, real-time admin dashboards in Next.js, and an automated LangChain-based resume assessment pipeline.

---

## Stack

**Languages:** Go, TypeScript, JavaScript, Python, SQL, Bash  
**Systems & Infrastructure:** Linux, Firecracker MicroVMs, KVM, Docker, Kubernetes, Terraform, AWS (EC2, S3, SQS)  
**Backend & Realtime:** REST APIs, GraphQL, WebSockets, gRPC, PostgreSQL, TimescaleDB, Redis Streams, BullMQ, LiveKit  
**Frontend:** React 19, Next.js, TanStack Router/Query, Tailwind CSS, Vite, Bun  

---

## Find me

GitHub → [github.com/chill-czar](https://github.com/chill-czar)  
LinkedIn → [linkedin.com/in/its-sarthak-chaudhari](https://www.linkedin.com/in/its-sarthak-chaudhari/)  
X / Twitter → [x.com/chill_czar](https://x.com/chill_czar)  
Email → [aisarthak2@gmail.com](mailto:aisarthak2@gmail.com)
