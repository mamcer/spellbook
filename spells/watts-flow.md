# Watts Flow: un handbook para equipos que desarrollan con IA

Cómo organizar el trabajo de un equipo de desarrollo cuando escribir código dejó de ser el cuello de botella.

Watts Flow nació en un equipo de producto que construye una plataforma de carga de vehículos eléctricos. Este handbook es la versión genérica: está pensado para que cualquier equipo lo pueda ver, buscar ideas y/o adaptar a su realidad. Los valores concretos (tamaños, límites, objetivos) son un punto de partida, no una ley. Están para calibrarlos con tus propios datos.

---

## 1. Por qué existe

Scrum parte de una premisa: lo escaso es la capacidad de escribir código. Story points, velocity y sprints con alcance fijo existen para administrar ese recurso.

Con herramientas de IA en el día a día, esa premisa se rompe. Tareas que antes llevaban días pueden salir en horas. Y cuando escribir código deja de ser escaso, las herramientas diseñadas para administrarlo se vuelven ruido.

Hay tres síntomas típicos de un equipo en esa situación:

1. **Ilusión de precisión.** Los puntos se negocian en cada planning y no predicen nada. Con IA, la relación entre esfuerzo y puntos se rompe por completo.
2. **El tiempo cede, el alcance no.** Cuando una feature crece, el sprint se estira o el ítem pasa al siguiente. Nadie pregunta si ese alcance extra vale la inversión.
3. **La velocity como señal de salud.** Un sprint lleno de bugs puede tener muy buena velocity y no mover ninguna métrica de producto.

Si en tu equipo una parte grande de lo comprometido en cada sprint pasa al siguiente, o si lo que se mergea tarda días o semanas en llegar a producción, probablemente estás administrando con mucho cuidado algo que ya no es tu restricción.

El nuevo cuello de botella está en todo lo que rodea al código:

- El trabajo abierto en paralelo.
- Los PRs esperando revisión.
- La coordinación entre servicios o equipos.
- La distancia entre "mergeado" y "en producción".

Watts Flow nos sirvió a nosotros para hacer visible y administrar eso.

Hay un segundo motivo, menos obvio. Muchos equipos "arreglan" su proceso a pulmón: alguien (casi siempre el manager) reacomoda sprints, persigue lo trabado y empuja cada release (también lo hicimos acá). Los números mejoran, pero el proceso no funciona: lo sostiene una persona. Eso no escala, y se cae el día que esa persona no está. Un buen proceso tiene que hacer visible lo trabado y repartir las decisiones sin depender de nadie en particular.

---

## 2. De dónde salen las ideas

Nada de esto es nuevo. Es una combinación de marcos que ya usan equipos que hicieron esta transición antes.

### Métricas DORA

DORA (DevOps Research and Assessment) es el programa de investigación detrás del libro *Accelerate*, hoy dentro de Google. Define cuatro métricas que miden la salud del sistema de entrega, no la velocidad de las personas:

- **Frecuencia de deploy:** cada cuánto se despliega a producción.
- **Lead time del cambio:** cuánto tarda un cambio desde el commit hasta producción.
- **Tasa de fallos por cambio (CFR, change failure rate):** qué porcentaje de los deploys causa una falla en producción.
- **Tiempo de recuperación:** cuánto se tarda en recuperarse de una falla.

Los umbrales de "elite" cambian según el año del reporte, así que conviene consultar el más reciente en dora.dev. La idea clave no cambia: si la frecuencia de deploy es baja, no significa que el equipo sea lento. Significa que algo entre el código y producción frena el flujo.

### Flow Metrics

Vienen de Mik Kersten, en *Project to Product*. Hacen visible dónde se traba el trabajo:

- **Velocidad de flujo:** ítems terminados por período.
- **Eficiencia de flujo:** tiempo activo sobre tiempo total. Incluye el tiempo esperando revisión.
- **Tiempo de flujo:** desde que un ítem se empieza hasta que está terminado.
- **Carga de flujo:** cuánto trabajo hay en curso.
- **Distribución de flujo:** proporción de features, bugs, deuda técnica y riesgo.

En muchos equipos el trabajo pasa la mayor parte del tiempo esperando, no en curso. Con IA esto se nota más: la implementación se achica y la espera pasa a dominar el ciclo.

### SPACE

