// RSA-OAEP-256
// RSAES OAEP using SHA-256 and MGF1 with SHA-256.
// Used to encrypt the CEK, producing the JWE Encrypted Key,
// or to use key agreement to agree upon the CEK.
// As described in https://tools.ietf.org/html/rfc7518

import * as assert from "assert";
import * as crypto from "crypto";
import fs from "fs";

function getTitles(that: any): string {
  return `${that!.title}-${that!.test!.title}`;
}

describe(`RSA-OAEP-256 algorithm`, () => {
  // Made with: openssl req -new -newkey rsa:2048 -sha256 -nodes -keyout RS256.pkey -subj "/C=US/ST=FL/L=DORAL/O=MICROSOFT/CN=AZURESDKTESTS"
  const private_key = fs.readFileSync("RS256.pkey", "utf-8");
  // Made with: openssl rsa -in RS256.pkey -pubout > RS256.pub
  const public_key = fs.readFileSync("RS256.pub", "utf-8");

  // Common properties,
  // these are needed by typescript as part of crypto.KeyObject
  const commonProperties = {
    passphrase: "",
    export: (): string | Buffer => "",
    // MGF1 might not be used by default, but in theory it can be added on top.
    // Read more at: https://en.wikipedia.org/wiki/Mask_generation_function
    padding: crypto.constants.RSA_PKCS1_OAEP_PADDING
  };

  it("envrypt and decrypt", async function(): Promise<void> {
    const value = getTitles(this);
    const encrypted = crypto.publicEncrypt(
      {
        key: public_key,
        type: "public",
        ...commonProperties
      } as crypto.KeyObject,
      Buffer.from(value, "utf8")
    );
    const decrypted = crypto.privateDecrypt(
      {
        key: private_key,
        type: "private",
        ...commonProperties
      } as crypto.KeyObject,
      encrypted
    );
    assert.equal(decrypted, value);
  });
});
