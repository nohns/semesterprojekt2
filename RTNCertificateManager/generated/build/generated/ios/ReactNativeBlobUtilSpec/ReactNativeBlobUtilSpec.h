/**
 * This code was generated by [react-native-codegen](https://www.npmjs.com/package/react-native-codegen).
 *
 * Do not edit this file as changes may cause incorrect behavior and will be lost
 * once the code is regenerated.
 *
 * @generated by codegen project: GenerateModuleObjCpp
 *
 * We create an umbrella header (and corresponding implementation) here since
 * Cxx compilation in BUCK has a limitation: source-code producing genrule()s
 * must have a single output. More files => more genrule()s => slower builds.
 */

#ifndef __cplusplus
#error This file must be compiled as Obj-C++. If you are importing it, you must change your file extension to .mm.
#endif
#import <Foundation/Foundation.h>
#import <RCTRequired/RCTRequired.h>
#import <RCTTypeSafety/RCTConvertHelpers.h>
#import <RCTTypeSafety/RCTTypedModuleConstants.h>
#import <React/RCTBridgeModule.h>
#import <React/RCTCxxConvert.h>
#import <React/RCTManagedPointer.h>
#import <ReactCommon/RCTTurboModule.h>
#import <optional>
#import <vector>

namespace JS {
  namespace NativeBlobUtils {
    struct Constants {

      struct Builder {
        struct Input {
          RCTRequired<NSString *> CacheDir;
          RCTRequired<NSString *> DocumentDir;
          RCTRequired<NSString *> DownloadDir;
          RCTRequired<NSString *> LibraryDir;
          RCTRequired<NSString *> MainBundleDir;
          RCTRequired<NSString *> MovieDir;
          RCTRequired<NSString *> MusicDir;
          RCTRequired<NSString *> PictureDir;
          RCTRequired<NSString *> ApplicationSupportDir;
          RCTRequired<NSString *> RingtoneDir;
          RCTRequired<NSString *> SDCardDir;
          RCTRequired<NSString *> SDCardApplicationDir;
          RCTRequired<NSString *> DCIMDir;
          RCTRequired<NSString *> LegacyDCIMDir;
          RCTRequired<NSString *> LegacyPictureDir;
          RCTRequired<NSString *> LegacyMusicDir;
          RCTRequired<NSString *> LegacyDownloadDir;
          RCTRequired<NSString *> LegacyMovieDir;
          RCTRequired<NSString *> LegacyRingtoneDir;
          RCTRequired<NSString *> LegacySDCardDir;
        };

        /** Initialize with a set of values */
        Builder(const Input i);
        /** Initialize with an existing Constants */
        Builder(Constants i);
        /** Builds the object. Generally used only by the infrastructure. */
        NSDictionary *buildUnsafeRawValue() const { return _factory(); };
      private:
        NSDictionary *(^_factory)(void);
      };

      static Constants fromUnsafeRawValue(NSDictionary *const v) { return {v}; }
      NSDictionary *unsafeRawValue() const { return _v; }
    private:
      Constants(NSDictionary *const v) : _v(v) {}
      NSDictionary *_v;
    };
  }
}
@protocol NativeBlobUtilsSpec <RCTBridgeModule, RCTTurboModule>

- (void)fetchBlobForm:(NSDictionary *)options
               taskId:(NSString *)taskId
               method:(NSString *)method
                  url:(NSString *)url
              headers:(NSDictionary *)headers
                 form:(NSArray *)form
             callback:(RCTResponseSenderBlock)callback;
- (void)fetchBlob:(NSDictionary *)options
           taskId:(NSString *)taskId
           method:(NSString *)method
              url:(NSString *)url
          headers:(NSDictionary *)headers
             body:(NSString *)body
         callback:(RCTResponseSenderBlock)callback;
- (void)createFile:(NSString *)path
              data:(NSString *)data
          encoding:(NSString *)encoding
           resolve:(RCTPromiseResolveBlock)resolve
            reject:(RCTPromiseRejectBlock)reject;
- (void)createFileASCII:(NSString *)path
                   data:(NSArray *)data
                resolve:(RCTPromiseResolveBlock)resolve
                 reject:(RCTPromiseRejectBlock)reject;
- (void)pathForAppGroup:(NSString *)groupName
                resolve:(RCTPromiseResolveBlock)resolve
                 reject:(RCTPromiseRejectBlock)reject;
- (NSString *)syncPathAppGroup:(NSString *)groupName;
- (void)exists:(NSString *)path
      callback:(RCTResponseSenderBlock)callback;
- (void)writeFile:(NSString *)path
         encoding:(NSString *)encoding
             data:(NSString *)data
    transformFile:(BOOL)transformFile
           append:(BOOL)append
          resolve:(RCTPromiseResolveBlock)resolve
           reject:(RCTPromiseRejectBlock)reject;
