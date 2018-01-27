// +build ignore

#include <stdio.h>
#include <stdlib.h>
#include <errno.h>

int main(int argc, char *argv[]) {
	FILE *f; long fs; char *buf;
	if (argc < 2) {
		printf("No file specified as first parameter.\n");
		return 1;
	}
	f = fopen(argv[1], "r");
	if (f == NULL) {
		perror(NULL);
		return errno;
	}
	fseek(f, 0L, SEEK_END);
	fs = ftell(f);
	fseek(f, 0L, SEEK_SET);
	buf = malloc(fs+1);
	fread(buf, 1, fs, f);
	fclose(f);
	fwrite(buf, 1, fs, stdout);
	free(buf);
}
