// ES256
// Digital Signature with ECDSA using P-256 and SHA-256
// As described in https://tools.ietf.org/html/rfc7518

import * as assert from "assert";
import * as crypto from "crypto";
import fs from "fs";

function getTitles(that: any): string {
  return `${that!.title}-${that!.test!.title}`;
}

describe(`ES256
  Digital Signature with Digital Signature with ECDSA using P-256 and SHA-256`, () => {
  // Made with: openssl ecparam -name secp256r1 -genkey -noout -out ES256.pkey
  const private_key = fs.readFileSync("ES256.pkey", "utf-8");

  // Made with: openssl ec -in ES256.pkey -pubout > ES256.pub
  const public_key = fs.readFileSync("ES256.pub", "utf-8");

  it("sign and verify", async function(): Promise<void> {
    const sign = crypto.createSign("SHA256");
    const value = getTitles(this);
    sign.write(value);
    sign.end();

    const signature = sign.sign(private_key);
    const verifier = crypto.createVerify("SHA256");
    verifier.update(value);
    verifier.end();
    assert.ok(verifier.verify(public_key, signature));
  });
});
