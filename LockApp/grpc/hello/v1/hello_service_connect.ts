// @generated by protoc-gen-connect-es v0.8.3 with parameter "target=ts"
// @generated from file hello/v1/hello_service.proto (package hello.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { GreetRequest, GreetRequest1, GreetResponse, GreetResponse1 } from "./hello_service_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service hello.v1.GreetService
 */
export const GreetService = {
  typeName: "hello.v1.GreetService",
  methods: {
    /**
     * @generated from rpc hello.v1.GreetService.Greet
     */
    greet: {
      name: "Greet",
      I: GreetRequest,
      O: GreetResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc hello.v1.GreetService.Greet1
     */
    greet1: {
      name: "Greet1",
      I: GreetRequest1,
      O: GreetResponse1,
      kind: MethodKind.Unary,
    },
  }
} as const;

