FROM nginxinc/nginx-unprivileged:1-bookworm AS runner

WORKDIR /app

COPY ./deployments/nginx.conf /etc/nginx/nginx.conf
COPY ./src/ ./

EXPOSE 8080

CMD ["nginx", "-g", "daemon off;"]
