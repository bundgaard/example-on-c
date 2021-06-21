#include <windows.h>

char *buffer;

void hej() {

buffer = (char*) malloc(2048*1024);

}

void cleanup() {

free(buffer);

}