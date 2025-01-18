// Copyright 2023-present Eser Ozvataf and other contributors. All rights reserved. Apache-2.0 license.
/// <reference lib="deno.ns" />

if (globalThis.Deno === undefined) {
  throw new Error("Deno is not defined");
}

export const deno = globalThis.Deno;
