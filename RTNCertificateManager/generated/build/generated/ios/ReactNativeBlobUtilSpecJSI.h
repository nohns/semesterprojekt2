/**
 * This code was generated by [react-native-codegen](https://www.npmjs.com/package/react-native-codegen).
 *
 * Do not edit this file as changes may cause incorrect behavior and will be lost
 * once the code is regenerated.
 *
 * @generated by codegen project: GenerateModuleH.js
 */

#pragma once

#include <ReactCommon/TurboModule.h>
#include <react/bridging/Bridging.h>

namespace facebook {
namespace react {

class JSI_EXPORT NativeBlobUtilsCxxSpecJSI : public TurboModule {
protected:
  NativeBlobUtilsCxxSpecJSI(std::shared_ptr<CallInvoker> jsInvoker);

public:
  virtual jsi::Object getConstants(jsi::Runtime &rt) = 0;
  virtual void fetchBlobForm(jsi::Runtime &rt, jsi::Object options, jsi::String taskId, jsi::String method, jsi::String url, jsi::Object headers, jsi::Array form, jsi::Function callback) = 0;
  virtual void fetchBlob(jsi::Runtime &rt, jsi::Object options, jsi::String taskId, jsi::String method, jsi::String url, jsi::Object headers, jsi::String body, jsi::Function callback) = 0;
  virtual jsi::Value createFile(jsi::Runtime &rt, jsi::String path, jsi::String data, jsi::String encoding) = 0;
  virtual jsi::Value createFileASCII(jsi::Runtime &rt, jsi::String path, jsi::Array data) = 0;
  virtual jsi::Value pathForAppGroup(jsi::Runtime &rt, jsi::String groupName) = 0;
  virtual jsi::String syncPathAppGroup(jsi::Runtime &rt, jsi::String groupName) = 0;
  virtual void exists(jsi::Runtime &rt, jsi::String path, jsi::Function callback) = 0;
  virtual jsi::Value writeFile(jsi::Runtime &rt, jsi::String path, jsi::String encoding, jsi::String data, bool transformFile, bool append) = 0;
  virtual jsi::Value writeFileArray(jsi::Runtime &rt, jsi::String path, jsi::Array data, bool append) = 0;
  virtual void writeStream(jsi::Runtime &rt, jsi::String path, jsi::String withEncoding, bool appendData, jsi::Function callback) = 0;
  virtual void writeArrayChunk(jsi::Runtime &rt, jsi::String streamId, jsi::Array withArray, jsi::Function callback) = 0;
  virtual void writeChunk(jsi::Runtime &rt, jsi::String streamId, jsi::String withData, jsi::Function callback) = 0;
  virtual void closeStream(jsi::Runtime &rt, jsi::String streamId, jsi::Function callback) = 0;
  virtual void unlink(jsi::Runtime &rt, jsi::String path, jsi::Function callback) = 0;
  virtual void removeSession(jsi::Runtime &rt, jsi::Array paths, jsi::Function callback) = 0;
  virtual jsi::Value ls(jsi::Runtime &rt, jsi::String path) = 0;
  virtual void stat(jsi::Runtime &rt, jsi::String target, jsi::Function callback) = 0;
  virtual void lstat(jsi::Runtime &rt, jsi::String path, jsi::Function callback) = 0;
  virtual void cp(jsi::Runtime &rt, jsi::String src, jsi::String dest, jsi::Function callback) = 0;
  virtual void mv(jsi::Runtime &rt, jsi::String path, jsi::String dest, jsi::Function callback) = 0;
  virtual jsi::Value mkdir(jsi::Runtime &rt, jsi::String path) = 0;
  virtual jsi::Value readFile(jsi::Runtime &rt, jsi::String path, jsi::String encoding, bool transformFile) = 0;
  virtual jsi::Value hash(jsi::Runtime &rt, jsi::String path, jsi::String algorithm) = 0;
  virtual void readStream(jsi::Runtime &rt, jsi::String path, jsi::String encoding, double bufferSize, double tick, jsi::String streamId) = 0;
  virtual void getEnvironmentDirs(jsi::Runtime &rt, jsi::Function callback) = 0;
  virtual void cancelRequest(jsi::Runtime &rt, jsi::String taskId, jsi::Function callback) = 0;
  virtual void enableProgressReport(jsi::Runtime &rt, jsi::String taskId, double interval, double count) = 0;
  virtual void enableUploadProgressReport(jsi::Runtime &rt, jsi::String taskId, double interval, double count) = 0;
  virtual jsi::Value slice(jsi::Runtime &rt, jsi::String src, jsi::String dest, double start, double end) = 0;
  virtual jsi::Value presentOptionsMenu(jsi::Runtime &rt, jsi::String uri, jsi::String scheme) = 0;
  virtual jsi::Value presentOpenInMenu(jsi::Runtime &rt, jsi::String uri, jsi::String scheme) = 0;
  virtual jsi::Value presentPreview(jsi::Runtime &rt, jsi::String uri, jsi::String scheme) = 0;
  virtual jsi::Value excludeFromBackupKey(jsi::Runtime &rt, jsi::String url) = 0;
  virtual void df(jsi::Runtime &rt, jsi::Function callback) = 0;
  virtual void emitExpiredEvent(jsi::Runtime &rt, jsi::Function callback) = 0;
  virtual jsi::Value actionViewIntent(jsi::Runtime &rt, jsi::String path, jsi::String mime, jsi::String chooserTitle) = 0;
  virtual jsi::Value addCompleteDownload(jsi::Runtime &rt, jsi::Object config) = 0;
  virtual jsi::Value copyToInternal(jsi::Runtime &rt, jsi::String contentUri, jsi::String destpath) = 0;
  virtual jsi::Value copyToMediaStore(jsi::Runtime &rt, jsi::Object filedata, jsi::String mt, jsi::String path) = 0;
  virtual jsi::Value createMediaFile(jsi::Runtime &rt, jsi::Object filedata, jsi::String mt) = 0;
  virtual jsi::Value getBlob(jsi::Runtime &rt, jsi::String contentUri, jsi::String encoding) = 0;
  virtual jsi::Value getContentIntent(jsi::Runtime &rt, jsi::String mime) = 0;
  virtual jsi::Value getSDCardDir(jsi::Runtime &rt) = 0;
  virtual jsi::Value getSDCardApplicationDir(jsi::Runtime &rt) = 0;
  virtual void scanFile(jsi::Runtime &rt, jsi::Array pairs, jsi::Function callback) = 0;
  virtual jsi::Value writeToMediaFile(jsi::Runtime &rt, jsi::String fileUri, jsi::String path, bool transformFile) = 0;

};

template <typename T>
class JSI_EXPORT NativeBlobUtilsCxxSpec : public TurboModule {
public:
  jsi::Value get(jsi::Runtime &rt, const jsi::PropNameID &propName) override {
    return delegate_.get(rt, propName);
  }

protected:
  NativeBlobUtilsCxxSpec(std::shared_ptr<CallInvoker> jsInvoker)
    : TurboModule("ReactNativeBlobUtil", jsInvoker),
      delegate_(static_cast<T*>(this), jsInvoker) {}

private:
  class Delegate : public NativeBlobUtilsCxxSpecJSI {
  public:
    Delegate(T *instance, std::shared_ptr<CallInvoker> jsInvoker) :
      NativeBlobUtilsCxxSpecJSI(std::move(jsInvoker)), instance_(instance) {}

