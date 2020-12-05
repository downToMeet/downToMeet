module.exports = {
  sections: [
    {
      name: "Main components",
      components: "src/components/**/*.js",
      exampleMode: "hide",
      usageMode: "expand",
      ignore: ["src/components/common/*.js"],
    },
    {
      name: "Shared components",
      components: "src/components/common/*.js",
      exampleMode: "hide",
      usageMode: "expand",
    },
  ],
  ignore: [
    "**/__tests__/**",
    "**/*.test.{js,jsx,ts,tsx}",
    "**/*.spec.{js,jsx,ts,tsx}",
    "**/*.d.ts",
    "**/components/**/assets/*Logo.js",
    "**/components/*test.js",
  ],
  getComponentPathLine(componentPath) {
    // Cries in Windows
    return componentPath.replace(/\\/g, "/");
  },
};
