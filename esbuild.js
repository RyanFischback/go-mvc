const esbuild = require("esbuild");
const { sassPlugin } = require("esbuild-sass-plugin");

esbuild
    .build({
        entryPoints: ["public/Application.tsx", "public/style.scss"],
        outdir: "public/assets",
        bundle: true,
        plugins: [sassPlugin()],
    })
    .then(() => console.log("⚡ Build complete! ⚡"))