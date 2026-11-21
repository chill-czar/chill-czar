# Sarthak Chaudhari

> **Systems & Agent Infrastructure Engineer** · Low-Level Runtimes, Linux Internals & Distributed Systems

Building isolated agent runtimes, microVM sandboxes, and high-throughput real-time backends. Core contributor to cloud-native virtualization, developer tooling, and agent SDK ecosystems.

[Portfolio](https://www.itsmesarthak.site/) · [GitHub](https://github.com/chill-czar) · [LinkedIn](https://www.linkedin.com/in/its-sarthak-chaudhari/) · [X / Twitter](https://x.com/chill_czar) · [Email](mailto:aisarthak2@gmail.com)

---

## 📌 Pinned Open Source Contributions

### 🟢 Active Upstream Contributions

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

### 🚀 Core Infrastructure & Agent Ecosystem Contributions

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