    jsi::Object getConstants(jsi::Runtime &rt) override {
      static_assert(
          bridging::getParameterCount(&T::getConstants) == 1,
          "Expected getConstants(...) to have 1 parameters");

      return bridging::callFromJs<jsi::Object>(
          rt, &T::getConstants, jsInvoker_, instance_);
    }
    void fetchBlobForm(jsi::Runtime &rt, jsi::Object options, jsi::String taskId, jsi::String method, jsi::String url, jsi::Object headers, jsi::Array form, jsi::Function callback) override {
      static_assert(
          bridging::getParameterCount(&T::fetchBlobForm) == 8,
          "Expected fetchBlobForm(...) to have 8 parameters");

      return bridging::callFromJs<void>(
          rt, &T::fetchBlobForm, jsInvoker_, instance_, std::move(options), std::move(taskId), std::move(method), std::move(url), std::move(headers), std::move(form), std::move(callback));
    }
    void fetchBlob(jsi::Runtime &rt, jsi::Object options, jsi::String taskId, jsi::String method, jsi::String url, jsi::Object headers, jsi::String body, jsi::Function callback) override {
      static_assert(
          bridging::getParameterCount(&T::fetchBlob) == 8,
          "Expected fetchBlob(...) to have 8 parameters");

      return bridging::callFromJs<void>(
          rt, &T::fetchBlob, jsInvoker_, instance_, std::move(options), std::move(taskId), std::move(method), std::move(url), std::move(headers), std::move(body), std::move(callback));
    }
    jsi::Value createFile(jsi::Runtime &rt, jsi::String path, jsi::String data, jsi::String encoding) override {
      static_assert(
          bridging::getParameterCount(&T::createFile) == 4,
          "Expected createFile(...) to have 4 parameters");

      return bridging::callFromJs<jsi::Value>(
          rt, &T::createFile, jsInvoker_, instance_, std::move(path), std::move(data), std::move(encoding));
    }
    jsi::Value createFileASCII(jsi::Runtime &rt, jsi::String path, jsi::Array data) override {
      static_assert(
          bridging::getParameterCount(&T::createFileASCII) == 3,
          "Expected createFileASCII(...) to have 3 parameters");

      return bridging::callFromJs<jsi::Value>(
          rt, &T::createFileASCII, jsInvoker_, instance_, std::move(path), std::move(data));
    }
    jsi::Value pathForAppGroup(jsi::Runtime &rt, jsi::String groupName) override {
      static_assert(
          bridging::getParameterCount(&T::pathForAppGroup) == 2,
          "Expected pathForAppGroup(...) to have 2 parameters");

      return bridging::callFromJs<jsi::Value>(
          rt, &T::pathForAppGroup, jsInvoker_, instance_, std::move(groupName));
    }
    jsi::String syncPathAppGroup(jsi::Runtime &rt, jsi::String groupName) override {
      static_assert(
          bridging::getParameterCount(&T::syncPathAppGroup) == 2,
          "Expected syncPathAppGroup(...) to have 2 parameters");

      return bridging::callFromJs<jsi::String>(
          rt, &T::syncPathAppGroup, jsInvoker_, instance_, std::move(groupName));
    }
    void exists(jsi::Runtime &rt, jsi::String path, jsi::Function callback) override {
      static_assert(
          bridging::getParameterCount(&T::exists) == 3,
          "Expected exists(...) to have 3 parameters");

      return bridging::callFromJs<void>(
          rt, &T::exists, jsInvoker_, instance_, std::move(path), std::move(callback));
    }
    jsi::Value writeFile(jsi::Runtime &rt, jsi::String path, jsi::String encoding, jsi::String data, bool transformFile, bool append) override {
      static_assert(
          bridging::getParameterCount(&T::writeFile) == 6,
          "Expected writeFile(...) to have 6 parameters");

      return bridging::callFromJs<jsi::Value>(
          rt, &T::writeFile, jsInvoker_, instance_, std::move(path), std::move(encoding), std::move(data), std::move(transformFile), std::move(append));
    }
    jsi::Value writeFileArray(jsi::Runtime &rt, jsi::String path, jsi::Array data, bool append) override {
      static_assert(
          bridging::getParameterCount(&T::writeFileArray) == 4,
          "Expected writeFileArray(...) to have 4 parameters");

      return bridging::callFromJs<jsi::Value>(
          rt, &T::writeFileArray, jsInvoker_, instance_, std::move(path), std::move(data), std::move(append));
    }
    void writeStream(jsi::Runtime &rt, jsi::String path, jsi::String withEncoding, bool appendData, jsi::Function callback) override {
      static_assert(
          bridging::getParameterCount(&T::writeStream) == 5,
          "Expected writeStream(...) to have 5 parameters");

      return bridging::callFromJs<void>(
          rt, &T::writeStream, jsInvoker_, instance_, std::move(path), std::move(withEncoding), std::move(appendData), std::move(callback));
    }
    void writeArrayChunk(jsi::Runtime &rt, jsi::String streamId, jsi::Array withArray, jsi::Function callback) override {
      static_assert(
          bridging::getParameterCount(&T::writeArrayChunk) == 4,
          "Expected writeArrayChunk(...) to have 4 parameters");

      return bridging::callFromJs<void>(
          rt, &T::writeArrayChunk, jsInvoker_, instance_, std::move(streamId), std::move(withArray), std::move(callback));
    }
    void writeChunk(jsi::Runtime &rt, jsi::String streamId, jsi::String withData, jsi::Function callback) override {
      static_assert(
          bridging::getParameterCount(&T::writeChunk) == 4,
          "Expected writeChunk(...) to have 4 parameters");

      return bridging::callFromJs<void>(
          rt, &T::writeChunk, jsInvoker_, instance_, std::move(streamId), std::move(withData), std::move(callback));
    }
    void closeStream(jsi::Runtime &rt, jsi::String streamId, jsi::Function callback) override {
      static_assert(
          bridging::getParameterCount(&T::closeStream) == 3,
          "Expected closeStream(...) to have 3 parameters");

      return bridging::callFromJs<void>(
          rt, &T::closeStream, jsInvoker_, instance_, std::move(streamId), std::move(callback));
    }
    void unlink(jsi::Runtime &rt, jsi::String path, jsi::Function callback) override {
      static_assert(
          bridging::getParameterCount(&T::unlink) == 3,
          "Expected unlink(...) to have 3 parameters");

      return bridging::callFromJs<void>(
          rt, &T::unlink, jsInvoker_, instance_, std::move(path), std::move(callback));
    }
    void removeSession(jsi::Runtime &rt, jsi::Array paths, jsi::Function callback) override {
      static_assert(
          bridging::getParameterCount(&T::removeSession) == 3,
          "Expected removeSession(...) to have 3 parameters");

      return bridging::callFromJs<void>(
          rt, &T::removeSession, jsInvoker_, instance_, std::move(paths), std::move(callback));
    }
    jsi::Value ls(jsi::Runtime &rt, jsi::String path) override {
      static_assert(
          bridging::getParameterCount(&T::ls) == 2,
          "Expected ls(...) to have 2 parameters");

      return bridging::callFromJs<jsi::Value>(
          rt, &T::ls, jsInvoker_, instance_, std::move(path));
    }
    void stat(jsi::Runtime &rt, jsi::String target, jsi::Function callback) override {
      static_assert(
          bridging::getParameterCount(&T::stat) == 3,
          "Expected stat(...) to have 3 parameters");

      return bridging::callFromJs<void>(
          rt, &T::stat, jsInvoker_, instance_, std::move(target), std::move(callback));
    }
    void lstat(jsi::Runtime &rt, jsi::String path, jsi::Function callback) override {
      static_assert(
          bridging::getParameterCount(&T::lstat) == 3,
          "Expected lstat(...) to have 3 parameters");

      return bridging::callFromJs<void>(
          rt, &T::lstat, jsInvoker_, instance_, std::move(path), std::move(callback));
    }
    void cp(jsi::Runtime &rt, jsi::String src, jsi::String dest, jsi::Function callback) override {
      static_assert(
          bridging::getParameterCount(&T::cp) == 4,
          "Expected cp(...) to have 4 parameters");

      return bridging::callFromJs<void>(
          rt, &T::cp, jsInvoker_, instance_, std::move(src), std::move(dest), std::move(callback));
    }
    void mv(jsi::Runtime &rt, jsi::String path, jsi::String dest, jsi::Function callback) override {
      static_assert(
          bridging::getParameterCount(&T::mv) == 4,
          "Expected mv(...) to have 4 parameters");

      return bridging::callFromJs<void>(
          rt, &T::mv, jsInvoker_, instance_, std::move(path), std::move(dest), std::move(callback));
    }
    jsi::Value mkdir(jsi::Runtime &rt, jsi::String path) override {
      static_assert(
          bridging::getParameterCount(&T::mkdir) == 2,
          "Expected mkdir(...) to have 2 parameters");

      return bridging::callFromJs<jsi::Value>(
          rt, &T::mkdir, jsInvoker_, instance_, std::move(path));
    }
    jsi::Value readFile(jsi::Runtime &rt, jsi::String path, jsi::String encoding, bool transformFile) override {
      static_assert(
          bridging::getParameterCount(&T::readFile) == 4,
          "Expected readFile(...) to have 4 parameters");

      return bridging::callFromJs<jsi::Value>(
          rt, &T::readFile, jsInvoker_, instance_, std::move(path), std::move(encoding), std::move(transformFile));
    }
    jsi::Value hash(jsi::Runtime &rt, jsi::String path, jsi::String algorithm) override {
      static_assert(
          bridging::getParameterCount(&T::hash) == 3,
          "Expected hash(...) to have 3 parameters");

      return bridging::callFromJs<jsi::Value>(
          rt, &T::hash, jsInvoker_, instance_, std::move(path), std::move(algorithm));
    }
    void readStream(jsi::Runtime &rt, jsi::String path, jsi::String encoding, double bufferSize, double tick, jsi::String streamId) override {
      static_assert(
          bridging::getParameterCount(&T::readStream) == 6,
          "Expected readStream(...) to have 6 parameters");

      return bridging::callFromJs<void>(
          rt, &T::readStream, jsInvoker_, instance_, std::move(path), std::move(encoding), std::move(bufferSize), std::move(tick), std::move(streamId));
    }
    void getEnvironmentDirs(jsi::Runtime &rt, jsi::Function callback) override {
      static_assert(
          bridging::getParameterCount(&T::getEnvironmentDirs) == 2,
          "Expected getEnvironmentDirs(...) to have 2 parameters");

      return bridging::callFromJs<void>(
          rt, &T::getEnvironmentDirs, jsInvoker_, instance_, std::move(callback));
    }
    void cancelRequest(jsi::Runtime &rt, jsi::String taskId, jsi::Function callback) override {
      static_assert(
          bridging::getParameterCount(&T::cancelRequest) == 3,
          "Expected cancelRequest(...) to have 3 parameters");

      return bridging::callFromJs<void>(
          rt, &T::cancelRequest, jsInvoker_, instance_, std::move(taskId), std::move(callback));
    }
    void enableProgressReport(jsi::Runtime &rt, jsi::String taskId, double interval, double count) override {
      static_assert(
          bridging::getParameterCount(&T::enableProgressReport) == 4,
          "Expected enableProgressReport(...) to have 4 parameters");

      return bridging::callFromJs<void>(
          rt, &T::enableProgressReport, jsInvoker_, instance_, std::move(taskId), std::move(interval), std::move(count));
    }
    void enableUploadProgressReport(jsi::Runtime &rt, jsi::String taskId, double interval, double count) override {
      static_assert(
          bridging::getParameterCount(&T::enableUploadProgressReport) == 4,
          "Expected enableUploadProgressReport(...) to have 4 parameters");

      return bridging::callFromJs<void>(
          rt, &T::enableUploadProgressReport, jsInvoker_, instance_, std::move(taskId), std::move(interval), std::move(count));
    }
    jsi::Value slice(jsi::Runtime &rt, jsi::String src, jsi::String dest, double start, double end) override {
      static_assert(
          bridging::getParameterCount(&T::slice) == 5,
          "Expected slice(...) to have 5 parameters");

      return bridging::callFromJs<jsi::Value>(
          rt, &T::slice, jsInvoker_, instance_, std::move(src), std::move(dest), std::move(start), std::move(end));
    }
    jsi::Value presentOptionsMenu(jsi::Runtime &rt, jsi::String uri, jsi::String scheme) override {
      static_assert(
          bridging::getParameterCount(&T::presentOptionsMenu) == 3,
          "Expected presentOptionsMenu(...) to have 3 parameters");

      return bridging::callFromJs<jsi::Value>(
          rt, &T::presentOptionsMenu, jsInvoker_, instance_, std::move(uri), std::move(scheme));
    }
    jsi::Value presentOpenInMenu(jsi::Runtime &rt, jsi::String uri, jsi::String scheme) override {
      static_assert(
          bridging::getParameterCount(&T::presentOpenInMenu) == 3,
          "Expected presentOpenInMenu(...) to have 3 parameters");

      return bridging::callFromJs<jsi::Value>(
          rt, &T::presentOpenInMenu, jsInvoker_, instance_, std::move(uri), std::move(scheme));
    }
    jsi::Value presentPreview(jsi::Runtime &rt, jsi::String uri, jsi::String scheme) override {
      static_assert(
          bridging::getParameterCount(&T::presentPreview) == 3,
          "Expected presentPreview(...) to have 3 parameters");

      return bridging::callFromJs<jsi::Value>(
          rt, &T::presentPreview, jsInvoker_, instance_, std::move(uri), std::move(scheme));
    }
    jsi::Value excludeFromBackupKey(jsi::Runtime &rt, jsi::String url) override {
      static_assert(
          bridging::getParameterCount(&T::excludeFromBackupKey) == 2,
          "Expected excludeFromBackupKey(...) to have 2 parameters");

      return bridging::callFromJs<jsi::Value>(
          rt, &T::excludeFromBackupKey, jsInvoker_, instance_, std::move(url));
    }
    void df(jsi::Runtime &rt, jsi::Function callback) override {
      static_assert(
          bridging::getParameterCount(&T::df) == 2,
          "Expected df(...) to have 2 parameters");

      return bridging::callFromJs<void>(
          rt, &T::df, jsInvoker_, instance_, std::move(callback));
    }
    void emitExpiredEvent(jsi::Runtime &rt, jsi::Function callback) override {
      static_assert(
          bridging::getParameterCount(&T::emitExpiredEvent) == 2,
          "Expected emitExpiredEvent(...) to have 2 parameters");

      return bridging::callFromJs<void>(
          rt, &T::emitExpiredEvent, jsInvoker_, instance_, std::move(callback));
    }
    jsi::Value actionViewIntent(jsi::Runtime &rt, jsi::String path, jsi::String mime, jsi::String chooserTitle) override {
      static_assert(
          bridging::getParameterCount(&T::actionViewIntent) == 4,
          "Expected actionViewIntent(...) to have 4 parameters");

      return bridging::callFromJs<jsi::Value>(
          rt, &T::actionViewIntent, jsInvoker_, instance_, std::move(path), std::move(mime), std::move(chooserTitle));
    }
    jsi::Value addCompleteDownload(jsi::Runtime &rt, jsi::Object config) override {
      static_assert(
          bridging::getParameterCount(&T::addCompleteDownload) == 2,
          "Expected addCompleteDownload(...) to have 2 parameters");

      return bridging::callFromJs<jsi::Value>(
          rt, &T::addCompleteDownload, jsInvoker_, instance_, std::move(config));
    }
    jsi::Value copyToInternal(jsi::Runtime &rt, jsi::String contentUri, jsi::String destpath) override {
      static_assert(
          bridging::getParameterCount(&T::copyToInternal) == 3,
          "Expected copyToInternal(...) to have 3 parameters");

      return bridging::callFromJs<jsi::Value>(
          rt, &T::copyToInternal, jsInvoker_, instance_, std::move(contentUri), std::move(destpath));
    }
    jsi::Value copyToMediaStore(jsi::Runtime &rt, jsi::Object filedata, jsi::String mt, jsi::String path) override {
      static_assert(
          bridging::getParameterCount(&T::copyToMediaStore) == 4,
          "Expected copyToMediaStore(...) to have 4 parameters");

      return bridging::callFromJs<jsi::Value>(
          rt, &T::copyToMediaStore, jsInvoker_, instance_, std::move(filedata), std::move(mt), std::move(path));
    }
    jsi::Value createMediaFile(jsi::Runtime &rt, jsi::Object filedata, jsi::String mt) override {
      static_assert(
          bridging::getParameterCount(&T::createMediaFile) == 3,
          "Expected createMediaFile(...) to have 3 parameters");

      return bridging::callFromJs<jsi::Value>(
          rt, &T::createMediaFile, jsInvoker_, instance_, std::move(filedata), std::move(mt));
    }
    jsi::Value getBlob(jsi::Runtime &rt, jsi::String contentUri, jsi::String encoding) override {
      static_assert(
          bridging::getParameterCount(&T::getBlob) == 3,
          "Expected getBlob(...) to have 3 parameters");

      return bridging::callFromJs<jsi::Value>(
          rt, &T::getBlob, jsInvoker_, instance_, std::move(contentUri), std::move(encoding));
    }
    jsi::Value getContentIntent(jsi::Runtime &rt, jsi::String mime) override {
      static_assert(
          bridging::getParameterCount(&T::getContentIntent) == 2,
          "Expected getContentIntent(...) to have 2 parameters");

      return bridging::callFromJs<jsi::Value>(
          rt, &T::getContentIntent, jsInvoker_, instance_, std::move(mime));
    }
    jsi::Value getSDCardDir(jsi::Runtime &rt) override {
      static_assert(
          bridging::getParameterCount(&T::getSDCardDir) == 1,
          "Expected getSDCardDir(...) to have 1 parameters");

      return bridging::callFromJs<jsi::Value>(
          rt, &T::getSDCardDir, jsInvoker_, instance_);
    }
    jsi::Value getSDCardApplicationDir(jsi::Runtime &rt) override {
      static_assert(
          bridging::getParameterCount(&T::getSDCardApplicationDir) == 1,
          "Expected getSDCardApplicationDir(...) to have 1 parameters");

      return bridging::callFromJs<jsi::Value>(
          rt, &T::getSDCardApplicationDir, jsInvoker_, instance_);
    }
    void scanFile(jsi::Runtime &rt, jsi::Array pairs, jsi::Function callback) override {
      static_assert(
          bridging::getParameterCount(&T::scanFile) == 3,
          "Expected scanFile(...) to have 3 parameters");

      return bridging::callFromJs<void>(
          rt, &T::scanFile, jsInvoker_, instance_, std::move(pairs), std::move(callback));
    }
    jsi::Value writeToMediaFile(jsi::Runtime &rt, jsi::String fileUri, jsi::String path, bool transformFile) override {
      static_assert(
          bridging::getParameterCount(&T::writeToMediaFile) == 4,
          "Expected writeToMediaFile(...) to have 4 parameters");

      return bridging::callFromJs<jsi::Value>(
          rt, &T::writeToMediaFile, jsInvoker_, instance_, std::move(fileUri), std::move(path), std::move(transformFile));
    }

  private:
    T *instance_;
  };

  Delegate delegate_;
};

} // namespace react
} // namespace facebook
