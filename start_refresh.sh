#!/bin/bash
# Este script es para iniciar reflex usando Bazel.

# Aquí se asume que tienes un target en Bazel que se llama 'reflex_binary'
# el cual está configurado para ejecutar el binario de reflex.
bazel run //:reflex_binary
