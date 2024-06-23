#!/bin/bash

SOURCE_DIR="../internal/core"
DESTINATION_DIR="../test/mock"

mkdir -p $DESTINATION_DIR

# Encontrar todos os arquivos que terminam com _repository.go nos subdiretórios de core
find $SOURCE_DIR -type f -name "*_repository.go" | while read file; do
    # Extrair o nome do arquivo sem a extensão
    filename=$(basename -- "$file")
    filename="${filename%.*}"

    # Gerar o mock usando mockgen
    mockgen -source="$file" -destination="$DESTINATION_DIR/mock_$filename.go" -package=mock
done
