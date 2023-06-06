// @generated by protoc-gen-connect-es v0.8.3 with parameter "target=ts"
// @generated from file lock/v1/lock_service.proto (package lock.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { GetLockStateRequest, GetLockStateResponse, SetLockStateRequest, SetLockStateResponse } from "./lock_service_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service lock.v1.LockService
 */
export const LockService = {
  typeName: "lock.v1.LockService",
  methods: {
    /**
     * @generated from rpc lock.v1.LockService.GetLockState
     */
    getLockState: {
      name: "GetLockState",
      I: GetLockStateRequest,
      O: GetLockStateResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc lock.v1.LockService.SetLockState
     */
    setLockState: {
      name: "SetLockState",
      I: SetLockStateRequest,
      O: SetLockStateResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;
