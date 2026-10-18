# Sarthak this side

> Applied AI Full-Stack Engineer · Agent Infrastructure, Voice AI & Distributed Systems

Building AI agent runtimes, WebRTC voice pipelines, and high-throughput real-time backends. Active open-source contributor to **Google ADK-Go**, **E2B**, and **SuperServe**.

[Portfolio](https://www.itsmesarthak.site/) · [GitHub](https://github.com/chill-czar) · [LinkedIn](https://www.linkedin.com/in/its-sarthak-chaudhari/) · [X / Twitter](https://x.com/chill_czar) · [Email](mailto:aisarthak2@gmail.com)

---

## What I've shipped

### 1. Collabify — Distributed Real-Time Collaboration Workspace
A high-concurrency collaborative workspace featuring real-time notes and interactive canvas.
- Built collaborative notes and whiteboard canvas with **Yjs CRDTs** over WebSockets and Redis Pub/Sub for deterministic state convergence.
- Implemented transactional outbox + **Redis Streams** chat pipeline with atomic Lua `INCR+XADD` sequencing and duplicate-write rejection.
- Engineered a TOCTOU-safe **AWS S3** vault with transactional quota reservations preventing concurrent upload races.  
**[Live Demo](https://collabifyy.com)** · **[Backend](https://github.com/official-collabify/collabify-server)** · **[Client](https://github.com/official-collabify/client)**

---

### 2. Devin — Autonomous AI Software Engineering Sandbox
An isolated microVM code-execution sandbox and orchestrator runtime for autonomous coding agents.
- Engineered a **Go-based Firecracker runtime** configuring ext4 rootfs, kernel, vCPU/RAM, virtio-net, and Unix sockets on bare metal.
- Implemented an in-guest daemon (`boxd`) communicating with the host over a custom `AF_VSOCK` HTTP transport for commands, code search, and streaming without exposed guest ports.
- Built TAP/iptables networking with MASQUERADE NAT, outbound HTTP CONNECT proxy, process-group termination, and orphan VM sweeping.  
**[GitHub](https://github.com/chill-czar/devin)**

---

### 3. Laura Voice — Real-Time Voice AI Intake Orchestrator
An end-to-end conversational voice AI intake system for high-stakes client triage and consultation.
- Cascaded voice pipeline: **LiveKit (WebRTC)** audio transport → **Deepgram Nova (STT)** → **OpenRouter LLMs** → **Deepgram Aura (TTS)** with sub-second turn latency.
- Deterministic 4-agent state machine (**Triage, Intake, Qualification, Scheduling**) where code validates phase boundaries and LLMs execute reasoning.
- Real-time observability: Server-Sent Events (SSE) stream off a Redis event bus into a Next.js live telemetry console.  
**[GitHub](https://github.com/chill-czar/laura-voice)**

---

### 4. ApexPerp — Perpetual Futures Trading Engine
A low-latency crypto perpetual futures exchange engine built from scratch.
- Hand-written in-memory orderbook matching engine with price-time priority.
- Event-driven pipeline using **Redis Streams** for durable write ordering and **WebSockets** for real-time depth/trade fanout.
- Automated PnL and liquidation engine fed by live Binance mark price feeds; 1m/1h/1d continuous aggregates in **TimescaleDB**.  
**[GitHub](https://github.com/chill-czar/apex-perp)**

---

## Open Source Contributions

* **[Google ADK-Go](https://github.com/google/adk-go)** *(5 Merged PRs)*:
  * Handled multi-part artifact loading when combined with multi-tool calls ([#1320](https://github.com/google/adk-go/pull/1320))
  * Prevented socket & goroutine leaks on context cancellation ([#1351](https://github.com/google/adk-go/pull/1351))
  * Deferred telemetry shutdown on launcher server failure ([#1470](https://github.com/google/adk-go/pull/1470))
  * Added memory store nil guards and allowed valid empty-text MCP tool responses ([#1349](https://github.com/google/adk-go/pull/1349), [#1353](https://github.com/google/adk-go/pull/1353))
* **[E2B Runtime](https://github.com/e2b-dev/runtime)** *(Merged PR)*:
  * Implemented non-positive timeout validation across sandbox `create`, `resume`, and `fork` APIs ([#3419](https://github.com/e2b-dev/runtime/pull/3419))
* **[SuperServe](https://github.com/superserve-ai/sandbox)** *(Merged PR)*:
  * Resolved quota alert serialization failure to guarantee Slack/email notifications ([#324](https://github.com/superserve-ai/sandbox/pull/324))

---

## Where I've worked

* **Cartigram** — Full-Stack Developer *(Dec 2025 – Present)*  
  Engineered multi-source authorization across 8 ERP modules, built a dynamic approval workflow engine, and implemented CRM lead ingestion pipelines.
* **Sheryians Coding School** — Full-Stack Developer Intern *(Jul 2025 – Oct 2025)*  
  Built candidate-management APIs, real-time admin dashboards in Next.js, and an automated LangChain-based resume assessment pipeline.

---

## Stack

* **Languages:** Go, TypeScript, JavaScript, Python, SQL, Bash
* **Applied AI & Agent Infra:** Google ADK, LiveKit (WebRTC), Deepgram (STT/TTS), Firecracker MicroVMs, Agent Sandboxing, LangChain, OpenRouter
* **Backend & Distributed Systems:** GraphQL Yoga, REST, WebSockets, gRPC, Hono, PostgreSQL, TimescaleDB, Redis Streams, BullMQ, Prisma
* **Cloud & DevOps:** Linux, Docker, KVM, Kubernetes, Terraform, AWS (EC2, S3, SQS), CI/CD
* **Frontend:** React 19, Next.js, TanStack Router/Query, Tailwind CSS, Yjs CRDTs, Vite, Bun

---

## Find me

Portfolio → [itsmesarthak.site](https://www.itsmesarthak.site/)  
GitHub → [github.com/chill-czar](https://github.com/chill-czar)  
LinkedIn → [linkedin.com/in/its-sarthak-chaudhari](https://www.linkedin.com/in/its-sarthak-chaudhari/)  
X / Twitter → [x.com/chill_czar](https://x.com/chill_czar)  
Email → [aisarthak2@gmail.com](mailto:aisarthak2@gmail.com)