Es un marco de GitHub y Microsoft Research (2021). Mide la productividad en cinco dimensiones: satisfacción, desempeño, actividad, comunicación y eficiencia. Su aporte principal es una advertencia: la actividad (commits, PRs) es una señal útil, pero nunca el objetivo. Una sola dimensión no dice nada, lo que importa es el perfil completo.

### Kanban

Viene del método de David J. Anderson para trabajo de conocimiento. De acá salen el flujo continuo y los límites de trabajo en curso (WIP, work in progress).

### Ley de Little

La formuló John Little en 1961: lead time = WIP / throughput. Si el throughput es constante y duplicás el trabajo en curso, duplicás lo que tarda cada ítem. Es la razón matemática detrás de los límites de WIP.

### Shape Up

Es el método de Basecamp (37signals). De acá sale el concepto de **appetite**: ciclos de tiempo fijo con alcance variable, y nada de backlog infinito. Lo que no entra en el ciclo se vuelve a evaluar, no se acumula.

### Qué tomamos de cada una

Watts Flow no inventa prácticas: elige algunas de cada metodología y las combina. Este es el origen de cada una:

- **Appetite, tiempo fijo con alcance variable, scope hammering:** Shape Up.
- **Planning de inversión en lugar de backlog infinito (lo que no entra se reevalúa, no se acumula):** Shape Up, con su idea de *betting table*.
- **Tablero de flujo continuo y límites de WIP:** Kanban (Anderson).
- **Fast Lane separada de Features:** las clases de servicio de Kanban, en particular la clase *expedite* para lo urgente.
- **Daily centrada en lo que está trabado y no en el estado de cada persona:** la daily de Kanban, que recorre el tablero buscando bloqueos en lugar de preguntar persona por persona.
- **Revisión mensual del flujo con datos:** la *operations review* de Kanban.
- **Por qué limitar el WIP:** Ley de Little.
- **Equilibrio entre Fast Lane y Features:** la distribución de flujo de las Flow Metrics (Kersten).
- **Done = en producción, deploy diario, todo lo mergeado listo para salir:** Continuous Delivery (Humble y Farley) y trunk-based development.
- **Feature flags para desplegar sin activar:** feature toggles (Pete Hodgson).
- **Métricas de entrega (frecuencia de deploy, lead time, tasa de fallos):** DORA.
- **No medir actividad como si fuera resultado:** SPACE.
- **Toda alerta tiene dueño y se lleva hasta una conclusión:** las prácticas de guardia e incidentes de SRE (Google).
- **Contrato primero entre servicios:** consumer-driven contracts (Ian Robinson).
- **Retro con una sola acción, sin culpables:** *Agile Retrospectives* (Derby y Larsen) y la cultura blameless de SRE.
- **Demo de lo que funciona y planning cada dos semanas:** lo que conservamos de Scrum, la sprint review y la cadencia.

---

## 3. El marco en una página

- **Kanban continuo** con límites de WIP para el trabajo diario.
- **Planning por appetite** cada dos semanas para el trabajo nuevo.
- **Pocas ceremonias, cortas y con un propósito.**
- **Dos cadencias explícitas** si tenés productos con ritmos de release distintos (por ejemplo, backend con entrega continua y mobile atado a los stores).
- **Done = en producción.**

No hay story points, ni velocity, ni burndown. No es una simplificación, es una decisión: medir algo que ya no determina el throughput distrae de lo que sí lo determina.

---

## 4. El tablero

```
INBOX → FAST LANE → FEATURE → EN REVISIÓN → PENDIENTE DE DEPLOY → DONE
```

Qué entra en cada columna:

- **Inbox:** todo lo que llega y todavía no se priorizó.
- **Fast Lane:** bugs, hotfixes, observabilidad, tareas de mantenimiento y deuda técnica chica.
- **Feature:** trabajo nuevo con appetite definido.
- **En revisión:** PR abierto, esperando revisión.
- **Pendiente de deploy:** mergeado, todavía no está en producción.
- **Done:** en producción. Si está detrás de un feature flag apagado también cuenta, porque el código ya está desplegado.

### Límites de WIP por persona

- **En curso (Fast Lane + Feature): máximo 2 ítems, y como mucho 1 de ellos es una Feature.**
- **En revisión: máximo 2 PRs propios.** Si tenés más, revisá los de otros antes de abrir uno nuevo.
- **Pendiente de deploy:** no tiene límite, pero es una señal. Si se acumulan ítems acá, faltan deploys. Y a quien tiene varias tarjetas en esta columna le conviene empujar el próximo.

