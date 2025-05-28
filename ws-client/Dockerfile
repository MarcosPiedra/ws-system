# Usa la imagen oficial de Nginx
FROM nginx:alpine

# Copia tu archivo HTML al directorio por defecto de Nginx
COPY index.html /usr/share/nginx/html/index.html

# Opcional: si tienes un archivo de configuración personalizado
# COPY default.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

# El contenedor inicia con nginx en modo foreground
CMD ["nginx", "-g", "daemon off;"]