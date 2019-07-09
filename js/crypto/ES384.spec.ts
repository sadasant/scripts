// ES384
// Digital Signature with ECDSA using P-384 and SHA-384
// As described in https://tools.ietf.org/html/rfc7518

import * as assert from "assert";
import * as crypto from "crypto";
import fs from "fs";

function getTitles(that: any): string {
  return `${that!.title}-${that!.test!.title}`;
}

describe(`ES384
  Digital Signature with Digital Signature with ECDSA using P-384 and SHA-384`, () => {
  // Made with: openssl ecparam -name secp384r1 -genkey -noout -out ES384.pkey
  const private_key = fs.readFileSync("ES384.pkey", "utf-8");

  // Made with: openssl ec -in ES384.pkey -pubout > ES384.pub
  const public_key = fs.readFileSync("ES384.pub", "utf-8");

  it("sign and verify", async function(): Promise<void> {
    const sign = crypto.createSign("SHA384");
    const value = getTitles(this);
    sign.write(value);
    sign.end();

    const signature = sign.sign(private_key);
    const verifier = crypto.createVerify("SHA384");
    verifier.update(value);
    verifier.end();
    assert.ok(verifier.verify(public_key, signature));
  });
});
