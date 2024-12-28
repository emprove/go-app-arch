### Go app 4-tier Architecture

This is a simple example to structure your HTTP + CLI application in Go using **N-tier Architecture**. It consists of 4 layers: Presentation, Application, Business (or Domain in DDD) and Infrastructure (data access). It is very simple and productive for small applications. The main idea of Layered Architecture (N-tier) is to make layers independent. It gives us interchangeability of implementations on the each layer, so we should utilize Interfaces and not concrete implementations.

Usecases within Application layer spans several Services from the Business layer. Because Services from Business layer should not call each other. We can jump over Application layer straight to Business if we don't need Usecases. It's a widely used practice to avoid using Layer as simple data-passing tunnel.

Folder structure not aligned with the architectural model. Many Go tutorials recommend naming your folders the same as your package name. I don’t see the need to make nested folders as long as the application is small.

### Diagram

![Diagram](diagram.jpg)

![Diagram description](diagram-desc.jpg)

> DISCLAIMER: This repository is a heavily stripped down version of the real-world application. Something might not work.
