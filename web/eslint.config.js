import js from '@eslint/js'
import pluginVue from 'eslint-plugin-vue'
import oxlint from 'eslint-plugin-oxlint'
import skipFormatting from '@vue/eslint-config-prettier/skip-formatting'

export default [
  {
    name: 'app/files-to-lint',
    files: ['**/*.{js,mjs,jsx,vue}'],
  },

  {
    name: 'app/files-to-ignore',
    ignores: ['**/dist/**', '**/dist-ssr/**', '**/coverage/**'],
  },

  js.configs.recommended,
  ...pluginVue.configs['flat/essential'],
  oxlint.configs['flat/recommended'],
  skipFormatting,
  {
    rules: {
      'vue/max-attributes-per-line': 0,
      'vue/no-v-model-argument': 0,
      'vue/multi-word-component-names': 'off',
      'no-lone-blocks': 'off',
      'no-extend-native': 'off',
      'no-unused-vars': ['error', { argsIgnorePattern: '^_' }]
    }
  }
]
