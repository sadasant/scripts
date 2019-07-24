import nodeResolve from "rollup-plugin-node-resolve";
import multiEntry from "rollup-plugin-multi-entry";
import cjs from "rollup-plugin-commonjs";
import sourcemaps from "rollup-plugin-sourcemaps";
import shim from "rollup-plugin-shim";

const baseConfig = {
  context: null,
  treeshake: false,
  input: "dist-esm/*.spec.js",
  output: {
    file: "dist-test/index.js",
    format: "umd",
    name: "cryptobrowsertest",
    sourcemap: true
  },
  preserveSymlinks: false,
  plugins: [
    sourcemaps(),
    // os is not used by the browser bundle, so just shim it
    shim({
      dotenv: `export function config() { }`,
      os: `
        export const type = 1;
        export const release = 1;
      `
    }),
    nodeResolve({
      mainFields: ["module", "browser"],
      preferBuiltins: false
    }),
    cjs({
      namedExports: {
        "assert/": ["ok", "equal", "strictEqual"]
      }
    }),
    multiEntry({ exports: false })
  ]
};

const inputs = [baseConfig];
export default inputs;
