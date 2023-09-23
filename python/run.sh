
#!/bin/bash


directory="$1"

# Execute o comando time com o seu comando personalizado aqui
python3 word_count.py "../$directory/dataset0/" "../$directory/dataset1/" "../$directory/dataset2/" "../$directory/dataset3/"
