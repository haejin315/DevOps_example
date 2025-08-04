FROM nginx:alpine

COPY --from=devops_example/react-builder:1.0 /build/dist /usr/share/nginx/html

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
