find -type f find -name '*.sh' | sed 's/.sh//g' | sort -r | tr -d './'

#-printf "%f\n"