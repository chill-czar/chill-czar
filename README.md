# Sarthak Chaudhari

> **Systems & Agent Infrastructure Engineer** · Low-Level Runtimes, Linux Internals & Distributed Systems

Building isolated agent runtimes, microVM sandboxes, and high-throughput real-time backends. Core contributor to cloud-native virtualization, developer tooling, and agent SDK ecosystems.

[Portfolio](https://www.itsmesarthak.site/) · [GitHub](https://github.com/chill-czar) · [LinkedIn](https://www.linkedin.com/in/its-sarthak-chaudhari/) · [X / Twitter](https://x.com/chill_czar) · [Email](mailto:aisarthak2@gmail.com)

---

## 📌 Open Source Contributions

> Repositories where I actively contribute and maintain upstream infrastructure, ordered by active engagements.

### 🟢 Active Upstream Engagements

* **[Lima VM](https://github.com/lima-vm/lima)** — *Linux virtual machines on macOS/Linux with container-first runtimes (`containerd`, `nerdctl`, Kubernetes)*
  * **[#5837](https://github.com/lima-vm/lima/pull/5837)** *(In Review)* — `limactl shell`: Reject `--workdir` and `--preserve-env` flags on Windows guests to eliminate silent execution discrepancies.
  * **[#5834](https://github.com/lima-vm/lima/pull/5834)** *(In Review)* — `pkg/hostagent`: Suppress spurious connection warnings and stream noise on clean teardown and expected socket EOF.

* **[KubeVirt](https://github.com/kubevirt/kubevirt)** — *Kubernetes Virtualization API and runtime to declare and run virtual machines on K8s*
  * **[#19982](https://github.com/kubevirt/kubevirt/pull/19982)** *(In Review)* — `client-go/log`: Migrate legacy, unmaintained `go-kit/log` to modern `logr` and `zap` structured loggers across client packages.
  * **[#19971](https://github.com/kubevirt/kubevirt/pull/19971)** *(In Review)* — `tls`: Enforce strict Common Name verification when TLS certificates are externally managed.
  * **[#19979](https://github.com/kubevirt/kubevirt/pull/19979)** — `pkg/hypervisor`: Modernize codebase idioms and enroll hypervisor packages into the `go fix` lint gate.

* **[Argo CD](https://github.com/argoproj/argo-cd)** — *Declarative, GitOps continuous delivery engine for Kubernetes*
  * Contributed upstream fixes and UX enhancements across resource filter shorthands, pod logs viewer controls, and rollback-aware sync safety.

---

### 🚀 Merged Infrastructure & Agent Ecosystem Contributions

* **[Google ADK-Go](https://github.com/google/adk-go)** — *Google's official Go Agent Development Kit (5 Merged PRs)*
  * **[#1351](https://github.com/google/adk-go/pull/1351)** *(Merged)* — `cmd/launcher/web`: Initialize telemetry prior to HTTP listener startup, preventing goroutine and socket leaks on context cancellation.
  * **[#1320](https://github.com/google/adk-go/pull/1320)** *(Merged)* — `loadartifactstool`: Support `load_artifacts` when invoked alongside concurrent multi-tool calls.
  * **[#1470](https://github.com/google/adk-go/pull/1470)** *(Merged)* — `cmd/launcher/web`: Defer telemetry shutdown on launcher startup failure to preserve diagnostic traces.
  * **[#1349](https://github.com/google/adk-go/pull/1349)** & **[#1353](https://github.com/google/adk-go/pull/1353)** *(Merged)* — Added memory store nil guards on part extraction and enabled valid empty-text MCP tool responses.

* **[Tracer Cloud / OpenSRE](https://github.com/Tracer-Cloud/opensre)** — *Autonomous open-source SRE agent for CI/CD reliability (5 Merged PRs, 15 Submitted)*
  * **[#7028](https://github.com/Tracer-Cloud/opensre/pull/7028)** *(Merged)* — Guardrails: Validate positive limit bounds in audit commands and return empty result sets on non-positive limits.
  * **[#7025](https://github.com/Tracer-Cloud/opensre/pull/7025)** *(Merged)* — Turn-Host: Detect overlapping non-reentrant session leases as busy to prevent race conditions during concurrent turns.
  * **[#7018](https://github.com/Tracer-Cloud/opensre/pull/7018)** *(Merged)* — CLI: Validate positive rate boundaries in health watcher streams.
  * **[#7001](https://github.com/Tracer-Cloud/opensre/pull/7001)** *(Merged)* — Cron: Honor global `--json` emission flags in background cron task logging.
  * **[#6997](https://github.com/Tracer-Cloud/opensre/pull/6997)** *(Merged)* — Work Scheduler: Report already-completed items and gracefully handle missing selectors in `work complete`.

* **[E2B Runtime](https://github.com/e2b-dev/runtime)** — *Secure cloud sandboxes and microVM execution runtime for AI agents*
  * **[#3419](https://github.com/e2b-dev/runtime/pull/3419)** *(Merged)* — API: Reject non-positive timeout inputs across sandbox `create`, `resume`, and `fork` endpoints.

* **[SuperServe](https://github.com/superserve-ai/superserve)** — *Cloud sandboxes and execution infrastructure for AI agents*
  * **[#332](https://github.com/superserve-ai/superserve/pull/332)** *(Active)* — Added Superserve sandbox execution adapter for the OpenAI Agents SDK.
  * **[#20](https://github.com/superserve-ai/canaries/pull/20)** *(Merged)* — Deployed UI lifecycle canaries across staging and production environments.

---

## 🛠️ Featured Systems & Projects

### 1. [Devin](https://github.com/chill-czar/devin) — Autonomous AI Engineering Sandbox
*An isolated microVM code-execution sandbox and orchestrator runtime for autonomous coding agents.*
- Built a **Go-based Firecracker runtime** configuring ext4 rootfs, guest kernel, vCPU/RAM quotas, virtio-net, and Unix sockets on bare metal.
- Implemented an in-guest daemon (`boxd`) communicating with the host over a custom `AF_VSOCK` HTTP transport for commands, code search, and streaming without exposing guest ports.
- Engineered TAP/iptables networking with MASQUERADE NAT, outbound HTTP CONNECT proxy, process-group termination, and orphan VM sweepers.

### 2. [Collabify](https://collabifyy.com) — Distributed Real-Time Collaboration Workspace
*A high-concurrency collaborative workspace featuring real-time notes and interactive canvas.*
- Built collaborative notes and whiteboard canvas with **Yjs CRDTs** over WebSockets and Redis Pub/Sub for deterministic state convergence.
- Implemented transactional outbox + **Redis Streams** chat pipeline with atomic Lua `INCR+XADD` sequencing and duplicate-write rejection.
- Engineered a TOCTOU-safe **AWS S3** vault with transactional quota reservations preventing concurrent upload races.  
*[Live Demo](https://collabifyy.com)* · *[Backend](https://github.com/official-collabify/collabify-server)* · *[Client](https://github.com/official-collabify/client)*

### 3. [Laura Voice](https://github.com/chill-czar/laura-voice) — Real-Time Voice AI Intake Orchestrator
*An end-to-end conversational voice AI intake system for high-stakes client triage and consultation.*
- Cascaded voice pipeline: **LiveKit (WebRTC)** audio transport → **Deepgram Nova (STT)** → **OpenRouter LLMs** → **Deepgram Aura (TTS)** with sub-second turn latency.
- Deterministic 4-agent state machine (**Triage, Intake, Qualification, Scheduling**) where code validates phase boundaries and LLMs execute reasoning.
- Real-time observability: Server-Sent Events (SSE) stream off a Redis event bus into a Next.js live telemetry console.

### 4. [ApexPerp](https://github.com/chill-czar/apex-perp) — Perpetual Futures Trading Engine
*A low-latency crypto perpetual futures exchange engine built from scratch in Go.*
- Hand-written in-memory orderbook matching engine with price-time priority.
- Event-driven pipeline using **Redis Streams** for durable write ordering and **WebSockets** for real-time depth/trade fanout.
- Automated PnL and liquidation engine fed by live Binance mark price feeds; 1m/1h/1d continuous aggregates in **TimescaleDB**.

---

## 💼 Experience

* **Cartigram** — Full-Stack Developer *(Dec 2025 – Present)*  
  Engineered multi-source authorization across 8 ERP modules, built a dynamic approval workflow engine, and implemented CRM lead ingestion pipelines.
* **Sheryians Coding School** — Full-Stack Developer Intern *(Jul 2025 – Oct 2025)*  
  Built candidate-management APIs, real-time admin dashboards in Next.js, and an automated LangChain-based resume assessment pipeline.

---

## ⚡ Technical Stack

* **Systems, Virtualization & Runtimes:** Go, Linux Internals, KVM, Firecracker MicroVMs, `AF_VSOCK`, cgroups v2, Docker, Containerd, Kubernetes, Bash
* **Agent Infrastructure & Voice AI:** Google ADK-Go, OpenAI Agents SDK, E2B, LiveKit (WebRTC), Deepgram (STT/TTS), LangChain, OpenRouter
* **Distributed Systems & Storage:** Redis Streams, Pub/Sub, PostgreSQL, TimescaleDB, gRPC, WebSockets, GraphQL Yoga, Hono, BullMQ, Prisma
* **Cloud & DevOps:** AWS (EC2, S3, SQS), Linux Networking (TAP, iptables, NAT), GitHub Actions CI/CD, Terraform
* **Frontend & Real-time Web:** React 19, Next.js, TanStack Router/Query, Tailwind CSS, Yjs CRDTs, Vite, TypeScript

---

## 📫 Find Me

Portfolio → [itsmesarthak.site](https://www.itsmesarthak.site/)  
GitHub → [github.com/chill-czar](https://github.com/chill-czar)  
LinkedIn → [linkedin.com/in/its-sarthak-chaudhari](https://www.linkedin.com/in/its-sarthak-chaudhari/)  
X / Twitter → [x.com/chill_czar](https://x.com/chill_czar)  
Email → [aisarthak2@gmail.com](mailto:aisarthak2@gmail.com)
