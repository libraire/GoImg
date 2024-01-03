###################
# BUILD FOR BUILD 
###################
FROM nginx:latest

# Set the working directory inside the container
WORKDIR /usr/share/nginx/html

# Copy the HTML source file from the host into the container
COPY nginx.conf /etc/nginx/nginx.conf   

# Expose port 80 for the NGINX server
EXPOSE 80
EXPOSE 443

# Start NGINX when the container starts
CMD ["nginx", "-g", "daemon off;"]