Una regla de la daily: **un ítem que lleva más de 24 horas sin revisión se nombra en la daily.**

### Por qué los límites de WIP no son opcionales

Por la Ley de Little, mucho WIP es la causa de un lead time alto, no su consecuencia. Antes de preguntar "¿por qué tarda tanto este PR?", conviene preguntar "¿cuántas cosas tiene abiertas esta persona?".

Con IA el riesgo es muy concreto: abrir cuatro PRs en paralelo y que los cuatro esperen revisión al mismo tiempo de las mismas personas. El WIP alto de una persona baja el throughput de todo el equipo.

### El tablero tiene que estar en la herramienta

Configurá las columnas reales y los límites de WIP en tu herramienta (Azure DevOps, Jira, Linear, la que sea) desde el día uno. Y verificá que la herramienta guarde el historial de los movimientos entre columnas. Algunas no lo hacen: en Azure DevOps, por ejemplo, el taskboard del sprint muestra las columnas, pero no guarda cuándo pasó cada ítem de una a otra. Sin ese historial no podés medir cuánto tiempo pasa el trabajo esperando revisión o deploy, y vas a necesitar tomar una foto diaria del tablero.

### El equilibrio entre Fast Lane y Features

Un punto de partida razonable es 40% del tiempo en Fast Lane y 60% en Features.

Si la Fast Lane supera el 60% durante más de dos semanas seguidas, el equipo está en modo reactivo. No es un problema de velocidad: algo estructural está roto, como deuda acumulada, un dominio frágil o tests que faltan. La respuesta no es "resolver los bugs más rápido". Es encontrar la causa raíz (¿siempre el mismo dominio?, ¿siempre el mismo tipo de error?) y convertirla en una Feature explícita.

---

## 5. Appetite

### La diferencia central

- **Estimación (Scrum):** la pregunta es "¿cuánto va a tardar?". El alcance queda fijo y el tiempo varía.
- **Appetite:** la pregunta es "¿cuánto vale la pena invertir en esto?". El tiempo queda fijo y el alcance varía.

Parecen la misma pregunta, pero producen respuestas completamente distintas. El appetite no es una estimación. Es una decisión de inversión que toman juntos producto y tecnología, antes de diseñar la solución.

### Tamaños

El ciclo es de dos semanas: 10 días hábiles. Si reservás ~40% para la Fast Lane, quedan ~6 días netos por persona para Features.

- **XS: hasta medio día.** Un ajuste de UI, un evento de analytics nuevo, un cambio de validación.
- **S: 1 a 2 días.** Un endpoint de configuración, una pantalla de filtros.
- **M: 3 a 5 días.** Un flujo nuevo completo, una UI con varias pantallas.
- **L: el ciclo entero de una persona (~6 días netos).** Una feature que cruza servicios o el rediseño de un flujo principal. Quien toma una L no lleva otras Features ese ciclo.
- **Más grande que una L:** no existe como ítem. Es un proyecto o está mal definido, y se descompone antes de llegar al tablero.

Una L solo entra si podés responder: ¿qué va a estar terminado y se va a poder mostrar a mitad del ciclo? Si no hay respuesta, el problema no es que sea complicado. Es que el alcance todavía no está claro.

¿Por qué XS/S/M/L y no Fibonacci? Fibonacci tiene sentido para estimar, porque la imprecisión crece con el tamaño. Pero el appetite no es una estimación: no necesita más precisión, necesita un techo claro. Una escala chica obliga a decidir si algo entra en el ciclo o no. Y aplicar Fibonacci al appetite reintroduce justo lo que se quiere evitar: ítems de "21 puntos" que nadie cuestiona porque la escala los permite.

### Cuando algo no entra en el appetite

El orden importa:

1. **Buscar un camino técnico más simple que llegue al mismo resultado de negocio.** Por ejemplo, un mensaje fijo y bien escrito en vez de plantillas configurables, o un envío directo con buen logging en vez de colas de reintento. Si el usuario ve prácticamente lo mismo, no es deuda técnica: es reconocer que el problema no necesitaba la solución sofisticada. Shape Up llama a esto *scope hammering*.
2. **Si igual no entra, recortar el alcance de negocio, pero a una porción que tenga valor por sí sola.** Por ejemplo, en un flujo de reservas: reservar y cancelar sin penalidad es una unidad completa, mientras que reservar y modificar dejando la cancelación "para después" es media feature.
3. **Lo que quedó afuera no pasa automáticamente al ciclo siguiente.** Es un problema nuevo que, cuando le llegue el turno, tendrá su propio appetite. A veces resulta que ya no vale la pena.

