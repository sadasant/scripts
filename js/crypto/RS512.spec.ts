// RS512
// RSASSA-PKCS1-v1_5 using SHA-512, as described in https://tools.ietf.org/html/rfc7518

import * as crypto from "crypto";
import * as assert from "assert";

async function process(processer: crypto.Cipher | crypto.Decipher, value: string): Promise<string> {
  let result = '';
  processer.on('readable', () => {
    let chunk;
    while (null !== (chunk = processer.read())) {
      result += chunk.toString('hex');
    }
  });
  processer.write(value);
  processer.end();
  return result;
}

function getTitles(that: any): string {
  return `${that!.title}-${that!.test!.title}`;
}
 
describe(`RS512
  RSASSA-PKCS1-v1_5 using SHA-512, as described in https://tools.ietf.org/html/rfc7518`, () => {

  const algorithm = "rsa_pkcs1_sha256";
  const password = "ornithorhynchus anatinus"

  it("encrypt and decrypt", async function(): Promise<void> {
    const value = getTitles(this);
    // A key of size 2048 bits or larger MUST be used with these algorithms.
    const key = crypto.scryptSync(password, 'salt', 24);
    // Use of an Initialization Vector (IV) of size 96 bits is REQUIRED with this algorithm.
    const iv = crypto.randomBytes(96);
    const cipher: crypto.Cipher = crypto.createCipheriv(algorithm, key, iv);
    const encrypted: string = await process(cipher, value);
    const decypher: crypto.Decipher = crypto.createDecipheriv(algorithm, key, iv);
    const decrypted: string = await process(decypher, encrypted);
    assert.equal(decrypted, value, "Unexpected decrypted value");
  });
  it("sign", async function(): Promise<void> {
  });
  it("verify", async function(): Promise<void> {
  });
  it("wrap", async function(): Promise<void> {
  });
  it("unwrap", async function(): Promise<void> {
  });
});
