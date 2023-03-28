import type { TurboModule } from "react-native/Libraries/TurboModule/RCTExport";
import { TurboModuleRegistry } from "react-native";

export interface Spec extends TurboModule {
  // Create a CSR by generating a key pair and returning the CSR in PEM format.
  createCSR(): Promise<string>;
  // Store the certificate, given in PEM format, on the device
  storeCertificate(certificate: string): Promise<void>;
}

export default TurboModuleRegistry.get<Spec>(
  "RTNCertificateManager"
) as Spec | null;
