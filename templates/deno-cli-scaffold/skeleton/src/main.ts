// Copyright 2023-present Eser Ozvataf and other contributors. All rights reserved. Apache-2.0 license.
export const add = (a: number, b: number): number => {
  return a + b;
};

// Learn more at https://deno.land/manual/examples/module_metadata#concepts
if (import.meta.main) {
  console.log("Add 2 + 3 =", add(2, 3));
}
