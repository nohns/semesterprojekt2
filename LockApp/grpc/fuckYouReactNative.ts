import ReactNativeBlobUtil from 'react-native-blob-util';

import {Message, ServiceType} from '@bufbuild/protobuf';

export interface fuckYouReactNative {
  service: ServiceType;
  baseUrl: string;
}

export class magicClient {
  private service: ServiceType;
  private baseUrl: string;

  constructor(input: fuckYouReactNative) {
    this.service = input.service;
    this.baseUrl = input.baseUrl;
  }

  async magic<I extends Message<I>, O extends Message<O>>(
    request: I,
  ): Promise<O | undefined> {
    const json = request.toJsonString();

    let method = null;
    for (const methodKey in this.service.methods) {
      if (methodKey === '') {
        throw new Error(
          'The type of ${this.service.typeName} has no methods that matches the request type ${request.constructor.name}',
        );
      }
      const t = this.service.methods[methodKey];

      if (t.I === request.constructor) {
        method = t.name;
        break;
      }
    }
    try {
      const res = await ReactNativeBlobUtil.fetch(
        'POST',
        `${this.baseUrl}/${this.service.typeName}/${method}`,
        {
          'content-type': 'application/json',
        },
        json,
      );
      return res.json();
    } catch (e) {
      console.log(e);
    }
  }
}
