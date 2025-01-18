// Copyright 2023-present Eser Ozvataf and other contributors. All rights reserved. Apache-2.0 license.
import { add } from "./main.ts";

Deno.bench("addSmall", () => {
  add(1, 2);
});

Deno.bench("addBig", () => {
  add(2 ** 32, 2 ** 32);
});
