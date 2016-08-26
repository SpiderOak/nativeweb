#include <Cocoa/Cocoa.h>

void *FetchURL(char *url);
char *StatusText(void *results);
int StatusCode(void *results);
long long ContentLength(void *results);
long DataBytesSize(void *results);
void *DataBytes(void *results);
