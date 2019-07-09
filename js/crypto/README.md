# Testing Node Crypto


## In case we need to generate the files manually:

In case you need to check if Node supports a curve: `console.log(crypto.getCurves())`;

```typescript
// const ecdh = crypto.createECDH('secp256k1');
ecdh.generateKeys();
const public_key = ecdh.getPublicKey();
const private_key = ecdh.getPrivateKey();

// Inspired by: https://github.com/nodejs/node/issues/26133
// And: https://gist.github.com/canterberry/bf190ae6402265751e51725be535a4e4
const privateStart = Buffer.from("302e0201010420", "hex");
const privateEnd = Buffer.from("a00706052b8104000a", "hex");
const privateKeyPem = `-----BEGIN EC PRIVATE KEY-----
${Buffer.concat([privateStart, private_key, privateEnd]).toString("base64")}
-----END EC PRIVATE KEY-----`;
    const publicKeyPem = `-----BEGIN PUBLIC KEY-----
${Buffer.from(`3056301006072a8648ce3d020106052b8104000a034200${public_key.toString('hex')}`, 'hex').toString('base64')}
-----END PUBLIC KEY-----`;

const sign = crypto.createSign("SHA256");
const value = getTitles(this);
sign.write(value);
sign.end();
const signature = sign.sign(private_key);

console.log('0002')
const verifier = crypto.createVerify("SHA256");
verifier.update(value);
verifier.end();
assert.ok(verifier.verify(public_key, signature));
```
