// RSA1_5
// RSAES-PKCS1-v1_5.
// Used to encrypt the CEK, producing the JWE Encrypted Key,
// or to use key agreement to agree upon the CEK.
// As described in https://tools.ietf.org/html/rfc7518

import * as assert from "assert";
import * as crypto from "crypto";
import fs from "fs";

function getTitles(that: any): string {
  return `${that!.title}-${that!.test!.title}`;
}

describe(`RSA1_5 algorithm`, () => {
  // Made with: openssl req -new -newkey rsa:2048 -sha512 -sigopt rsa_padding_mode:pss -nodes -keyout PS512.pkey -subj "/C=US/ST=FL/L=DORAL/O=MICROSOFT/CN=AZURESDKTESTS"
  const private_key = fs.readFileSync("PS512.pkey", "utf-8");
  // Made with: openssl rsa -in PS512.pkey -pubout > PS512.pub
  const public_key = fs.readFileSync("PS512.pub", "utf-8");

  it("envrypt and decrypt", async function(): Promise<void> {
    const value = getTitles(this);
    const encrypted = crypto.publicEncrypt(public_key, Buffer.from(value, "utf8"));
    const decrypted = crypto.privateDecrypt(private_key, encrypted);
    assert.equal(decrypted, value);
  });
});
