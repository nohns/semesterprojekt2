/* import {GreetRequest, GreetResponse} from './hello/v1/hello_service_pb';
import {MethodKind} from '@bufbuild/protobuf'; */
import {createPromiseClient} from '@bufbuild/connect';
import {GreetService} from './hello/v1/hello_service_connect';
import {createConnectTransport} from '@bufbuild/connect-web';
import {GreetResponse} from './hello/v1/hello_service_pb';
import {GreetRequest} from './hello/v1/hello_service_pb';
import {magicClient} from './fuckYouReactNative';

import 'fast-text-encoding';

const ok = new magicClient({
  service: GreetService,
  baseUrl: 'http://localhost:8080',
});

export async function greet() {
  const response = new GreetResponse();
  response.greeting = 'okokok';
  response.greeting1 = 'okokok';

  const yes = await ok.magic(response);
  console.log(yes);
}
