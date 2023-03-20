/* import {GreetRequest, GreetResponse} from './hello/v1/hello_service_pb';
import {MethodKind} from '@bufbuild/protobuf'; */

import {GreetService} from './hello/v1/hello_service_connect';

import {GreetRequest1} from './hello/v1/hello_service_pb';
import {GreetRequest} from './hello/v1/hello_service_pb';
import {magicClient} from './fuckYouReactNative';

import 'fast-text-encoding';

//magic client only supports json
const ok = new magicClient({
  service: GreetService,
  baseUrl: 'http://localhost:8080',
});

export async function greet() {
  const request = new GreetRequest1();
  request.greeting = 'Fuck you';

  const yes = await ok.magic(request);

  console.log(yes);

  const request1 = new GreetRequest();
  request1.greeting = 'Fuck you twice';
  const yes1 = await ok.magic(request1);
  console.log(yes1);
}
