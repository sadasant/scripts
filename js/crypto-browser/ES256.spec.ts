// ES256
// Digital Signature with ECDSA using P-256 and SHA-256
// As described in https://tools.ietf.org/html/rfc7518

// PLEASE GO HERE FOR MORE EXAMPLES: https://github.com/diafygi/webcrypto-examples/blob/master/index.html

import * as assert from "assert";
import * as asn1js from "asn1js";
import "@types/pkijs";
import ECPrivateKey from "pkijs/src/ECPrivateKey";
import ECPublicKey from "pkijs/src/ECPublicKey";

// source: http://stackoverflow.com/a/11058858
function str2ab(str: string): ArrayBuffer {
  const bytes = new Uint8Array(str.length);
  for (let i = 0; i < str.length; i++) {
    bytes[i] = str.charCodeAt(i);
  }
  return bytes.buffer;
}

function pem2jwk(key: string, type: string): any {
  const binaryKey = str2ab(atob(key));
  const asn1 = asn1js.fromBER(binaryKey);

  const result: {
    kty: string;
    crv: string;
    ext: boolean;
    x?: string;
    y?: string;
    d?: string;
  } = {
    kty: "EC",
    crv: "P-256",
    ext: true
  };

  let pkiKey: ECPrivateKey | ECPublicKey;
  let jsonKey: any;

  switch (type) {
    case "private":
      pkiKey = new ECPrivateKey({ schema: asn1.result });
      jsonKey = pkiKey.toJSON();
      result.x = jsonKey.x;
      result.y = jsonKey.y;
      result.d = jsonKey.d;
      break;
    default:
      pkiKey = new ECPublicKey({ schema: asn1.result });
      jsonKey = pkiKey.toJSON();
      result.x = jsonKey.x;
      result.y = jsonKey.y;
  }

  return result;
}

function getTitles(that: any): string {
  return `${that!.title}-${that!.test!.title}`;
}

describe(`ES256
  Digital Signature with ECDSA using P-256 and SHA-256`, () => {
  // Made with: openssl ecparam -name secp256r1 -genkey -noout -out ES256.pkey
  const private_key = `MHcCAQEEII6YeFTAdfDDgY1Y5YiGFuZfKRkGvUTqjCkJjAVg4aAhoAoGCCqGSM49AwEHoUQDQgAEst7RUTMHf5wE9KYG2eZpZpGISbJEnlK0KOrFMuGvdbf9aoqiG9pwF+qjHm8g06I/27LPXeyW/OtxYgveWA8kNw==`;

  // Made with: openssl ec -in ES256.pkey -pubout > ES256.pub
  const public_key = `MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEst7RUTMHf5wE9KYG2eZpZpGISbJEnlK0KOrFMuGvdbf9aoqiG9pwF+qjHm8g06I/27LPXeyW/OtxYgveWA8kNw==`;

  it("sign and verify", async function(): Promise<void> {
    const value = getTitles(this);
    const abValue = str2ab(value);

    const jwkPrivateKey = pem2jwk(private_key, "private");
    
    const ecGenParams: EcKeyGenParams = {
        name: "ECDSA",
        namedCurve: "P-256"
    }
    const ecdsaParams: EcdsaParams = {
        name: "ECDSA",
        hash: { name: "SHA-256" }
    }

    const privateKey = await window.crypto.subtle.importKey(
      "jwk",
      jwkPrivateKey,
      ecGenParams,
      false,
      ["sign"]
    );

    const publicKey = await window.crypto.subtle.importKey(
      "spki",
      str2ab(atob(public_key)),
      ecGenParams,
      false,
      ["verify"]
    );

    const signature = await window.crypto.subtle.sign(
      ecdsaParams,
      privateKey,
      abValue
    );

    const isValid = await window.crypto.subtle.verify(
      ecdsaParams,
      publicKey,
      signature,
      abValue
    );

    assert.ok(isValid);
  });
});
