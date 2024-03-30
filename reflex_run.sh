#!/bin/bash
# Este script utiliza Reflex para monitorear cambios en archivos Go y luego ejecuta Bazel para reconstruir y reiniciar la aplicación.
# -r '\.go$': Monitorea todos los archivos que terminen en .go
# -s: Detiene la ejecución del script en el primer error (opcional).
# -- sh -c 'comando': Ejecuta el comando especificado cuando se detecta un cambio.

$(location de reflex) -r '\.go$' -s -- sh -c 'bazel run //:main'
