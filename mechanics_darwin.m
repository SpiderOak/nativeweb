#import <Foundation/Foundation.h>
#import "mechanics_darwin.h"

@interface SOGetDataResults : NSObject

@property (copy) NSData *data;
@property (copy) NSURLResponse *resp;
@property (copy) NSError *error;

@end

@implementation SOGetDataResults

@end

/*
actually runs the URL operations for us to get data.
INPUTS:
   char *url: C String of the URL we want to read.

OUTPUTS:
   SOGetDataResults: Object with the results of our get, because C doesn't have multiple returns.
*/
SOGetDataResults* getData(NSString *url) {
    NSURL *my_url = [NSURL URLWithString: url];
    
    NSURLSessionConfiguration *defaultConfigObject = [NSURLSessionConfiguration defaultSessionConfiguration];
    
    NSURLSession *delegateFreeSession = [NSURLSession sessionWithConfiguration: defaultConfigObject
                                                                      delegate: nil
                                                                 delegateQueue: [NSOperationQueue mainQueue]];

    __block BOOL runningURL = YES;
    __block NSData *outputData;
    __block NSURLResponse *outputResp;
    __block NSError *outputError;
    
    [[delegateFreeSession dataTaskWithURL: my_url
                        completionHandler: ^(NSData *data, NSURLResponse *response, NSError *error) {
                outputData = data;
                outputResp = response;
                outputError = error;
                runningURL = NO;
            }] resume];
    
    NSRunLoop *theRL = [NSRunLoop currentRunLoop];
    while (runningURL && [theRL runMode: NSDefaultRunLoopMode beforeDate: [NSDate distantFuture]]);

    SOGetDataResults *retrResults = [[SOGetDataResults alloc] init];

    retrResults.data = outputData;
    retrResults.resp = outputResp;
    retrResults.error = outputError;
    
    return retrResults;
}


void *FetchURL(char *url) {
    SOGetDataResults *results;

    NSString *urlString = [NSString stringWithUTF8String: url];
    results = getData(urlString);
    return results;

}

char *StatusText(void *results) {
    SOGetDataResults *res = results;
    return (char *)[[NSHTTPURLResponse localizedStringForStatusCode: [(NSHTTPURLResponse *)res.resp statusCode]] UTF8String];
}

int StatusCode(void *results) {
    SOGetDataResults *res = results;
    return [(NSHTTPURLResponse *)res.resp statusCode];
}

long long ContentLength(void *results) {
    SOGetDataResults *res = results;
    return [res.resp expectedContentLength];
}

long DataBytesSize(void *results) {
    SOGetDataResults *res = results;

    return [res.data length];
}

void *DataBytes(void *results) {
    SOGetDataResults *res = results;

    return [res.data bytes];
}
