// PS256
// Digital Signature with RSASSA-PSS using SHA-256 and MGF1 with SHA-256
// As described in https://tools.ietf.org/html/rfc7518

import * as assert from "assert";
import * as crypto from "crypto";
import fs from "fs";

function getTitles(that: any): string {
  return `${that!.title}-${that!.test!.title}`;
}

describe(`PS256
  Digital Signature with RSASSA-PSS using SHA-256 and MGF1 with SHA-256`, () => {
  // Based on https://tools.ietf.org/html/rfc7518,
  // a key of size 2048 bits or larger MUST be used with these algorithms.

  // From: https://stackoverflow.com/a/17795535
  // Apparenlty, MGF1 is the default for openssl.

  // From: https://crypto.stackexchange.com/a/32565
  // > OpenSSL supports 3 RSA paddings other than none: pkcs1 (more exactly the
  //   type-1 scheme of PKCS#1 through v1.5, now retronymed RSASSA-PKCS1-v1_5),
  //   pss, and x931. dgst -sign for RSA defaults to pkcs.
  // > Specify -sigopt rsa_padding_mode:pss to get randomized results.

  // Made with: openssl req -new -newkey rsa:2048 -sha256 -sigopt rsa_padding_mode:pss -nodes -keyout PS256.pkey -subj "/C=US/ST=FL/L=DORAL/O=MICROSOFT/CN=AZURESDKTESTS"
  const private_key = fs.readFileSync("PS256.pkey", "utf-8");

  // Made with: openssl rsa -in PS256.pkey -pubout > PS256.pub
  const public_key = fs.readFileSync("PS256.pub", "utf-8");

  it("sign and verify", async function(): Promise<void> {
    // For RSA keys, the default algorithm is RSASSA-PKCS1-v1_5
    const sign = crypto.createSign("sha256");
    const value = getTitles(this);
    sign.write(value);
    sign.end();
    const signature = sign.sign({
      key: private_key,
      passphrase: "",
      padding: crypto.constants.RSA_PKCS1_PSS_PADDING
    });

    const verifier = crypto.createVerify("sha256");
    verifier.update(value);
    verifier.end();
    assert.ok(
      verifier.verify(
        {
          key: public_key,
          padding: crypto.constants.RSA_PKCS1_PSS_PADDING
        },
        signature
      )
    );
  });
});
