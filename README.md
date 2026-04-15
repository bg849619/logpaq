# LogPAQ

A distributed contact logger with weak-consistency. Logging nodes reach eventual consistency with the rest of the nodes, but are capable of operating independently. No single node is central to a log.

## Design Constraints

- No significant network configuration should be needed. If nodes are in the same Layer 2 broadcast domain, the nodes should be able to connect.
- Nodes should be able to operate independently against the most recently known log without significant merge conflict.

## AI Usage

Claude and GitHub Copilot were used in the architecture/design of this application, and often used to create boilerplate for the implementation of that architecture. It is the intention of the developers to understand and verify the LLM output used. Output is continually refined as the models are consulted regarding their initial solution, until a high quality implementation is achieved.

Agent use has been scope-limited to smaller software components to increase the likelihood of a high-quality solution and to make it easier for developers to validate output within the larger architecture.