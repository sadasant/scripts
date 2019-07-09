// ES512
// Digital Signature with ECDSA using P-521 and SHA-512
// As described in https://tools.ietf.org/html/rfc7518

import * as assert from "assert";
import * as crypto from "crypto";
import fs from "fs";

function getTitles(that: any): string {
  return `${that!.title}-${that!.test!.title}`;
}

describe(`ES512
  Digital Signature with RSASSA-PSS using SHA-512 and MGF1 with SHA-512`, () => {
  // From: https://superuser.com/a/1103530
  // Made with: openssl ecparam -name secp521r1 -genkey -noout -out ES512.pkey
  const private_key = fs.readFileSync("ES512.pkey", "utf-8");

  // Made with: openssl ec -in ES512.pkey -pubout > ES512.pub
  const public_key = fs.readFileSync("ES512.pub", "utf-8");

  it("sign and verify", async function(): Promise<void> {
    const sign = crypto.createSign("SHA512");
    const value = getTitles(this);
    sign.write(value);
    sign.end();

    const signature = sign.sign(private_key);
    const verifier = crypto.createVerify("SHA512");
    verifier.update(value);
    verifier.end();
    assert.ok(verifier.verify(public_key, signature));
  });
});
