// ES256
// Digital Signature with ECDSA using P-256 and SHA-256
// As described in https://tools.ietf.org/html/rfc7518

import * as assert from "assert";

// source: http://stackoverflow.com/a/11058858
function str2ab(str: string): ArrayBuffer {
  const bytes = new Uint8Array(str.length);
  for (let i = 0; i < str.length; i++) {
    bytes[i] = str.charCodeAt(i);
  }
  return bytes.buffer;
}

function getTitles(that: any): string {
  return `${that!.title}-${that!.test!.title}`;
}

describe(`ES256
  Digital Signature with ECDSA using P-256 and SHA-256`, () => {
  // Made with: openssl ecparam -name secp256r1 -genkey -noout -out ES256.pkey
  const private_key = `MHcCAQEEII6YeFTAdfDDgY1Y5YiGFuZfKRkGvUTqjCkJjAVg4aAhoAoGCCqGSM49AwEHoUQDQgAEst7RUTMHf5wE9KYG2eZpZpGISbJEnlK0KOrFMuGvdbf9aoqiG9pwF+qjHm8g06I/27LPXeyW/OtxYgveWA8kNw==`

  // Made with: openssl ec -in ES256.pkey -pubout > ES256.pub
  const public_key = `MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEst7RUTMHf5wE9KYG2eZpZpGISbJEnlK0KOrFMuGvdbf9aoqiG9pwF+qjHm8g06I/27LPXeyW/OtxYgveWA8kNw==`

  it("sign and verify", async function(): Promise<void> {
    const value = getTitles(this);
    const abValue = str2ab(value);

    console.log("0001");
    const privateKey = await window.crypto.subtle.importKey(
      "pkcs8",
      str2ab(atob(private_key)),
      {
        name: "ECDSA",
        namedCurve: "P-256"
      } as EcKeyGenParams,
      false,
      ['sign']
    );

    console.log("0002");
    const publicKey = await window.crypto.subtle.importKey(
      "spki",
      str2ab(atob(public_key)),
      {
        name: "ECDSA",
        namedCurve: "P-256"
      } as EcKeyGenParams,
      false,
      ['verify']
    );

    console.log("0003");
    const signature = await window.crypto.subtle.sign({
      name: "ECDSA",
      hash: { name: "SHA-256" }
    } as EcdsaParams, privateKey, abValue);

    console.log("0004");
    const isValid = await window.crypto.subtle.verify({
      name: "ECDSA",
      hash: { name: "SHA-256" }
    } as EcdsaParams, publicKey, signature, abValue);

    console.log("0005");
    assert.ok(isValid);
  });
});
