#include <stdio.h>

const char* hello_string() {
	return "Hello, world.\n";
}

void hello() {
	printf("%s", hello_string());
}