Qué no es appetite:

- No es "terminamos lo que podamos y el resto pasa al próximo sprint". Eso es Scrum con otro nombre.
- No es acumular atajos que generan deuda. Un atajo aceptable se puede defender: "esto es más simple Y es correcto para el problema actual". Si el atajo es "lo hacemos mal ahora y lo arreglamos después", es deuda, no appetite.

### Capacidad: cuándo dejar de sumar ítems al ciclo

En la planning, sumá los appetites que asume cada persona usando estos valores:

- XS = 0.5 días
- S = 1.5 días
- M = 4 días
- L = 6 días

**Cuando la suma de una persona llega a ~6 días, ya no se le asignan más Features en ese ciclo.** El cálculo es por persona, no por equipo, porque los límites de WIP también son por persona.

Ejemplos:

- 1 M + 1 S = 5.5 días: entra un XS más.
- 1 M + 1 M = 8 días: no entra. Una de las dos se va al backlog o se recorta.
- 1 L = 6 días: ciclo completo.

La planning termina cuando cada persona tiene un ciclo lleno pero no sobrecargado. Si después de eso quedan ítems sin asignar, van al backlog. Es el resultado esperado, no un fracaso.

Si un ítem pasa al ciclo siguiente sin terminar (sin contar los adelantos marcados a propósito), casi siempre es un problema de la planning, no de la ejecución. Hay dos causas típicas: el appetite fue demasiado optimista (si un mismo dominio entra siempre corto, hay que calibrarlo) o se comprometió más de lo que había. Si durante la ejecución algo supera su appetite, la conversación es sobre el alcance, no sobre extender el ciclo.

---

## 6. Ceremonias

### Daily: 10 minutos

Una sola pregunta: **¿qué está en riesgo de no avanzar hoy?**

No hay actualizaciones de estado, eso ya está en el tablero. Si algo está trabado hace más de un día, se nombra. Si hay un bloqueo entre servicios o equipos, se resuelve ahí o se le asigna dueño en el momento.

Si nadie tiene nada trabado, la daily dura dos minutos. El silencio es una buena señal.

### Planning: cada dos semanas, 60 minutos

- **Revisión (15 minutos):** qué llegó a producción y qué no, y por qué. No es una retro, es contexto para planificar mejor.
- **Planning (45 minutos):** qué tiene que poder hacer el sistema al final del ciclo que hoy no puede. Cada ítem sale con appetite, responsable y dependencias. Si un ítem cruza dos equipos o servicios, primero se acuerda el contrato. Lo que no tiene un appetite claro no entra.

### Demo: cada dos semanas, 30 minutos

Se muestra lo que está en producción o listo para release, funcionando. Si no se puede mostrar, no cuenta como entregado. "Está casi" no es una demo.

### Retro: cada dos semanas, 45 minutos

Tres preguntas fijas:

1. ¿Qué mejoró concretamente con IA en este ciclo?
2. ¿Qué generó retrabajo o deuda? (Sin culpables, se busca aprender como sistema.)
3. ¿Dónde se trabó el flujo?

Sale una sola acción, con responsable y un resultado medible. Una acción ejecutada vale más que tres olvidadas.

### Revisión mensual del flujo: 30 minutos

Se miran datos, sin personas en la discusión: el diagrama de flujo acumulado (CFD, cumulative flow diagram), dónde está el cuello de botella y si subió la tasa de fallos. Se busca fricción estructural, no culpables.

---

## 7. Métricas

### Qué medir

- **Lead time de Fast Lane, desde que se empieza hasta producción.** Objetivo: menos de 24 horas.
- **Lead time de Features, desde que se empieza hasta producción.** Objetivo: que no supere el appetite del ítem más un día de margen para revisión y deploy, y que una L no supere el ciclo.
- **Frecuencia de deploy.** Objetivo para servicios con entrega continua: al menos uno por día hábil. Leela junto con el lead time: frecuencia baja con lead time estable puede ser una feature grande en curso, y frecuencia baja con lead time en subida es un bloqueo real.
- **Tasa de fallos por cambio.** Objetivo: menos de 5%. Cuando sube, buscá el patrón. Si es siempre el mismo dominio, hay deuda localizada o falta un test. Si está repartida entre repos, el problema es de proceso: review insuficiente o contratos sin acordar.
- **Tiempo de espera en revisión.** Objetivo: menos de 4 horas en Fast Lane y menos de 24 horas en Features.
- **WIP en curso por persona.** Objetivo: 2 o menos.

