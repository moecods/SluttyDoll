CQRS (Command Query Responsibility Segregation) is a pattern that separates read and write operations in an application. Instead of having one service or function handle both, CQRS divides them into:

1. Command Side (Write operations) – Responsible for modifying data (Create, Update, Delete).
2. Query Side (Read operations) – Responsible for retrieving data (Read).
   This separation improves scalability, security, and maintainability, especially in high-traffic applications.