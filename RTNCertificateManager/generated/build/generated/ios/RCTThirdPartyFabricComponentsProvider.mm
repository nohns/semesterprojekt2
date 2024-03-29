
/**
 * This code was generated by [react-native-codegen](https://www.npmjs.com/package/react-native-codegen).
 *
 * Do not edit this file as changes may cause incorrect behavior and will be lost
 * once the code is regenerated.
 *
 * @generated by GenerateRCTThirdPartyFabricComponentsProviderCpp
 */

// OSS-compatibility layer

#import "RCTThirdPartyFabricComponentsProvider.h"

#import <string>
#import <unordered_map>

Class<RCTComponentViewProtocol> RCTThirdPartyFabricComponentsProvider(const char *name) {
  static std::unordered_map<std::string, Class (*)(void)> sFabricComponentsClassMap = {


    {"RNCSafeAreaProvider", RNCSafeAreaProviderCls}, // safeareacontext,
    {"RNCSafeAreaView", RNCSafeAreaViewCls}, // safeareacontext

    {"RNSFullWindowOverlay", RNSFullWindowOverlayCls}, // rnscreens,
    {"RNSScreenContainer", RNSScreenContainerCls}, // rnscreens,
    {"RNSScreen", RNSScreenCls}, // rnscreens,
    {"RNSScreenNavigationContainer", RNSScreenNavigationContainerCls}, // rnscreens,
    {"RNSScreenStackHeaderConfig", RNSScreenStackHeaderConfigCls}, // rnscreens,
    {"RNSScreenStackHeaderSubview", RNSScreenStackHeaderSubviewCls}, // rnscreens,
    {"RNSScreenStack", RNSScreenStackCls}, // rnscreens,
    {"RNSSearchBar", RNSSearchBarCls}, // rnscreens

    {"RNSVGCircle", RNSVGCircleCls}, // rnsvg,
    {"RNSVGClipPath", RNSVGClipPathCls}, // rnsvg,
    {"RNSVGDefs", RNSVGDefsCls}, // rnsvg,
    {"RNSVGEllipse", RNSVGEllipseCls}, // rnsvg,
    {"RNSVGForeignObject", RNSVGForeignObjectCls}, // rnsvg,
    {"RNSVGGroup", RNSVGGroupCls}, // rnsvg,
    {"RNSVGImage", RNSVGImageCls}, // rnsvg,
    {"RNSVGSvgView", RNSVGSvgViewCls}, // rnsvg,
    {"RNSVGLinearGradient", RNSVGLinearGradientCls}, // rnsvg,
    {"RNSVGLine", RNSVGLineCls}, // rnsvg,
    {"RNSVGMarker", RNSVGMarkerCls}, // rnsvg,
    {"RNSVGMask", RNSVGMaskCls}, // rnsvg,
    {"RNSVGPath", RNSVGPathCls}, // rnsvg,
    {"RNSVGPattern", RNSVGPatternCls}, // rnsvg,
    {"RNSVGRadialGradient", RNSVGRadialGradientCls}, // rnsvg,
    {"RNSVGRect", RNSVGRectCls}, // rnsvg,
    {"RNSVGSymbol", RNSVGSymbolCls}, // rnsvg,
    {"RNSVGText", RNSVGTextCls}, // rnsvg,
    {"RNSVGTextPath", RNSVGTextPathCls}, // rnsvg,
    {"RNSVGTSpan", RNSVGTSpanCls}, // rnsvg,
    {"RNSVGUse", RNSVGUseCls}, // rnsvg
  };

  auto p = sFabricComponentsClassMap.find(name);
  if (p != sFabricComponentsClassMap.end()) {
    auto classFunc = p->second;
    return classFunc();
  }
  return nil;
}