### Qué dejar de medir

- **Story points y velocity:** miden el esfuerzo, que ya no es el recurso escaso.
- **Commits por día y líneas de código:** con IA cualquiera genera mil líneas. Son una señal de actividad, no de resultado.

### La IA dentro de las métricas

Hay que separar dos cosas que suelen mezclarse:

- **Adopción:** qué porcentaje del trabajo se hace mayormente con IA. Se mide marcando cada PR.
- **Impacto:** cuánto cambió el throughput contra una línea de base previa a la adopción. Sin línea de base no hay impacto medible, solo una sensación.

La regla que importa: **si el throughput sube pero la tasa de fallos también sube, la IA está generando velocidad sin criterio.** El equipo está aceptando código sin entenderlo, y eso es peor que ir más lento.

### Cómo hacer que esto sea medible

- **Cada PR referencia su ítem de trabajo** en el título o la descripción. Sin ese vínculo no se puede calcular el lead time automáticamente.
- **Cada PR generado mayormente con IA lleva una etiqueta** (por ejemplo, `ai-assisted`). No te apoyes solo en el trailer `Co-Authored-By` de los commits: si hacés squash merge, se pierde.
- **La revisión tiene que ser real.** Si la política de branches permite que el autor apruebe su propio PR, el "tiempo de revisión" no mide nada.
- **Guardá el historial de deploys.** Muchas herramientas borran las corridas viejas de pipelines. Si no las guardás, en seis meses no vas a poder medir la frecuencia de deploy.

---

## 8. Trabajar con IA

- **Un archivo de contexto por repositorio** (por ejemplo `CLAUDE.md`) que funcione como contrato vivo: arquitectura, patrones, restricciones. Si un PR introduce un patrón nuevo, actualizarlo es parte de la definición de terminado (DoD, definition of done). Sin ese contexto, la IA genera código plausible pero inconsistente.
- **Prompts de arranque por dominio:** un párrafo con el contexto mínimo de los dominios más activos (qué servicios intervienen, qué contratos no se pueden tocar).
- **Los tests se piden desde el primer prompt,** no se agregan al final. Si el patrón de tests está documentado, la IA lo sigue en lugar de inventar uno.
- **Revisión asistida:** en los PRs complejos, quien revisa lo hace con la IA abierta y le pregunta qué casos borde no están cubiertos. No reemplaza el criterio humano, lo amplifica.
- **Iteraciones cortas:** la IA funciona mejor con pedidos acotados y contexto preciso que con "hacé todo esto de una".
- **Lo que la IA no decide:** contratos entre servicios. Puede implementarlos una vez acordados, pero el acuerdo es entre personas.

---

## 9. Dependencias entre servicios o equipos

La IA acelera la implementación, pero no acelera sola la coordinación. Para cualquier cambio que toque un contrato entre dos servicios:

1. Se define el contrato (endpoint, payload, evento) en un documento o en un PR solo de documentación.
2. Ambos lados implementan en paralelo contra ese contrato.
3. Hay una prueba de integración antes de llegar a producción.

Un ítem que toca un contrato se planifica con los dos responsables en la misma planning. No se empieza hasta que ambos lados tienen claro el contrato.

---

## 10. Acuerdos de equipo

- **Desplegar es responsabilidad de todos.** Nadie debería ser el cuello de botella para llevar código a producción.
- **Cada persona es dueña de su código en producción.** El día que tu cambio llega a producción, lo monitoreás: métricas, logs o una verificación manual. Si algo anda mal, sos el primero en enterarte.
- **Toda alerta tiene dueño.** En horario laboral, alguien la toma explícitamente ("lo miro") y la lleva hasta una conclusión: un bug priorizado, una situación excepcional que no necesita fix o una alerta que hay que recalibrar. Ninguna queda ignorada ni abierta sin cierre. Si el equipo es chico, se rota.
- **Lo que se mergea está listo para producción.** O está completo, o está detrás de un feature flag apagado. Si un cambio vive días en la rama principal sin camino a producción, hay un problema de planning o de alcance.

---

## 11. Cómo empezar en tu equipo

No hace falta adoptar todo de golpe. Este es el orden que recomendaría, según lo que aprendimos:

