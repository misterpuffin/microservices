# Todo 1: Add mail-service to compose file 

> Goal: Create a compose config for mail service together with mailhog

### mail-service
- locate the Dockerfile and use the correct context to build
- take note of the following environment variables needed:
  MAIL_DOMAIN: localhost
  MAIL_HOST: mailhog
  MAIL_PORT: 1025
  MAIL_ENCRYPTION: none
  MAIL_USERNAME: ""
  MAIL_PASSWORD: ""
  FROM_NAME: "Your Name"
  FROM_ADDRESS: guy@gmail.com

### mailhog
- Use the `mailhog:mailhog:latest` image
- Use `ports` to expose the SMTP server port on 1025, and the frontend port on any port you like
- the default SMTP server starts on port 1025
- the HTTP server starts on port 8025


