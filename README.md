# LogPAQ

A distributed contact logger with weak-consistency. Logging nodes reach eventual consistency with the rest of the nodes, but are capable of operating independently. No single node is central to a log.

## Design Constraints

- No significant network configuration should be needed. If nodes are in the same Layer 2 broadcast domain, the nodes should be able to connect.
- Nodes should be able to operate independently against the most recently known log without significant merge conflict.

## AI Usage

- Claude and GitHub Copilot were used in the architecture/design of the logging engine. However, they did not generate any code *directly*. LLMs were treated as consultants to this software, and their outputs were refined further given design constraints and edge cases.