1. **Medí la línea de base antes de cambiar nada.** Lead time hasta producción, frecuencia de deploy, cuánto pasa de un sprint al siguiente y tasa de fallos. Sin esto, en tres meses no vas a poder saber si mejoró.
2. **Redefiní Done como "en producción"** y agregá la columna "Pendiente de deploy". Es el cambio más barato y el que más rápido muestra dónde se traba el trabajo.
3. **Configurá el tablero real en tu herramienta,** con los límites de WIP. No lo dejes como un acuerdo verbal.
4. **Cambiá la daily** a una sola pregunta.
5. **Pasá la planning a appetite.** Arrancá con los tamaños de este handbook y calibralos después de dos o tres ciclos.
6. **Arreglá lo que hace posible medir:** referencia al ítem en cada PR, etiqueta de IA, revisión real y el historial de deploys guardado.
7. **Revisá las métricas una vez por mes** y ajustá los valores.

Errores que conviene evitar:

- Configurar el tablero sin verificar que guarde el historial de columnas: el proceso funciona, pero no se puede medir dónde espera el trabajo.
- Permitir que los PRs se aprueben con el voto de su autor: el tiempo de revisión deja de significar algo.
- Celebrar más deploys sin mirar la proporción de hotfixes: la velocidad puede esconder un problema de calidad.
- Comparar solo contra el peor momento: si venís de una crisis o de un período de fiestas, mostrá también el período intermedio. Si no, los resultados se inflan y se pierde credibilidad.
- Confundir un proceso que funciona con uno que alguien sostiene a mano: si el manager hace una parte grande de los movimientos del tablero (reacomodar ítems, reasignar, cambiar estados), los números pueden verse bien sin que el proceso funcione. Medí quién mueve el tablero, no solo cómo se mueve.
- Contar como "no terminado" lo que se adelantó a propósito: si un ítem entra el último día del ciclo sabiendo que sigue en el próximo, no es un compromiso incumplido. Marcalo (por ejemplo, con `[Extra]` en el título) y excluilo del carry-over.

---

## Glosario

- **Appetite:** cuánto tiempo vale la pena invertir en un problema. El tiempo es fijo y el alcance se ajusta.
- **CFD (cumulative flow diagram):** diagrama de flujo acumulado. Muestra cuántos ítems hay en cada columna a lo largo del tiempo.
- **CFR (change failure rate):** porcentaje de deploys que causan una falla en producción.
- **DoD (definition of done):** lo que tiene que cumplir un ítem para considerarse terminado.
- **DORA (DevOps Research and Assessment):** programa de investigación sobre el desempeño de la entrega de software.
- **Lead time:** tiempo desde que se empieza un ítem (o desde el primer commit) hasta que está en producción.
- **Throughput:** ítems terminados por período.
- **WIP (work in progress):** trabajo en curso, empezado y no terminado.

---

## Referencias

- *Accelerate*, de Nicole Forsgren, Jez Humble y Gene Kim (2018). La investigación DORA. Reportes actualizados en https://dora.dev
- *SPACE: A Framework for Understanding Developer Productivity*, de Forsgren y otros (ACM Queue, 2021). https://queue.acm.org/detail.cfm?id=3454124
- *Project to Product*, de Mik Kersten (2018). El Flow Framework.
- *Kanban: Successful Evolutionary Change for Your Technology Business*, de David J. Anderson (2010).
- *A Proof for the Queuing Formula L = λW*, de John D. C. Little (Operations Research, 1961). La Ley de Little.
- *Shape Up*, de Ryan Singer (Basecamp, 2019). Gratis en https://basecamp.com/shapeup
- *Continuous Delivery*, de Jez Humble y David Farley (2010).
- *Trunk Based Development*, de Paul Hammant y otros. https://trunkbaseddevelopment.com
- *Feature Toggles (aka Feature Flags)*, de Pete Hodgson (2017). https://martinfowler.com/articles/feature-toggles.html
- *Site Reliability Engineering*, de Betsy Beyer y otros (Google, 2016). Gratis en https://sre.google/books
- *Consumer-Driven Contracts: A Service Evolution Pattern*, de Ian Robinson (2006). https://martinfowler.com/articles/consumerDrivenContracts.html
- *Agile Retrospectives*, de Esther Derby y Diana Larsen (2006).
- *The Scrum Guide*, de Ken Schwaber y Jeff Sutherland. Lo que conservamos y lo que dejamos. https://scrumguides.org
