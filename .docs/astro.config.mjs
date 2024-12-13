// @ts-check
import { defineConfig } from "astro/config";
import starlight from "@astrojs/starlight";

// https://astro.build/config
export default defineConfig({
  site: "https://saulova.github.io",
  base: "seam",
  integrations: [
    starlight({
      title: "🪡 Seam API Gateway",
      social: {
        github: "https://github.com/saulova/seam",
      },
      customCss: ["./src/styles/color-theme.css"],
      sidebar: [
        {
          label: "Guides",
          items: [{ label: "Quick Start", slug: "" }],
        },
        {
          label: "Components",
          autogenerate: { directory: "components" },
        },
        {
          label: "Plugins",
          autogenerate: { directory: "plugins" },
        },
      ],
      expressiveCode: {
        themes: ["github-dark-default", "github-light-default"],
      },
    }),
  ],
});
