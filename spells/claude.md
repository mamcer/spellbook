# Anatomy of the .claude/ Folder

Summary based on the article by [Avi Chawla](https://blog.dailydoseofds.com/p/anatomy-of-the-claude-folder).

## Overview
The `.claude/` folder acts as the **control center** for Claude Code, defining instructions, custom commands, permission rules, and session memory.

---

## 1. Dual Locations
*   **Project-level (`your-project/.claude/`)**: Holds team-wide configurations. Should be committed to Git so every team member shares the same rules and commands.
*   **Global (`~/.claude/`)**: Stores personal preferences, machine-local state, session history, and auto-memory.

---

## 2. Core Files & Structure

### CLAUDE.md (The Brain)
*   The most critical file; loaded directly into the system prompt.
*   **Recommendations**:
    *   Include build/test/lint commands, architectural decisions, and naming patterns.
    *   Keep it under **200 lines** to maintain instruction adherence.
*   **CLAUDE.local.md**: Personal overrides for a specific project (gitignored).

### .claude/rules/ (Modular Instructions)
*   Splits a giant `CLAUDE.md` into focused files (e.g., `testing.md`, `api-conventions.md`).
*   **Path-scoping**: Use YAML frontmatter to activate rules only for specific directories (e.g., `src/api/`).

### .claude/commands/ (Custom Slash Commands)
*   Each `.md` file becomes a `/project:name` command.
*   Supports `$ARGUMENTS` and shell command execution using the `!` backtick syntax.

### .claude/skills/ (Auto-invoked Workflows)
*   Unlike commands, skills are triggered **automatically** by Claude when a task matches the skill's description (defined in YAML frontmatter).
*   Skills are "packages" that can bundle supporting files.

### .claude/agents/ (Subagent Personas)
*   Defines specialized specialists with isolated context windows.
*   Reduces context pollution in the main session.
*   Allows restricting tool access (e.g., read-only) and selecting specific models (e.g., using Haiku for cheap exploration).

### settings.json (Permissions)
*   **Allow list**: Commands that run without asking for confirmation (e.g., `npm run`, `git` read-only).
*   **Deny list**: Explicitly blocked commands or files (e.g., `rm -rf`, `.env`).
*   **Default**: If not listed, Claude asks for permission.

---

## 3. Global Directory Anatomy (`~/.claude/`)
*   `CLAUDE.md`: Global coding principles and preferred styles across all projects.
*   `projects/`: Stores session transcripts and **auto-memory** (architecture insights, observed patterns).
*   `commands/` & `skills/`: Personal tools available in every repository (prefix: `/user:`).

---

## 4. Suggested Implementation Path
1.  **Run `/init`**: Generate a starter `CLAUDE.md`.
2.  **Configure `settings.json`**: Define allow/deny rules for your stack.
3.  **Create Commands**: Automate frequent workflows (e.g., code review).
4.  **Modularize Rules**: As the project grows, split `CLAUDE.md` into `rules/`.
5.  **Global Prefs**: Set personal defaults in `~/.claude/CLAUDE.md`.
