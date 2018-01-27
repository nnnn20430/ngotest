// +build ignore

#include <stdio.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <unistd.h>
#include <stdlib.h>
#include <errno.h>

int main(int argc, char *argv[]) {
	int fd; long fs; char *buf;
	if (argc < 2) {
		printf("No file specified as first parameter.\n");
		return 1;
	}
	fd = open(argv[1], O_RDONLY);
	if (fd == -1) {
		perror(NULL);
		return errno;
	}
	fs = lseek(fd, 0, SEEK_END);
	lseek(fd, 0, SEEK_SET);
	buf = malloc(fs+1);
	read(fd, buf, fs);
	close(fd);
	write(STDOUT_FILENO, buf, fs);
	free(buf);
}
