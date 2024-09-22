import { defineConfig, Plugin } from "vite"
import { viteSingleFile } from "vite-plugin-singlefile"

/**
 * @param newFilename {string}
 * @returns {import('vite').Plugin}
 */
function renameIndexPlugin(filename: string): Plugin {
  if (!filename) {
		throw new Error('filename is required')
	}

  return {
    name: 'renameIndex',
    enforce: 'post',
    generateBundle(options, bundle) {
      const indexHtml = bundle['index.html']
      indexHtml.fileName = filename
    },
  }
}

export default defineConfig({
	plugins: [
		viteSingleFile(),
		renameIndexPlugin('template.html.tmpl'),
	],
	build: {
		outDir: "../pkg/pwdify/",
	}
})