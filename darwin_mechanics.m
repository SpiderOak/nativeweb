#import <Foundation/Foundation.h>
#include "darwin_mechanics.h"

    
NSData *getData(char *url) {
    NSURL *my_url = [NSURL URLWithString: [NSString stringWithUTF8String:url]];

    NSURLSessionConfiguration *defaultConfigObject = [NSURLSessionConfiguration defaultSessionConfiguration];
    
    NSURLSession *delegateFreeSession = [NSURLSession sessionWithConfiguration: defaultConfigObject
                                                                      delegate: nil
                                                                 delegateQueue: [NSOperationQueue mainQueue]];

    __block BOOL runningURL = YES;
    __block NSData *outputData;

    NSRunLoop *theRL = [NSRunLoop currentRunLoop];

    NSLog(@"Starting URLSession");
    [[delegateFreeSession dataTaskWithURL: my_url
                        completionHandler: ^(NSData *data, NSURLResponse *response, NSError *error) {
                NSLog(@"Got response %@ with error %@.\n", response, error);
                outputData = [NSData dataWithData: data];
                runningURL = NO;
            }] resume];

    NSLog(@"Starting run loop...");
    while (runningURL && [theRL runMode: NSDefaultRunLoopMode beforeDate: [NSDate distantFuture]]);
    NSLog(@"Run loop complete!");

    return outputData;
}


char *FetchURL(char *url) {
    NSData *retBody;

    retBody = getData(url);

    return [retBody bytes];
}
