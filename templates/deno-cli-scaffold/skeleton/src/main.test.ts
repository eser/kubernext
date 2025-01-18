// Copyright 2023-present Eser Ozvataf and other contributors. All rights reserved. Apache-2.0 license.
import { assertEquals } from "$std/assert/mod.ts";
import { describe, it } from "$std/testing/bdd.ts";
import { add } from "./main.ts";

describe("add", () => {
  it("basic", () => {
    assertEquals(add(2, 3), 5);
  });
});
