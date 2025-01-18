// Copyright 2023-present Eser Ozvataf and other contributors. All rights reserved. Apache-2.0 license.
/// <reference lib="deno.unstable" />

import { deno } from "../deps.ts";

if (deno.openKv === undefined) {
  throw new Error("Deno.openKv() is not defined");
}

export class Connection {
  static instance: Connection = new Connection();
  kv: Deno.Kv | undefined = undefined;

  async getKv() {
    if (this.kv === undefined) {
      this.kv = await deno.openKv();
    }

    return this.kv;
  }
}
