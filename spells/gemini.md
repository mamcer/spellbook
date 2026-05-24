# gemini

Esta guía detalla la estructura y el propósito del ecosistema de configuración de Gemini CLI, el cual permite personalizar el comportamiento, la seguridad y las habilidades de la IA en tus proyectos.

## 1. Niveles de Configuración

Gemini sigue una jerarquía de carga para combinar preferencias globales con reglas específicas del proyecto:

1.  **Nivel Global (`~/.gemini/`)**: Define tu configuración base, modelos preferidos y `GEMINI.md` global que se aplica a todos los repositorios.
2.  **Nivel de Proyecto (`./.gemini/` y `GEMINI.md`)**: Configuración específica del repositorio, compartida con el equipo a través de Git.

---

## 2. Estructura de Archivos y Carpetas

Jerarquía recomendada para un proyecto optimizado con Gemini:

```text
tu-proyecto/
├── GEMINI.md                # Instrucciones principales (Equivalente a CLAUDE.md)
└── .gemini/
    ├── settings.json        # Configuración del proyecto (modelos, límites, UI)
    ├── system.md            # (Opcional) Sobrescritura del System Prompt base
    ├── commands/            # Comandos slash personalizados (/nombre) en formato TOML
    │   └── test.toml
    ├── policies/            # Reglas de seguridad y ejecución de herramientas (TOML)
    │   └── security.toml
    └── skills/              # Paquetes de "habilidades" activables bajo demanda
        └── sql-expert/
            └── SKILL.md
```

---

## 3. Componentes Clave

### GEMINI.md (El Corazón del Contexto)
Es el archivo que Gemini lee primero. Proporciona las reglas de juego del proyecto.
- **Tip:** Puedes usar `@archivo.md` dentro de `GEMINI.md` para importar otros archivos de documentación y mantener este archivo limpio.
- **Ejemplo Práctico:**
  ```markdown
  # Proyecto: Nostalgia Go
  
  ## Arquitectura
  - Basado en Clean Architecture (Internal/Domain, Internal/App).
  - Framework: Gin-gonic.
  
  ## Convenciones
  - Los handlers de Gin deben validar el input con `ShouldBindJSON`.
  - Usar interfaces para inyección de dependencias.
  ```

### .gemini/settings.json (Ajustes de Motor)
Define qué modelo usar y cómo deben comportarse las herramientas.
- **Ejemplo:**
  ```json
  {
    "model": {
      "name": "gemini-2.0-flash",
      "temperature": 0.2
    },
    "tools": {
      "maxOutputLines": 500
    }
  }
  ```

### .gemini/policies/ (Control de Herramientas)
A diferencia de otros agentes, Gemini usa un motor de políticas TOML para decidir qué puede hacer la IA sin preguntar.
- **Ejemplo de Política:**
  ```toml
  [[rules]]
  toolName = "run_shell_command"
  pattern = "go test .*"
  decision = "allow"
  priority = 10

  [[rules]]
  toolName = "read_file"
  pattern = ".env"
  decision = "deny"
  priority = 100
  ```

### .gemini/commands/ (Comandos Slash)
Se definen en archivos TOML y permiten automatizar prompts complejos.
- **Ejemplo (`/lint.toml`):**
  ```toml
  description = "Corre el linter y pide a Gemini que sugiera correcciones"
  prompt = """
  Ejecuta el linter del proyecto:
  !{golangci-lint run}
  Basado en los errores anteriores, propón un plan de corrección.
  """
  ```

---

## 4. Comparativa: Gemini vs Claude

| Característica | Claude Code | Gemini CLI |
| :--- | :--- | :--- |
| **Archivo Contexto** | `CLAUDE.md` | `GEMINI.md` |
| **Carpeta Config** | `.claude/` | `.gemini/` |
| **Formato Comandos**| Markdown (`.md`) | TOML (`.toml`) |
| **Seguridad** | `settings.json` (allow/deny) | Policy Engine (TOML con Regex) |
| **Habilidades** | Skills (Carpetas) | Agent Skills (`activate_skill`) |

---

## 5. Tips y Trucos para Gemini

