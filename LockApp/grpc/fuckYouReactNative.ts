import ReactNativeBlobUtil from 'react-native-blob-util';

import {Message, ServiceType} from '@bufbuild/protobuf';

export interface fuckyou {
  service: ServiceType;
  baseUrl: string;
}

export class magicClient {
  private service: ServiceType;
  private baseUrl: string;

  constructor(input: fuckyou) {
    this.service = input.service;
    this.baseUrl = input.baseUrl;
  }

  async magic<I extends Message<I>, O extends Message<O>>(response: O) {
    const json = response.toJsonString();
    try {
      const res = await ReactNativeBlobUtil.fetch(
        'POST',
        `${this.baseUrl}/${this.service.typeName}/${this.service.methods.greet.name}`,
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
