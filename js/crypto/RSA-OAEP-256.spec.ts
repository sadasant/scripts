// RSA-OAEP-256
// RSAES OAEP using SHA-256 and MGF1 with SHA-256.
// Used to encrypt the CEK, producing the JWE Encrypted Key,
// or to use key agreement to agree upon the CEK.
// As described in https://tools.ietf.org/html/rfc7518

describe.skip(`RSA-OAEP-256 algorithm`, () => {
  // Node <= 12.6.0 doesn't support RSA-OAEP extensions, see:
  // - https://github.com/nodejs/node/pull/28335
  // - https://github.com/nodejs/node/pull/28335
});
