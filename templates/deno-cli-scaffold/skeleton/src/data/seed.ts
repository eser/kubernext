// Copyright 2023-present Eser Ozvataf and other contributors. All rights reserved. Apache-2.0 license.
import { ulid } from "$ulid/mod.ts";
import { Connection } from "./connection.ts";

export const seed = async () => {
  const kv = await Connection.instance.getKv();

  const profile = {
    id: ulid(),
    slug: "eser",
    name: "Eser Ozvataf",
    url: "https://eser.live",
  };

  await kv.set(["profile", "eser"], profile);
};

if (import.meta.main) {
  await seed();
}
