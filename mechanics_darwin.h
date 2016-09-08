#include <Cocoa/Cocoa.h>

void *OpenSession();
void *FetchURL(void *session, char *url);
char *StatusText(void *results);
int StatusCode(void *results);
long long ContentLength(void *results);
long DataBytesSize(void *results);
void *DataBytes(void *results);
void Release(void *obj);