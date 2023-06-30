## gate is an API gateway with request tracking, approval and denial capabilities

<p>It tracks and manages incoming requests, applies permission checks, and grants or denies access to the requested external API based on defined rules or manual approval. This API gateway is designed to enforce access control, implement security measures, and centralize request management in a controlled and auditable manner.</p>

### How it works?

<p>It provides a middleware layer to control access and enforce restrictions based on predefined rules.</p>

<p>The key components of the project are as follows:</p>

- Middleware: The project utilizes middleware functions implemented using the Gin framework. These middlewares intercept incoming requests and perform necessary checks before allowing or denying access.

- Permission Checks: The project incorporates permission checks to determine whether a request is allowed to access the external API. It maintains a PermissionMap that stores permissions for specific requests. Each incoming request is matched against the PermissionMap to determine if it has permission to proceed.

- Rate Limiting: The project aims to implement rate limiting to restrict the number of requests that can be made to the external API within a specified timeframe. It employs a counter mechanism to track the number of requests and enforce the defined limit.

- Request Logging: Each incoming request is logged into a request.log file. The log entry includes the timestamp, request method, request path, and the status of access (e.g., APPROVED or DENIED).

- Environmental Configuration: The project supports environmental configuration using a .env file. This allows for easy customization and adjustment of various parameters, such as the external API endpoint, rate limit values, and more.
