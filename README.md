# NativeWeb

# What is it?

NativeWeb aims to provide a client similar to net/http.Client, with
all transport and HTTP tasks provided by the native system HTTP
libraries. For use with client software, this gets us additional
built-in functionality that comes for free with the native HTTP
clients.

# Why is it?

The native Go HTTP client lacks features sometimes useful for client
software, such as NTLM/Negotiate authentication and automatic proxy
support on Windows and being able to
[power up the cellular radio on iOS](https://developer.apple.com/library/ios/documentation/NetworkingInternet/Conceptual/NetworkingTopics/Articles/UsingSocketsandSocketStreams.html#//apple_ref/doc/uid/CH73-SW4). 

# How is it built?

We aim to provide a client interface that implements a substantial 
