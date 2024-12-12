---
title: Quick Start
description: Quick Start page.
tableOfContents: false
---

Seam is a lightweight and high-performance API Gateway built with Go, leveraging the <a href="https://github.com/gofiber/fiber">Fiber</a> web framework. Designed for modern microservices architectures, Seam helps manage traffic routing, load balancing, authentication, and monitoring in a simple and efficient way.  

Follow these steps to quickly set up and run the **Seam API Gateway**:

### Prerequisites

- **Docker**: Ensure Docker is installed and running on your system. You can download it [here](https://www.docker.com/).

### Installation

1. **Clone the Repository**  
   Begin by cloning the Seam repository to your local machine:

   ```sh
   git clone https://github.com/saulova/seam.git
   cd seam
   ```

1. **Configure Settings**  
   Rename the sample configuration file and customize it to match your requirements:

   ```sh
   mkdir -p ./.configs
   cp -r ./.configs.sample/* ./.configs/
   ```

   Open the `.configs` file in your favorite text editor and update the settings as needed.

1. **Run the Docker Container**
    - **Option 1**
        1. **Using Docker Compose**  
            Start the Seam server by running the Docker compose:

            ```sh
            docker compose up -d
            ```

    - **Option 2**
        1. **Build the Docker Image**  
            Build the Seam Docker image:

            ```sh
            docker build --tag 'seam' .
            ```

        1. **Run the Docker Container**  
            Start the Seam server by running the Docker container:

            ```sh
            docker run -p 8090:8090 --detach 'seam'
            ```

5. **Access the API Gateway**  
   The Seam server is now running and accessible on `http://localhost:8090`.

### Verifying the Installation

1. **Health Check**  
   If health check routes are enabled in your configuration, verify the server's readiness:

   - **Liveness Route**: `http://localhost:8090/health/live`
   - **Readiness Route**: `http://localhost:8090/health/ready`

2. **Log Monitoring**  
   Check the container logs to ensure the server is running as expected:
   ```sh
   docker logs [CONTAINER_ID]
   ```
   Replace `[CONTAINER_ID]` with the ID of the running Seam container.

### Next Steps

- **Load Plugins**: Add plugins to extend the gateway's functionality.
- **Configure Routes**: Define your routes, middlewares, actions, and storages `.yaml` config files in the `.configs/` folder.
- **Deploy**: Once configured, deploy Seam to your production environment.
