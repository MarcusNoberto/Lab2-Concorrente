#!/bin/bash

# Verifica se o Gradle está instalado
if ! command -v gradle &> /dev/null; then
    echo "Gradle não encontrado. Certifique-se de que o Gradle está instalado e configurado corretamente."
    exit 1
fi

# Executa o comando 'gradle runApp'
gradle runApp