- (void)writeFileArray:(NSString *)path
                  data:(NSArray *)data
                append:(BOOL)append
               resolve:(RCTPromiseResolveBlock)resolve
                reject:(RCTPromiseRejectBlock)reject;
- (void)writeStream:(NSString *)path
       withEncoding:(NSString *)withEncoding
         appendData:(BOOL)appendData
           callback:(RCTResponseSenderBlock)callback;
- (void)writeArrayChunk:(NSString *)streamId
              withArray:(NSArray *)withArray
               callback:(RCTResponseSenderBlock)callback;
- (void)writeChunk:(NSString *)streamId
          withData:(NSString *)withData
          callback:(RCTResponseSenderBlock)callback;
- (void)closeStream:(NSString *)streamId
           callback:(RCTResponseSenderBlock)callback;
- (void)unlink:(NSString *)path
      callback:(RCTResponseSenderBlock)callback;
- (void)removeSession:(NSArray *)paths
             callback:(RCTResponseSenderBlock)callback;
- (void)ls:(NSString *)path
   resolve:(RCTPromiseResolveBlock)resolve
    reject:(RCTPromiseRejectBlock)reject;
- (void)stat:(NSString *)target
    callback:(RCTResponseSenderBlock)callback;
- (void)lstat:(NSString *)path
     callback:(RCTResponseSenderBlock)callback;
- (void)cp:(NSString *)src
      dest:(NSString *)dest
  callback:(RCTResponseSenderBlock)callback;
- (void)mv:(NSString *)path
      dest:(NSString *)dest
  callback:(RCTResponseSenderBlock)callback;
- (void)mkdir:(NSString *)path
      resolve:(RCTPromiseResolveBlock)resolve
       reject:(RCTPromiseRejectBlock)reject;
- (void)readFile:(NSString *)path
        encoding:(NSString *)encoding
   transformFile:(BOOL)transformFile
         resolve:(RCTPromiseResolveBlock)resolve
          reject:(RCTPromiseRejectBlock)reject;
- (void)hash:(NSString *)path
   algorithm:(NSString *)algorithm
     resolve:(RCTPromiseResolveBlock)resolve
      reject:(RCTPromiseRejectBlock)reject;
- (void)readStream:(NSString *)path
          encoding:(NSString *)encoding
        bufferSize:(double)bufferSize
              tick:(double)tick
          streamId:(NSString *)streamId;
- (void)getEnvironmentDirs:(RCTResponseSenderBlock)callback;
- (void)cancelRequest:(NSString *)taskId
             callback:(RCTResponseSenderBlock)callback;
- (void)enableProgressReport:(NSString *)taskId
                    interval:(double)interval
                       count:(double)count;
- (void)enableUploadProgressReport:(NSString *)taskId
                          interval:(double)interval
                             count:(double)count;
- (void)slice:(NSString *)src
         dest:(NSString *)dest
        start:(double)start
          end:(double)end
      resolve:(RCTPromiseResolveBlock)resolve
       reject:(RCTPromiseRejectBlock)reject;
- (void)presentOptionsMenu:(NSString *)uri
                    scheme:(NSString *)scheme
                   resolve:(RCTPromiseResolveBlock)resolve
                    reject:(RCTPromiseRejectBlock)reject;
- (void)presentOpenInMenu:(NSString *)uri
                   scheme:(NSString *)scheme
                  resolve:(RCTPromiseResolveBlock)resolve
                   reject:(RCTPromiseRejectBlock)reject;
- (void)presentPreview:(NSString *)uri
                scheme:(NSString *)scheme
               resolve:(RCTPromiseResolveBlock)resolve
                reject:(RCTPromiseRejectBlock)reject;
- (void)excludeFromBackupKey:(NSString *)url
                     resolve:(RCTPromiseResolveBlock)resolve
                      reject:(RCTPromiseRejectBlock)reject;
- (void)df:(RCTResponseSenderBlock)callback;
- (void)emitExpiredEvent:(RCTResponseSenderBlock)callback;
- (void)actionViewIntent:(NSString *)path
                    mime:(NSString *)mime
            chooserTitle:(NSString *)chooserTitle
                 resolve:(RCTPromiseResolveBlock)resolve
                  reject:(RCTPromiseRejectBlock)reject;
- (void)addCompleteDownload:(NSDictionary *)config
                    resolve:(RCTPromiseResolveBlock)resolve
                     reject:(RCTPromiseRejectBlock)reject;
- (void)copyToInternal:(NSString *)contentUri
              destpath:(NSString *)destpath
               resolve:(RCTPromiseResolveBlock)resolve
                reject:(RCTPromiseRejectBlock)reject;
- (void)copyToMediaStore:(NSDictionary *)filedata
                      mt:(NSString *)mt
                    path:(NSString *)path
                 resolve:(RCTPromiseResolveBlock)resolve
                  reject:(RCTPromiseRejectBlock)reject;
