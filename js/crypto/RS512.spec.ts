// RS512
// Digital Signature with RSASSA-PKCS1-v1_5 using SHA-512
// As described in https://tools.ietf.org/html/rfc7518

import * as assert from "assert";
import * as crypto from "crypto";
import fs from "fs";

function getTitles(that: any): string {
  return `${that!.title}-${that!.test!.title}`;
}
 
describe(`RS512
  Digital Signature with RSASSA-PKCS1-v1_5 using SHA-512`, () => {

  // A key of size 2048 bits or larger MUST be used with these algorithms.
  // Made with: openssl req -new -newkey rsa:2048 -sha512 -nodes -keyout RS512.pkey -subj "/C=US/ST=FL/L=DORAL/O=MICROSOFT/CN=AZURESDKTESTS" 
  const private_key = fs.readFileSync('RS512.pkey', 'utf-8');

  // Made with: openssl rsa -in RS512.pkey -pubout > RS512.pub
  const public_key = fs.readFileSync('RS512.pub', 'utf-8')

  it("sign and verify", async function(): Promise<void> {
    // For RSA keys, the default algorithm is RSASSA-PKCS1-v1_5
    const sign = crypto.createSign('sha512');
    const value = getTitles(this);
    sign.write(value);
    sign.end();
    const signature = sign.sign(private_key);

    const verifier = crypto.createVerify('sha512');
    verifier.update(value);
    verifier.end();
    assert.ok(verifier.verify(public_key, signature));
  });

});
