#include <stdio.h>
#include <stdlib.h>
#include <libascart.h>

int main(int argc, char *argv[]){
    int rv,i,j;
    ascart *art;
    if (argc != 4){
        printf("Wrong number of args %d\n",argc);
        exit(EXIT_FAILURE);
    }
    rv=Image2Ascii(argv[1], argv[2], argv[3], &art);
    if (rv < 0){
        printf("Bad rv %d. Inputs \"%s\" \"%s\" \"%s\"",rv,argv[1], argv[2], argv[3]);
        exit(EXIT_FAILURE);
    }
    for (i=0;i<art->h;i++){
        for(j=0;j<art->w;j++){
            putchar(art->art[i*art->w+j]);
        }
        putchar('\n');
    }
    exit(EXIT_SUCCESS);
}
