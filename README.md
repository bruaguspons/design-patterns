# 🧠 Design Patterns - TypeScript, Go & Python

Este repositorio contiene implementaciones de distintos **patrones de
diseño** en:

- TypeScript
- Go
- Python

El objetivo es mostrar cómo se implementan los mismos patrones en
distintos lenguajes y comparar diferencias en sintaxis, tipado y
filosofía.

------------------------------------------------------------------------

# 📚 Patrones incluidos

- Adapter
- Builder
- Factory
- Observer
- Singleton
- Strategy

------------------------------------------------------------------------

# 1️⃣ Adapter

## 📌 ¿Qué problema resuelve?

El patrón **Adapter** permite que dos interfaces incompatibles trabajen
juntas.

Es útil cuando: - Queremos reutilizar una clase existente - Su interfaz
no coincide con la que nuestro sistema espera - No podemos modificar la
clase original

## 🧠 Idea principal

El Adapter actúa como un **traductor** entre dos interfaces.

Client → Adapter → Adaptee

## 🧩 Cuándo usarlo

-   Integración con librerías externas
-   Migraciones progresivas de sistemas
-   Unificación de múltiples APIs bajo una interfaz común

------------------------------------------------------------------------

# 2️⃣ Builder

## 📌 ¿Qué problema resuelve?

El patrón **Builder** permite construir objetos complejos paso a paso.

Es útil cuando: - Un objeto tiene muchos parámetros opcionales -
Queremos evitar constructores enormes - Queremos controlar el proceso de
construcción

## 🧠 Idea principal

Separar: - 🏗️ La construcción del objeto
- 📦 La representación final

Permite crear distintas representaciones usando el mismo proceso.

## 🧩 Ventajas

-   Mejora la legibilidad
-   Evita constructores telescópicos
-   Permite objetos inmutables

------------------------------------------------------------------------

# 3️⃣ Factory

## 📌 ¿Qué problema resuelve?

El patrón **Factory** encapsula la lógica de creación de objetos.

En vez de usar `new` directamente, delegamos la creación a una fábrica.

## 🧠 Idea principal

Delegar la responsabilidad de instanciación a una clase o función
especializada.

Client → Factory → Concrete Product

## 🧩 Cuándo usarlo

-   Cuando la lógica de creación es compleja
-   Cuando queremos desacoplar cliente de implementación concreta
-   Cuando el tipo a crear depende de condiciones

------------------------------------------------------------------------

# 4️⃣ Observer

## 📌 ¿Qué problema resuelve?

El patrón **Observer** define una dependencia uno-a-muchos entre
objetos.

Cuando un objeto cambia su estado, todos sus observadores son
notificados automáticamente.

## 🧠 Idea principal

Subject → notify() → Observers

Los observadores se suscriben y reciben actualizaciones.

## 🧩 Cuándo usarlo

-   Sistemas de eventos
-   UI reactiva
-   Arquitecturas basadas en eventos
-   Pub/Sub

------------------------------------------------------------------------

# 5️⃣ Singleton

## 📌 ¿Qué problema resuelve?

El patrón **Singleton** garantiza que exista una única instancia de una
clase y provee un punto global de acceso a ella.

## 🧠 Idea principal

-   Constructor privado (o controlado)
-   Instancia estática compartida

## 🧩 Cuándo usarlo

-   Configuración global
-   Logger
-   Pool de conexiones
-   Cache compartido

⚠️ Advertencia: Puede dificultar testing si se usa en exceso.

------------------------------------------------------------------------

# 6️⃣ Strategy

## 📌 ¿Qué problema resuelve?

El patrón **Strategy** permite definir una familia de algoritmos,
encapsularlos y hacerlos intercambiables.

Evita múltiples `if` o `switch` para cambiar comportamiento.

## 🧠 Idea principal

Context → Strategy Interface → Concrete Strategies

El comportamiento se puede cambiar dinámicamente.

## 🧩 Cuándo usarlo

-   Múltiples formas de resolver un problema
-   Algoritmos intercambiables
-   Lógica seleccionable en runtime

------------------------------------------------------------------------

# 📌 Notas

-   Cada patrón está implementado de forma simple y didáctica.
-   No se busca código orientado a producción, sino claridad conceptual.
-   Se prioriza la comprensión del patrón por sobre la optimización.
