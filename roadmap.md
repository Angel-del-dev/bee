# Bee Capability Roadmap

This roadmap defines progressively harder tasks that Bee should be able to complete using different local models.

The goal is not to benchmark raw model intelligence, but to measure practical software engineering capabilities inside a real repository.

---

# Level 1 — Repository Awareness

Basic understanding of a codebase.

## Configuration Awareness

Task:
- Read and obey a `.bee` configuration file.
- Adapt generated code to project-specific rules.

Success Criteria:
- Generated code follows all configured conventions.
- No violation of explicitly defined project rules.

---

## TODO Extraction

Task:
- Scan the repository.
- Find every TODO, FIXME, HACK and NOTE comment.
- Generate a centralized TODO.md.

Success Criteria:
- No TODO is missed.
- Entries are grouped by file.

---

## Repository Summary

Task:
- Analyze an unknown repository.
- Generate a concise overview.

Success Criteria:
- Correctly identifies purpose.
- Lists main technologies.
- Explains project structure.

---

# Level 2 — Documentation

Repository understanding and communication.

## Full Documentation Generation

Task:
- Generate complete project documentation.

Success Criteria:
- Installation guide.
- Usage examples.
- Architecture overview.
- API documentation.

---

## Code Explanation

Task:
- Explain any selected file.

Success Criteria:
- Correctly describes purpose.
- Explains important algorithms.
- Identifies dependencies.

---

# Level 3 — Development Tasks

Perform real engineering work.

## Feature Implementation

Task:
- Implement a feature described in natural language.

Success Criteria:
- Builds successfully.
- Follows project conventions.
- Existing tests continue passing.

---

## Bug Fixing

Task:
- Fix a reported issue.

Success Criteria:
- Root cause identified.
- Fix implemented.
- No regression introduced.

---

## Refactoring

Task:
- Improve code quality without changing behavior.

Success Criteria:
- Functionality preserved.
- Complexity reduced.
- Readability improved.

---

# Level 4 — Environment Interaction

Operate beyond source code.

## Controlled Command Execution

Task:
- Execute shell commands with user approval.

Success Criteria:
- Never executes without permission.
- Correctly interprets command output.
- Can continue workflows using results.

---

## Build Recovery

Task:
- Fix build failures.

Success Criteria:
- Detects failing step.
- Applies correction.
- Build succeeds.

---

## Dependency Management

Task:
- Upgrade dependencies safely.

Success Criteria:
- Project still builds.
- Breaking changes documented.

---

# Level 5 — Project Creation

Create complete artifacts.

## Landing Page Generation

Task:
- Generate a complete landing page for the current project.

Success Criteria:
- Modern design.
- Installation instructions.
- Feature showcase.
- Responsive layout.

---

## Project Bootstrap

Task:
- Create a new project from a high-level description.

Success Criteria:
- Functional structure.
- Working build system.
- Documentation included.

---

## Example Generation

Task:
- Generate practical examples for users.

Success Criteria:
- Executable examples.
- Demonstrates key features.

---

# Level 6 — Repository Ownership

Act as a maintainer.

## Technical Debt Audit

Task:
- Analyze repository quality.

Success Criteria:
- Detects risks.
- Prioritizes issues.
- Produces actionable report.

---

## Architecture Review

Task:
- Review overall architecture.

Success Criteria:
- Identifies weaknesses.
- Suggests improvements.
- Justifies recommendations.

---

## Continuous Documentation Maintenance

Task:
- Keep documentation synchronized with source code.

Success Criteria:
- Detects outdated docs.
- Updates affected sections.

---

# Level 7 — Language Engineering

Advanced software engineering tasks.

## DSL Creation

Task:
- Design a domain-specific language.

Success Criteria:
- Grammar definition.
- Parser implementation.
- Documentation.

---

## DSL Maintenance

Task:
- Evolve an existing DSL.

Success Criteria:
- Backward compatibility.
- Migration guides.
- Updated tooling.

---

## Compiler Development

Task:
- Create a compiler or interpreter.

Success Criteria:
- Valid parsing.
- Correct execution.
- Error reporting.

---

# Level 8 — Autonomous Engineering

Maximum capability benchmark.

## Large Feature Delivery

Task:
- Implement a feature requiring modifications across multiple modules.

Success Criteria:
- Correct planning.
- Correct implementation.
- Tests updated.
- Documentation updated.

---

## Repository Modernization

Task:
- Upgrade a legacy project to modern standards.

Success Criteria:
- Existing functionality preserved.
- Modern tooling adopted.
- Migration documented.

---

## Independent Project Evolution

Task:
- Propose, implement and document a meaningful improvement.

Success Criteria:
- Improvement is useful.
- Implementation is correct.
- Changes are justified.