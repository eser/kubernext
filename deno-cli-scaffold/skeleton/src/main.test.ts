import { assertEquals } from "$std/assert/mod.ts";
import { describe, it } from "$std/testing/bdd.ts";
import { add } from "./main.ts";

describe("add", () => {
  it("basic", () => {
    assertEquals(add(2, 3), 5);
  });
});