- (void)createMediaFile:(NSDictionary *)filedata
                     mt:(NSString *)mt
                resolve:(RCTPromiseResolveBlock)resolve
                 reject:(RCTPromiseRejectBlock)reject;
- (void)getBlob:(NSString *)contentUri
       encoding:(NSString *)encoding
        resolve:(RCTPromiseResolveBlock)resolve
         reject:(RCTPromiseRejectBlock)reject;
- (void)getContentIntent:(NSString *)mime
                 resolve:(RCTPromiseResolveBlock)resolve
                  reject:(RCTPromiseRejectBlock)reject;
- (void)getSDCardDir:(RCTPromiseResolveBlock)resolve
              reject:(RCTPromiseRejectBlock)reject;
- (void)getSDCardApplicationDir:(RCTPromiseResolveBlock)resolve
                         reject:(RCTPromiseRejectBlock)reject;
- (void)scanFile:(NSArray *)pairs
        callback:(RCTResponseSenderBlock)callback;
- (void)writeToMediaFile:(NSString *)fileUri
                    path:(NSString *)path
           transformFile:(BOOL)transformFile
                 resolve:(RCTPromiseResolveBlock)resolve
                  reject:(RCTPromiseRejectBlock)reject;
- (facebook::react::ModuleConstants<JS::NativeBlobUtils::Constants::Builder>)constantsToExport;
- (facebook::react::ModuleConstants<JS::NativeBlobUtils::Constants::Builder>)getConstants;

@end
namespace facebook {
  namespace react {
    /**
     * ObjC++ class for module 'NativeBlobUtils'
     */
    class JSI_EXPORT NativeBlobUtilsSpecJSI : public ObjCTurboModule {
    public:
      NativeBlobUtilsSpecJSI(const ObjCTurboModule::InitParams &params);
    };
  } // namespace react
} // namespace facebook
inline JS::NativeBlobUtils::Constants::Builder::Builder(const Input i) : _factory(^{
  NSMutableDictionary *d = [NSMutableDictionary new];
  auto CacheDir = i.CacheDir.get();
  d[@"CacheDir"] = CacheDir;
  auto DocumentDir = i.DocumentDir.get();
  d[@"DocumentDir"] = DocumentDir;
  auto DownloadDir = i.DownloadDir.get();
  d[@"DownloadDir"] = DownloadDir;
  auto LibraryDir = i.LibraryDir.get();
  d[@"LibraryDir"] = LibraryDir;
  auto MainBundleDir = i.MainBundleDir.get();
  d[@"MainBundleDir"] = MainBundleDir;
  auto MovieDir = i.MovieDir.get();
  d[@"MovieDir"] = MovieDir;
  auto MusicDir = i.MusicDir.get();
  d[@"MusicDir"] = MusicDir;
  auto PictureDir = i.PictureDir.get();
  d[@"PictureDir"] = PictureDir;
  auto ApplicationSupportDir = i.ApplicationSupportDir.get();
  d[@"ApplicationSupportDir"] = ApplicationSupportDir;
  auto RingtoneDir = i.RingtoneDir.get();
  d[@"RingtoneDir"] = RingtoneDir;
  auto SDCardDir = i.SDCardDir.get();
  d[@"SDCardDir"] = SDCardDir;
  auto SDCardApplicationDir = i.SDCardApplicationDir.get();
  d[@"SDCardApplicationDir"] = SDCardApplicationDir;
  auto DCIMDir = i.DCIMDir.get();
  d[@"DCIMDir"] = DCIMDir;
  auto LegacyDCIMDir = i.LegacyDCIMDir.get();
  d[@"LegacyDCIMDir"] = LegacyDCIMDir;
  auto LegacyPictureDir = i.LegacyPictureDir.get();
  d[@"LegacyPictureDir"] = LegacyPictureDir;
  auto LegacyMusicDir = i.LegacyMusicDir.get();
  d[@"LegacyMusicDir"] = LegacyMusicDir;
  auto LegacyDownloadDir = i.LegacyDownloadDir.get();
  d[@"LegacyDownloadDir"] = LegacyDownloadDir;
  auto LegacyMovieDir = i.LegacyMovieDir.get();
  d[@"LegacyMovieDir"] = LegacyMovieDir;
  auto LegacyRingtoneDir = i.LegacyRingtoneDir.get();
  d[@"LegacyRingtoneDir"] = LegacyRingtoneDir;
  auto LegacySDCardDir = i.LegacySDCardDir.get();
  d[@"LegacySDCardDir"] = LegacySDCardDir;
  return d;
}) {}
inline JS::NativeBlobUtils::Constants::Builder::Builder(Constants i) : _factory(^{
  return i.unsafeRawValue();
}) {}
