package com.rtncertificatemanager;

import androidx.annotation.Nullable;
import com.facebook.react.bridge.NativeModule;
import com.facebook.react.bridge.ReactApplicationContext;
import com.facebook.react.module.model.ReactModuleInfoProvider;
import com.facebook.react.TurboReactPackage;

import java.util.Collections;
import java.util.List;

public class CertificateManagerPackage extends TurboReactPackage {

  @Nullable
  @Override
  public NativeModule getModule(String name, ReactApplicationContext reactContext) {
          return null;
  }

  @Override
  public ReactModuleInfoProvider getReactModuleInfoProvider() {
      return null;
  }
}

node LockApp/node_modules/react-native/scripts/generate-codegen-artifacts.js \
  --path LockApp/ \
  --outputPath RTNCertificateManager/generated/