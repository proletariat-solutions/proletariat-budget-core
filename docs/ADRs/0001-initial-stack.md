# Architectural Decision Record: Technology Stack Selection

## Status

Accepted (unanymously, i'm the only dev)

## Date

2025-05-30

## Context

When starting the Proletariat Budget application, we needed to select a technology stack that would provide a solid foundation for building a reliable, maintainable, and scalable financial management system. The key components of this stack include:

1. Backend programming language and framework
2. Database management system
3. Frontend framework

We needed to consider factors such as:
- Development speed and efficiency
- Cross-platform compatibility
- Performance characteristics
- Team expertise and learning curve
- Long-term maintainability
- Cloud deployment options
- Community support and ecosystem

## Decision

We have decided to use the following technology stack for our supported client & server apps:

1. **Backend**: Go (Golang)
2. **Database**: MySQL
3. **Frontend**: HTMX with server-side rendered HTML

## Rationale

### Go (Golang) for Backend

1. **Simplicity**: Go's straightforward syntax and minimal feature set reduce cognitive overhead and make the codebase more approachable for new developers.

2. **Cross-platform Compatibility**: Go compiles to native binaries for multiple architectures (x86, ARM, etc.) without dependencies, making deployment flexible across different environments.

3. **Plugin Architecture**: Go's plugin system allows for modular code organization and potential runtime extensions, supporting a more adaptable application structure.

4. **Development Speed**: Go's fast compilation, built-in testing framework, and straightforward concurrency model accelerate the development cycle.

5. **Performance**: Go's compiled nature and efficient garbage collection provide excellent performance characteristics for API services.

6. **Strong Standard Library**: Go's comprehensive standard library reduces dependency on third-party packages for common functionality.

### MySQL for Database

1. **Cross-platform Availability**: MySQL runs on virtually all operating systems and architectures, aligning with our cross-platform goals.

2. **Simplicity**: MySQL offers a straightforward setup, administration, and query language that reduces operational complexity.

3. **Cloud Compatibility**: Using MySQL allows us to seamlessly migrate to Amazon Aurora DB on AWS if needed, without changing our application code due to compatible adapters in Go.

4. **Maturity and Reliability**: MySQL is a battle-tested database with decades of production use across countless applications.

5. **Transactional Support**: MySQL's ACID compliance ensures data integrity for financial transactions.

6. **Scalability Options**: MySQL provides various replication and clustering options to scale as our user base grows.

### HTMX for Frontend

1. **Simplicity**: HTMX allows building dynamic web applications using HTML attributes, eliminating the need for complex JavaScript frameworks and reducing cognitive overhead.

2. **Server-Side Rendering**: By leveraging server-side rendering with Go templates, we maintain a single source of truth for business logic and reduce client-server complexity.

3. **Minimal JavaScript**: HTMX requires minimal custom JavaScript, reducing bundle sizes, build complexity, and potential security vulnerabilities.

4. **Progressive Enhancement**: HTMX enhances standard HTML forms and links, ensuring the application remains functional even if JavaScript fails to load.

5. **Fast Development**: With HTMX, frontend development becomes primarily about HTML templating, which aligns well with Go's template system and reduces context switching.

6. **Performance**: Server-side rendering provides faster initial page loads and better SEO, while HTMX's lightweight nature ensures minimal client-side overhead.

7. **Accessibility**: HTML-first approach naturally promotes better accessibility practices compared to JavaScript-heavy SPAs.

## Consequences

### Positive

1. **Unified Development Experience**: Both backend logic and frontend rendering are handled in Go, reducing context switching and maintaining consistency.

2. **Deployment Simplicity**: Single binary deployment with embedded templates eliminates the need for separate frontend build processes and hosting.

3. **Performance**: Server-side rendering provides excellent initial load times, and HTMX's minimal footprint ensures fast interactions.

4. **Maintainability**: HTML templates are easier to understand and maintain than complex JavaScript component hierarchies.

5. **Security**: Reduced client-side JavaScript surface area minimizes potential XSS vulnerabilities.

6. **SEO-Friendly**: Server-side rendering ensures content is immediately available to search engines. Even though... we won't need SEO given it's a local app (for now).

### Negative

1. **Limited Offline Capabilities**: Server-dependent architecture means limited functionality when offline compared to SPAs.

2. **Real-time Features**: Implementing real-time features like live updates requires additional consideration with WebSockets or Server-Sent Events.

3. **Complex UI Interactions**: Very complex client-side interactions might require custom JavaScript, potentially breaking the HTMX paradigm.

4. **Learning Curve**: Team members familiar with modern JavaScript frameworks may need to adjust to the HTML-first approach.

5. **Third-party Integrations**: Some third-party services designed for SPAs might require additional integration work.

## Alternatives Considered

### Backend Alternatives

1. **Node.js/Express**: Would provide JavaScript consistency across the stack but lacks Go's performance characteristics and type safety. Also, dev hates JS.

2. **Java/Spring**: Offers robust enterprise features but has a steeper learning curve and more verbose development process.

3. **Python/Django or Flask**: Would enable rapid development but might face performance challenges at scale compared to Go. Also, the dev does not know Python :)

### Database Alternatives

1. **PostgreSQL**: Offers more advanced features than MySQL but might be more complex for our current needs.

2. **MongoDB**: Would provide schema flexibility but cloud providers for this are way more expensive, and the footprint to start an instance is larger than MySQL.

3. **SQLite**: Would be simpler and way more faster, but the adapter would only work for this database while MySQL would also work for cloud providers like AWS or GCP.

### Frontend Alternatives

1. **Angular**: Provides a comprehensive framework with strong typing but introduces significant complexity and build processes that may be overkill for our needs.

2. **React**: Offers a large ecosystem and flexibility but requires more decisions about additional libraries, build tools, and state management.

3. **Vue.js**: Provides a gentler learning curve but still requires a separate build process and JavaScript expertise.

4. **Svelte**: Offers excellent performance but has a smaller ecosystem and requires compilation steps.

5. **Plain HTML/CSS/JS**: Would be simpler but lacks the dynamic capabilities needed for a modern financial application.

## Implementation Notes

1. **Template System**: We'll use Go's built-in `html/template` package for server-side rendering with HTMX attributes.

2. **API Design**: We'll design endpoints that return HTML fragments for HTMX requests and full pages for direct navigation.

3. **Database Access**: We'll use a clean repository pattern in Go to abstract database operations, making potential future database migrations easier.

4. **State Management**: Application state will be managed server-side with session storage, reducing client-side complexity.

5. **Deployment Strategy**: We'll containerize the application as a single binary with embedded templates for consistent deployment across environments.

6. **Progressive Enhancement**: Core functionality will work without JavaScript, with HTMX providing enhanced user experience.

## References

- [Go Documentation](https://golang.org/doc/)
- [MySQL Documentation](https://dev.mysql.com/doc/)
- [HTMX Documentation](https://htmx.org/docs/)
- [Go HTML Templates](https://pkg.go.dev/html/template)
- [AWS Aurora DB Compatibility](https://aws.amazon.com/rds/aurora/mysql-features/)