1.  **JIT (Just-In-Time) Context**: Si tienes un `GEMINI.md` dentro de una subcarpeta (ej: `internal/db/GEMINI.md`), Gemini solo cargará esas instrucciones cuando estés trabajando en esa carpeta. ¡Úsalo para proyectos grandes!
2.  **Variables en Comandos**: Puedes usar `#{selection}` o `#{file}` en tus comandos personalizados para pasar el código seleccionado o el archivo actual.
3.  **Activación de Skills**: No satures la memoria. Usa `skills/` para conocimientos que Gemini no necesita todo el tiempo (ej: una guía de migración compleja) y actívalos solo cuando sea necesario con `activate_skill`.

---
## Further Reading
- `/help` dentro de Gemini CLI para ver la documentación interactiva.
- Repositorio oficial de Gemini CLI (Sección de Configuración).
- Guía de `Policy Engine` para control granular de seguridad.

## Example

```toml
# Gemini CLI Context: Nostalgia, Spellbook & Go Ecosystem

## User Profile: Senior Manager & Tech Architect
- **Role:** Senior Manager with a strong technical background.
- **Interests:** Clean Architecture, Modern Microservices, REST API best practices, and robust Infrastructure (IaC, CI/CD, Observability).
- **Languages:** 
  - **Golang (Primary):** Strategic focus for all new features and migrations. Prefers idiomatic Go, Gin for APIs, and Cobra for CLIs.
  - **C# (.NET):** Maintaining legacy personal projects (`nmovies`, `ncookbook`, `nostalgia` C# version) but actively migrating core logic to Go.
- **Communication Style:** Direct, professional, and focus on high-signal architectural decisions.

## Project Portfolio & Evolution

### Nostalgia (C# -> Go Migration)
- **Status:** Active migration via `feature/golang-migration` branch.
- **Go Structure (`nostalgia/golang`):** Follows Clean Architecture/Hexagonal patterns.
  - `internal/domain`: Pure business logic and entities.
  - `internal/app`: Use cases and application services.
  - `internal/infra`: Data persistence (SQL), external APIs, and configuration.
  - `internal/api`: REST handlers using Gin.
  - `cmd/`: Entry points for API and CLI (using Cobra).
- **C# Structure:** Classic Onion Architecture (Core, Application, Data, Api).

### Go Ecosystem (`gostalgia`, `gmovies`)
- **Philosophy:** Lightweight, high-performance services. 
- **Tech Stack:** 
  - **Framework:** Gin-gonic for HTTP.
  - **CLI:** Cobra for command-line tools.
  - **Persistence:** MySQL/MariaDB and SQLite.
  - **CI/CD:** Jenkins (`Jenkinsfile` present in most repos).

### Legacy C# Projects (`nmovies`, `ncookbook`)
- **Status:** Maintenance mode. 
- **Architecture:** Standard .NET Clean Architecture with Entity Framework.

### Spellbook (Documentation)
- **Goal:** Every migration pattern or Go idiom discovered should be documented here.
- **Tech:** Docsify.

## Core Directives for Gemini

### 1. The "Go First" Rule
- For any new feature request across the portfolio, default to proposing a Go-based implementation within the `golang/` directory or a new Go repository.
- Support the `nostalgia` migration by identifying C# logic in `src/` and proposing its Go equivalent in `golang/internal/`.
- Best Practices: Follow the industry best practices and then the specific language best practices, to write idiomatic code. No shortcuts allowed here. 

### 2. Architectural Integrity (Go Context)
- In Go projects, strictly separate `internal/domain` from `internal/infra`.
- Use interfaces in `domain` or `app` to define dependencies that `infra` implements.
- Avoid global state; use dependency injection (passing structs/interfaces).

### 3. Modern REST & Infra
- **APIs:** Ensure Gin handlers are clean, use proper binding/validation, and return consistent JSON error structures.
- **Docker:** Maintain `Dockerfile.api` and `Dockerfile.cli` patterns. Ensure `docker-compose.yml` orchestrates Go services alongside existing databases.

### 4. Documentation & Sync
- When working on Go migrations, update the `spellbook` with "Lessons Learned" or "C# to Go Patterns."

## Preferred Workflow
1. **Research:** Map the C# logic in `nostalgia/src/` before migrating to `nostalgia/golang/`.
2. **Strategy:** Propose the Go structure first.
3. **Execution:** Surgical implementation in Go, followed by unit tests in the new structure.
4. **Validation:** Ensure `go test ./...` passes and Jenkins-compatible structures are maintained.
```