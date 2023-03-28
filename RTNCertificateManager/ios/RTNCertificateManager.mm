#import "RTNCertificateManager.h"

@implementation RTNCertificateManager

RCT_EXPORT_MODULE()

- (void)createCSR:(RCTPromiseResolveBlock)resolve
           reject:(RCTPromiseRejectBlock)reject {

    NSString *result = @"Hello from Objective-C!";
    resolve(result);
}
- (void)storeCertificate:(NSString *)certificate
                 resolve:(RCTPromiseResolveBlock)resolve
                  reject:(RCTPromiseRejectBlock)reject {
    resolve(NULL);
}

- (std::shared_ptr<facebook::react::TurboModule>)getTurboModule:
    (const facebook::react::ObjCTurboModule::InitParams &)params
{
    return std::make_shared<facebook::react::NativeCertificateManagerSpecJSI>(params);
}

@end


