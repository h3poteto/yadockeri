const path = require('path')

module.exports = {
  outputDir: path.resolve(__dirname, 'assets'),
  indexPath: path.resolve(__dirname, 'templates', 'index.html'),
  configureWebpack: {
    entry: path.resolve(__dirname, 'frontend', 'main.ts'),
  },
  publicPath: '/assets'
}
