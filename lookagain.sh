find -type f find -name '*.sh' -printf "%f\n" | sed 's/.sh//g' | sort -r 
 
